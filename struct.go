package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"sync"
	"time"
	"unicode/utf8"
)

var IsLetter = regexp.MustCompile(`^[0-9+-/*()]+$`).MatchString

func validation(equation string) (int, int, error) { // проверяем на валидность выражение
	x := 0
	y := 0
	U := 0
	D := 0
	P := 0
	M := 0
	err := IsLetter(equation)
	if !err {
		return 0, 0, errors.New("incorrect input")
	}
	for i := 0; i < utf8.RuneCountInString(equation); i++ {
		if string(equation[i]) == "(" {
			x++
		}
		if string(equation[i]) == ")" {
			y++
		}
		if string(equation[i]) == "*" {
			U++
		}
		if string(equation[i]) == "/" {
			D++
		}
		if string(equation[i]) == "+" {
			P++
		}
		if string(equation[i]) == "-" {
			M++
		}
		if i < utf8.RuneCountInString(equation)-1 {
			if string(equation[i]) == "*" && (string(equation[i+1]) == "+" || string(equation[i+1]) == "-" || string(equation[i+1]) == "*" || string(equation[i+1]) == "/") {
				return 0, 0, errors.New("incorrect input")
			}
			if string(equation[i]) == "/" && (string(equation[i+1]) == "+" || string(equation[i+1]) == "-" || string(equation[i+1]) == "*" || string(equation[i+1]) == "/") {
				return 0, 0, errors.New("incorrect input")
			}
			if string(equation[i]) == "+" && (string(equation[i+1]) == "+" || string(equation[i+1]) == "-" || string(equation[i+1]) == "*" || string(equation[i+1]) == "/") {
				return 0, 0, errors.New("incorrect input")
			}
			if string(equation[i]) == "-" && (string(equation[i+1]) == "+" || string(equation[i+1]) == "-" || string(equation[i+1]) == "*" || string(equation[i+1]) == "/") {
				return 0, 0, errors.New("incorrect input")
			}
			if string(equation[i]) == "(" && (string(equation[i+1]) == ")" || string(equation[i+1]) == "+" || string(equation[i+1]) == "*" || string(equation[i+1]) == "/") {
				return 0, 0, errors.New("incorrect input")
			}
		} else if string(equation[0]) == "*" || string(equation[0]) == "/" {
			return 0, 0, errors.New("incorrect input")
		}
	}
	if x != y || (U == 0 && D == 0 && P == 0 && M == 0) {
		return 0, 0, errors.New("incorrect input")
	}
	return x, y, nil
}

func addendum_otvet(equation string, ID int) { // добавляем выражение в ответ
	f2, _ := os.Open("otvet.txt")
	buffer2 := make([]byte, 2048)
	_, _ = f2.Read(buffer2)
	sch := 0
	for t, b := range buffer2 { // смотрим количество выражений в ответе
		if b == 10 {
			sch++
		}
		if b == 0 {
			buffer2 = buffer2[:t]
			break
		}
	}
	if sch == 4 { // если в ответе больше 5 выражение, убираем первое
		t := 0
		p := 0
		flag_delete := true
		for i, n := range buffer2 {
			if n == 32 {
				p = i
			}
			if n == 10 {
				if string(buffer2[p+1:i]) != "adopted" {
					buffer := buffer2[i+1:]
					buffer2 = buffer2[:t]
					for _, by := range buffer {
						buffer2 = append(buffer2, by)
					}
					flag_delete = false
				}
				t = i + 1
			}
			if i == len(buffer2)-1 {
				if string(buffer2[p:i]) != "adopted" {
					buffer2 = buffer2[t : i+1]
					flag_delete = false
				}
				t = i + 1
			}
			if flag_delete == false {
				break
			}
		}
		file2, _ := os.Create("otvet.txt")
		_, _ = file2.WriteString(string(buffer2))
	}
	f, _ := os.Open("otvet.txt")
	buffer := make([]byte, 2048)
	_, _ = f.Read(buffer)
	file, err := os.OpenFile("otvet.txt", os.O_APPEND, 0600)
	if err != nil {
		fmt.Println("ID:", ID, err)
	}
	if buffer[0] == 0 { // добавляем новое выражение
		_, err = file.WriteString(strconv.Itoa(ID) + " " + equation + " adopted")
		if err != nil {
			fmt.Println("ID:", ID, "\n", err)
		}
	} else {
		_, err = file.WriteString("\n" + strconv.Itoa(ID) + " " + equation + " adopted")
		if err != nil {
			fmt.Println("ID:", ID, err)
		}
	}
}

func addendum_save(equation string, ID, time_OperationU, time_OperationD, time_OperationP, time_OperationM int) { // добавляем выражение в базу агентов
	file, _ := os.Open("save.txt")
	buffer := make([]byte, 2048)
	_, _ = file.Read(buffer)
	f, err := os.OpenFile("save.txt", os.O_APPEND, 0600)
	if err != nil {
		fmt.Println("ID:", ID, err)
	}
	if buffer[0] == 0 {
		_, err = f.WriteString(strconv.Itoa(ID) + " " + strconv.Itoa(time_OperationU) + " " + strconv.Itoa(time_OperationD) + " " + strconv.Itoa(time_OperationP) + " " + strconv.Itoa(time_OperationM) + " " + equation)
		if err != nil {
			fmt.Println("ID:", ID, "\n", err)
		}
	} else {
		_, err = f.WriteString("\n" + strconv.Itoa(ID) + " " + strconv.Itoa(time_OperationU) + " " + strconv.Itoa(time_OperationD) + " " + strconv.Itoa(time_OperationP) + " " + strconv.Itoa(time_OperationM) + " " + equation)
		if err != nil {
			fmt.Println("ID:", ID, err)
		}
	}
}

func change_otvet(ID int, equation, otvet string, err2 error) { // меняем статус выражения в ответе
	file, _ := os.Open("otvet.txt")
	buffer := make([]byte, 2048)
	_, _ = file.Read(buffer)
	for i, b := range buffer {
		if b == 0 {
			buffer = buffer[:i]
			break
		}
	}
	nachalo := 0
	st := ""
	flag := true
	gl := 0
	for i, byt := range buffer { // пробегаем по ответам
		if byt == 10 || i == len(buffer)-1 {
			st = string(buffer[nachalo:i])
			probel := 0
			for t, s := range st {
				if string(s) == " " && probel == 0 { // смотрим id выражения, если не совпадает идем на следующее выражение
					if string(st[:t]) != strconv.Itoa(ID) {
						break
					} else {
						probel = t
					}
				} else if string(s) == " " && probel != 0 { // иначе меняем статус выражения
					if equation == st[probel+1:t] {
						if err2 == nil {
							st = st[:probel+1] + (equation + "=" + otvet) + " ok"
						} else {
							st = st[:probel+1] + equation + " " + fmt.Sprint(err2)
						}
						flag = false
					}
					break
				}
			}
			if flag == true {
				nachalo = i + 1
			}
		}
		if flag == false {
			if i == len(buffer)-1 {
				gl = i + 1
			} else {
				gl = i
			}
			break
		}
	}
	byts := make([]byte, 0)
	for i, b := range buffer[:nachalo] {
		if i == nachalo {
			break
		}
		byts = append(byts, b)
	}
	for _, i := range st {
		byts = append(byts, byte(i))
	}
	for _, i := range buffer[gl:] {
		byts = append(byts, i)
	}
	f, _ := os.Create("otvet.txt")
	_, _ = f.Write(byts) // перезаписываем ответы
}

func change_save(equation string, ID int) { // удаляем выражение с базы агентов
	file, _ := os.Open("save.txt")
	buffer := make([]byte, 2048)
	_, _ = file.Read(buffer)
	for i, b := range buffer {
		if b == 0 {
			buffer = buffer[:i]
			break
		}
	}
	nachalo := 0
	st := ""
	flag := true
	byts := make([]byte, 0)
	for i, byt := range buffer { // пробегаемя по базе агентов
		if byt == 10 || i == len(buffer)-1 {
			if i == len(buffer)-1 {
				i++
			}
			st = string(buffer[nachalo:i])
			for t, s := range st {
				if string(s) == " " {
					if string(st[:t]) != strconv.Itoa(ID) { // id не совпадает идем на следующее выражение
						break
					} else {
						for t2 := utf8.RuneCountInString(st) - 1; t2 >= 0; t2-- { // удаляем выражение
							if string(st[t2]) == " " {
								if string(st[t2+1:]) == equation {
									for b := 0; b < nachalo; b++ {
										byts = append(byts, buffer[b])
									}
									for b := i + 1; b < len(buffer); b++ {
										byts = append(byts, buffer[b])
									}
									flag = false
									break
								}
							}
						}
					}
				}
			}
			if flag == true {
				nachalo = i + 1
			}
		}
		if flag == false {
			break
		}
	}
	if len(byts) > 0 {
		if byts[len(byts)-1] == 10 {
			byts = byts[:len(byts)-1]
		}
	}
	f, _ := os.Create("save.txt")
	_, _ = f.Write(byts) // обновляем базу агентов
}

func proverka() int { // проверяем есть ли у нас данные в базе агентов с запуском программы
	f, _ := os.Open("save.txt")
	buffer := make([]byte, 2048)
	_, _ = f.Read(buffer)
	mx := sync.Mutex{}
	var equation string
	var time_OperationU int
	var time_OperationD int
	var time_OperationP int
	var time_OperationM int
	var ID int
	if buffer[0] != 0 { // если выражений есть отправляем их агентам
		nachalo := 0
		for i, b := range buffer {
			if b == 0 {
				buffer = buffer[:i]
				break
			}
		}
		for i, b := range buffer {
			st := ""
			if b == 10 || i == len(buffer)-1 {
				if i != len(buffer)-1 {
					st = string(buffer[nachalo:i])
				} else {
					st = string(buffer[nachalo : i+1])
				}
				nach := 0
				sch := 0
				for t, s := range st {
					if string(s) == " " && sch == 0 {
						ID, _ = strconv.Atoi(st[:t])
						nach = t + 1
						sch++
					} else if string(s) == " " && sch == 1 {
						time_OperationU, _ = strconv.Atoi(st[nach:t])
						nach = t + 1
						sch++
					} else if string(s) == " " && sch == 2 {
						time_OperationD, _ = strconv.Atoi(st[nach:t])
						nach = t + 1
						sch++
					} else if string(s) == " " && sch == 3 {
						time_OperationP, _ = strconv.Atoi(st[nach:t])
						nach = t + 1
						sch++
					} else if string(s) == " " && sch == 4 {
						time_OperationM, _ = strconv.Atoi(st[nach:t])
						equation = string(st[t+1:])
						nach = t + 1
						sch++
						break
					}
				}
				go func(equation string, ID, time_OperationU, time_OperationD, time_OperationP, time_OperationM int) {
					fmt.Println("ID:", ID, "adopted")
					otvet, err := Orchestrator(ID, time_OperationU, time_OperationD, time_OperationP, time_OperationM, equation)
					mx.Lock()
					change_save(equation, ID)
					change_otvet(ID, equation, otvet, err)
					mx.Unlock()
				}(equation, ID, time_OperationU, time_OperationD, time_OperationP, time_OperationM)
				nachalo = i + 1
				time.Sleep(1 * time.Second)
			}
		}
	}
	return ID
}

func check_dlin_save() bool { // проверяем количество выражений в базе агентов
	f, _ := os.Open("save.txt")
	buffer := make([]byte, 2048)
	_, _ = f.Read(buffer)
	sch := 0
	for _, b := range buffer {
		if b == 10 {
			sch++
		}
		if b == 0 {
			break
		}
	}
	if sch == 4 {
		return false
	}
	return true
}

func max_ID() int { // смотрим какой id самый большой в ответах
	f, _ := os.Open("otvet.txt")
	buffer := make([]byte, 2048)
	_, _ = f.Read(buffer)
	if buffer[0] == 0 {
		return 1
	}
	max_id, _ := strconv.Atoi(string(buffer[0]))
	flag := false
	t := 0
	for i, b := range buffer {
		if b == 10 {
			t = i
			flag = true
		}
		if flag == true {
			if string(b) == " " {
				max, _ := strconv.Atoi(string(buffer[t+1 : i]))
				if max_id < max {
					max_id = max
				}
				flag = false
			}
		}
	}
	return max_id
}

func number_operations(equation string) (int, int, int, int) { // смотрим какие операции имеются в выражении
	operatoinU := 0
	operatoinD := 0
	operatoinP := 0
	operatoinM := 0
	for _, s := range equation {
		if string(s) == "+" {
			operatoinP++
		}
		if string(s) == "-" {
			operatoinM++
		}
		if string(s) == "*" {
			operatoinU++
		}
		if string(s) == "/" {
			operatoinD++
		}
	}
	return operatoinU, operatoinD, operatoinP, operatoinM
}

func check_to_repeat(expression string) bool { // проверка на повторное выражение
	f, _ := os.Open("save.txt")
	buffer := make([]byte, 2048)
	_, _ = f.Read(buffer)
	sch := 0
	t := 0
	if buffer[0] != 0 {
		for i, b := range buffer {
			if string(b) == " " {
				sch++
				if sch == 5 {
					t = i
				}
			}
			if sch == 5 && b == 10 || b == 0 {
				st := string(buffer[t+1 : i])
				if st == expression {
					return false
				}
				sch = 0
			}
		}
	}
	return true
}
