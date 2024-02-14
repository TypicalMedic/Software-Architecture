FROM debian:12
WORKDIR /app
EXPOSE 8080
COPY ./cmd/spms/web-app ./web-app
COPY ./cmd/spms/config.yml ./config.yml
CMD ./web-app