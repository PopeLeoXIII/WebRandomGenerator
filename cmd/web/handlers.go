package web

import (
	"fmt"
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
	//if r.URL.Path != "/" {
	//	fmt.Print("n")
	//	r.ParseForm() //анализ аргументов,
	//	fmt.Println(r.Form)  // ввод информации о форме на стороне сервера
	//	fmt.Println("path", r.URL.)
	//	fmt.Println("scheme", r.URL.Scheme)
	//	fmt.Println(r.Form["url_long"])
	//	for k, v := range r.Form {
	//		fmt.Println("key:", k)
	//		fmt.Println("val:", strings.Join(v, ""))
	//	}
	//	http.NotFound(w, r)
	//	return
	//}

	if r.Method == http.MethodPost {
		fmt.Print("p")
		homePost(w, r)
		return
	}

	if r.Method == http.MethodGet {
		fmt.Print("g")
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
	r.ParseForm()
	data := getViewData(r.Form)
	arr, genErr := genRand(data)

	if genErr == nil {
		data.Result = true
		data.Message = "Генерация успешна"
		data.Arr = arr
	} else {
		data.Result = false
		data.Message = genErr.Error()
		data.Arr = nil
	}

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
