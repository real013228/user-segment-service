# User-segment service
Сервис, хранящий пользователя и сегменты, в которых он состоит и позволяющий выполнять CRUD операции над ними.

## Содержание
- [User-segment service](#user-segment-service)
  - [Содержание](#содержание)
  - [Использование](#использование)
    - [На Windows](#на-windows)
    - [На Linux или Mac](#на-linux-или-mac)
  - [Разработка](#разработка)
    - [Требования](#требования)
    - [Установка зависимостей](#установка-зависимостей)
  - [To do](#to-do)
  - [Команда проекта](#команда-проекта)

## Использование
Склонируйте репозиторий с помощью команды:
```sh
$ git clone https://github.com/real013228/user-segment-service.git
```
Далее, для запуска, в корне проекта запустите следующие команды:
### На Windows
```sh
$ go build
$ .\user-segment-service.exe
```
или
```sh
$ CompileDaemon -command="./user-segment-service"
```

### На Linux или Mac
```sh
$ go build && ./user-segment-service
```

## Разработка

### Требования
Для установки и запуска проекта, необходимы модули, прописанные в go.mod файле.

### Установка зависимостей
Для установки зависимостей, выполните команды:
```sh
$ go get github.com/githubnemo/CompileDaemon
$ go install github.com/githubnemo/CompileDaemon
$ go get github.com/joho/godotenv
$ go get -u github.com/gin-gonic/gin
$ go get -u gorm.io/gorm
$ go get -u gorm.io/driver/postgres
```

## To do
- [x] Добавить крутое README
- [x] Реализовать однослойку
- [ ] Реализовать многослойную и расширяемую архитектуру

## Команда проекта

- [Кирилл Саввинов](https://t.me/sakir0132) —  Backend developer