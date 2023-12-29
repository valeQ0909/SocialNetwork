import { createRouter, createWebHistory } from 'vue-router'

import LoginIndexView from "../views/LoginRegister/LoginIndexView"
import RegisterIndexView from "../views/LoginRegister/RegisterIndexView"
import HomeIndexView from "../views/home/HomeIndexView"

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
     path:"/home",
     name:"home_index",
     component:HomeIndexView,
     meta:{
        requestAuth:true,
     }
    },


]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router