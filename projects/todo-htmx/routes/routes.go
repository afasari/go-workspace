package routes

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"html/template"
	"todo-htmx/model"

	"github.com/gorilla/mux"
)

func sendTodos(w http.ResponseWriter) {
	todos, err := model.GetAllTodos()
	if err != nil {
		fmt.Println("could not get all todos from db", err)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	err = tmpl.ExecuteTemplate(w, "Todos", todos)
	if err != nil {
		fmt.Println("could not execute template", err)
		return
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	todos, err := model.GetAllTodos()
	if err != nil {
		fmt.Println("could not get all todos from db", err)
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	err = tmpl.Execute(w, todos)
	if err != nil {
		fmt.Println("could not execute template", err)
	}
}

func markTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		fmt.Println("could not parse id", err)
	}

	err = model.MarkDone(id)
	if err != nil {
		fmt.Println("could not update todo", err)
	}

	sendTodos(w)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("could not parsing form", err)
	}

	err = model.CreateTodo(r.FormValue("todo"))
	if err != nil {
		fmt.Println("could not create todo", err)
	}

	sendTodos(w)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		fmt.Println("could not parse id", err)
	}

	err = model.Delete(id)
	if err != nil {
		fmt.Println("could not delete", err)
	}

	sendTodos(w)
}

func SetupAndRun() {
	mux := mux.NewRouter()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/todo/{id}", markTodo).Methods("PUT")
	mux.HandleFunc("/todo/{id}", deleteTodo).Methods("DELETE")
	mux.HandleFunc("/create", createTodo).Methods("POST")

	log.Fatal(http.ListenAndServe(":5000", mux))
}
