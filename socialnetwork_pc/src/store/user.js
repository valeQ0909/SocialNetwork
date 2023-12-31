import axios from "axios";
import router from '../router';

export default {
    // state用来存放数据
    state(){
        return{
            id: "",
            username: "",
            nickname:"",
            photo: "",
            token: "",
            identity: "",
            is_login: false,
            pulling_info:true,   // 是否正在云端拉取信息
        }
    },

    // 当state中的数据需要经过处理后再使用时，可以使用getters配置
    getters:{

    },

    // 更新state中的数据的状态
    mutations:{
        updateUser(state, user) {
            state.id = user.data.id;
            state.username = user.data.username;
            state.nickname = user.data.nickname;
            state.photo = user.data.photo;
            state.identity = user.data.identity;
            state.is_login = user.is_login;
        },
        updateToken(state, token) {
            state.token = token;
        },
        logout(state){
            state.id = "";
            state.username = "";
            state.photo = "";
            state.identity = "";
            state.token = "";
            state.is_login = false;
        },
        updatePullingInfo(state, pulling_info){
            state.pulling_info = pulling_info;
        }
    },

    // actions里面放置的是一个一个的方法，用于响应组件中的“动作”
    actions:{
        // 获取信息
        getinfo(context, data){
            axios({
                headers: {
                  Authorization:"Bearer " + context.state.token,
                },
                method: "GET",
                url: context.rootState.backBaseUrl + "user/info/",
            }).then((resp)=>{
                  console.log("resp: " + resp.data.error)

                  if(resp.data.error_message === "success") {
                      context.commit("updateUser", {
                          ...resp,
                          is_login: true,
                        });
                        data.success(resp);
                  }
                  else {
                      data.error(resp);
                  }
            });
          },
  
          // 登录
          login(context, data){
              axios({
                  header:{
                    'Content-Type':'application/x-www-form-urlencoded'
                  },
                  method: "POST",
                  url: context.rootState.backBaseUrl + "user/token/",
                  data: {
                    'username': data.username,
                    'password': data.password,
                  },
                }).then((resp) => {
                    if (resp.data.error_message === "success") {
                      localStorage.setItem("jwt_token", resp.data.token);
                      context.commit("updateToken", resp.data.token);
                      data.success(resp);
                    } else {
                        data.account = "";
                        data.password = "";
                        data.error(resp);
                    }
                });
          },
  
          //退出登录
          logout(context) {
              localStorage.removeItem("jwt_token");
              context.commit("logout");
              router.push({name: "login_index" });
          },
    }

}