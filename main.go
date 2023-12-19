package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:12345@tcp(127.0.0.1:3306)/apideneme")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	router := mux.NewRouter()

	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", getPostById).Methods("GET")
	router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")
	router.HandleFunc("/posts", createPost).Methods("POST")

	fmt.Println("5555 PORT AÇIK")
	http.ListenAndServe(":5555", router)
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var posts []Post
	result, err := db.Query("SELECT id, title from posts")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var post Post
		err := result.Scan(&post.ID, &post.Title)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, post)
	}
	json.NewEncoder(w).Encode(posts)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare("INSERT INTO posts(title) VALUES(?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	title := keyVal["title"]

	_, err = stmt.Exec(title)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New post was created")
}

func getPostById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT * FROM posts WHERE id = ?", params["id"])

	if err != nil {
		panic(err)
	}

	defer result.Close()

	var post Post

	for result.Next() {
		err := result.Scan(&post.ID, &post.Title)
		if err != nil {
			panic(err)
		}
	}
	json.NewEncoder(w).Encode(post)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	stmt, err := db.Prepare("UPDATE posts SET title = ? WHERE id = ?")

	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	paramsMap := make(map[string]string)
	json.Unmarshal(body, &paramsMap)
	newTitle := paramsMap["title"]

	_, err = stmt.Exec(newTitle, params["id"])

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Update post")
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	stms, err := db.Prepare("DELETE FROM posts WHERE id = ?")

	if err != nil {
		panic(err)
	}

	_, err = stms.Exec(params["id"])

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Deleted post %s", params["id"])
}
