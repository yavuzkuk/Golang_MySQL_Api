// package main

// import (
// 	"apideneme/controller"
// 	"fmt"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// func main() {
// 	router := mux.NewRouter()
// 	// database.ConnectDB()

// 	router.HandleFunc("/post", controller.GetPosts).Methods("GET")

// 	// router.HandleFunc("/post/:id", controller.GetPost).Methods("GET")

// 	fmt.Println("Server açıldı:  5555 PORTU")
// 	http.ListenAndServe(":5555", router)
// }

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
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
