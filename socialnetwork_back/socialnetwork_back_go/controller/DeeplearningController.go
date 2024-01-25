package controller

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"
)

type Predict struct {
	Code    int    `json:"code"`
	ErrMsg  string `json:"err_msg"`
	Predict string `json:"predict"`
}

// MnistResponse 手写数字识别
type MnistResponse struct {
	Response
	Predict string `json:"predict,omitempty"`
}

// ImagePredict 图片识别
func ImagePredict(c *gin.Context) {
	// 参考博客https://www.cnblogs.com/l199616j/p/15248673.html
	// *multipart.FileHeader转reader类型文件
	header, _ := c.FormFile("image_file")
	imgFile, _ := header.Open()
	category := c.PostForm("category")

	fmt.Println("imgFile: ", imgFile)
	fmt.Println("category: ", category)

	// 请求地址
	modelurl := "http://127.0.0.1:3001/deeplearning/" + category
	// 参数名称
	paramName := "image_file"
	// 创建请求体
	body := &bytes.Buffer{}
	//创建一个multipart类型的写文件
	writer := multipart.NewWriter(body)
	//使用给出的属性名paramName和文件名filePath创建一个新的form-data头
	part, err := writer.CreateFormFile(paramName, "mnist")
	if err != nil {
		fmt.Println(" post err=", err)
	}
	//将源复制到目标，将file写入到part   是按默认的缓冲区32k循环操作的，不会将内容一次性全写入内存中,这样就能解决大文件的问题
	_, err = io.Copy(part, imgFile)
	err = writer.Close()
	if err != nil {
		fmt.Println(" post err=", err)
	}
	request, err := http.NewRequest("POST", modelurl, body)
	request.Header.Add("S-COOKIE2", "a=2l=310260000000000&m=460&n=00")
	//writer.FormDataContentType() ： 返回w对应的HTTP multipart请求的Content-Type的值，多以multipart/form-data起始
	request.Header.Set("Content-Type", writer.FormDataContentType())
	//设置host，只能用request.Host = “”，不能用request.Header.Add(),也不能用request.Header.Set()来添加host
	//request.Host = "api.shouji.com"
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100
	clt := http.Client{
		Timeout:   10 * time.Second,
		Transport: t,
	}
	defer clt.CloseIdleConnections()

	res, err := clt.Do(request)
	defer func() {
		res.Body.Close() //http返回的response的body必须close,否则就会有内存泄露
	}()
	if err != nil {
		fmt.Println("err is ", err)
	}
	body1, err1 := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll  err is ", err1)
		return
	}
	// 从bytes中解析json
	var predict Predict
	json.Unmarshal(body1, &predict)
	fmt.Println("predict:    ", predict)

	c.JSON(http.StatusOK, MnistResponse{
		Response: Response{StatusCode: 0, StatusMsg: "识别成功"},
		Predict:  predict.Predict,
	})

}
