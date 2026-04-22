package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Деление на ноль")
	}
	return a / b, nil
}

func main() {
	res1, err1 := divide(10.5, 2.0)
	if err1 != nil {
		fmt.Println("Ошибка:", err1)

	} else {
		fmt.Println("Все ок: 10.5 / 2.0 =", res1)
	}

	res2, err2 := divide(7.0, 0.0)
	if err2 != nil {
		fmt.Println("Ошибка:", err2)

	} else {

		fmt.Println("Все ок: 7.0 / 0.0 =", res2)

	}
}
