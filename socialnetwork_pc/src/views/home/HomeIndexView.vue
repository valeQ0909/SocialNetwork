<template>
    <div class="container">
        <!--左边部分：导航栏-->
        <div class="page_left">
            <ul>
                <li>推荐</li>
                <li>热榜</li>
                <li>广场</li>
                <li>树洞</li>
                <li>二手市场</li>
                <li>失物招领</li>
                <li>计算机</li>
                <li>招聘</li>
            </ul>
        </div>

        <!--中间部分：帖子的容器-->
        <div class="page_mid">
            <Post></Post>
            <Post></Post>
            <Post></Post>
            <Post></Post>
        </div>

        <!--右边部分，暂时不知道放什么东西-->
        <div class="page_right">

        </div>
    </div>

</template>

<script>
import Post from '../../components/Post.vue';

export default{
    components:{
        Post,
    },

    setup(){

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
.page_left{
    width: 10vw;
    height: 70vh;
    margin-left: 6vw;
    position: fixed;
    top: 11vh;
    z-index: 1111;
    background-color: white;
    
}
.page_left ul{
    padding-top: 20px;
}
.page_left ul li {
    padding-left: 10px;
    margin-top: 20px;
    font-size: 26px;
    border-radius: 10px;
    cursor: pointer;  /*鼠标悬停变小手*/
    color:rgba(0, 0, 0,0.6);
}
.page_left ul li:hover{
    background-color: rgb(209, 243, 249);
}

.page_mid{
    float: left;
    margin-left: 19vw;
    height: 100%;
    width: 60vw;
    background-color: white;
}
.page_right{
    float: left;
}
</style>