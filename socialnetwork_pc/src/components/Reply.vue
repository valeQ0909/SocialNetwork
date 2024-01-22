<template>
    <div class="reply-container">
        <!--回复用户头像名称部分-->
        <div class="avatar-name">
            <div class="avatar"  >
                <img :src="commenter.avatar" alt="avatar"  @click.stop="personalpage" />
            </div>
            <div class="box">
                <div class="name"><p>{{commenter.username}}</p></div>
                <div class="datetime"><p>{{fmt_publish_time}}</p></div>
            </div>
        </div>
        <!--回复用户，回复内容部分-->
        <div class="reply-content">{{comment_text}}</div>
    </div>
</template>

<script>
export default{
    props:["commenter","comment_text","fmt_publish_time"],
    setup(){


        return{

        }
    }
}
</script>

<style scoped>
.reply-container{
    margin-top: 2vh;
    box-sizing: border-box;
    padding-left: 10vh;
    width: 100%;
}
.avatar-name{
    margin-top: 4vh;
    width: 100%;
    display: flex;
}
.avatar-name .avatar{
    float: left;
    width: 5vh;
    height: 5vh;
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
    font-size: 15px;
    color: rgba(0,0,0,0.5);
}
.avatar-name .box .datetime{
    font-size: 11px;
    color: rgba(0,0,0,0.7);
    margin-top: 0.6vh;
}
.reply-content{
    margin-top: 1vh;
    margin-left: 8vh;
}
</style>