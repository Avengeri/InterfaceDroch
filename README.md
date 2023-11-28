# Программа для тренировки замены БД.

## Реализуется по щелчку пальцев заменой одной переменной 

### Порядок запуска программы:

Для работы с БД типа map:
- В файле main.go вставьте аргумент `store_new.StoreMap` в соответствующее поле функции `store_new.NewStore`
- Запустите функцию `main` в файле `main.go`
- Посмотрите в консоль и вы все увидите

Для работы с БД типа PostgresQL:

1. Установите Docker на свой компьютер. Инструкцию по установке можно найти [здесь](https://www.docker.com/)
2. Проверьте установлен ли `Docker Compose` с помощью команды  `docker compose version`. Если он не установлен, то не мои проблемы, решите сами этот вопрос! :thinking:	
3. В файле main.go вставьте аргумент `store_new.StorePostgres` в соответствующее поле функции `store_new.NewStore`
4. Убедитесь что вы находитесь в корне проекта и введите команду `make up` для запуска БД и `make down` для остановки БД
5. Запустите функцию `main` в файле `main.go`

### Если не дошло, то вот этот блок кода куда надо вставить все, что я писал выше

```Go
func main() {
	ctx := context.Background()
	storage, _ := store_new.NewStore(store_new.StoreMap)
	if storage == nil {
		log.Fatal("storage is nil")
		return
	}
```

> Если вы хотите пожаловаться, то обязательно пишите [сюда](https://t.me/zak47) 

![Суслик](https://www.pngitem.com/pimgs/m/285-2854983_gopher-golang-russian-gophercon-russia-2018-hd-png.png)

