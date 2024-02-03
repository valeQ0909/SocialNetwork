<template>
    <div class="container">
        <!--左边部分：导航栏-->
        <div class="page_left">
            <ul>
                <li :class="category[0]" @click="chose_category(0)">推荐</li>
                <li :class="category[1]" @click="chose_category(1)">热榜</li>
                <li :class="category[2]" @click="chose_category(2)">关注</li>
                <li :class="category[3]" @click="chose_category(3)">广场</li>
                <li :class="category[4]" @click="chose_category(4)">树洞</li>
                <li :class="category[5]" @click="chose_category(5)">Golang</li>
                <li :class="category[6]" @click="chose_category(6)">机器学习</li>
                <li :class="category[7]" @click="chose_category(7)">深度学习</li>
                <li :class="category[8]" @click="chose_category(8)">CV</li>
                <li :class="category[9]" @click="chose_category(9)">NLP</li>
            </ul>
        </div>

        <!--中间部分：帖子的容器-->
        <div class="page_mid">
            <Post v-for="post in postlist" :key="post.id"
                                           :id="post.id"
                                           :author="post.author"
                                           :category="post.category"
                                           :comment_count="post.comment_count"
                                           :favorite_count="post.favorite_count"
                                           :post_text="post.post_text"
                                           :fmt_publish_time="post.fmt_publish_time"
                                           :is_favorite="post.is_favorite"
                                           :watch_count="post.watch_count"
            >
            </Post>
        </div>

        <!--右边部分，暂时不知道放什么东西-->
        <div class="page_right">

        </div>
    </div>

</template>

<script>
import Post from '../../components/Post.vue';
import axios from "axios";
import {onMounted, ref} from "vue"
export default{
    components:{
        Post,
    },

    setup(){
        let postlist = ref([])
        let category_page = ref()

        // 导航栏属性相关
        let category = ref(["chose_category","","","","","","","","",""])

        const getpostlist = () =>{
            let date = new Date();
            let token = localStorage.getItem("jwt_token")
            axios({
                headers:{
                   Authorization: token,
                  'Content-Type':'application/x-www-form-urlencoded'
                },
                method: 'POST',
                url: "http://127.0.0.1:3000/socialnetwork/post/feed/",
                data: {              
                  'category': "广场",
                  'last_time': date.getDate()
                }                
              }).then(response => {
                    postlist.value = response.data.post_list
                    console.log("feed response: ", response)
                    console.log("postlist: ", postlist.value)
            });
        }

        // 导航栏属性相关
        const chose_category = (num) =>{
            let categoryPage = ["推荐", "热榜", "关注","广场", "树洞", "Golang", "机器学习", "深度学习", "CV", "NLP"]
            for(let i = 0; i < 10; i ++){
                category.value[i] = ""
            }
            category.value[num] = "chose_category"
            category_page.value = categoryPage[num]

            let token = localStorage.getItem("jwt_token")
            axios({
                headers:{
                   Authorization: token,
                  'Content-Type':'application/x-www-form-urlencoded'
                },
                method: 'POST',
                url: "http://127.0.0.1:3000/socialnetwork/post/feed/",
                data: {              
                  'category': category_page.value,
                }                
              }).then(resp => {
                    postlist.value = resp.data.post_list
                    console.log("change category response: ", resp.data)
            });
        }

        onMounted(()=>{
            category_page.value = "推荐"
            getpostlist()
        })

        return{
            postlist,
            category,
            category_page,
            getpostlist,
            chose_category,
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
.page_left{
    width: 10vw;
    height: 84vh;
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
    color:rgba(0, 0, 0, 0.6);
}
.page_left ul li:hover{
    background-color: rgb(247,248,250);
}
.page_left ul .chose_category{
    color: rgb(30,128,255);
    background-color: rgb(234,242,255);
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