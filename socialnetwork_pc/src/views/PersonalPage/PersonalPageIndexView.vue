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
                    <div class="num"><p>315</p></div>
                    <div class="text"><p>贴子</p></div>
                </div>
                <div class="box">
                    <div class="num"><p>7.3k</p></div>
                    <div class="text"><p>喜欢</p></div>
                </div>
                <div class="box">
                    <div class="num"><p>6.4k</p></div>
                    <div class="text"><p>粉丝</p></div>
                </div>

            </div>
            <div class="follow-message">
                <div class="follow"><p>关注</p></div>
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
import {onMounted, onUpdated, ref} from "vue"

export default{
    components:{
        Post,
    },

    setup(){
        let postlist = ref([])
        let avatar = ref()
        let username = ref()
        let signature = ref()

        const getUserInfo = () => {
            let token = localStorage.getItem("jwt_token")
            axios({
                headers:{
                    Authorization: token,
                    'Content-Type':'application/x-www-form-urlencoded'
                },
                method: 'GET',
                url: "http://127.0.0.1:3000/socialnetwork/user/token/",       
            }).then(response => {
                    avatar.value = response.data.user.avatar
                    signature.value = response.data.user.signature
                    username.value = response.data.user.username
                    ///let token = localStorage.getItem("jwt_token")
                    console.log("username: ", username.value)
                    axios({
                        headers:{
                            Authorization: token,
                            'Content-Type':'application/x-www-form-urlencoded'
                        },
                        method: 'POST',
                        url: "http://127.0.0.1:3000/socialnetwork/post/postlist/",
                        data: {              
                        'username': username.value,
                        }                
                    }).then(response => {
                            console.log("response: ", response)
                            postlist.value = response.data.post_list
                    });
            });
        }

        const getpostlist = () =>{

        }

        onMounted(()=>{
            getUserInfo()
            getpostlist()
        })
        onUpdated(()=>{
           
        })

        return{
            postlist,
            avatar,
            signature,
            username,
            getUserInfo,
            getpostlist
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
</style>