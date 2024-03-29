package service

import (
	"log"
	"socialnetwork_back_go/models"
	"time"
)

type CommentServiceImpl struct {
	UserService
}

// CountByPostId 计算帖子下有多少评论
func (csi *CommentServiceImpl) CountByPostId(postId int64) (int64, error) {
	count, err := models.CommentCount(postId)
	return count, err
}

// SendComment 发送评论
func (csi *CommentServiceImpl) SendComment(commenterId int64, postId int64, parentCommentId int64, commentText string) error {
	newComment := models.TableComment{
		CommenterId:     commenterId,
		PostId:          postId,
		ParentCommentId: parentCommentId,
		CommentText:     commentText,
		PublishTime:     time.Now(),
		Cancel:          0,
	}
	err := models.InsertComment(newComment)
	if err != nil {
		log.Println("发送评论失败")
	}
	return err
}

// DeleteComment 删除评论
// 施工ing
func (csi *CommentServiceImpl) DeleteComment(commentId int64) error {
	return nil
}

// GetFmtCommentList 获取评论列表
func (csi *CommentServiceImpl) GetFmtCommentList(postId int64, userId int64) ([]FmtComment, error) {
	var commentList []models.TableComment
	commentList, _ = models.GetCommentList(postId, 0) //获取评论列表

	var fmtCommentList []FmtComment
	err := csi.refactorComments(&fmtCommentList, &commentList, userId)

	if err != nil {
		log.Println("GetCommentList方法失败")
	}
	return fmtCommentList, err
}

// GetFmtReplyList 获取帖子回复列表
func (csi *CommentServiceImpl) GetFmtReplyList(postId int64, parentCommentId int64, userId int64) ([]FmtComment, error) {
	var replyList []models.TableComment
	replyList, _ = models.GetCommentList(postId, parentCommentId)
	var fmtReplyList []FmtComment
	err := csi.refactorComments(&fmtReplyList, &replyList, userId)

	if err != nil {
		log.Println("GetReplyList方法失败")
	}
	return fmtReplyList, err
}

// refactorComments 将comment格式化为FmtComment
func (csi *CommentServiceImpl) refactorComments(result *[]FmtComment, data *[]models.TableComment, userId int64) error {
	for _, temp := range *data {
		var fmtComment FmtComment
		// 将fmtPost进行组装，添加想要的信息
		csi.createFmtComment(&fmtComment, &temp, userId)
		*result = append(*result, fmtComment)
	}
	return nil
}
func (csi *CommentServiceImpl) createFmtComment(fmtComment *FmtComment, comment *models.TableComment, userId int64) {
	usi := UserServiceImpl{}
	fmtComment.TableComment = *comment
	fmtComment.FmtPublishTime = comment.PublishTime.Format("2006-01-02 15:04:05") // 输出: 2019-04-30 14:43:26

	// 查询comment的评论人
	var err error
	fmtComment.Commenter, err = usi.GetFmtUserByIdWithCurId(comment.CommenterId, userId)
	if err != nil {
		log.Printf("方法usi.GetFmtUserByIdWithCurId(comment.CommenterId, userId) 失败：%v", err)
	} else {
		log.Printf("方法usi.GetFmtUserByIdWithCurId(comment.CommenterId, userId) 成功")
	}

	// 查询comment的回复数量
	var replyList []models.TableComment
	replyList, _ = models.GetCommentList(comment.PostId, comment.Id) //查询自己的reply列表
	fmtComment.ReplyCount = int64(len(replyList))
}
