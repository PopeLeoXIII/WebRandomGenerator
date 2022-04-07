package web

import (
	"log"
	"net/http"
)

// main Регистрируем обработчики и соответствующие URL-шаблоны, запускаем сервер
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
