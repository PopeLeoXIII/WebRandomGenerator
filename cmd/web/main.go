package main

import (
	"log"
	"net/http"
)

// main Регистрируем обработчики и соответствующие URL-шаблоны, запускаем сервер
func main() {

	//Инициализируем обработчик для главной страницы в маршрутизаторе servemux
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	//Инициализируем FileServer, он будет обрабатывать запросы к файлам из папки "./ui/static"
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	//обработчика для запросов на "/static/"
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//Запускам веб сервер на порте 8080
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
