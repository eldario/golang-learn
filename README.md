# golang-learn
Learning golang

Each step is a new directory in root

### Task1

Add first "HelloWord" program

### Task2

Added task3 with list and sorter package

### Task3

Added task3 with SortedMap

### Task4

Added functionality:

|Package name|Method name(s)|Description|
|---:|---|---|
|**pipe**|`NewPipe(ch1, ch2 chan string, func(msg string) string)`|соединяет два канала, все что попадает в ch1 выводит в ch2 с примененной функцией|
|**fanIn**|`NewFanIn(chansIn []chan string, out chan string)`|всё что попадает в каналы chansIn должно попасть в out|
|**fanOut**|`NewFanOut(in chan string, chansOut []chan string)`|всё что попадает в канал in должно попасть в каналы out|
|**bufferedChan**|`NewChan(in, bufferSize int)`|все что не вместилось в буфер (если медленно читаем) - отбрасываем|
