<template>
    <div class="container">
        <!--左边部分-->
        <div class="page-left">
            <div class="avatar"><img :src="imgStr"/></div>
            <p class="tip">图片限制10Mb <span class="error">{{errorStr}}</span></p>
            <input type="file" name="image" accept="image/*" @change='uploadImg' class="header-upload-btn">
            <div class="update-avatar" @click="updateImg">更新头像</div>
            <div class="username-box">
                <div class="username">用户名：</div>
                <input v-model="username" class="input-username" type="text"/>
            </div>
            <div class="signature-box">
                <div class="signature">个性签名：</div>
                <input v-model="signature" class="input-signature" type="text"/>
            </div>
            <div class="update-btn" @click="updatePersonalInfo">更新信息</div>
        </div>

        <!--右边部分：用来撑开页面-->
        <div class="page-right">

        </div>
    </div>
</template>

<script>
import axios from "axios";
import {ref, onMounted} from "vue"

export default{

    setup(){
        let avatar = ref()
        let username = ref()
        let signature = ref()
        let user_id = ref()
        
        // 上传图片相关变量定义
        let errorStr = ref()
        let buffer = ref()
        let imgStr = ref()

        const getUserInfo = () => {
            let token = localStorage.getItem("jwt_token")
            axios({
                headers:{
                    Authorization: token,
                    'Content-Type':'application/x-www-form-urlencoded'
                },
                method: 'GET',
                url: "http://127.0.0.1:3000/socialnetwork/user/token/",       
            }).then(response => {
                    imgStr.value = response.data.user.avatar
                    signature.value = response.data.user.signature
                    username.value = response.data.user.username
                    user_id.value = response.data.user.id
            });
        }

        const uploadImg = (e) => { // 上传本地图片
            let file = e.target.files[0]
            buffer.value = file;
            console.log(file)
            // 获取图片的大小，做大小限制有用
            let imgSize = file.size
            // console.log(imgSize)

            // 比如上传头像限制5M大小，这里获取的大小单位是b
            if (imgSize <= 5 * 1024* 1024) {
                // 图片大小符合要求
                errorStr.value = ""
                // base64方法
                let reader = new FileReader()
                reader.readAsDataURL(file) // 读出 base64
                reader.onloadend = function () {
                    //图片的 base64 格式, 可以直接当成 img 的 src 属性值
                    var dataURL = reader.result
                    //console.log(dataURL)
                    imgStr.value = dataURL
                }
            } else {
                // 图片大小超出限制
                errorStr.value = '图片大小超出范围'
            }
        }

        const updateImg = () => {
            let token = localStorage.getItem("jwt_token")
            // console.log("file: ", buffer.value)
            axios({
                headers:{
                    Authorization: token,
                    'Content-Type':'multipart/form-data'
                },
                method: 'POST',
                url: "http://127.0.0.1:3000/socialnetwork/user/updateavatar/",
                data: {              
                    'image_file': buffer.value,
                }                
            }).then(response => {
                    localStorage.setItem("avatar", response.data.avatar)
                    window.location.reload();
            });
        }

        const updatePersonalInfo = () => {
            let token = localStorage.getItem("jwt_token")
            axios({
                headers:{
                    Authorization: token,
                    'Content-Type':'application/x-www-form-urlencoded'
                },
                method: 'POST',
                url: "http://127.0.0.1:3000/socialnetwork/user/updatepersonalinfo/",
                data: {              
                    'username': username.value,
                    'signature': signature.value
                }                
            }).then(response => {
                    console.log("response: ", response)
                    username.value = response.data.username
                    signature.value = response.data.signature
                });
        }

        onMounted(()=>{
            getUserInfo()
        })
        return{
            avatar,
            username,
            signature,
            errorStr,
            imgStr,
            buffer,
            getUserInfo,
            uploadImg,
            updateImg,
            updatePersonalInfo,
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

.page-left{
    width: 46vw;
    height: 80vh;
    margin-left: 27vw;
    position: fixed;
    top: 11vh;
    z-index: 1111;
    background-color: white;
}
.page-left .avatar{
    width: 10vw;
    height: 10vw;
    margin-left: 18vw;
    margin-top: 2vh;
    box-sizing: border-box;
}
.page-left .avatar img{
    height: 100%;
    width: 100%;
}
.page-left .tip{
    margin-left: 18vw;
}
.page-left .header-upload-btn{
    margin-left: 18vw;
    width: 20vw;
    height: 6vh;
    margin-top: 10px;
    cursor: pointer;
}
.update-avatar{
    width: 8vw;
    height: 2.5vw;
    margin-left: 19vw;
    font-size: 16px;
    border-radius: 5px;
    display: flex;
    justify-content: center;
    align-items: center;
    cursor: pointer;
    color: white;
    background-color: rgb(30,128,255);
}
.update-avatar:hover{
    background-color: rgb(122, 175, 245);
}

.username-box{
    height: 6vh;
    margin-top: 5vh;
    margin-left: 10vw;
}
.username-box .username{
    float: left;
    width: 7vw;
    font-size: 20px;
}
.username-box .input-username{
    float: left;
    height: 4vh;
    width: 16vw;
    font-size: 18px;
}

.signature-box{
    height:6vh;
    margin-left: 10vw;
}
.signature-box .signature{
    float: left;
    width: 7vw;
    font-size: 20px;
}
.signature-box .input-signature{
    float: left;
    height: 4vh;
    width: 16vw;
    font-size: 16px;
}
.update-btn{
    width: 8vw;
    height: 2.5vw;
    margin-left: 19vw;
    margin-top: 4vh;
    font-size: 16px;
    border-radius: 5px;
    display: flex;
    justify-content: center;
    align-items: center;
    cursor: pointer;
    color: white;
    background-color: rgb(30,128,255);
}
.update-btn:hover{
    background-color: rgb(122, 175, 245);
}



.page-right{
    float: left;
    width: 2vw;
    height: 90vh;
}
</style>