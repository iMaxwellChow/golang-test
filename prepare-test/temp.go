package main

import (
	"io"
	"log"
	"net/http"
)

//type HomeController struct {
//	beego.Controller
//}
//
//func (this *HomeController) Get() {
//	this.Ctx.WriteString("Hello world!")
//}

func main() {
	////注册路由
	//beego.Router("/", &HomeController{})
	//beego.Run()

	//设置路由
	http.HandleFunc("/", sayHello) //sayHello 这个函数做为参数传递进去

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world,this is version 1.")
}
