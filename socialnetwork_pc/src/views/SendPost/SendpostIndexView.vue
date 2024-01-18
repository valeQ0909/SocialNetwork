<template>
    <div class="container">
        <!--左边部分-->
        <div class="page-left">

        </div>

        <!--中间部分-->
        <div class="page-mid">
            <div style="border: 1px solid #ccc">
                <Toolbar
                    style="border-bottom: 1px solid #ccc"
                    :editor="editorRef"
                    :defaultConfig="toolbarConfig"
                    :mode="mode"
                />
                <Editor
                    style="height: 500px; overflow-y: hidden;"
                    v-model="valueHtml"
                    :defaultConfig="editorConfig"
                    :mode="mode"
                    @onCreated="handleCreated"
                />
            </div>
            <div class="power">
                <div class="send" @click="sendpost"><p>发布</p></div>
                <div class="save"><p>保存</p></div>
            </div>

        </div>

    </div>
</template>

<script>
import axios from "axios";
import '@wangeditor/editor/dist/css/style.css' // 引入 css

import { onBeforeUnmount, ref, shallowRef, onMounted, onBeforeUpdate } from 'vue'
import { Editor, Toolbar } from '@wangeditor/editor-for-vue'
export default{
    components: { 
        Editor, 
        Toolbar 
    },

    setup(){
        // 编辑器实例，必须用 shallowRef
        const editorRef = shallowRef()
        // 内容 HTML
        const valueHtml = ref('<p></p>')
        // 模拟 ajax 异步获取内容
        onMounted(() => {
            setTimeout(() => {
                valueHtml.value = '<p></p>'
            }, 1500)
        })
        const toolbarConfig = {}
        const editorConfig = { placeholder: '请输入内容...' }
        // 组件销毁时，也及时销毁编辑器
        onBeforeUnmount(() => {
            const editor = editorRef.value
            if (editor == null) return
            editor.destroy()
        })
        const handleCreated = (editor) => {
           editorRef.value = editor // 记录 editor 实例，重要！
        }

        // ----------vale的代码部分--------------
        let post_text = ""
        const sendpost = () =>{
            console.log("postcontent: ", valueHtml.value)
            post_text = editorRef.value.getText()
            console.log("posttext: ", post_text)
            let token = localStorage.getItem("jwt_token")
            axios({
                headers:{
                   Authorization: token,
                  'Content-Type':'application/x-www-form-urlencoded'
                },
                method: 'POST',
                url: "http://127.0.0.1:3000/socialnetwork/post/sendpost",
                data: {              
                  'post_html': valueHtml.value,
                  'post_text': post_text,
                  'category': "广场"
                }                
              }).then(response => {
                    if(response.data.status_code == 0){
                        alert("发帖成功")
                    }
                    else{
                        alert("发帖失败")
                    }
                });
        }

        onBeforeUpdate(() => { // 获取 editor ，必须等待它渲染完之后
            const editor = editorRef.value // 获取 editor ，必须等待它渲染完之后
            editor.handleTab = () => editor.insertText('    ')

        })
        
        return {
        post_text,
        editorRef,
        valueHtml,
        mode: 'default', // 或 'simple'
        toolbarConfig,
        editorConfig,
        handleCreated,
        sendpost,
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
    padding-bottom: 10vh;
    background-color: rgb(242,243,245);
    overflow:hidden; /*父盒子内有的任一级子盒子有浮动会导致父盒子无法高度自适应（即被子盒子撑开）*/
}
.page-left{  /*之后会在这里放头像啥的 */
    width: 10vw;
    height: 70vh;
    margin-left: 8vw;
    position: fixed;
    top: 11vh;
    z-index: 1111;
}
.page-mid{
    float: left;
    margin-left: 19vw;
    width: 56vw;
    background-color: white;
}

.power{
    width: 100%;
    margin-top: 3vh;
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