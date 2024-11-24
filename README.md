[![Go](https://img.shields.io/badge/-Go-464646?style=flat-square&logo=Go)](https://go.dev/)

# log_analyzer
# Анализатор логов

---
## Описание проекта
Проект представляет собой анализатор логов, который собирает и анализирует данные, и выводит текстовый отчёт в выбранном формате с анализом логов.

---
## Технологии
* Go 1.23.0
* DDD (Domain Driven Design)
* Tests

---
## Запуск проекта

**1. Клонировать репозиторий:**
```
git clone https://github.com/KazikovAP/log_analyzer
```

**2. Запустить анализатор логов:**
```
go run cmd/log_analyzer/main.go --path https://raw.githubusercontent.com/elastic/examples/master/Common%20Data%20Formats/nginx_logs/nginx_logs
```

## Пример ответа
```
Загружено логов: 51462
Загружено логов после фильтрации: 51462

#### Общая информация

|        Метрика        |       Значение     |
|:---------------------:|-------------------:|
|       Файл(-ы)        |                url |
|    Начальная дата     |                  - |
|     Конечная дата     |                  - |
|  Количество запросов  |              51462 |
| Средний размер ответа |         659509.51b |
|  95p размера ответа   |              1768b |

#### Запрашиваемые ресурсы

| Ресурс                              | Количество |
|:-----------------------------------:|-----------:|
| HEAD /downloads/product_1 HTTP/1.1  |         13 |
| HEAD /downloads/product_2 HTTP/1.1  |         70 |
| GET /downloads/product_3 HTTP/1.1   |         73 |
| GET /downloads/product_1 HTTP/1.1   |      30272 |
| GET /downloads/product_2 HTTP/1.1   |      21034 |

#### Коды ответа

| Код |                 Имя                 | Количество |
|:---:|:-----------------------------------:|-----------:|
| 416 | Requested Range Not Satisfiable     |          4 |
| 304 | Not Modified                        |      13330 |
| 200 | OK                                  |       4028 |
| 404 | Not Found                           |      33876 |
| 206 | Partial Content                     |        186 |
| 403 | Forbidden                           |         38 |
```

```
go run cmd/log_analyzer/main.go --path https://raw.githubusercontent.com/elastic/examples/master/Common%20Data%20Formats/nginx_logs/nginx_logs --from 2015-05-18T00:00:00Z --to 2015-05-30T00:00:00Z --format adoc
```

```
Загружено логов: 51462
Загружено логов после фильтрации: 34335

=== Общая информация

| Метрика                  | Значение     |
|--------------------------|--------------|
| Файл(-ы)                 |          url |
| Начальная дата           |   18.05.2015 |
| Конечная дата            |   30.05.2015 |
| Количество запросов      |        34335 |
| Средний размер ответа    |   718259.85b |
| 95-й перцентиль размера  |        1768b |

=== Запрашиваемые ресурсы

| Ресурс                              | Количество |
|-------------------------------------|------------|
| GET /downloads/product_1 HTTP/1.1   |      20294 |
| HEAD /downloads/product_2 HTTP/1.1  |         39 |
| GET /downloads/product_3 HTTP/1.1   |         52 |
| GET /downloads/product_2 HTTP/1.1   |      13950 |

=== Коды ответа

| Код |                 Имя                 | Количество |
|-----|-------------------------------------|------------|
| 404 | Not Found                           |      22402 |
| 304 | Not Modified                        |       8971 |
| 200 | OK                                  |       2808 |
| 403 | Forbidden                           |         21 |
| 206 | Partial Content                     |        129 |
| 416 | Requested Range Not Satisfiable     |          4 |
```

---
## Разработал:
[Aleksey Kazikov](https://github.com/KazikovAP)

---
## Лицензия:
[MIT](https://opensource.org/licenses/MIT)
