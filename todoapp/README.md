# TODO APP

Телеграм бот, который позволяет управлять вашими задачами.

## Должен содержать следующий функционал

1. Создать новую задачу:
    - Указать описание задачи
    - Задать приоритет задачи
    - Установить время/интервал когда бот должен уведомить о задаче (если требуется)
2. Взять задачу в работу
3. Завершить задачу
4. Просмотреть список актуальных задач с возможностью получения задач по времени/приоритету и времени
5. Просмотреть список завершённых задач
6. Получить статистику за указанный период времени (количество выполненных задач, общее время)

## Бот должен сам

1. Уведомлять о задаче

## Этапы выполнения

1. Проектирование БД
2. Подключение telegram API, разобраться с API
3. Определиться, какие методы нужны для приложения. Сделать пакет для взаимодействия с БД.
4. Некоторые задачи требуется выполнять периодически, можно использовать https://github.com/go-co-op/gocron