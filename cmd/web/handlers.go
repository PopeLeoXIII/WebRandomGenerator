package main

import (
	"html/template"
	"net/http"
	"net/url"
)

// ViewData Данные для общения с html страницей
type ViewData struct {
	NInput    string
	MaxInput  string
	GenNInput string
	Result    bool
	Message   string
	Arr       []int
}

// home Главный обработчик
func home(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		homePost(w, r)
		return
	}

	if r.Method == http.MethodGet {
		homeGet(w, r)
		return
	}
}

// homeGet Обработчик для запросов GET к главной странице
func homeGet(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/index.html")
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

// getViewData Преобразует данные принятые формой в ViewData
func getViewData(values url.Values) ViewData {
	n := values.Get("nInput")
	max := values.Get("maxInput")
	genN := values.Get("genNInput")

	return ViewData{n, max, genN, false, "", nil}

}

// homePost Обработчик для запросов POST к главной странице
func homePost(w http.ResponseWriter, r *http.Request) {
	//Получаем двнные из формы и создаем экземпляр ViewData
	r.ParseForm()
	data := getViewData(r.Form)

	//Генерируем случайные числа и проверяем не произошла ли ошибка
	arr, genErr := genRandFromViewData(data)
	if genErr == nil {
		//Заполняем поля ViewData сгенерированными числами
		data.Result = true
		data.Message = "Генерация успешна"
		data.Arr = arr
	} else {
		//Не получилось сгенерировать числа, обрабатываем ошибку
		data.Result = false
		data.Message = genErr.Error()
		data.Arr = nil
	}

	//Отправляем заполненную ViewData на клиент
	tmpl, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
	}
}
