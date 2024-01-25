import { createRouter, createWebHistory } from 'vue-router'

import LoginIndexView from "../views/LoginRegister/LoginIndexView"
import RegisterIndexView from "../views/LoginRegister/RegisterIndexView"
import HomeIndexView from "../views/Home/HomeIndexView"
import PostdetailIndexView from "../views/PostDetail/PostdetailIndexView"
import PersonalPageIndexView from "../views/PersonalPage/PersonalPageIndexView"
import UserInfoIndexView from "../views/UserInfo/UserInfoIndexView"
import SendpostIndexView from "../views/SendPost/SendpostIndexView"
import ChangePersonalInfoIndexView from  "../views/ChangePersonalInfo/ChangePersonalInfoIndexView"
import AiHomeIndexView from "../views/AiHome/AiHomeIndexView"

const routes = [
    {
     path: "/",
     name: "home",
     redirect: "/home",
     meta: {
        requestAuth: true,
     }
    },
    {
     path:"/login",
     name:"login_index",
     component: LoginIndexView,
     meta:{
        requestAuth: true,
     }
    },
    {
     path:"/register",
     name:"register_index",
     component: RegisterIndexView,
     meta:{
        requestAuth: true,
     }
    },
    {
      path:"/postdetail",  // 帖子详情页面
      name:"postdetail_index",
      component: PostdetailIndexView,
      meta:{
         requestAuth: false,
      }
    },
    {
      path:"/sendpost",
      name:"sendpost_index",
      component: SendpostIndexView,
      meta:{
         requestAuth: false,
      }
    },
    {
      path:"/personalpage", // 个人信息详情页
      name:"personalpage_index",
      component: PersonalPageIndexView,
      meta:{
         requestAuth: true,
      }
    },
    {
      path:"/changepersonalinfo",
      name:"changepersonalinfo_index",
      component:ChangePersonalInfoIndexView,
      meta:{
         requestAuth: true,
      }
    },
    {
      path: "/userinfo",
      name: "userinfo_index",
      component: UserInfoIndexView,
      meta:{
         requestAuth: false,
      }
    },
    {
      path: "/aihome",
      name: "aihome_index",
      component: AiHomeIndexView,
      meta:{
         requestAuth: false,
      }
    },
    {
     path:"/home",
     name:"home_index",
     component:HomeIndexView,
     meta:{
        requestAuth:false,
     }
    },


]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router