package learn01_初识

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello 王兰花")
}

func learn01() {
	// 初识
	r := gin.Default()
	r.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "hello 王兰花",
		})
	})

	r.Run()
}

//learn02 模板
func learn02() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("server start failed", err)
		return
	}
}

// learn03 自定义模板函数
func learn03(writer http.ResponseWriter, request *http.Request) {
	//定义模板
	t := template.New("自定义模板函数.tmpl")
	myfunc := func(s string) string {
		return "hello " + s
	}
	t.Funcs(template.FuncMap{"myfunc": myfunc})
	//解析模板
	_, err := t.ParseFiles("./自定义模板函数.tmpl")
	if err != nil {
		fmt.Printf("parse failed %v\n", err)
		return
	}
	//渲染
	t.Execute(writer, nil)
}
func learn04(writer http.ResponseWriter, request *http.Request) {
	//模板嵌套
	t, err := template.ParseFiles("./模板嵌套.tmpl", "./2.tmpl")
	if err != nil {
		fmt.Printf("parse failed %v\n", err)
		return
	}
	t.Execute(writer, nil)
}
func learn05() {
	//模板继承
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/home", home)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("server start failed", err)
		return
	}
}

func home(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("./base.tmpl", "./home.tmpl")
	if err != nil {
		fmt.Printf("parse failed %v\n", err)
		return
	}
	t.Execute(writer, "琉紫")
}

func index2(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("./base.tmpl", "./index2.tmpl") //需要先解析根模板
	if err != nil {
		fmt.Printf("parse failed %v\n", err)
		return
	}
	t.Execute(writer, "莉莉娅娜")
}

type user struct {
	Age       int
	Name      string
	Hobbylist []string
}

func index(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("parse template failed %v\n", err)
		return
	}
	err = t.Execute(writer, user{Name: "王兰花", Age: 18, Hobbylist: []string{
		"琉紫",
		"艾丽卡",
		"莉莉娅娜",
	}})
	if err != nil {
		fmt.Printf("execute template failed %v\n", err)
	}
}
func main() {
	http.HandleFunc("/", learn03)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("server start failed", err)
		return
	}
	//learn05()
}
