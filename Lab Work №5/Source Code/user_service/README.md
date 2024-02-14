## Как запустить приложение
Для запуска приложения требуется склонировать данный репозиторий и находиться в его корневой папке. Приложение развернуто при помощи Docker:

### **Запуск приложения:**


1. Склонировать данный репозиторий 
2. Открыть терминал (cmd, git bash, powershell, terminal)
3. Зайти в корневую папку проекта
4. Собрать docker образы приложения следуюими командами:
    
        docker build --pull --rm -f build/package/run.Dockerfile -t ggghfffg/paps:run .
        docker build --pull --rm -f build/package/build.Dockerfile -t ggghfffg/paps:build .
        docker build --pull --rm -f build/package/mysql.Dockerfile -t ggghfffg/paps:mysql .

5. Запустить контейнер сборки приложений:

        docker run --name docker-build-web-app-container -it -v <путь/к/корневой/папке/проекта>:/app ggghfffg/paps:build
        docker rm docker-build-web-app-container
6. Запустить контейнеры для полноценной работы приложения:

        docker-compose -f build/package/docker-compose.yml up -d
7. Если схема БД еще не была импортирована (при первом запуске), то после успешной настройки MySQL в контейнере (в логах будет сообщение _ready for connections_) требуется выполнить следующие команды:

        docker exec -i mysql-container mysql -u root --password=root
        source db_baseline.sql
        exit


> Т.к. контейнеру с MySQL требуется время для настройки, контейнер веб-приложения может несколько раз перезапуститься, ожидая открытия портов MySQL. При успешном запуске приложения в логах контейнера будет сообщение _App started!_

> Приложение будет доступно по адресу `localhost:8080`

### **Доступные методы**

> Описание методов представлено в лабораторной №4

- /meetings
- /meetings/filter?student={student_id}
- /meeting/{meeting_id}
- /meeting/add

## Над проектом работала

[@Mmmmmmm1](https://gitlab.com/Mmmmmmm1) (Мария Плетнева) - Backend Developer (cool.maru.pletneva@yandex.ru);

