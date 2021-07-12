package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	m "jeopardy-api/model"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func homeHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func uploadCSV(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(32 << 20)

	csvFile, handler, err := r.FormFile("file")

	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	uploadedFile, err := os.Open(handler.Filename)
	if err != nil {
		panic(err)
	}
	defer uploadedFile.Close()

	recs := csv.NewReader(uploadedFile)

	_, err1 := recs.Read()
	if err1 != nil {
		panic(err)
	}

	records, err := recs.ReadAll()
	if err != nil {
		panic(err)
	}

	var qst m.Question
	var qsts []m.Question

	for _, rec := range records {
		qst.ShowNumber, _ = strconv.Atoi(rec[0])
		qst.Round = strings.Replace(strings.Split(rec[1], " ")[0], "!", "", -1)
		qst.Category = rec[2]
		qst.Value, _ = strconv.Atoi(strings.Replace(rec[3], "$", "0", -1))
		qst.Question = rec[4]
		qst.Answer = rec[5]
		qsts = append(qsts, qst)
	}

	DBConn := "root:<password>@/QuestionsDB?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(DBConn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.CreateInBatches(&qsts, 1000)
}

func allQuestions(w http.ResponseWriter, r *http.Request) {
	DBConn := "root:<password>@/QuestionsDB?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(DBConn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	vars := mux.Vars(r)
	numberEntries, _ := strconv.Atoi(vars["num"])

	var questions []m.Question

	db.Limit(numberEntries).Find(&questions)

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	enc.Encode(questions)
}

func questionByID(w http.ResponseWriter, r *http.Request) {
	DBConn := "root:<password>@/QuestionsDB?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(DBConn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var question []m.Question
	db.Where("id = ?", id).First(&question)

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	enc.Encode(question)
}

func questionsByValue(w http.ResponseWriter, r *http.Request) {
	DBConn := "root:<password>@/QuestionsDB?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(DBConn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	vars := mux.Vars(r)
	value, _ := strconv.Atoi(vars["value"])

	var questions []m.Question
	db.Where("value = ?", value).Find(&questions)

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	enc.Encode(questions)
}

func questionsByCategory(w http.ResponseWriter, r *http.Request) {
	DBConn := "root:<password>@/QuestionsDB?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(DBConn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	vars := mux.Vars(r)
	category, _ := vars["category"]

	var questions []m.Question
	db.Where("category like ?", "%"+category+"%").Find(&questions)

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	enc.Encode(questions)
}

func questionsByRound(w http.ResponseWriter, r *http.Request) {
	DBConn := "root:<password>@/QuestionsDB?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(DBConn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	vars := mux.Vars(r)
	round, _ := vars["round"]

	var questions []m.Question
	db.Where("round = ?", round).Find(&questions)

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	enc.Encode(questions)
}

func questionsByRoundAndCategory(w http.ResponseWriter, r *http.Request) {
	DBConn := "root:<password>@/QuestionsDB?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(DBConn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	vars := mux.Vars(r)
	round, _ := vars["round"]
	category, _ := vars["category"]

	var questions []m.Question
	db.Where("round = ? AND category like ?", round, "%"+category+"%").Find(&questions)

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	enc.Encode(questions)
}

func addQuestion(w http.ResponseWriter, r *http.Request) {
	DBConn := "root:<password>@/QuestionsDB?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(DBConn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	vars := mux.Vars(r)
	shownumber, _ := strconv.Atoi(vars["shownumber"])
	round, _ := vars["round"]
	category, _ := vars["category"]
	value, _ := strconv.Atoi(vars["value"])
	ques, _ := vars["question"]
	answer, _ := vars["answer"]

	q := &m.Question{
		Model:      gorm.Model{},
		ShowNumber: shownumber,
		Round:      round,
		Category:   category,
		Value:      value,
		Question:   ques,
		Answer:     answer,
	}

	db.Create(&q)

}

func deleteQuestion(w http.ResponseWriter, r *http.Request) {
	DBConn := "root:<password>@/QuestionsDB?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(DBConn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var question []m.Question
	db.Where("id = ?", id).Delete(&question)

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	enc.Encode(question)
}
