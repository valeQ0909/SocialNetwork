package service

import (
	"log"
	"socialnetwork_back_go/config"
	"socialnetwork_back_go/models"
	"time"
)

type PostServiceImpl struct {
}

// Feed 通过传入时间戳，当前用户的id返回对应的帖子切片数组，以及帖子数组中最早的发布时间
// 施工ing
func (psi *PostServiceImpl) Feed(category string, lastTime time.Time, userId int64) ([]FmtPost, error) {
	//创建对应返回视频的切片数组，提前将切片的容量设置好，可以减少切片扩容的性能
	postsList := make([]FmtPost, 0, config.PostCount)

	//根据传入的时间，获得传入时间前n个视频，可以通过config.postCount来控制
	var posts []models.TablePost
	var err error
	if category == "推荐" || category == "热榜" || category == "关注" {
		posts, err = models.GetPostsByLastTime(lastTime)
	} else {
		posts, err = models.GetPostsDescByCategory(category)
	}

	if err != nil {
		log.Println("方法Feed获取按时间排序的视频失败: ", err)
		return nil, err
	}

	//将models层内的video加工成FmtVideo
	err = psi.refactorPosts(&postsList, &posts, userId)

	return postsList, nil
}

// GetFmtPostByPostId 传入帖子id获取对应的FmtPost对象
func (psi *PostServiceImpl) GetFmtPostByPostId(postId int64, userId int64) (FmtPost, error) {
	fmtPost := FmtPost{}
	//从数据库中查询数据，如果查询不到数据，就直接失败返回，后续流程就不需要执行了
	post, err := models.GetPostByPostId(postId)
	if err != nil {
		log.Println("方法GetFmtPostByPostId失败")
		return fmtPost, err
	}
	//插入从数据库中查到的数据
	psi.createFmtPost(&fmtPost, &post, userId)
	return fmtPost, nil
}

// SendPost 发送帖子
func (psi *PostServiceImpl) SendPost(authorId int64, category string, postHtml string, postText string) error {
	err := models.InsertTablePost(authorId, category, postHtml, postText)

	if err != nil {
		log.Println("帖子信息入库失败：", err)
		return err
	}

	log.Println("帖子信息入库成功")
	return nil
}

// GetPostList 获取具体用户帖子的列表
func (psi *PostServiceImpl) GetPostList(userId int64, curId int64) ([]FmtPost, error) {
	fmtPostsList := make([]FmtPost, 0, config.PostCount)
	posts, err := models.GetPostsByAuthorId(userId)

	err = psi.refactorPosts(&fmtPostsList, &posts, curId)
	if err != nil {
		log.Println("格式化Posts时发生错误：", err)
	}
	return fmtPostsList, nil
}

// PostCount 获取具体用户帖子的数量
func (psi *PostServiceImpl) PostCount(userId int64) (int64, error) {
	postIds, _ := models.GetPostIdList(userId)
	return int64(len(postIds)), nil
}

// 将Post格式化为FmtPost
func (psi *PostServiceImpl) refactorPosts(result *[]FmtPost, data *[]models.TablePost, userId int64) error {
	for _, temp := range *data {
		var fmtPost FmtPost
		// 将fmtPost进行组装，添加想要的信息
		psi.createFmtPost(&fmtPost, &temp, userId)
		*result = append(*result, fmtPost)
	}
	return nil
}

// 将fmtPost进行组装，添加想要的信息,插入从数据库中查到的数据
func (psi *PostServiceImpl) createFmtPost(fmtPost *FmtPost, post *models.TablePost, userId int64) {
	usi := UserServiceImpl{}
	lsi := LikeServiceImpl{}
	csi := CommentServiceImpl{}
	wsi := WatchServiceImpl{}

	fmtPost.TablePost = *post
	fmtPost.FmtPublishTime = post.PublishTime.Format("2006-01-02 15:04:05") // 输出: 2019-04-30 14:43:26
	//截取postText
	text := []rune(post.PostText) // 参考博客https://juejin.cn/post/6991734032534880270
	strText := ""
	if len(text) > 200 {
		text = text[0:200]
		strText = string(text)
		strText = strText + "......"
	} else {
		strText = string(text)
	}

	fmtPost.PostText = strText

	var err error
	fmtPost.Author, err = usi.GetFmtUserByIdWithCurId(post.AuthorId, userId)
	if err != nil {
		log.Println("方法createFmtPost失败")
	}

	// 查询帖子点赞总数
	fmtPost.FavoriteCount, _ = lsi.PostFavouriteCount(post.Id)
	// 查询当前用户是否点赞过这条帖子
	fmtPost.IsFavorite, _ = lsi.IsFavourite(post.Id, userId)
	// 查询帖子的评论数量
	fmtPost.CommentCount, _ = csi.CountByPostId(post.Id)
	// 查询帖子的浏览数量
	fmtPost.WatchCount, _ = wsi.PostWatchCnt(post.Id)
}
