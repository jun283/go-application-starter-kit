package main

import (
	"net/http"
	"text/template"
)

type TodoPageData struct {
	PageTitle string
	Todos     []T
}

type T struct {
	Name     string
	Position string
	Office   string
	Age      int32
}

func SimpleHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.tmpl"))

	data := todoPageDatafunc()

	err := tmpl.Execute(w, data)
	if err != nil {
		errLog.Println(err)
	}
}

func todoPageDatafunc() TodoPageData {

	employees := []T{
		T{"Airi Satou", "Accountant", "Tokyo", 33},
		T{"Angelica Ramos", "Chief Executive Officer (CEO)", "London", 44},
		T{"Ashton Cox", "Junior Technical Author", "San Francisco", 66},
		T{"Brenden Wagner", "Software Engineer", "New York", 61},
		T{"Bruno Nash", "Integration Specialist", "New York", 55},
	}

	data := TodoPageData{
		PageTitle: "List of employees",
		Todos:     employees,
	}

	return data
}
