package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"tmpl/model"
)

func HandleInfo(w http.ResponseWriter, r *http.Request) {
	htmlByte, err := os.ReadFile("./template/index.tmpl")
	if err != nil {
		fmt.Println("read html failed, err", err)
		return
	}
	info := func(arg string) (string, error) {
		msg := fmt.Sprintf("Hello %s. Your infomation is as follow:", arg)
		return msg, nil
	}
	t, err := template.New("index").Funcs(template.FuncMap{"info": info}).Parse(string(htmlByte))

	if err != nil {
		fmt.Println("parse template failed, err:", err)
		return
	}
	choi := &model.User{
		Uname: "choi",
		Email: "cwy@tdyh.com.cn",
		Dept:  "qa",
		Phone: "13672808880",
	}
	t.Execute(w, choi)
}
