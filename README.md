# lms
 Предварительно скачав IDE и Golang на ваш компьютер, нужно скачать в одну директорию 5 файлов из этого репозитория (server.go, agent.go, struct.go, otvet.txt, save.txt, active_vorker.txt).
 
 После открываем файлы otvet.txt и active_vorker.txt (В otvet.txt вы будете видеть id выражения, само выражение и его статус. А в active_vorker.txt вы будете видеть количество активных воркеров на данный момент).
 
 Далее запускаем файлы server.go, agent.go, struct.go.
 
 В консоли вводим выражение(больше ничего вводить не надо!) и нажимаем enter, а вот дальше уже вводим время, которое вы хотите, чтобы выполнялась такая-то операция такое-то количество времени, какая именно это операция будет написано.
 
 Программа поддерживает ()*/+- и любые числа
 
 Примеры выражений, которые вы можете использоват при проверке:
 
     (90-2)*8+24
     
     90-8*2
     
     (40+35*(40-38)/2)+20
     
     *9+2
     
     9/(20-2*10)
     
  Завершить работу программы можно закрыв файлы(завершить их работу) или вводом "end" без кавычек вместо выражения.
