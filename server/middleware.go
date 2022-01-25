package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type task struct {
	Id   int
	Desc string
	Done bool
}

// postgre db object
var db *sql.DB

// helper function to help check for errors
func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// needed for postgre database connection
const (
	host   = "localhost"
	port   = 5432
	dbname = "todolist"
)

func db_Init() {

	// loads environment variables
	err := godotenv.Load(".env")
	Check(err)

	user := os.Getenv("pg_user")
	password := os.Getenv("pg_password")

	// prepares statement for db connection
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// connects to db and initialises db variable
	cur_db, err := sql.Open("postgres", dbinfo)
	Check(err)
	db = cur_db
}

//not in use
func Register(username string, password string) (bool, error) {
	rows, err := db.Query(`
	  select username
	  from login.user
	  where username = $1`, username)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	var count int = 0
	for rows.Next() {
		count += 1
	}
	if count != 0 {
		//username taken
		return false, nil
	} else {
		db.QueryRow(`
		insert into login.user
		values ($1, $2, $3)`, username, password, "'[]'")
		return true, nil
	}
}

//not in use
func Login(username string, password string) (bool, error) {
	rows, err := db.Query(`
	  select username
	  from login.user
	  where username = $1 AND password = $2`, username, password)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	fmt.Println(rows)

	var count int = 0
	for rows.Next() {
		count += 1
	}
	if count == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

// retrieves the list of tasks
// returns json encoded list 
func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	payload := getTask()
	json.NewEncoder(w).Encode(payload)
}

// helper function
func getTask() []task {
	rows, err := db.Query(`
	  select *
	  from login.tasks
	  order by done, id`)

	Check(err)
	defer rows.Close()

	var tasks []task

	var id int
	var desc string
	var done bool

	for rows.Next() {
		rows.Scan(&desc, &done, &id)
		tasks = append(tasks, task{id, desc, done})
	}

	return tasks
}

// creates a new task
// replies with the id of the new task created
func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	task := r.FormValue("task")
	fmt.Println(r.FormValue("task"))

	id := createTask(task)
	json.NewEncoder(w).Encode(id)
}

// helper function
// returns id of created task
func createTask(desc string) int {
	var id int

	// creates task and stores id of created task
	err := db.QueryRow(`
	insert into login.tasks
	values ($1, $2)
	returning id`, desc, false).Scan(&id)
	Check(err)

	return id
}

// toggles completion of tasks
func ToggleTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	Check(err)

	toggleTask(id)
}

// helper function
func toggleTask(id int) {
	rows, err := db.Query(`
	update login.tasks
	set done = not done
	where id = $1`, id)
	Check(err)
	defer rows.Close()
}

// deletes a task with the given id
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	Check(err)

	deleteTask(id)
}

// helper function
func deleteTask(id int) {
	rows, err := db.Query(`
	delete from login.tasks
	where id = $1`, id)
	Check(err)
	defer rows.Close()
}
