<template>
    <div class="post"  @click="postdetail">
        <!--头像、名称区-->
        <div class="avatar_name">
            <div class="avatar"  >
                <img :src="author.avatar" alt="avatar"  @click.stop="personalpage" />
            </div>
            <div class="name">
                {{author.username}}
            </div>
            <div class="datetime">
                {{fmt_publish_time}}
            </div>
        </div>

        <!--文章摘要区-->
        <div class="summary">{{post_text}}</div>

        <!--浏览量、点赞数功能区-->
        <div class="power">
            <div class="like">
                <img :src="require('../assets/images/' + like)" @click.stop="likepost"/>
                <div class="text">{{favorite_count}}</div>
            </div>

            <div class="comment">
                <img src="../assets/images/comment.png"/>
                <div class="text">{{comment_count}}</div>
            </div>

            <div class="watched">
                <img src="../assets/images/watched.png"/>
                <div class="text">2.2k</div>
            </div>

        </div>

    </div>
</template>

<script>
import { useRouter } from "vue-router";
import { ref } from "vue";
export default{
    props:['id','author','category','comment_count','favorite_count','post_text','fmt_publish_time'],
    setup(props){
        const router = useRouter();
        let like = ref("like.png")  //关于动态切换图片的解决方案https://blog.csdn.net/tangshiyilang/article/details/134637734

        const postdetail = () =>{
            console.log("post id: ", props.id)
            router.push({
                name: 'postdetail_index',
                query:{
                    id: props.id,
                }
            });
            
        }
        const personalpage = () => {
            console.log("username: ", props.author.username)
            router.push({
                name: 'userinfo_index',
                query:{
                    username: props.author.username,
                }
            });
        }
        const likepost = () => {
            if(like.value == "like.png"){
                like.value = "liked.png"
            }else{
                like.value = "like.png"
            }
        }


        return{
            like,
            postdetail,
            personalpage,
            likepost
        }
    }
}
</script>

<style scoped>
.post{
    width: 60vw;
    height: 34vh;
    position: relative;
    box-sizing: border-box;/*加padding或者margin后盒子会被撑大，加上这句盒子就不会被撑大了*/
    padding-left: 10px;
    padding-top: 10px;
    border-top-style: solid;
    border-top-width: 0.5px;
    border-top-color: rgba(0,0,0,0.2);
    cursor: pointer;  /*鼠标悬停变小手*/
}

.avatar_name{
    width: 100%;
    height: 7vh;
   
}
.avatar_name .avatar{
    float: left;
    width: 7vh;
    height: 7vh;
    box-sizing: border-box;
    padding-left: 10px;
    padding-top: 10px;
    margin-right: 10px;
}
.avatar_name .avatar img{
    height: 100%;
    width: 100%;
    border-radius: 50px;
}
.avatar_name .name{
    font-size: 18px;
}

.avatar_name .datetime{
    font-size: 12px;
    color: rgba(0, 0, 0, 0.7);
    margin-top: 1vh;
}

.summary{
    box-sizing: border-box;
    padding-left: 10px;
    padding-top: 10px;
    padding-right: 20px;
    font-size: 18px;
    cursor: pointer;  /*鼠标悬停变小手*/
}
.power{
    width: 100%;
    height: 4vh;
    position:absolute;  /*父盒子position: relative; 子盒子position:absolute;即可让子盒子固定在父盒子底部*/
    bottom: 0px;
    box-sizing: border-box;
}
.power .like{
    float: left;
    width: 8vh;
    height: 8vh;
}
.power .like img{
    float: left;
    width: 3vh;
    height: 3vh;
    
}

.power .comment{
    float: left;
    width: 8vh;
    height: 8vh;
    margin-left: 3vh;
}
.power .comment img{
    float: left;
    width: 3vh;
    height: 3vh;
}
.power .watched{
    float: left;
    width: 8vh;
    height: 8vh;
    margin-left: 3vh;
}
.power .watched img{
    float: left;
    width: 3vh;
    height: 3vh;
}
.power .text{
    float: left;
    margin-left: 0.5vh;
    margin-top: 0.1vh;
    font-size: 2vh;
}
</style>