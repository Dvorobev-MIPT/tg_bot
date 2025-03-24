# telegram bot by PSQL

## Запуск

### Создание/обновление БД
Для создания/обновления Базы данных предлагаю воспользоваться простым скриптом https://github.com/Dvorobev-MIPT/wiki-sql-parser

Однако, это занимает длительное время (порядка 40 минут)

Поэтому, как альтернативный вариант имеется заранее собранная база данных /tg_bot/example_data_base/
Для ее заполнения потребуется запустить готовый SQL скрипт: tg_bot/example_data_base/create.sql. Но вам потребуется написать полный путь к файлам внутри этого скрипта (строки 67-72).

### Загрузка библиотек
Кроме того, потребуеются загрузить внешние библитеки. Для этого из папки проекта в терминале пропишите следующие 4 команды:
go mod init TG-bot
go get -u github.com/texttheater/golang-levenshtein/levenshtein
go get -u github.com/go-telegram-bot-api/telegram-bot-api/v5
go get -u github.com/lib/pq

### Создание бота
Последнее что необходимо сделать - зарегистрировать бота и ввести локальные данные в код:
1. Для регистрации необходимо написать боту @BotFather /start, /newbot и проследовав инструкции, предложенной ботом, создать своего бота. Так же он вышлет токен, который потребуется позже.
2. Ввести параметры для подключения к psql в tg_bot/data_base/params.go
3. Ввести токен в /tg_bot/main.go в строке 10

Поздравляю вы создали своего бота!

## Возможности бота

Он умеет работать с 3 командами:
/start - активирует бота и выводит информацию о командах
/help - выводит информацию о командах
/letter X - выводит всех преподавателей чьи фамилии начинаются с этой буквы, X - любая буква

Остальные команды распознаются как ввод ФИО

Так же стоит заметить, что при вводе любой команды важен регистр букв.

В случае ошибки бот умеет подсказывать 3 наиболее близкие варианта

## Пример работы

![alt text](https://github.com/Dvorobev-MIPT/tg_bot/images/blob/main/images/start_help.png)
![alt text](https://github.com/Dvorobev-MIPT/tg_botimages/blob/main/images/letter_fio.png)
![alt text](https://github.com/Dvorobev-MIPT/tg_botimages/blob/main/images/fio.png)


## БД используемая в боте
![alt text](https://github.com/Dvorobev-MIPT/tg_bot/blob/main/images/Example.jpg)
