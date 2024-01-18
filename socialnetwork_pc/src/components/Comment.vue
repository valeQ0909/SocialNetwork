<template>
    <div class="comment-container">
        <!--评论用户头像名称部分-->
        <div class="avatar-name">
            <div class="avatar"  >
                <img src="../assets/images/cat_3.jpg" alt="avatar"  @click.stop="personalpage" />
            </div>
            <div class="box">
                <div class="name"><p>哈哈</p></div>
                <div class="datetime"><p>2024-01-18</p></div>
            </div>
        </div>
        <!--评论用户，评论内容部分-->
        <div class="comment-content">
            你这不行，还得看我
        </div>
        <!--功能图标区-->
        <div class="power">
            <div class="zan">
                <img src="../assets/images/zan.png"/>
                <div class="text"><p>12</p></div>
            </div>
            <div class="pinglun">
                <img src="../assets/images/comment.png">
                <div class="text"><p>评论</p></div>
            </div>
        </div>
    </div>
</template>

<script>
export default{
    setup(){

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
}
.power .zan{
    float: left;
    width: 10vh;
    height: 10vh;
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
    height: 10vh;
}
.power .pinglun img{
    float: left;
    width: 3vh;
    height: 3vh;
    cursor: pointer;
}
.power .text{
    float: left;
    margin-left: 0.5vh;
    color: rgba(0,0,0,0.7);
}
</style>