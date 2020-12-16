package handlers

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func (m *Mappa) Test(writer http.ResponseWriter, request *http.Request) {
	file, err := os.Open("../task3/files/some_text.txt")

	if err != nil {
		panic(err.Error())
	}

	defer file.Close()

	content := bufio.NewScanner(file)
	var paragraphNumber int
	type m2 struct {
		number int
		text   string
	}

	for content.Scan() {
		paragraphNumber++
		jsonStr := m2{paragraphNumber, content.Text()}
		values, _ := json.Marshal(jsonStr)

		//req, _ := http.NewRequest("POST", "http://127.0.0.1:4000/text", bytes.NewReader(values))

		resp, _ := http.Post("http://localhost:4001/text", "", bytes.NewReader(values))

		//client := &http.Client{}
		//resp, _ := client.Do(req)
		fmt.Println("", resp)
		if paragraphNumber == 3 {
			break
		}
	}

}

//
//func a() {
//	file, err := os.Open("../../task3/files/some_text.txt")
//
//	if err != nil {
//		panic(err.Error())
//	}
//
//	defer file.Close()
//
//	content := bufio.NewScanner(file)
//	var paragraphNumber uint32
//
//	for content.Scan() {
//		paragraphNumber++
//
//		(content.Text(), paragraphNumber)
//	}
//}
