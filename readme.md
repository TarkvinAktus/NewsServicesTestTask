# Тестовое задание на написание двух сервисов на Golang

### Реализация
    - Все порты и ключи хранятся в конфигурационных файлах
    - В случае отсутствия нужной записи возвращён код 404
    - В случае иных проблем на стороне storage сервиса код 500
### Обращение ко второму микросервису 
| Method | Route | Params |
| ------- | ------- |------- |
| GET |  "/" | title |
### Пример запроса
```
localhost:8080?title=england
```
### Пример ответа для title = "england"  
```
Title - england With text - some news text
```