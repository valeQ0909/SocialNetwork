package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

/*
http请求:
· 发起百度搜索的GET请求：“http://www.baidu.com/s?wd=ReganYue”，打印回复的内容
· 对https://httpbin.org/post发起post请求，携带表单数据(application/x-www-form-urlencoded),建值自拟(strings.NewReader("money=0&hobby=撩妹")),打印回复的内容
· 表单数据类型application/x-www-form-urlencoded，表单读取API：strings.NewReader("money=0&hobby=撩妹")
*/
func main() {
	resp, err := http.Get("http://127.0.0.1:3001/deeplearning/mnist")
	HandleError(err, "http.Get")

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")
	fmt.Println(string(bytes))

}

func HandleError(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}
