import axios from "axios";
import router from '../router';

export default {
    // state用来存放数据
    state(){
        return{
            id: "",
            username: "",
            avatar: "",
            signature: "",
            token: "",
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
            state.id = user.id;
            state.username = user.username;
            state.avatar = user.avatar;
            state.signature = user.signature;
        },
        updateIsLogin(state, is_login){
            state.is_login = is_login
        },
        updateToken(state, token) {
            state.token = token;
        },
        logout(state){
            state.id = "";
            state.username = "";
            state.photo = "";
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
                  Authorization: context.state.token,
                },
                method: "GET",
                url: context.rootState.backBaseUrl + "/user/info/",
            }).then((resp)=>{
                  console.log("sotre getinfo resp: ", resp)
                  if(resp.data.status_code === 0) {
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
                  headers:{
                    'Content-Type':'application/x-www-form-urlencoded'
                  },
                  method: "POST",
                  url: context.rootState.backBaseUrl + "/user/login/",
                  data: {
                    'username': data.username,
                    'password': data.password,
                  },
                }).then((resp) => {
                    if (resp.data.status_code === 0) {
                      console.log("resp: ", resp)
                      localStorage.setItem("jwt_token", resp.data.token);
                      localStorage.setItem("avatar", resp.data.user.avatar);
                      context.commit("updateUser", resp.data.user);
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
              router.push({name: "login_index"});
          },
    }

}