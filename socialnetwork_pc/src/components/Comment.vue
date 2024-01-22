<template>
    <div class="comment-container">
        <!--评论用户头像名称部分-->
        <div class="avatar-name">
            <div class="avatar"  >
                <img :src="commenter.avatar" alt="avatar"  @click.stop="personalpage" />
            </div>
            <div class="box">
                <div class="name"><p>{{commenter.username}}</p></div>
                <div class="datetime"><p>{{publish_time}}</p></div>
            </div>
        </div>
        <!--评论用户，评论内容部分-->
        <div class="comment-content">{{comment_text}}</div>
        <!--功能图标区-->
        <div class="power">
            <div class="zan">
                <img src="../assets/images/zan.png"/>
                <div class="text"><p>12</p></div>
            </div>
            <div class="pinglun" @click="reply">
                <img src="../assets/images/comment.png" >
                <div class="text"><p>评论</p></div>
            </div>
        </div>

        <Reply v-show="reply_state" v-for="reply in replylist" :key="reply.id"
                                                               :commenter="reply.commenter"
                                                               :comment_text="reply.comment_text"
                                                               :fmt_publish_time="reply.fmt_publish_time"
        >
        </Reply>


        <div class="reply" v-show="reply_state">
                <textarea class="reply-area" :placeholder="placeholder" v-model="reply_text"></textarea>
                <div class="btn" @click="sendreply"><p>回复</p></div>
        </div>
    </div>
</template>

<script>
import axios from "axios"
import {onMounted, ref} from "vue"
import Reply from "./Reply.vue"
export default{
    props:["id","post_id", "commenter", "publish_time", "comment_text"],  // 这里id是comment的id
    components:{
        Reply
    },

    setup(props){
        let reply_state = ref()
        let placeholder = ref()
        let reply_text = ref()
        placeholder.value = "回复@"+props.commenter.username
        reply_state.value = false

        // reply相关
        let replylist = ref([])

        const reply = () =>{
            if(reply_state.value == false){
                reply_state.value = true
            } else {
                reply_state.value = false
            }
        }

        const getReplyList = () => {
            let token = localStorage.getItem("jwt_token")
            axios({
                headers:{
                    Authorization: token,
                    'Content-Type':'application/x-www-form-urlencoded'
                },
                method: 'POST',
                url: "http://127.0.0.1:3000/socialnetwork/comment/getreplylist/",
                data: {
                    'post_id': props.post_id,
                    "parent_comment_id":props.id,
                }
                }).then(resp => {
                    replylist.value = resp.data.comment_list
                });
        }
        
        const sendreply = () =>{
            let token = localStorage.getItem("jwt_token")
            axios({
                headers:{
                    Authorization: token,
                    'Content-Type':'application/x-www-form-urlencoded'
                },
                method: 'POST',
                url: "http://127.0.0.1:3000/socialnetwork/comment/sendcomment/",
                data: {
                    "post_id": props.post_id,
                    "parent_comment_id":props.id,
                    "comment_text":reply_text.value
                }
                }).then(resp => {
                    console.log("reply resp: ", resp)
                });
        }

        onMounted(()=>{
            getReplyList()
        })

        return{
            reply_state,
            placeholder,
            reply_text,
            replylist,
            sendreply,
            getReplyList,
            reply,
        }
    }
}
</script>

<style scoped>
.comment-container{
    margin-top: 5vh;
    width: 100%;
}
.avatar-name{
    width: 100%;
    display: flex;
}
.avatar-name .avatar{
    float: left;
    width: 7vh;
    height: 7vh;
    padding-left: 10px;
    padding-top: 2px;
    margin-right: 10px;
}
.avatar-name .avatar img{
    height: 100%;
    width: 100%;
    border-radius: 50px;
}
.avatar-name .box{
    float: left;
    width: 40vw;
}
.avatar-name .box .name{
    font-size: 18px;
    color: rgba(0,0,0,0.5);
}
.avatar-name .box .datetime{
    font-size: 12px;
    color: rgba(0,0,0,0.7);
    margin-top: 1vh;
}
.comment-content{
    margin-top: 1.5vh;
    margin-left: 10vh;
    font-size: 18px;
}
/*评论功能区*/
.power{
    width: 100%;
    margin-top: 2vh;
    margin-left: 10vh;
    display: flex;
}
.power .zan{
    float: left;
    width: 10vh;
    height: 4vh;
}
.power .zan img{
    float: left;
    width: 3vh;
    height: 3vh;
    cursor: pointer;
}
.power .pinglun{
    float: left;
    width: 10vh;
    height: 4vh;
    cursor: pointer;
}
.power .pinglun img{
    float: left;
    width: 3vh;
    height: 3vh;
}
.power .text{
    float: left;
    margin-left: 0.5vh;
    color: rgba(0,0,0,0.7);
}

.reply{
    width: 46vw;
    height: 10vh;
    margin-left: 10vh;
}
.reply .reply-area{
    font-size: 18px;
    width: 46vw;
    height: 7vh;
    resize: none;
    /* 去除边框 */
    border: none;
    /* 去除选中后的边框 */
    outline:none;
    border-radius: 10px;
    background-color: rgb(242,243,245);
    margin-top: 4vh;
}
.reply .btn{
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
.reply .btn:hover{
    background-color: rgb(171,205,255);
}
</style>