package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type randomNumbers struct {
	arr []int
	err error
}

var result *randomNumbers

//toIndex Обработчик для запроса /ws/get
func toIndex(w http.ResponseWriter, r *http.Request) {
	q := strings.Split(r.URL.RawQuery, "&")
	fmt.Println("get: ", r.URL.RawQuery, len(q))
	//возвращает страницу indexWebsocket.html если в запросе нет параметров
	if len(q) == 1 {
		http.ServeFile(w, r, "./ui/html/indexWebsocket.html")
		return
	}

	//если запрос с параметрами
	if len(q) == 3 {
		//Получаем genContext из url запроса
		context, err := urlToGenContext(r.URL)
		if err != nil {
			//Сохраняем неудачный результат и ошибку
			result = &randomNumbers{nil, err}
		} else {
			//Сохраняем полученый срез случайных чисел
			arr, err := genRandFromGenContext(context)
			result = &randomNumbers{arr, err}
			//Данные готовы к отправке клиенту
			fmt.Println("g: success")
		}
	}
}

//unsortWS Обработчик для запроса /wsocket
func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	//Если запрос с параметрами, то вызываем unsortWS
	q := strings.Split(r.URL.RawQuery, "&")
	if len(q) == 3 {
		unsortWS(w, r)
		return
	}

	//upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client Connected")

	//websocket готов к передаче данных, делаем паузу чтобы успеть сгенерировать случайные числа
	fmt.Println("ws: open")
	time.Sleep(time.Second / 10)
	defer ws.Close()

	//Проверям посчитан ли срез со случайными числами
	if result.err != nil {
		sendStr(ws, err.Error())
		return
	}

	//Отправляем случайные числа клиенту
	fmt.Println("ws: send")
	for _, randomNumber := range result.arr {
		str := strconv.Itoa(randomNumber)
		sendStr(ws, str)
	}

	//Стераем отправленные данные
	result = &randomNumbers{nil, errors.New("Ответ не сгенерирован")}
}
