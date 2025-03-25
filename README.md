# telegram bot by PSQL

@test_wiki_1971_bot - мой бот

## Запуск

### Создание/обновление БД
Для начала надо создать базу данных: CREATE DATABASE phystech_wiki;

Далее для заполнения/обновления Базы данных предлагаю воспользоваться простым скриптом https://github.com/Dvorobev-MIPT/wiki-sql-parser

Однако, это занимает длительное время (порядка 40 минут)

Поэтому, как альтернативный вариант имеется заранее собранная база данных /tg_bot/example_data_base/

Для ее заполнения потребуется запустить готовый SQL скрипт: tg_bot/example_data_base/create.sql. Но вам потребуется написать полный путь к файлам внутри этого скрипта (строки 67-72).

### Загрузка библиотек
Кроме того, потребуеются загрузить внешние библитеки. Для этого из папки проекта в терминале пропишите следующие 4 команды:

1. go mod init TG-bot
2. go get -u github.com/texttheater/golang-levenshtein/levenshtein
3. go get -u github.com/go-telegram-bot-api/telegram-bot-api/v5
4. go get -u github.com/lib/pq

### Создание бота
Последнее что необходимо сделать - зарегистрировать бота и ввести локальные данные в код:

1. Для регистрации необходимо написать боту @BotFather /start, /newbot и проследовав инструкции, предложенной ботом, создать своего бота. Так же он вышлет токен, который потребуется позже.
2. Ввести параметры для подключения к psql в tg_bot/data_base/params.go
3. Ввести токен в /tg_bot/main.go в строке 10

Поздравляю вы создали своего бота!

## Возможности бота

Он умеет работать с 3 командами:

1. /start - активирует бота и выводит информацию о командах
2. /help - выводит информацию о командах
3. /letter X - выводит всех преподавателей чьи фамилии начинаются с этой буквы, X - любая буква

Остальные команды распознаются как ввод ФИО

Так же стоит заметить, что при вводе любой команды важен регистр букв.

В случае ошибки бот умеет подсказывать 3 наиболее близкие варианта

## Пример работы

![alt text](https://github.com/Dvorobev-MIPT/tg_bot/blob/main/images/start_help.png)
![alt text](https://github.com/Dvorobev-MIPT/tg_bot/blob/main/images/fio.png)
![alt text](https://github.com/Dvorobev-MIPT/tg_bot/blob/main/images/letter_fio.png)


## БД используемая в боте
![alt text](https://github.com/Dvorobev-MIPT/tg_bot/blob/main/images/Example.jpg)

