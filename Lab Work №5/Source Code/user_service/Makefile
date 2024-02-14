build-docker-images: ## собираем образы для будущих сборки и запуска приложения
	docker build --pull --rm -f build/package/run.Dockerfile -t ggghfffg/paps:run .	
	docker build --pull --rm -f build/package/build.Dockerfile -t ggghfffg/paps:build .
	docker build --pull --rm -f build/package/mysql.Dockerfile -t ggghfffg/paps:mysql .

build-in-docker: ## запускаем контейнер с биндом к корневой папке хоста, в нем собираются приложения и сохраняются на хосте
	docker run --name docker-build-web-app-container -it -v $(PWD):/app ggghfffg/paps:build
	docker rm docker-build-web-app-container
 
run-in-docker: ## запускаем контейнер запуска приложения и бд
	docker-compose -f build/package/docker-compose.yml up -d

run-commands-to-build-go: ## вызываются при запуске контейнера сборки приложения
	go mod download
	go build -o ./cmd/spms/web-app ./cmd/spms/

