package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type randomNumbers struct {
	arr []int
	err error
}

var result randomNumbers

func toIndex(w http.ResponseWriter, r *http.Request) {
	q := strings.Split(r.URL.RawQuery, "&")
	fmt.Println("r: ", r.URL.RawQuery, len(q))
	if len(q) == 1 {
		http.ServeFile(w, r, "./ui/html/indexWebsocket.html")
		return
	}

	if len(q) == 3 {
		data := ViewData{r.URL.Query().Get("n"),
			r.URL.Query().Get("max"),
			r.URL.Query().Get("genn"),
			false, "", nil}
		context, err := checkData(data)
		if err != nil {
			result = randomNumbers{nil, err}
		} else {
			arr, err := genRandFromGenContext(context)
			result = randomNumbers{arr, err}
		}
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	defer ws.Close()

	// helpful log statement to show connections
	log.Println("Client Connected")
	err = ws.WriteMessage(1, []byte("Hi Client! "))
	if err != nil {
		log.Println(err)
	}

	if result.err != nil {
		err = ws.WriteMessage(1, []byte(result.err.Error()))
		if err != nil {
			log.Println(err)
		}
	}

	for _, randomNumber := range result.arr {
		str := strconv.Itoa(randomNumber)
		if err := ws.WriteMessage(1, []byte(str)); err != nil {
			log.Println(err)
			return
		}

	}
}
