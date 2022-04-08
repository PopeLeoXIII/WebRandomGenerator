package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// main Регистрируем обработчики и соответствующие URL-шаблоны, запускаем сервер
func main() {

	//Инициализируем обработчик для главной страницы в маршрутизаторе servemux
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/ws", toIndex)
	mux.HandleFunc("/wsocket", wsEndpoint)

	//Инициализируем FileServer, он будет обрабатывать запросы к файлам из папки "./ui/static"
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	//обработчика для запросов на "/static/"
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//Запускам веб сервер на порте 8080
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
