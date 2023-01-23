### Telegram Bot на Golang

> https://www.youtube.com/playlist?list=PLFAQFisfyqlWDwouVTUztKX2wUjYQ4T3l

+ Получает на вход ссылки
+ Хранит их
+ Периодически присылает рандомную ссылку на прочтение

+ https://core.telegram.org/api

Архитектура.

main.go -> consumer // работает постоянно

EventProseccor -> Fetcher & Processor // части можно заменить
Fetcher - получает события 
Processor - обрабатывает события

Fetcher -> Client // обрабатывает события от тг
Processor -> Storage - хранит данные.

