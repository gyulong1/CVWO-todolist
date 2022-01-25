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

  var db *sql.DB

  // Check is a helper that terminates the program with err.Error() logged in
  // case err is not nil.
  func Check(err error) {
	if err != nil {
	  log.Fatal(err)
	}
  }

  const (
	host     = "localhost"
	port     = 5432
	dbname   = "todolist"
  )

  func Init(){
	err := godotenv.Load(".env")
	Check(err)

	user     := os.Getenv("pg_user")
	password := os.Getenv("pg_password")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

	cur_db, err := sql.Open("postgres", psqlInfo)
	Check(err)
	db = cur_db
}

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

  func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := getTask()
	json.NewEncoder(w).Encode(payload)
}

  func getTask() ([]task) {
	rows, err:= db.Query(`
	  select *
	  from login.tasks`)

	Check(err)
	defer rows.Close()

	var tasks []task

	for rows.Next() {
		var id int
		var desc string
		var done bool
		rows.Scan(&desc, &done, &id);
		tasks = append(tasks, task{id, desc, done})
	  }
	  return tasks
  }

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

  func createTask(desc string) (int) {
	var id int
	err := db.QueryRow(`
	insert into login.tasks
	values ($1, $2)
	returning id`, desc, false).Scan(&id)
	Check(err)
	  return id
	}

	func ToggleTask(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		fmt.Println(id)
		Check(err)
		toggleTask(id)
	}

  func toggleTask(id int) () {
	rows, err := db.Query(`
	update login.tasks
	set done = not done
	where id = $1`, id)
	Check(err)
  	defer rows.Close()
  }


  func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	Check(err)
	deleteTask(id)
	json.NewEncoder(w).Encode(id)
	// json.NewEncoder(w).Encode("Task not found")

}

func deleteTask(id int) () {
	rows, err := db.Query(`
	delete from login.tasks
	where id = $1`, id)
	Check(err)
  	defer rows.Close()
}