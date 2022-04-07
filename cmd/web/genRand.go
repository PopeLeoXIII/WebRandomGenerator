package web

import (
	"errors"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

// gen генерирует случайные числа
func gen(c chan<- int, maxValue int) {
	//id := rand.Intn(maxValue)
	for {
		randomNumber := 1 + rand.Intn(maxValue)
		//fmt.Println(el, id)
		c <- randomNumber
	}
}

// counter следит за уникальность сгенерированных чисел
func counter(c <-chan int, n int) []int {
	counter := make(map[int]int)
	result := make([]int, 0, n)

	for i := 0; i < n; {
		randomNumber := <-c
		if counter[randomNumber] == 0 {
			counter[randomNumber] = 1
			result = append(result, randomNumber)
			//fmt.Println(el)
			i++
		}
	}
	return result
}

// genContext Данные необходимые для генерации случайных чисел (параметры генератора)
type genContext struct {
	n, maxValue, genN, cb int
}

// checkData Проверяет правильность поступивших данных
func checkData(data ViewData) (genContext, error) {
	n, e1 := strconv.Atoi(data.NInput)
	maxValue, e2 := strconv.Atoi(data.MaxInput)
	genN, e3 := strconv.Atoi(data.GenNInput)

	context := genContext{n, maxValue, genN, 10}
	if e1 != nil || e2 != nil || e3 != nil {
		return context, errors.New("Введите числа")
	}

	if n <= 0 || maxValue <= 0 || genN <= 0 {
		return context, errors.New("Параметры должны быть > 0")
	}

	if n > maxValue {
		return context, errors.New("Количество чисел не может превышать максимальное значение")
	}
	return context, nil
}

// genRand Запускает генерацию случайных чисел
func genRand(data ViewData) ([]int, error) {
	context, err := checkData(data)
	if err != nil {
		return nil, err
	}
	rand.Seed(time.Now().UnixNano())
	var c chan int = make(chan int, context.cb)

	for i := 0; i < context.genN; i++ {
		go gen(c, context.maxValue)
	}

	arr := counter(c, context.n)
	sort.Ints(arr)
	return arr, nil
}
