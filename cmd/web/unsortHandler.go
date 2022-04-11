package main

import (
	"github.com/gorilla/websocket"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

//sendStr Отправляет строку через websocket
func sendStr(ws *websocket.Conn, str string) {
	if err := ws.WriteMessage(1, []byte(str)); err != nil {
		log.Println(err)
	}
}

//urlToGenContext Собирает из строки url запроса genContext
func urlToGenContext(u *url.URL) (genContext, error) {
	data := ViewData{u.Query().Get("n"),
		u.Query().Get("max"),
		u.Query().Get("genn"),
		false, "", nil}
	return checkData(data)
}

//unsort Обработчик для запроса /ws, возвращает страницу indexUnsort.html заполненую дефолтными значениями
func unsort(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/indexUnsort.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	//создаем вариант ViewData заполненый дефолтными значениями
	defaultData := ViewData{"10", "10", "1", false, "", nil}

	err = ts.Execute(w, defaultData)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
	}
}

//unsortWS Обработчик для запроса /wsocket с параметрами
func unsortWS(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	defer ws.Close()

	//Получаем genContext из url запроса
	context, err := urlToGenContext(r.URL)
	//Если преобразовать не получилось, то отправляем клиенту сообщение ошибки
	if err != nil {
		sendStr(ws, err.Error())
	} else {
		//Создаем канал для приема сгенерированных случайных чисел
		var c chan int = make(chan int, context.cb)

		//Запускаем genN генераторов
		for i := 0; i < context.genN; i++ {
			go gen(c, context.maxValue)
		}

		//Отправляем клиенту n случайных чисел
		for i := 0; i < context.n; i++ {
			randomNumber := <-c
			sendStr(ws, strconv.Itoa(randomNumber))
		}
	}

	//tmpl, err := template.ParseFiles("./ui/html/indexUnsort.html")
	//data := ViewData{r.URL.Query().Get("n"),
	//	r.URL.Query().Get("max"),
	//	r.URL.Query().Get("genn"),
	//	true, "Успешная генерация", nil}
	//err = tmpl.Execute(w, data)

}
