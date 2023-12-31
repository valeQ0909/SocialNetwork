<template>
    <div class="allpage">
      <div class="motto">
        <div>Think Different</div>
        <div style="height:16px;"></div>
        <div>Here's to the crazy ones.</div>
        <div>The misfits.</div>
        <div>The rebels.</div>
        <div>The troublemakers.</div>
        <div>The round pegs in the square holes.</div>
        <div>The ones who see things differently.</div>
        <div>They're not fond of rules </div>
        <div>And they have no respect for the status quo.</div>
        <div>You can praise them, quote them, disagree with them, disbelieve them, glorify or vilify them.</div>
        <div>About the only thing that you can't do is ignore them.</div>
        <div>Because they change things.</div>
        <div>They invent. </div>
        <div>They imagine. </div>
        <div>They heal.</div>
        <div>They explore.</div>
        <div>They create.</div>
        <div>They inspire.</div>
        <div>They push the human race forward.</div>
        <div>Maybe they have to be crazy.</div>
        <div>How else can you stare at an empty canvas and see a work of art? Or sit in silence and hear a song that's never been written? </div>
        <div>Or gaze at a red planet and see a laboratory on wheels?</div>
        <div>We make tools for these kinds of people.</div>
        <div>While some may see them as the crazy ones, we see genius.</div>
        <div>Because the people who are crazy enough to think that they can change the world, are the ones who do.</div>
        <div style="height:16px;"></div>
        <div style="margin-left:350px"> Apple   --1997</div>
      </div>
  
      <div class="panel">
        <div class="content">
          <div class="switch">
            <h1 id="login">登录</h1>
          </div>
  
          <form>
            <div id="userName" class="input" aria-placeholder="用户名">
              <input v-model="username" type="text" />
            </div>
            <div id="password" class="input" aria-placeholder="密码">
              <input v-model="password" type="password" />
            </div>
  
            <p>
              <router-link id="signUp" replace to="/register" class="input"
                >注册</router-link
              >
              <span>|</span>
              <a id="forget" href="/index/Retrieve" class="input">忘记密码？</a>
            </p>
            <div class="error-message">{{error_message}}</div>
            <button class="button" @click.prevent="login">登录</button>
          </form>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import { useStore } from 'vuex'
  import { ref } from 'vue'
  import router from '../../router';
  
  export default {
      setup(){
        const store = useStore();
        let username = ref('');
        let password = ref('');
        let error_message = ref('');
  
        const jwt_token = localStorage.getItem("jwt_token");
        if(jwt_token){
            store.commit("updateToken", jwt_token);
            store.dispatch("getinfo", {
                success(){
                  store.commit("updatePullingInfo", false);
                },
                error(){
                    store.commit("updatePullingInfo", false);
                }
            })
        }
        else {
            store.commit("updatePullingInfo", false);
        }
  
        const login = () => {
              store.dispatch("login", {
                  username: username.value,
                  password: password.value,
                  success() {
                      store.dispatch("getinfo", {
                          success() {
                              if(store.state.user.identity === "管理员"){
                                router.push({name: "kucun_index" });
                              }
                              else if(store.state.user.identity === "采购员"){
                                router.push({name: "caigou_index" });
                              }
                              else{
                                router.push({name: "caiwub_index" });
                              }
                              
                          },
                          error(){
                          
                          }
                      })                    
                  },
                  error(){
                      error_message.value = "用户名或者密码错误"
                  }
              })
          }
  
        return{
          username,
          password,
          error_message,
          login,
        }
  
      }
  
  
  
  };
  </script>
  
  <style scoped>
  .allpage {
    width: 100%;
    height: 100%;
  }
  .motto{
    float: left;
    width: 500px;
    margin-left: 30px;
    margin-top: 10px;
    font-size: 16px;
  }
  .motto .span{
    display: block;
    margin-left: 20px;
  }
  
  .panel {
    float: left;
    margin-top: 250px;
    margin-left: 250px;
    transform: translate(-50%, -50%);
    background: rgba(0, 0, 0, 0.8);
    display: flex;
    justify-content: center;
    width: 400px;
    padding: 40px;
    box-sizing: border-box;
    box-shadow: 0 15px 25px rgba(0, 0, 0, 0.5);
    border-radius: 10px;
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
  
  .input#signUp {
    color: #7d74ff;
    font-size: 0.8rem;
    text-decoration: none;
  }
  
  .input#signUp:hover {
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
    height: 1.6rem;
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
  
  div.error-message {
      color: red;
  }
  </style>
  