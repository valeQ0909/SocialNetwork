<template>
    <div class="container">
        <div class="logo">
            <img src="../assets/images/logo.png"  @click="gohome" alt="logo" />
        </div>
        
        <div class="motto">
            <span class="motto-item"
                v-for="(item, index) in motto"
                :key="index"
                v-show="n === index"
            >
                {{ item }}
            </span>
        </div>

  
        <div class="avatar"  @mouseover="avatarshowshortcut" @mouseleave="avatarhideshortcut">
            <img :src="avatar" alt="avatar"/>
        </div>

    </div>

    <!--头像功能区-->
    <div class="userpower" @mouseover="powershowshortcut" @mouseleave="powerhideshortcut" v-if="showpower">
       <div class="mypersonalpage selection">
         <router-link replace to="/sendpost"><p>写帖子</p></router-link>
       </div>
       <div class="sendpost selection">
         <router-link replace to="/userinfo"><p>我的主页</p></router-link>
       </div>
       <div class="logout selection" >
         <p @click="logout">退出登录</p>
       </div>
    </div>

  </template>
  
<script>
console.log("乘风好去，长空万里，直下看山河。\n斫去桂婆娑，人道是、清光更多。\n                   -- 辛弃疾");
import {ref, reactive, onMounted } from "vue"
import { useStore } from 'vuex';
import { useRouter } from "vue-router";

export default {
    setup(){
        const store = useStore();
        const router = useRouter();
        let currentpage = ref();
        let avatar = ref(); // 头像地址


        const theme1 = reactive({
            color: 'rgb(141, 139, 139)'
        });
        const theme2 = reactive({
            color: 'rgb(141, 139, 139)'
        });

        let onpower = ref(false);
        onpower.value = false;
        let showpower = ref(false);
        showpower.value = false;
        let onavatar = ref(false);
        onavatar.value = false;
        const avatarshowshortcut = () => {
            onavatar.value = true;
            showpower.value = true;
        };
        const avatarhideshortcut = () => {
            onavatar.value = false;
        };
        const powershowshortcut = () =>{
            onpower.value = true;
        };
        
        const powerhideshortcut = () =>{
            onpower.value = false;
            showpower.value = false;
 
        };
        const gohome = () =>{
            router.push({
                name: 'home_index',
            });
        }
        const logout = () => {
          store.dispatch("logout");
        }
        const getAvatar = () =>{
            avatar.value = localStorage.getItem("avatar");
            if(avatar.value == null){ //如果本地没有存放头像地址
                avatar.value = "http://127.0.0.1:3000/static/images/defaultavatar.jpg"
            }
        }


        setInterval(() => {
            currentpage.value = router.currentRoute.value.name;
            if (currentpage.value === "kucun_index" || currentpage.value === "caigou_index" || currentpage.value === "caiwub_index"){
                theme1.color = "white"
                theme2.color = "rgb(141, 139, 139)"
            }
            else if(currentpage.value === "dinghuo_index" || currentpage.value === "caiwua_index"){
                theme2.color = "white"
                theme1.color = "rgb(141, 139, 139)"
            }
        }, 10);

        onMounted(()=>{
            getAvatar()
        })

        return {
            theme1,
            theme2,
            onpower,
            onavatar,
            showpower,
            currentpage,
            avatar,
            avatarshowshortcut,
            avatarhideshortcut,
            powershowshortcut,
            powerhideshortcut,
            gohome,
            logout,
            getAvatar
       };
    },

    data() {
      return {
        motto: [
            "我见青山多妩媚，料青山见我应如是。",
            "不恨古人吾不见，恨古人不见吾狂耳。",
            "我最怜君中宵舞，道“男儿到死心如铁”。看试手，补天裂。",
            "吾庐小，在龙蛇影外，风雨声中。",
            "乘风好去，长空万里，直下看山河。",
            "斫去桂婆娑，人道是、清光更多。",
            "鸿鹄一再高举，天地睹方圆。",
        ],
        n: 0,
      };
    },

    methods: {
        go() {
        //从开始改变变量n
        setInterval(() => {
            this.n++;
            if (this.n === this.motto.length) this.n = 0;
        }, 6000);
        },
    },

    mounted(){
        this.go();
    }

  };
</script>
  
<style scoped>
.router-link-active{
    text-decoration: none;
}
a {
    text-decoration: none;
}
.container{
    height: 10vh;
    width: 100%;
    background-color: white;
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    z-index: 1111;
}
.container .logo{
    float: left;
    height: 5vh;
    width: 5vh;
    margin-left: 5vw;
    margin-top: 2.5vh;
    cursor: pointer;

}
.container .logo img{
    width: 100%;
    height: 100%;
}


.container .motto{
    color: rgb(0, 0, 0);
    float: left;
    width: 500px;
    margin-left: 10vw;
    margin-top: 3vh;
}

.container .item1{
    cursor: pointer;
    float: left;
    margin-left: 5vw;
    margin-top: 2vh;
    color: v-bind('theme1.color');
    font-size: 4vh;
}
.container .item2{
    cursor: pointer;  
    float: left;
    margin-left: 8vw;
    margin-top: 2vh;
    color: v-bind('theme2.color');;
    font-size: 4vh;
}
.container .item:hover{
    color: white;
}

.container .avatar{
    margin-top: 1vh;
    margin-right: 5vw;
    float: right;
    height: 8vh;
    width: 8vh;
}

.container .avatar img{
    width: 100%;
    height: 100%;
    border-radius: 50px;
}


.userpower{
  background-color: #eaeaea;
  height: 150px;
  width: 100px;
  z-index: 1001;
  position: fixed;
  right: 55px;
}

.userpower p{
  text-align: center;
  font-size: 14px;
}

.userpower .mypersonalpage{
 margin-top: 20px;
}

.userpower .changeinformation{
  margin-top: 10px;
}

.userpower .logout{
  cursor: pointer;  /*鼠标悬停变小手*/
  margin-top: 10px;
}

.userpower .selection:hover{
  border-radius: 30px;
  background-color: rgb(172, 182, 182);
}
</style>