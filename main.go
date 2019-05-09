package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

//connection
func connect() *sql.DB {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/lara-vue")
	if err != nil {
		panic(err.Error())
	}

	return db
}

//start of api

//Category struct for database
type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CountTopics int    `json:"numberOfTopics"`
}

//User struct
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

//Topic struct
type Topic struct {
	ID         int    `json:"id"`
	CategoryID int    `json:"category_id"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	Views      int    `json:"views"`
	Times      string `json:"time"`
	UserID     int    `json:"user_id"`
	User       User   `json:"user"`
}

//Topics s
type Topics []Topic

//Categories struct for database
type Categories []Category

func returnAllCategories(w http.ResponseWriter, r *http.Request) {
	var allCate []Category
	var cate Category

	/*categories := Categories{
		Category{ID: 1, Name: "ABC", Description: "ABCDEF"},
		Category{ID: 2, Name: "ABC", Description: "ABCDEF"},
	}*/

	//excute querry

	db := connect()
	defer db.Close()
	results, err := db.Query("SELECT categories.id,name,description,Count(topics.category_id) FROM categories LEFT JOIN topics ON categories.id = topics.category_id GROUP BY category_id;")
	for results.Next() {
		// for each row, scan the result into our cate composite object
		err = results.Scan(&cate.ID, &cate.Name, &cate.Description, &cate.CountTopics)
		if err != nil {
			fmt.Println("dadsadsbad baskdbakdj baksjd bakjd b")
			panic(err.Error()) // proper error handling instead of panic in your app

		} else {
			allCate = append(allCate, cate)
		}
		// and then print out the cate's Name attribute
		log.Println(cate.Name)
	}

	fmt.Println("Endpoint Hit: returnAllCategories")

	json.NewEncoder(w).Encode(allCate)
}

func returnListTopics(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Println(key)
	var allTopics []Topic
	var topic Topic
	var user User

	/*categories := Categories{
		Category{ID: 1, Name: "ABC", Description: "ABCDEF"},
		Category{ID: 2, Name: "ABC", Description: "ABCDEF"},
	}*/

	//excute querry

	db := connect()
	defer db.Close()
	querry := "SELECT id,category_id,title,body,views,created_at,user_id FROM topics WHERE category_id=" + key
	fmt.Println(querry)
	results, err := db.Query(querry)
	for results.Next() {
		// for each row, scan the result into our topic composite object
		err = results.Scan(&topic.ID, &topic.CategoryID, &topic.Title, &topic.Body, &topic.Views, &topic.Times, &topic.UserID)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app

		} else {
			querry = "SELECT ID,name,email FROM users WHERE id=" + strconv.Itoa(topic.UserID)
			row := db.QueryRow(querry)
			err = row.Scan(&user.ID, &user.Name, &user.Email)
			topic.User = user
			allTopics = append(allTopics, topic)
		}
		// and then print out the topic's Name attribute
		log.Println(topic.Title)
	}

	fmt.Println("Endpoint Hit: returnAllTopics")

	json.NewEncoder(w).Encode(allTopics)
}

//end of api

//renderer

//Page title and page content
type Page struct {
	Title string
	Body  []byte
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles(tmpl + ".html")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Println(tmpl + ".html")
	t.Execute(w, "abc")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "view")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	myRouter.HandleFunc("/api/categories", returnAllCategories)
	myRouter.HandleFunc("/api/categories/{id}/topics", returnListTopics)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {

	//open connection
	/*db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/lara-vue")
	if err != nil {
		panic(err.Error())
	}
	results, err := db.Query("SELECT id, name, description FROM categories")
	fmt.Println(results)
	for results.Next() {
		var cate Category
		// for each row, scan the result into our cate composite object
		err = results.Scan(&cate.ID, &cate.Name, &cate.Description)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the cate's Name attribute
		log.Println(cate)
	}*/
	handleRequests()

}
