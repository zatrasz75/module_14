package main

import (
	"errors"
	"fmt"
)

func main() {

	var razmer1, razmer2 int
	fmt.Println("Введите размер первого массива:")
	_, err := fmt.Scanln(&razmer1)
	if err != nil {
		fmt.Println(errors.New("проверьте типы входных параметров"))
	}
	ar1 := make(map[string]interface{}, razmer1) // Инициализируем 1 мапу

	fmt.Println("Введите размер второго массива:")
	_, err = fmt.Scanln(&razmer2)
	if err != nil {
		fmt.Println(errors.New("проверьте типы входных параметров"))
	}
	ar2 := make(map[string]interface{}, razmer2) // Инициализируем 2 мапу

	fmt.Println("Заполните элементами первый массив:")
	arrayMap(ar1, razmer1)
	fmt.Println("Заполните элементами второй массив:")
	arrayMap(ar2, razmer2)

	fmt.Println("Список одинаковых элементов:", identicalItems(ar1, ar2))

}

// Заполняет мапу заданного размера элементами
func arrayMap(m map[string]interface{}, arrayRazmer int) {
	for i := 0; i < arrayRazmer; i++ {
		var elem string
		_, err := fmt.Scanln(&elem)
		if err != nil {
			fmt.Println(errors.New("вы не ввели никаких элементов, повторите заново"))
		}
		m[elem] = nil
	}
}

// Выводит одинаковые элементы из двух массивов
func identicalItems(ar1 map[string]interface{}, ar2 map[string]interface{}) []string {
	var identical []string
	if len(ar1) < len(ar2) {
		for key := range ar1 {
			if _, ok := ar2[key]; ok {
				identical = append(identical, key)
			}
		}
	} else {
		for key := range ar2 {
			if _, ok := ar1[key]; ok {
				identical = append(identical, key)
			}
		}
	}
	return identical
}
