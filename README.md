# WB Tech L0 task

Задание L0 курса "Горутиновый golang"

## Задание
Необходимо разработать демонстрационный сервис с простейшим интерфейсом, отображающий данные о заказе. Модель данных можно найти в конца задания.

Что нужно сделать:
- Развернуть локально PostgreSQL
   - Создать свою БД
   - Настроить своего пользователя
   - Создать таблицы для хранения полученных данных
- Разработать сервис
    - Реализовать подключение к брокерам и подписку на топик orders в Kafka
    - Полученные данные записывать в БД
    - Реализовать кэширование полученных данных в сервисе (сохранять in memory)
    - В случае прекращения работы сервиса необходимо восстанавливать кэш из БД
    - Запустить http-сервер и выдавать данные по id из кэша
- Разработать простейший интерфейс отображения полученных данных по id заказа

## Стек технологий
- PostgreSQL, Apache Kafka
- Используемые библиотеки:
  - [Gin](https://github.com/gin-gonic/gin) для роутинга
  - [GORM](https://github.com/go-gorm/gorm) для маппинга таблиц в структуры
  - [Sarama](https://github.com/IBM/sarama) для работы с Apache Kafka
  - [zap](https://github.com/uber-go/zap) для логгирования

## Запуск
1) Зайти в папку deployments
```console
cd deployments
```
3) Запустить через Docker Compose
```console
docker compose up -d
```
