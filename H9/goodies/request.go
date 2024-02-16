package goodies

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetParcel() {
	res, err := http.Get("http://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res.Body)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	sb := string(body)
	log.Println(sb)
}

func PostParcel() {
	data := map[string]interface{}{
		"title":  "Dionysius",
		"body":   "Raka",
		"userid": 1,
	}

	requestJson, err := json.Marshal(data)
	client := &http.Client{}
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("POST", "http://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(requestJson))

	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}
