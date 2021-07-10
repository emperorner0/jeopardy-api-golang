package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	// "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	ShowNumber int    `json:"show_number"`
	Round      string `json:"round"`
	Category   string `json:"category"`
	Value      int    `json:"value"`
	Question   string `json:"question"`
	Answer     string `json:"answer"`
}

func csvToStruct(s string) *[]Question {

	csvfile, err := os.Open(s)
	if err != nil {
		log.Fatalln("Couldn't Open CSV", err)
	}
	defer csvfile.Close()

	r := csv.NewReader(csvfile)

	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var qst Question
	var qsts []Question

	for _, rec := range records {
		qst.ShowNumber, _ = strconv.Atoi(rec[0])
		qst.Round = rec[1]
		qst.Category = rec[2]
		qst.Value, _ = strconv.Atoi(strings.Replace(rec[3], "$", "0", -1))
		qst.Question = rec[4]
		qst.Answer = rec[5]
		qsts = append(qsts, qst)
	}
	return &qsts
}

func jsonEncode(q []Question) []byte {
	buf := new(bytes.Buffer)

	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(q)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return buf.Bytes()
}

func main() {

	// DBConn := "root:Haloking12!@tcp(127.0.0.1:3306)/QuestionsDB"
	// db := mysql.Open(DBConn)

	filePath := "test.csv"

	qsts := csvToStruct(filePath)

	fmt.Println(string(jsonEncode(*qsts)))

}
