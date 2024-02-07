FROM debian:12
WORKDIR /app
EXPOSE 8080
CMD cd ./cmd/spms; ./web-app