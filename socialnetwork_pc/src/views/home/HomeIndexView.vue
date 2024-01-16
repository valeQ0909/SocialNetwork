<template>
    <div class="container">
        <!--左边部分：导航栏-->
        <div class="page-left">
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
        <div class="page-mid">

        </div>

        <!--右边部分，暂时不知道放什么东西-->
        <div class="page-right">

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
.container{
    width: 100%;
    height: 100%;
    margin-top: 11vh;
}
.page-left{
    width: 10vw;
    height: 70vh;
    margin-left: 6vw;
    border-width: 0.1px;
    border-style: solid;
    position: fixed;
    top: 11vh;
    z-index: 1111;
}
.page-left ul{
    padding-top: 20px;
}
.page-left ul li {
    padding-left: 10px;
    margin-top: 20px;
    font-size: 26px;
    border-radius: 10px;
    cursor: pointer;  /*鼠标悬停变小手*/
    color:rgba(0, 0, 0,0.6);
}
.page-left ul li:hover{
    background-color: rgb(209, 243, 249);
}


.page-mid {
    float: left;
    margin-left: 19vw;
    height: 110vh;
    width: 60vw;
    background-color: aqua;
}

.page-right{
    float: left;
}
</style>