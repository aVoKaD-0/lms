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
	_, _ = f.WriteString("0") // запишем, что активных воркеров нету
	id := proverka()
	id++
	max := max_ID()
	if id < max {
		id = max + 1
	} // смотрим какой id дать вражению
	mx := sync.Mutex{}
	var equation string
	for equation != "end" { // чтобы выйти из цикла введите вместо выражения "end" без кавычек или завершите работу программы)
		equation = ""
		var time_OperationU = 1
		var time_OperationD = 1
		var time_OperationP = 1
		var time_OperationM = 1
		flag := check_dlin_save() // проверяет сколько работают агентов
		if flag == false {
			fmt.Println("все агенты заняты, ожидайте")
			for flag == false {
				flag = check_dlin_save()
			}
			fmt.Println("Продолжайте)")
		}
		fmt.Println("Введите уравнение:")
		fmt.Scanln(&equation)
		operatoinU, operatoinD, operatoinP, operatoinM := number_operations(equation) // проверяем какие операции имеются в выражении
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
		if check_to_repeat(equation) {
			go func(equation string, ID, time_OperationU, time_OperationD, time_OperationP, time_OperationM int) { // так называемый демон
				fmt.Println("ID:", ID, "adopted")
				mx.Lock()
				addendum_otvet(equation, ID)
				addendum_save(equation, ID, time_OperationU, time_OperationD, time_OperationP, time_OperationM) // добавление в базу агентов и ответа
				mx.Unlock()
				if time_OperationU == 0 || time_OperationD == 0 || time_OperationP == 0 || time_OperationM == 0 { // если нету никаких операций возвращаем ошибку
					mx.Lock()
					change_save(equation, ID)
					change_otvet(ID, equation, "", errors.New("Not enough time"))
					mx.Unlock()
				} else { // иначе отправляем оркестратору
					otvet, err := Orchestrator(ID, time_OperationU, time_OperationD, time_OperationP, time_OperationM, equation)
					mx.Lock()
					change_save(equation, ID)
					change_otvet(ID, equation, otvet, err) // записываем решение в базу ответа и удаляем с базы агента
					mx.Unlock()
				}
			}(equation, id, time_OperationU, time_OperationD, time_OperationP, time_OperationM)
		} else {
			mx.Lock()
			addendum_otvet(equation, id)
			change_otvet(id, equation, "", errors.New("already in progress"))
			mx.Unlock()
		}
		time.Sleep(1 * time.Second) // останавливаем ввод, чтобы не было никаких проблем)
		id++
	}
}
