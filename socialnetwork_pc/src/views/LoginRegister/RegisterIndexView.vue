<template>
    <div class="container">
      <div class="panel">
        <div class="content">
  
          <div class="switch">
            <h1 id="signUp">注册</h1>
          </div>
  
          <form>
            <div id="userName" class="input" aria-placeholder="账号">
              <input v-model="username" type="text" />
            </div>
            <div id="password" class="input" aria-placeholder="密码">
              <input v-model="password_1" type="password" />
            </div>
            <div id="repeat" class="input" aria-placeholder="确认密码">
              <input v-model="password_2" type="password" />
            </div>
  
            <router-link id="forget" replace to="/login" class="input">
              <p>已有账户？点击登录</p>
            </router-link>
            <select class="identity" v-model="identity" @change="hhh">
              <option class="item" value="manager">管理员</option>
              <option class="item" value="buyer">采购员</option>
              <option class="item" value="finance">财务员</option>
            </select>
  
            <button @click="register" type="button">注册</button>
          </form>
  
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import axios from "axios";
  import router from '../../router'
  import { ref } from "vue";
  
  
  export default {
  
      setup(){
        let identity = ref("manager")
        let username = ref("")
        let password_1 = ref("")
        let password_2 = ref("")
  
        const hhh = () => {
          console.log(identity.value)
        }
  
        const check = (str) => {
            if(str.length < 3 || str.length >10){
              alert("账号建议使用3-12位的阿拉伯数字")
              return false;
            }
  
            if (username.value === '' || password_1.value === '' || password_2.value === ''){
              alert("账号或密码不能为空")
              return false;
            }
            
            for(let i = 0; i < str.length; i ++)
              if(str[i] > '9' || str[i] < '0'){
                alert("账号建议使用3-12位的阿拉伯数字")
                return false;
              }
  
            return true;
        }
  
        const register = () => {
          let state = false;
          state = check(username.value);
          console.log("身份:", identity.value)
  
          if(state){
            console.log("username", username.value);
            console.log("password_1", password_1.value);
            console.log("password_2", password_2.value);
  
            let identityStr = "";
            identityStr = "管理员";
  
            if(identity.value === "manager"){
              identityStr = "管理员";
            }
            else if(identity.value === "buyer"){
              identityStr = "采购员"
            }
            else if(identity.value === "finance"){
              identityStr = "财务员"
            }
  
            console.log(identityStr);
  
            axios({
                header:{
                  'Content-Type':'application/x-www-form-urlencoded'
                },
                method: 'POST',
                url: "http://127.0.0.1:3000/user/register/",
                data: {
                    'username': username.value,
                    'password': password_1.value,
                    'confirmedPassword': password_2.value,
                    'identity': identityStr,
                }
                }).then(response => {
                    if(response.data.code == 200){
                        console.log(response)
                        router.push({name:"home_index"});
                    }
                    else{
                      username.value = "";
                      password_1.value = "";
                      password_2.value = "";
                      console.log(response)
                    }
                });
            }
  
        }
  
        return{
          username,
          password_1,
          password_2,
          identity,
          check,
          register,
          hhh,
        }
  
      }
  
  
  
  };
  </script>
  
  <style scoped>
  .container {
    margin-left: 550px;
    position: relative;
    width: 24rem;
  }
  
  .panel {
    position: absolute;
    top: 40%;
    left: 55%;
    transform: translate(-50%, -50%);
    background: rgba(0, 0, 0, 0.8);
    display: flex;
    justify-content: center;
    width: 400px;
    padding: 40px;
    box-sizing: border-box;
    box-shadow: 0 15px 25px rgba(0, 0, 0, 0.5);
    border-radius: 10px;
    margin-top: 300px;
  }
  
  .switch h1 {
    text-align: center;
    font-size: 1.4rem;
    color: rgba(125, 116, 255, 0.8);
    border-bottom: rgba(125, 116, 255, 0.8) solid 2px;
    cursor: default;
  }
  
  .input input {
    outline: none;
    width: 100%;
    border: none;
    background: none;
    border-bottom: 0.1rem solid #7d74ff;
    color: rgba(37, 215, 202, 0.84);
    font-size: 1rem;
  }
  
  .input::after {
    content: attr(aria-placeholder);
    position: absolute;
    left: 0;
    top: -20%;
    font-size: 1.1rem;
    color: rgba(125, 116, 255, 0.44);
    transition: 0.3s;
  }
  
  .input.focus::after {
    top: -70%;
    font-size: 1rem;
  }
  
  .input#forget {
    color: #7d74ff;
    font-size: 0.8rem;
    text-decoration: none;
  }
  
  .input#forget:hover {
    color: rgba(138, 143, 255, 0.4);
  }
  
  .input#login {
    color: #7d74ff;
    font-size: 0.8rem;
    text-decoration: none;
  }
  
  .input#login:hover {
    color: rgba(138, 143, 255, 0.4);
  }
  
  form p {
    text-align: center;
  }
  
  form span {
    color: #7d74ff;
    font-size: 0.8rem;
    cursor: default;
  }
  
  form {
    width: 12rem;
    margin: 1rem 0 0;
  }
  
  form .input {
    position: relative;
    opacity: 1;
    width: 100%;
    margin: 2rem 0 0;
    height: 42px;
  }
  
  form .input#userName {
    margin: 3rem 0 0;
  }
  
  form .input#password {
    height: 2rem;
  }
  
  form .identity{
    cursor: pointer;
  
    margin-top: 2rem;
    margin-left: 3rem;
    width: 5rem;
    height: 2rem;
  
  }
  
  
  form button {
    display: block;
    border: none;
    outline: none;
    margin: 2rem 61px 0;
    width: 56px;
    height: 56px;
    border-radius: 50%;
    background: linear-gradient(90deg, #8a8fff, rgb(216, 174, 255));
    box-shadow: 0 0 8px #8a8fff;
    cursor: pointer;
  }
  
  form button:hover {
    border: none;
    outline: none;
    margin: 2rem -7px 0;
    width: 100%;
    height: 3.5rem;
    border-radius: 3rem;
    background: linear-gradient(
      90deg,
      rgba(138, 143, 255, 0.75),
      rgba(216, 174, 255, 0.75)
    );
    box-shadow: 0 0 8px #8a8fff;
    cursor: pointer;
    color: rgba(0, 0, 0, 0.6);
    transition: 0.4s;
  }
  </style>