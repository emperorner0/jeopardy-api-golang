package main

import (
	m "jeopardy-api/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func handleRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandle)
	r.HandleFunc("/upload", uploadCSV).Methods("POST")
	r.HandleFunc("/questions/{num}", allQuestions)
	r.HandleFunc("/question/{id}", questionByID)
	r.HandleFunc("/questionsByValue/{value}", questionsByValue)
	r.HandleFunc("/questionsByCategory/{category}", questionsByCategory)
	r.HandleFunc("/questionsByRound/{round}", questionsByRound)
	r.HandleFunc("/questionsByRoundAndCategory/{round}/{category}", questionsByRoundAndCategory)
	r.HandleFunc("/addQuestion/{shownumber}/{round}/{category}/{value}/{question}/{answer}", addQuestion).Methods("POST")
	r.HandleFunc("/deleteQuestion/{id}", deleteQuestion).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func dbIntialize() {
	DBConn := "root:<password>!@/QuestionsDB?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(DBConn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&m.Question{})
}
func main() {

	dbIntialize()
	handleRoutes()

}
