<template>
    <div class="container">

        <div class="page-mid">
            <div style="border: 1px solid #ccc">
                <v-md-editor v-model="markdown_text"
                             :tab-size=table_size
                             :include-level=nav_level
                             height="88vh"
                >
                </v-md-editor>
            </div>

            <div class="power">
                <div class="select">
                    <input class="category" type="text" v-model="category" readonly="readonly" @click="chooseOption">
                    <div class="option" v-show="option_state">
                        <ul>
                            <li @click="choseCategory(0)">广场</li>
                            <li @click="choseCategory(1)">树洞</li>
                            <li @click="choseCategory(2)">二手市场</li>
                            <li @click="choseCategory(3)">失物招领</li>
                            <li @click="choseCategory(4)">计算机</li>
                            <li @click="choseCategory(5)">招聘</li>
                        </ul>
                    </div>
                </div>

                <div class="send" @click="sendpost"><p>发布</p></div>
                <div class="save"><p>保存</p></div>
            </div>

        </div>

    </div>
</template>

<script>
import axios from "axios";
import '@wangeditor/editor/dist/css/style.css' // 引入 css

import {ref} from 'vue'

export default{
    setup(){
        // markdown组件相关
        let markdown_text = ref("")
        let table_size = ref(2) // tab键
        let nav_level = ref([1,2,3])  //导航栏显示的标题等级
   
        // ----------vale的代码部分--------------
        const sendpost = () =>{
            let token = localStorage.getItem("jwt_token")
            axios({
                headers:{
                   Authorization: token,
                  'Content-Type':'application/x-www-form-urlencoded'
                },
                method: 'POST',
                url: "http://127.0.0.1:3000/socialnetwork/post/sendpost/",
                data: {              
                  'post_markdown': markdown_text.value,
                  'category': category.value
                }                
              }).then(response => {
                    if(response.data.status_code == 0){
                        alert("发帖成功")
                    }
                    else{
                        alert("发帖失败")
                    }
                });

            console.log("text: ", markdown_text.value)
        }


        // 下拉框相关
        let category = ref("广场")
        let option_state = ref(false)
        let category_dict = ["广场", "树洞", "二手市场", "失物招领", "计算机", "招聘"]
        const chooseOption = () =>{
            if(option_state.value == false){
                option_state.value = true
            } else {
                option_state.value = false
            }
        }
        
        const choseCategory = (num) =>{
            category.value = category_dict[num]
            option_state.value = false
        }

        return {
            option_state,
            category,
            markdown_text,
            nav_level,
            table_size,
            sendpost,
            chooseOption,
            choseCategory,
        };
  
    }
}
</script>

<style scoped>
.container{
    width: 100%;
    height: 100%;
    margin-top: 10vh;
    padding-top: 1vh;
    padding-bottom: 25vh;
    background-color: rgb(242,243,245);
    overflow:hidden; /* 父盒子内有的任一级子盒子有浮动会导致父盒子无法高度自适应（即被子盒子撑开）*/
}
.page-mid{
    float: left;
    width: 100%;
    margin-left: 1vh;
    background-color: white;
}

.power{
    width: 100%;
    margin-top: 3vh;
}
.power .select{
    height: 3vw;
    float: left;
    margin-left: 2vw;
}
.power .category{
    display: flex;
    cursor: pointer;
    outline: none;
    width: 8vw;
    height: 3vw;
    font-size: 18px;
    border-color: rgba(0, 0,0, 0.1);
    border-radius: 5px;
    padding-left: 1vh;
} 
.power.option {
    width: 180px;
    border-style: solid;
    border-color: #e2b5b5;
 }
.power .option li {
    padding-left: 1vh;
    height: 2vw;
    font-size: 18px;
    list-style: none;
    cursor: pointer;
    background-color: white;
}
.power .option li:hover{
    background-color: rgb(171,205,255);
} 

.power .send{
    float: right;
    display: flex;
    justify-content: center;
    align-items: center;
    width: 8vw;
    height: 3vw;
    margin-right: 1vw;
    margin-bottom: 3vh;
    border-radius: 5px;
    background-color: rgb(79, 136, 222);
    color: white;
    cursor: pointer;
}
.power .send:hover{
    background-color: rgb(171,205,255);
}
.power .save{
    float: right;
    display: flex;
    justify-content: center;
    align-items: center;
    width: 8vw;
    height: 3vw;
    margin-right: 3vw;
    margin-bottom: 3vh;
    border-radius: 5px;
    background-color: rgb(79, 136, 222);
    color: white;
    cursor: pointer;
}
.power .save:hover{
    background-color: rgb(171,205,255);;
}
</style>