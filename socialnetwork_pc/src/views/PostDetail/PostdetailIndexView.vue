<template>
    <div class="container">
        <!--左边部分：点赞-->
        <div class="page-left">
            <div class="like">
                <img :src="require('../../assets/images/' + like)" @click.stop="likepost"/>
                <div class="text"></div>
            </div>
        </div>

        <!--中间部分：帖子具体内容区-->
        <div class="page-mid">
            <div class="postdetail">
                <div class="author">
                    <div class="zuozhe"><p>作者：</p></div>
                    <img class="avatar" :src="avatar"/>
                    <div class="box-author"><p>{{authorname}}</p></div>
                    <div class="datetime"><p>{{publish_time}}</p></div>
                </div>
                <div class="info" v-html="post_html"></div>
            </div>
            <div class="comment">
                <div class="box">评论 {{comment_count}}</div>
                <div class="avatar-comment">
                    <img class="avatar" src="../../assets/images/cat_2.jpg"/>
                    <div class="comment-box">
                        <textarea class="comment-area" :class="focus_state" @focus="comment_focus" @blur="lost_focus" placeholder="您可以在这里发表评论" v-model="comment_text"></textarea>
                        <div class="btn" @click="sendcomment"><p>发送</p></div>
                    </div>
                </div>
                <Comment v-for="comment in commentlist" :key="comment.id"
                                                        :id="comment.id"
                                                        :post_id="post_id"
                                                        :commenter="comment.commenter"
                                                        :reply_count="comment.reply_count"
                                                        :comment_text="comment.comment_text"
                                                        :publish_time="comment.fmt_publish_time"
                >
                </Comment>
            </div>
        </div>

        <!--发帖人信息部分-->
        <div class="page-right">
            <div class="avatar-name">
                <div class="avatar">
                    <img :src="avatar"/>
                </div>
                <div class="name">{{authorname}}</div>
                <div class="signature">{{signature}}</div>
            </div>

            <div class="number">
                <div class="box">
                    <div class="num"><p>{{post_cnt}}</p></div>
                    <div class="text"><p>贴子</p></div>
                </div>
                <div class="box">
                    <div class="num"><p>{{like_cnt}}</p></div>
                    <div class="text"><p>喜欢</p></div>
                </div>
                <div class="box">
                    <div class="num"><p>{{follower_cnt}}</p></div>
                    <div class="text"><p>粉丝</p></div>
                </div>
            </div>

            <div class="follow-message" v-show="!is_me">
                <div :class="follow_state" @click="followPoster"><p>{{p_follow}}</p></div>
                <div class="message"><p>私信</p></div>
            </div>

        </div>
    </div>
</template>

<script>
import axios from "axios"
import { ref,onMounted } from "vue";
import  Comment  from "../../components/Comment.vue";
import { useRouter } from "vue-router";

export default{
    components: {
        Comment,
    },

    setup(){
        let like = ref("like.png")  //关于动态切换图片的解决方案https://blog.csdn.net/tangshiyilang/article/details/134637734
        let focus_state = ref("pre_focus") //评论文本框是否聚焦
        const router = useRouter()
        let post_id = ref()
        post_id.value = router.currentRoute.value.query.id

        // postdetail相关数据
        let author_id = ref()
        let authorname = ref()
        let avatar = ref()
        let signature = ref()
        let publish_time = ref()
        let favorite_count = ref()
        let comment_count = ref()
        let post_html = ref()
        //let category = ref()  以后再加这个功能

        //comment相关数据、
        let commentlist = ref()
        let comment_text = ref()

        // 关注相关数据
        let is_me = ref()
        let is_follow = ref()
        let p_follow = ref()
        p_follow.value = "关注"
        let follow_actionType = 0
        let follow_state = ref()

        // 点赞相关数据
        let is_like = ref()
        let like_actionType = 0


        //用户info相关数据
        let post_cnt = ref()
        let like_cnt = ref()
        let follower_cnt = ref()


        const comment_focus = () =>{
            focus_state.value = "after_focus"
        }
        const lost_focus = () =>{
            focus_state.value = "pre_focus"
        }

        const likepost = () => {
            if(is_like.value==true){
                like_actionType = 1
            } else {
                like_actionType = 0
            }
            let token = localStorage.getItem("jwt_token")
            axios({
                headers:{
                    Authorization: token,
                    'Content-Type':'application/x-www-form-urlencoded'
                },
                method: 'POST',
                url: "http://127.0.0.1:3000/socialnetwork/favorite/action/",
                data: {
                    'post_id': post_id.value,
                    'action_type': like_actionType,
                }
                }).then(resp => {
                    console.log("resp: ", resp)
                    if(like.value == "like.png"){
                        like.value = "liked.png"
                    }else{
                        like.value = "like.png"
                    }
                });
        }

        const followPoster = () => {
            if(is_follow.value==true){
                follow_actionType = 1
            } else {
                follow_actionType = 0
            }
            let token = localStorage.getItem("jwt_token")
            axios({
                headers:{
                    Authorization: token,
                    'Content-Type':'application/x-www-form-urlencoded'
                },
                method: 'POST',
                url: "http://127.0.0.1:3000/socialnetwork/follow/action/",
                data: {
                    "user_id": author_id.value,
                    "action_type": follow_actionType,
                }
                }).then(resp => {
                    console.log("followPoster resp: ", resp)
                    if(resp.data.status_code == 0){
                        if(is_follow.value){ //如果之前是关注状态，现在就是未关注状态
                            follow_state.value = "follow"
                            p_follow.value = "关注"
                            is_follow.value = false
                        } else { //如果之前是未关注状态，现在就是关状态
                            follow_state.value = "followed"
                            p_follow.value = "已关注"
                            is_follow.value = true
                        }
                    }
                });
        }

        //获取帖子的详细数据
        const getPostDetail = () => {
            let token = localStorage.getItem("jwt_token")
            axios({
                headers:{
                    Authorization: token,
                    'Content-Type':'application/x-www-form-urlencoded'
                },
                method: 'POST',
                url: "http://127.0.0.1:3000/socialnetwork/post/postdetail/",
                data: {
                    'post_id': post_id.value,
                }
                }).then(resp => {
                    console.log("postDetail: ", resp)
                    author_id.value = resp.data.post.author_id
                    authorname.value = resp.data.post.author.username
                    avatar.value = resp.data.post.author.avatar
                    signature.value = resp.data.post.author.signature
                    publish_time.value = resp.data.post.fmt_publish_time
                    post_cnt.value = resp.data.post.author.post_cnt
                    like_cnt.value = resp.data.post.author.like_cnt
                    follower_cnt.value = resp.data.post.author.follower_cnt

                    post_html.value = resp.data.post.post_html
                    // 点赞相关
                    if(resp.data.post.author.is_me == true){  // 这里json解析过后会自动变成bool类型
                        is_me.value = true
                    } else {
                        is_me.value = false
                    }
                    is_like.value = resp.data.post.is_favorite
                    if(resp.data.post.is_favorite == true){ // 如果喜欢过该帖子，则like图标显示红心
                        like.value = "liked.png"
                    }
                    // 关注相关
                    is_follow.value = resp.data.post.author.is_follow
                    if(resp.data.post.author.is_follow == true){ // 如果关注过发帖人，则显示已关注
                        p_follow.value = "已关注"
                        follow_state.value = "followed"
                    } else {
                        p_follow.value = "关注"
                        follow_state.value = "follow"
                    }
                });
        }

        const getCommentList = () => {
            let token = localStorage.getItem("jwt_token")
            axios({
                headers:{
                    Authorization: token,
                    'Content-Type':'application/x-www-form-urlencoded'
                },
                method: 'POST',
                url: "http://127.0.0.1:3000/socialnetwork/comment/getcommentlist/",
                data: {
                    'post_id': post_id.value,
                }
                }).then(resp => {
                    console.log("commentlist: ",resp.data.comment_list)
                    commentlist.value = resp.data.comment_list
                });
        }

        const sendcomment = () => {
            let token = localStorage.getItem("jwt_token")
            axios({
                headers:{
                    Authorization: token,
                    'Content-Type':'application/x-www-form-urlencoded'
                },
                method: 'POST',
                url: "http://127.0.0.1:3000/socialnetwork/comment/sendcomment/",
                data: {
                    'post_id': post_id.value,
                    'parent_comment_id': 0,
                    'comment_text': comment_text.value,
                }
                }).then(resp => {
                    comment_text.value = ""
                    alert("发表评论成功！")
                    console.log("send comment resp: ",resp.data);
                });
        }


        onMounted(()=>{
            getPostDetail()
            getCommentList()
        })

        return{
            like,
            focus_state,
            author_id,
            authorname,
            avatar,
            signature,
            publish_time,
            favorite_count,
            comment_count,
            post_html,
            comment_text,
            commentlist,
            post_id,
            is_me,
            is_like,
            is_follow,
            p_follow,
            follow_state,
            post_cnt,
            like_cnt,
            follower_cnt,
            likepost,
            followPoster,
            comment_focus,
            lost_focus,
            getPostDetail,
            sendcomment,
            getCommentList,
        }
        
    }
}
</script>

<style scoped>
.pre_focus{
  height: 10vh;
  border-radius: 10px;
}

.after_focus{
  height: 20vh;
  border-radius: 10px;
}
.container{
    width: 100%;
    height: 100%;
    margin-top: 10vh;
    padding-top: 1vh;
    background-color: rgb(242,243,245);
    overflow:hidden; /*父盒子内有的任一级子盒子有浮动会导致父盒子无法高度自适应（即被子盒子撑开）*/
}
.page-left{
    width: 10vw;
    height: 70vh;
    margin-left: 8vw;
    position: fixed;
    top: 11vh;
    z-index: 1111;
}
.page-left .like{
    height: 8vh;
    width: 8vh;
    border-radius: 50px;
    margin-top: 20vh;
    margin-left: 12vh;
    background-color: white;
    cursor: pointer;
}
.page-left .like img{
    margin-top: 1vh;
    margin-left: 1vh;
    height: 6vh;
    width: 6vh;
}


.page-mid{
    float: left;
    margin-left: 19vw;
    width: 56vw;
}
.page-mid .postdetail{
    width: 100%;
    height:100%;
    background-color: white;
}
.page-mid .postdetail .author{
    height: 10vh;
    padding-top: 5vh;
}
.page-mid .postdetail .author .zuozhe{
    float: left;
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 24px;
    margin-left: 2vw;
    margin-top: 0.5vw;
    color: rgba(0,0,0,0.7);
}
.page-mid .postdetail .author img{
    display: block;
    float: left;
    width: 3vw;
    height: 3vw;
    border-radius: 50px;
}
.page-mid .postdetail .author .box-author{
    float: left;
    height: 2vw;
    font-size: 24px;
    display: flex;
    justify-content: center;
    align-items: center;
    color: rgb(0, 200, 255);
    margin-left: 1vw;
    margin-top: 0.5vw;
}
.page-mid .postdetail .author .datetime{
    float: left;
    font-size: 14px;
    margin-left: 1vw;
    margin-top: 1vw;
    color: rgba(0,0,0,0.7);
}
.page-mid .postdetail .info{
    margin-left: 2vw;
    margin-right: 2vw;
    font-size: 18px;
    padding-bottom: 20vh;
}
.page-mid .comment{
    width: 56vw;
    margin-top: 10vh;
    padding-left: 2vw;
    padding-bottom: 10vh;
    margin-bottom: 10vh;
    box-sizing: border-box;
    background-color: white;
}
.page-mid .comment .box{
    padding-top: 2vw;
    font-size: 24px;
}
.page-mid .comment .avatar-comment{
    width: 54vw;
    margin-top: 4vh;
    display: flex;
}
.page-mid .comment .avatar-comment .avatar{
    display: block;
    float: left;
    width: 3vw;
    height: 3vw;
    border-radius: 50px;
}
.comment .avatar-comment .comment-box{
    float: left;
    width: 48vw;
    margin-left: 1vw;
    border-radius: 10px;
    background-color: rgb(242,243,245);
}
.comment .avatar-comment .comment-box .comment-area{
    font-size: 18px;
    width: 47vw;
    margin-left: 0.5vw;
    margin-top: 0.5vw;
    resize: none;
    /* 去除边框 */
    border: none;
    /* 去除选中后的边框 */
    outline:none;
    background-color: rgb(242,243,245);
}
.comment .avatar-comment .comment-box .btn{
    float: right;
    display: flex;
    justify-content: center;
    align-items: center;
    width: 5vw;
    height: 2vw;
    margin-right: 0.5vw;
    margin-bottom: 0.5vw;
    border-radius: 5px;
    background-color: rgb(30,128,255);
    color: white;
    cursor: pointer;
}
.comment .avatar-comment .comment-box .btn:hover{
    background-color: rgb(171,205,255);
}

.page-right{
    width: 20vw;
    border-radius: 2px;
    position: fixed;
    top: 11vh;
    right: 2vw;
    padding-bottom: 5vh;
    z-index: 1000;
    background-color: white;
}

.avatar-name{
    width: 100%;
    height: 10vh;
}
.avatar-name .avatar{
    float: left;
    width: 8vh;
    height: 8vh;
    box-sizing: border-box;
    padding-left: 10px;
    padding-top: 10px;
    margin-right: 10px;
}
.avatar-name .avatar img{
    height: 100%;
    width: 100%;
    border-radius: 50px;
}
.avatar-name .name{
    font-size: 20px;
}
.avatar-name .signature{
    font-size: 14px;
    color: rgba(0, 0, 0, 0.7);
    margin-top: 1vh;
}
.number{
    width: 100%;
    height: 8vh;
    box-sizing: border-box;
    display: flex;
    justify-content: center;
    align-items: center;
}
.number .box{
    float: left;
    width: 6.7vw;
    text-align: center;
}
.number .box .num{
    font-size: 18px;
}
.number .box .text{
    font-size: 16px;
    color: rgba(0, 0, 0, 0.7);
}
.follow-message{
    width: 100%;
    height: 100%;
    margin-top: 1vw;
    
}
.follow-message .follow{
    float: left;
    width: 8vw;
    height: 2.5vw;
    margin-left: 1vw;
    font-size: 16px;
    border-radius: 5px;
    display: flex;
    justify-content: center;
    align-items: center;
    cursor: pointer;
    color: white;
    background-color: rgb(30,128,255);
}
.follow-message .followed{
    float: left;
    width: 8vw;
    height: 2.5vw;
    margin-left: 1vw;
    font-size: 16px;
    border-radius: 5px;
    display: flex;
    justify-content: center;
    align-items: center;
    cursor: pointer;
    color: rgb(122, 180, 255);
    background-color: rgb(216, 233, 255);
}
.follow-message .follow:hover{
    background-color: rgb(122, 180, 255);
}
.follow-message .followed:hover{
    background-color: rgb(207, 227, 254);
}
.follow-message .message{
    float: left;
    width: 8vw;
    height: 2.5vw;
    font-size: 16px;
    border-radius: 5px;
    margin-left: 2vw;
    display: flex;
    justify-content: center;
    align-items: center;
    cursor: pointer;
    color: rgb(122, 180, 255);
    background-color: rgb(216, 233, 255);
}
.follow-message .message:hover{
    background-color: rgb(207, 227, 254);
}
</style>