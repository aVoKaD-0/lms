package main

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	f, _ := os.Create("active_vorker.txt")
	_, _ = f.WriteString("0")
	id := proverka()
	id++
	max := max_ID()
	if id < max {
		id = max + 1
	}
	mx := sync.Mutex{}
	for id > 0 {
		var equation string
		var time_OperationU = 1
		var time_OperationD = 1
		var time_OperationP = 1
		var time_OperationM = 1
		flag := proverka_dlin()
		if flag == false {
			fmt.Println("все агенты заняты, ожидайте")
			for flag == false {
				flag = proverka_dlin()
			}
			fmt.Println("Продолжайте)")
		}
		fmt.Println("Введите уравнение:")
		fmt.Scanln(&equation)
		operatoinU, operatoinD, operatoinP, operatoinM := number_operations(equation)
		if operatoinU != 0 {
			fmt.Println("Введите время работы для умножения(сек):")
			fmt.Scanln(&time_OperationU)
		}
		if operatoinD != 0 {
			fmt.Println("Введите время работы для деления(сек):")
			fmt.Scanln(&time_OperationD)
		}
		if operatoinP != 0 {
			fmt.Println("Введите время работы для плюса(сек):")
			fmt.Scanln(&time_OperationP)
		}
		if operatoinM != 0 {
			fmt.Println("Введите время работы для минуса(сек):")
			fmt.Scanln(&time_OperationM)
		}
		go func(equation string, ID, time_OperationU, time_OperationD, time_OperationP, time_OperationM int) {
			fmt.Println("ID:", ID, "adopted")
			mx.Lock()
			addendum_otvet(equation, ID)
			addendum_save(equation, ID, time_OperationU, time_OperationD, time_OperationP, time_OperationM)
			mx.Unlock()
			if time_OperationU == 0 || time_OperationD == 0 || time_OperationP == 0 || time_OperationM == 0 {
				mx.Lock()
				change_save(equation, ID)
				change_otvet(ID, equation, "", errors.New("Not enough time"))
				mx.Unlock()
			} else {
				otvet, err := Orchestrator2(ID, time_OperationU, time_OperationD, time_OperationP, time_OperationM, equation)
				mx.Lock()
				change_save(equation, ID)
				change_otvet(ID, equation, otvet, err)
				mx.Unlock()
			}
		}(equation, id, time_OperationU, time_OperationD, time_OperationP, time_OperationM)
		time.Sleep(1 * time.Second)
		id++
	}
}
