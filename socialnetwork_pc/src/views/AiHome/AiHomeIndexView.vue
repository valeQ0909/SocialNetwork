<template>
    <div class="container">
        <!--左边部分-->
        <div class="page-left">

        </div>

        <!--中间部分-->
        <div class="page-mid">
            <div class="predict-img"><img :src="imgStr"/></div>
            <p class="tip">图片限制10Mb <span class="error">{{errorStr}}</span></p>
            <input type="file" name="image" accept="image/*" @change='uploadImg' class="image-upload-btn">
            <div class="category-box">
                <div class="category">识别任务：</div>
                <input v-model="category" class="input-category" type="text"/>
            </div>
            <div class="predict-box">
                <div class="predict"><p>识别结果：               {{predict}}</p></div>
            </div>
            <div class="update-btn" @click="imgPredict">更新信息</div>
        </div>

        <!--右边部分-->
        <div class="page-right">

        </div>
    </div>
</template>

<script>
import {ref} from "vue"
import axios from "axios";

export default{

    setup(){
        // 上传图片相关变量定义
        let errorStr = ref()
        let buffer = ref()
        let imgStr = ref()
        imgStr.value = "http://127.0.0.1:3000/static/images/avatar.jpg"

        // 识别任务
        let category = ref()
        let predict = ref()


        const uploadImg = (e) => { // 上传本地图片
            let file = e.target.files[0]
            buffer.value = file;
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

        const imgPredict = () =>{
            let token = localStorage.getItem("jwt_token")
            // console.log("file: ", buffer.value)
            axios({
                headers:{
                    Authorization: token,
                    'Content-Type':'multipart/form-data'
                },
                method: 'POST',
                url: "http://127.0.0.1:3000/socialnetwork/aihome/",
                data: {              
                    "image_file": buffer.value,
                    "category": category.value,
                }                
            }).then(response => {
                    predict.value = response.data.predict
                    console.log("预测结果为：", response)
            });
        }

        return{
            imgStr,
            buffer,
            errorStr,
            category,
            predict,
            uploadImg,
            imgPredict,
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
/*页面——左边部分 */
.page-left{
    width: 20vw;
    height: 30vh;
    margin-left: 1vw;
    position: fixed;
    top: 11vh;
    z-index: 1111;
    background-color: white;
}
/*页面——中间部分 */
.page-mid{
    float: left;
    margin-left: 26vw;
    margin-top: 1vh;
    height: 100%;
    width: 46vw;
    background-color: white;
}
/*
.page-left{
    width: 46vw;
    height: 80vh;
    margin-left: 27vw;
    position: fixed;
    top: 11vh;
    z-index: 1111;
    background-color: white;
}*/
.page-mid .predict-img{
    width: 10vw;
    height: 10vw;
    margin-left: 18vw;
    margin-top: 2vh;
    box-sizing: border-box;
}
.page-mid .predict-img img{
    height: 100%;
    width: 100%;
}
.page-mid .tip{
    margin-left: 18vw;
}
.page-mid .image-upload-btn{
    margin-left: 18vw;
    width: 20vw;
    height: 6vh;
    margin-top: 10px;
    cursor: pointer;
}
.category-box{
    height: 6vh;
    margin-top: 5vh;
    margin-left: 10vw;
}
.category-box .category{
    float: left;
    width: 7vw;
    font-size: 20px;
}
.category-box .input-category{
    float: left;
    height: 4vh;
    width: 16vw;
    font-size: 18px;
}
.predict-box{
    height: 6vh;
    margin-top: 5vh;
    margin-left: 10vw;
}
.predict-box .predict{
    float: left;
    width: 15vw;
    font-size: 20px;
}

.update-btn{
    width: 8vw;
    height: 2.5vw;
    margin-left: 19vw;
    margin-top: 4vh;
    margin-bottom: 4vh;
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


/*页面——右边部分 */
.page-right{
    float: left;
    width: 2vw;
    height: 90vh;
}
</style>