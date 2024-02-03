<template>
    <div class="container">
        <!--左边部分-->
        <div class="page-left">
            <div class="avatar-name">
                <div class="avatar">
                    <img :src="avatar"/>
                </div>
                <div class="name">{{username}}</div>
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
                <div :class="follow_state" @click="followUser"><p>{{p_follow}}</p></div>
                <div class="message"><p>私信</p></div>
            </div>
        </div>

        <!--中间部分-->
        <div class="page-mid">
            <Post v-for="post in postlist" :key="post.id"
                                           :id="post.id"
                                           :author="post.author"
                                           :category="post.category"
                                           :comment_count="post.comment_count"
                                           :favorite_count="post.favorite_count"
                                           :post_text="post.post_text"
                                           :fmt_publish_time="post.fmt_publish_time"
            >
            </Post>
        </div>

        <!--右边部分：用来撑开页面-->
        <div class="page-right">

        </div>

    </div>
</template>

<script>
import Post from '../../components/Post.vue';
import axios from "axios";
import {onMounted, ref} from "vue"
import { useRouter } from "vue-router";

export default{
    components:{
        Post,
    },

    setup(){
        let postlist = ref([])
        let avatar = ref()
        let username = ref()
        let signature = ref()
        const router = useRouter()
        username = router.currentRoute.value.query.username

        let userId = ref()

        // 关注相关
        let is_me = ref()
        let is_follow = ref()
        let p_follow = ref()
        p_follow.value = "关注"
        let follow_actionType = 0
        let follow_state = ref()

        // 用户info相关
        let post_cnt = ref()
        let like_cnt = ref()
        let follower_cnt = ref()


        const getUserInfo = () => {
            let token = localStorage.getItem("jwt_token")
            axios({
                headers:{
                    Authorization: token,
                    'Content-Type':'application/x-www-form-urlencoded'
                },
                method: 'POST',
                url: "http://127.0.0.1:3000/socialnetwork/user/info/",
                data: {              
                'username': username,
                }                
            }).then(resp => {
                    userId.value = resp.data.user.id
                    avatar.value = resp.data.user.avatar
                    signature.value = resp.data.user.signature
                    is_me.value = resp.data.user.is_me
                    is_follow.value = resp.data.user.is_follow
                    post_cnt.value = resp.data.user.post_cnt
                    like_cnt.value = resp.data.user.like_cnt
                    follower_cnt.value = resp.data.user.follower_cnt
                    if(is_follow.value){
                            follow_state.value = "followed"
                            p_follow.value = "已关注"
                            is_follow.value = true
                    } else {
                            follow_state.value = "follow"
                            p_follow.value = "关注"
                            is_follow.value = false
                    }
            });
        }

        const getpostlist = () =>{
            let token = localStorage.getItem("jwt_token")
            axios({
                headers:{
                    Authorization: token,
                    'Content-Type':'application/x-www-form-urlencoded'
                },
                method: 'POST',
                url: "http://127.0.0.1:3000/socialnetwork/post/postlist/",
                data: {              
                'username': username,
                }                
            }).then(resp => {
                    console.log("response: ", resp)
                    postlist.value = resp.data.post_list
            });
        }


        // 关注用户
        const followUser = () =>{
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
                    "user_id": userId.value,
                    "action_type": follow_actionType,
                }
                }).then(resp => {
                    console.log("followPoster resp: ", resp)
                    if(resp.data.status_code == 0){
                        if(is_follow.value){ //如果之前是关注状态，现在就是未关注状态
                            follow_state.value = "follow"
                            p_follow.value = "关注"
                            is_follow.value = false
                            follower_cnt.value = follower_cnt.value - 1
                        } else { //如果之前是未关注状态，现在就是关注状态
                            follow_state.value = "followed"
                            p_follow.value = "已关注"
                            is_follow.value = true
                            follower_cnt.value = follower_cnt.value + 1
                        }
                    }
                });
        }


        onMounted(()=>{
            getUserInfo()
            getpostlist()
        })

        return{
            postlist,
            avatar,
            signature,
            username,
            userId,
            is_me,
            is_follow,
            p_follow,
            follow_state,
            post_cnt,
            like_cnt,
            follower_cnt,
            getUserInfo,
            getpostlist,
            followUser
        }
    }
}
</script>

<style scoped>
.container{
    width: 100%;
    height: 100%;
    margin-top: 10vh;
    padding-top: 1vh;
    background-color: rgb(242,243,245);
    overflow:hidden; /*父盒子内有的任一级子盒子有浮动会导致父盒子无法高度自适应（即被子盒子撑开）*/
}
.page-left{
    width: 20vw;
    height: 30vh;
    margin-left: 1vw;
    position: fixed;
    top: 11vh;
    z-index: 1111;
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
.follow-message .follow:hover{
    background-color: rgb(122, 175, 245);
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

.page-mid{
    float: left;
    margin-left: 23vw;
    height: 100%;
    width: 60vw;
    background-color: white;
}
.page-right{
    float: left;
    width: 2vw;
    height: 90vh;
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