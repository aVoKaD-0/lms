# lms
 Предварительно скачав IDE и Golang на ваш компьютер, нужно скачать в одну директорию 5 файлов из этого репозитория (server.go, agent.go, struct.go, otvet.txt, save.txt, active_vorker.txt).
 
 После открываем файлы otvet.txt и active_vorker.txt (В otvet.txt вы будете видеть id выражения, само выражение и его статус. А в active_vorker.txt вы будете видеть количество активных воркеров на данный момент).

 в консоли пишем go run server.go agent.go struct.go
 (запускаем сразу всё 3 файла) 
 
 В консоли вводим выражение(больше ничего вводить не надо!) и нажимаем enter, а вот дальше уже вводим время, которое вы хотите, чтобы выполнялась такая-то операция такое-то количество времени, какая именно это операция будет написано.
 
 Программа поддерживает ()*/+- и любые целые числа
 
 Примеры выражений, которые вы можете использоват при проверке:
 
     (90-2)*8+24
     
     90-8*2
     
     (40+35*(40-38)/2)+20
     
     *9+2
     
     9/(20-2*10)
     
Завершить работу программы можно закрыв файлы(завершить их работу) или вводом "end" без кавычек вместо выражения.
  
ДОКУМЕНТАЦИЯ

     Пользователь отправляет выражение и время выполнение операции демону (горутина), демон проверяет выражение, 
     если всё верно написано записывает выражение, id и статус в два разных бд (txt файлы). После он отпраляет выражение 
     и время оркестратору, тот в свою очереь проверяет выражение на наличие скобок, создает агента (горутину), 
     который решает выражение (может решать выражение без скобок, если оркестратором найдены скобки, он 
     отправляет сперва выражение в скобках агенту и только потом всё выражение). Агент в свою очередь ищет что-бы 
     выполнить, находит одну операцию и отправляет её воркеру (горутина), тот её решает и возвращает. 
     После оркестратор возвращает ответ демону, демон записывает ответ в бд (txt файл).

![документация](https://github.com/aVoKaD-0/lms/assets/139006972/3100791c-c220-48fe-925a-a4c712886bc7)

тг @aVoKaD_0

