package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/complete/", CompleteTaskFunc)
	http.HandleFunc("/delete/", DeleteTaskFunc)
	http.HandleFunc("/deleted/", ShowTrashTaskFunc)
	http.HandleFunc("/trash/", TrashTaskFunc)
	http.HandleFunc("/edit/", EditTaskFunc)
	http.HandleFunc("/completed/", ShowCompletedTasksFunc)
	http.HandleFunc("/restore/", RestoreTaskFunc)
	http.HandleFunc("/add/", AddTaskFunc)
	http.HandleFunc("/update/", UpdateTaskFunc)
	http.HandleFunc("/search/", SearchTaskFunc)
	http.HandleFunc("/login", GetLogin)
	http.HandleFunc("/register", PostRegister)
	http.HandleFunc("/admin", HandleAdmin)
	http.HandleFunc("/add_user", PostAddUser)
	http.HandleFunc("/change", PostChange)
	http.HandleFunc("/logout", HandleLogout)
	http.HandleFunc("/", ShowAllTaskFunc)

	http.Handle("/static/", http.FileServer(http.Dir("public")))
	log.Print("running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func CompleteTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "complets tasks GET"
	} else {
		message = "complete tasks POST"
	}
	w.Write([]byte(message))
}

func DeleteTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "delete task GET"
	} else {
		message = "delete task POST"
	}
	w.Write([]byte(message))
}

func ShowTrashTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "show deleted GET"
	} else {
		message = "show deleted POST"
	}
	w.Write([]byte(message))
}

func TrashTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "show trash GET"
	} else {
		message = "show trash POST"
	}
	w.Write([]byte(message))
}
func EditTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "edit task GET"
	} else {
		message = "edit task POST"
	}
	w.Write([]byte(message))
}

func ShowCompletedTasksFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "show completed tasks GET"
	} else {
		message = "show completed tasks  POST"
	}
	w.Write([]byte(message))
}

func RestoreTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "restore task GET"
	} else {
		message = "restore task POST"
	}
	w.Write([]byte(message))
}

func AddTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "add task GET"
	} else {
		message = "add task POST"
	}
	w.Write([]byte(message))
}
func UpdateTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "update task GET"
	} else {
		message = "update task POST"
	}
	w.Write([]byte(message))
}

func SearchTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "search task GET"
	} else {
		message = "search task POST"
	}
	w.Write([]byte(message))
}

func GetLogin(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "login GET"
	} else {
		message = "login POST"
	}
	w.Write([]byte(message))
}

func PostRegister(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "register GET"
	} else {
		message = "register POST"
	}
	w.Write([]byte(message))
}

func HandleAdmin(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "handle admin GET"
	} else {
		message = "handle admin POST"
	}
	w.Write([]byte(message))
}

func PostAddUser(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "add user GET"
	} else {
		message = "add user POST"
	}
	w.Write([]byte(message))
}

func PostChange(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "post change GET"
	} else {
		message = "post change POST"
	}
	w.Write([]byte(message))
}

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "handle logout GET"
	} else {
		message = "handle logout POST"
	}
	w.Write([]byte(message))
}

func ShowAllTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "all pending tasks GET"
	} else {
		message = "all pending tasks POST"
	}
	w.Write([]byte(message))
}
