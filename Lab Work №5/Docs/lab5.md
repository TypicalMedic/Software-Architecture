# CI для сервисов
## Разделение приложения на микросервисы
Т.к. на данный момент приложение имеет только модуль базы данных и API для нескольких методов, то разделим приложение по составляющей базы данных:

 1. Сервис управления базы данных пользователей.
 2. Сервис управления базы данных встреч.
 
 Все запросы будут приниматься на 1 сервис, который в свою очередь по надобности вызывает 2.
 
## Создание Docker контейнеров
 Для обоих сервисов создается run.Dockerfile, который будет копировать собранное приложение и конфигурацию в контейнер и запускать бинарный файл. Бинарные файлы будут собираться посредством CI:

	FROM  debian:12
	WORKDIR  /app
	COPY  ./cmd/spms/web-app  ./web-app
	COPY  ./cmd/spms/config.yml  ./config.yml
	CMD  ./web-app

Для миграции структур баз данных создается еще один mysql.Dockerfile, который просто копирует SQL запрос в контейнер для последующей миграции:
	
	FROM  mysql:latest
	COPY  ./db_baseline.sql  ./db_baseline.sql

Для доступа к собранным образам они будут отправляться на DockerHub в репозиторий paps с различными тегами: 

 - run – для сервиса пользователей
 - run-meeting – для сервиса встреч
 - mysql – для базы данных пользователей
 - mysql-meeting – для базы данных встреч
 
 ## Сборка контейнеров с docker-compose
Для создания среды, где сервисы могут обращаться к друг-другу или базам данным без каких-либо проблем используется docker-compose. Создается файл docker-compose.yml, в котором указываются зависимости контейнеров друг от друга, пробрасываются порты, указываются используемые тома, и другая конфигурация:

	version: '2.1'  
	name: user-container
	services:
		user-mysql:
			image: ggghfffg/paps:mysql
			container_name: mysql-container-user
			volumes:
				- papsDB:/var/lib/mysql
			environment:
				MYSQL_ROOT_PASSWORD: root
				MYSQL_DATABASE: userdb
				MYSQL_USER: admin
				MYSQL_PASSWORD: adminapp
		web-app-container:
			image: ggghfffg/paps:run
			build: .
			restart: on-failure  
			depends_on: 
				- user-mysql
			links:
				- user-mysql
			ports:
				- 8080:8080
		meeting-mysql:
			image: ggghfffg/paps:mysql-meeting
			container_name: mysql-container-meeting
			volumes:
				- papsDB2:/var/lib/mysql
			environment:
				MYSQL_ROOT_PASSWORD: root
				MYSQL_DATABASE: meeting
				MYSQL_USER: admin
				MYSQL_PASSWORD: adminapp
	meeting-service-container:
		image: ggghfffg/paps:run-meeting
		build: .
		restart: on-failure 
		depends_on: 
			- meeting-mysql
		links:
			- meeting-mysql
		volumes:
			papsDB:
				external: true
			papsDB2:
				external: true

После полной настройки окружения (миграции структуры бд, создания внешних томов) все сервисы стабильно работают:

![enter image description here](https://github.com/TypicalMedic/Software-Architecture/tree/LabWork5/Lab%20Work%20№5/Docs/img/12.png)

## Проверка взаимодействия сервисов

Для проверки были введены логи вызовов методов API для сервиса встреч, также логи ведутся при любых взаимодействиях с базами данных.

В качестве проверки используется метод /meetings:

![enter image description here](https://github.com/TypicalMedic/Software-Architecture/tree/LabWork5/Lab%20Work%20№5/Docs/img/13.png)

В логах контейнеров же отобразится информация вызовов базы данных и метода API:

![enter image description here](https://github.com/TypicalMedic/Software-Architecture/tree/LabWork5/Lab%20Work%20№5/Docs/img/14.png)

![enter image description here](https://github.com/TypicalMedic/Software-Architecture/tree/LabWork5/Lab%20Work%20№5/Docs/img/15.png)

## Настройка непрерывной интеграции

В рамках непрерываной интеграции будут собираться бинарные файлы сервисов и docker образы для сервисов и баз данных. В рамках практики пайплайн будет запускаться при любых пушах и пулл реквестах. Для этого создан файл .github/workflows/build.yml:

	name: CI

	# Controls when the workflow will run
	on:
	  # Triggers the workflow on push or pull request events but only for the "main" branch
	  push:
	    branches: [ "*" ]
	  pull_request:
	    branches: [ "*" ]
	env:
	 LOGIN: ${{ secrets.DOCKER_LOGIN }}
	 NAME: ${{ secrets.DOCKER_NAME }}
	# A workflow run is made up of one or more jobs that can run sequentially or in parallel
	jobs:
	  build-sevices:  
	    # The type of runner that the job will run on
	    runs-on: ubuntu-latest
	    steps:
	      # Checks-out your repository under $GITHUB_WORKSPACE, so your job can access it
	      - uses: actions/checkout@v3
	      - name: Set up Go
	        run: |
	          mkdir /home/runner/.cache
	          mkdir /home/runner/.cache/go-build
	          mkdir /home/runner/.cache/go-mod
	          mkdir /home/runner/go
	          sudo apt install golang-go
	          GOPATH="/home/runner/go"
	          GOCACHE="/home/runner/.cache/go-build"
	          GOMODCACHE="/home/runner/.cache/go-mod"
	          echo GOPATH=$GOPATH >> $GITHUB_ENV
	          echo GOCACHE=$GOCACHE >> $GITHUB_ENV
	          echo GOMODCACHE=$GOMODCACHE >> $GITHUB_ENV
	      - name: Verify dependencies
	        run: |
	          cd "Lab Work №5"/"Source Code"/meeting_service
	          go mod verify
	          cd ../user_service
	          go mod verify
	      - name: Build user service
	        run: |
	          cd "Lab Work №5"/"Source Code"/user_service            
	          go mod download
	          go build -o ./cmd/spms/web-app ./cmd/spms/
	      - name: Build meeting service
	        run: |
	          cd "Lab Work №5"/"Source Code"/meeting_service
	          go mod download
	          go build -o ./cmd/spms/web-app ./cmd/spms/
	      - uses: actions/upload-artifact@master
	        with:
	          name: user-service
	          path: Lab Work №5/Source Code/user_service/cmd/spms/web-app          
	      - uses: actions/upload-artifact@master
	        with:
	          name: meeting-service
	          path: Lab Work №5/Source Code/meeting_service/cmd/spms/web-app
	  # This workflow contains a single job called "build"
	  build-docker-images:
	    needs: build-sevices
	    # The type of runner that the job will run on
	    runs-on: ubuntu-latest

	    # Steps represent a sequence of tasks that will be executed as part of the job
	    steps:
	      # Checks-out your repository under $GITHUB_WORKSPACE, so your job can access it
	      - uses: actions/checkout@v3
	      - uses: actions/download-artifact@master
	        with:
	          name: user-service
	          path: Lab Work №5/Source Code/user_service/cmd/spms
	      - uses: actions/download-artifact@master
	        with:
	          name: meeting-service
	          path: Lab Work №5/Source Code/meeting_service/cmd/spms
	      - name: give executing rights to the services
	        run: |
	          chmod +x ./"Lab Work №5"/"Source Code"/user_service/cmd/spms/web-app          
	          chmod +x ./"Lab Work №5"/"Source Code"/meeting_service/cmd/spms/web-app
	      - name: check current dir
	        run: ls
	      - name: Login to docker.io
	        run:  echo ${{ secrets.DOCKER_PWD }} | docker login -u ${{ secrets.DOCKER_LOGIN }} --password-stdin
	      - name: Build image for user service
	        run: |
	          cd "Lab Work №5"/"Source Code"/user_service
	          docker build --pull --rm -f build/package/run.Dockerfile -t $LOGIN/$NAME:run .
	      - name: Build image for user database
	        run: |
	          cd "Lab Work №5"/"Source Code"/user_service
	          docker build --pull --rm -f build/package/mysql.Dockerfile -t $LOGIN/$NAME:mysql .
	      - name: Build image for meeting service
	        run: |
	          cd "Lab Work №5"/"Source Code"/meeting_service
	          docker build --pull --rm -f build/package/run.Dockerfile -t $LOGIN/$NAME:run-meeting .
	      - name: Build image for meeting database
	        run: |
	          cd "Lab Work №5"/"Source Code"/meeting_service
	          docker build --pull --rm -f build/package/mysql.Dockerfile -t $LOGIN/$NAME:mysql-meeting .
	      - name: Push all builded images to docker.io
	        run: |
	          docker push $LOGIN/$NAME:run
	          docker push $LOGIN/$NAME:run-meeting
	          docker push $LOGIN/$NAME:mysql
	          docker push $LOGIN/$NAME:mysql-meeting

Для того, чтобы войти в аккаунт Docker и загрузить образы требуются логин, пароль и название репозитория. Для этого созданы переменные в GitHub secrets для сокрытия данных:

![enter image description here](https://github.com/TypicalMedic/Software-Architecture/tree/LabWork5/Lab%20Work%20№5/Docs/img/16.png)

Теперь для проверки работы пайплайна внесем изменения в lab5.md:

![enter image description here](https://github.com/TypicalMedic/Software-Architecture/tree/LabWork5/Lab%20Work%20№5/Docs/img/4.png)

Состояние образов на DockerHub:

![enter image description here](https://github.com/TypicalMedic/Software-Architecture/tree/LabWork5/Lab%20Work%20№5/Docs/img/3.png)

При пуше запускается пайплайн с двумя джобами:

![enter image description here](https://github.com/TypicalMedic/Software-Architecture/tree/LabWork5/Lab%20Work%20№5/Docs/img/6.png)
![enter image description here](https://github.com/TypicalMedic/Software-Architecture/tree/LabWork5/Lab%20Work%20№5/Docs/img/7.png)

После выполнения всех шагов пайплайн успешно завершается.

![enter image description here](https://github.com/TypicalMedic/Software-Architecture/tree/LabWork5/Lab%20Work%20№5/Docs/img/10.png)

Образы на DockerHub обновились, теперь их можно собрать локально при помощи docker-compose.yml:

![enter image description here](https://github.com/TypicalMedic/Software-Architecture/tree/LabWork5/Lab%20Work%20№5/Docs/img/11.png)

