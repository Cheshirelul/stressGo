package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func post(cantidadInsert int, btn int) {

	for i := 0; i < cantidadInsert; i++ {
		url := "http://192.168.0.103:8000/api/postjson"
		fmt.Println("URL:>", url)

		var jsonStr = []byte(`{"idDispositivo":"` + strconv.Itoa(i) + `","btn":"` + strconv.Itoa(btn) + `","sw":"0"}`)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(body))
	}
}

func main() {

	go post(100000, 1)
	go post(140000, 3)
	go post(13000, 4)
	go post(200000, 6)
	go post(30000, 7)
	post(500000, 2)
	fmt.Println("programa finalizo")
}
