package main

import (
	"errors"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

// genContext Данные необходимые для генерации случайных чисел (параметры генератора)
type genContext struct {
	n, maxValue, genN, cb int
}

// gen генерирует случайные числа
func gen(c chan<- int, maxValue int) {
	for {
		randomNumber := 1 + rand.Intn(maxValue)
		c <- randomNumber
	}
}

// counter следит за уникальность сгенерированных чисел
func counter(c <-chan int, n int) []int {
	//создаем отображение для определения уникальности генерируемых случайных чисел
	counter := make(map[int]int)
	//создаем срез для хранения генерируемых случайных чисел
	result := make([]int, 0, n)

	for i := 0; i < n; {
		randomNumber := <-c
		//если полученное число уникально, добавляем его в срез
		if counter[randomNumber] == 0 {
			counter[randomNumber] = 1
			result = append(result, randomNumber)
			i++
		}
	}
	return result
}

// checkData Проверяет правильность поступивших данных
func checkData(data ViewData) (genContext, error) {
	//конвертируем параметры в int
	n, errN := strconv.Atoi(data.NInput)
	maxValue, errMax := strconv.Atoi(data.MaxInput)
	genN, errGenN := strconv.Atoi(data.GenNInput)

	context := genContext{n, maxValue, genN, 10}

	//Проверяем все ли параметры являются числаи
	if errN != nil || errMax != nil || errGenN != nil {
		return context, errors.New("Введите числа")
	}

	//Проверяем все ли параметры > 0
	if n <= 0 || maxValue <= 0 || genN <= 0 {
		return context, errors.New("Параметры должны быть > 0")
	}

	//Максимально возможное значение генерируемых уникальных случайных чисел не должно превышать их количества
	if n > maxValue {
		return context, errors.New("Количество чисел не может превышать максимальное значение")
	}
	return context, nil
}

func genRandFromGenContext(context genContext) ([]int, error) {
	//Генерируем зерно для rand
	rand.Seed(time.Now().UnixNano())
	//Создаем канал в которм будут передаваться генерируемые случайные числа
	var c chan int = make(chan int, context.cb)

	//Запускаем genN генераторов
	for i := 0; i < context.genN; i++ {
		go gen(c, context.maxValue)
	}

	//Вызываем counter который будет обрабатывать генерируемые случайные числа
	arr := counter(c, context.n)
	//Сортируем полученный срез случайных чисел
	sort.Ints(arr)
	return arr, nil
}

// genRand Запускает генерацию случайных чисел
func genRandFromViewData(data ViewData) ([]int, error) {
	//Проверяем входные данные
	context, err := checkData(data)
	if err != nil {
		return nil, err
	}

	return genRandFromGenContext(context)
}
