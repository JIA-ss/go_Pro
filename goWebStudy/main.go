package main

import (
	"fmt"
	"html/template"
	"net/http"
)

/*
	1. w -> 响应流写入器
	2. r -> 请求对象的指针
*/
func GetParam(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"func myWeb start handle request！")

	//将请求主体解析为表单，获得POST Form表单数据
	r.ParseForm()

	for k,v := range r.URL.Query() {
		fmt.Println("key: ",k,"\tvalue: ",v[0])
	}

	for k,v := range r.PostForm {
		fmt.Println("key: ",k,"\tvalue: ",v[0])
	}

	fmt.Fprintf(w,"处理完毕")
}

func UseTemp(w http.ResponseWriter, r *http.Request)  {
	t, _ := template.ParseFiles("./templates/index.html")

	data := map[string]string{
		"name": "joshua",
		"someStr": "咳咳",
	}
	//template包的核心功能就是将HTML字符串解析暂存起来，然后调用的时候，替换掉html字符串中的{{}}里面的内容
	t.Execute(w,data)
}

func main() {
	//匹配路由和处理函数   -> http://localhost:8080/ -> myWeb
	http.HandleFunc("/getparam",GetParam)
	http.HandleFunc("/usetemp", UseTemp)
	fmt.Println("服务器即将开启，访问地址 http://localhost:8080")

	//指定相对路径./static 为文件服务路径
	staticHandle := http.FileServer(http.Dir("./static"))
	//将/js/路径下的请求匹配到 ./static/js/下
	http.Handle("/js/",staticHandle)
	//或者直接写成↓
	//http.Handle("/js/",http.FileServer(http.Dir("./static")))

	//http包还提供 http.StripPrefix 剥开前缀, 如下就是对于/css的访问，直接定位到./static文件夹下
	http.Handle("/css/",http.StripPrefix("/css/",http.FileServer(http.Dir("./static"))))




	//第一个参数是指定监听的端口号，第二个是指定处理请求的handler(nil默认为ServeMux)
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		fmt.Println("服务器开启错误",err)
	}

}
