# Go-Survey

This application is a simple survey where the go app run on `localhost:8080` and when you send a `POST` request as a `JSON` format it will be added automatically in the `MONGODB` database.

`NOTE:` In the first 2 options make sure you already have `MONGODB` [Installed](https://www.mongodb.com/docs/manual/administration/install-community/) and running on `localhost:27017`

## 1. Build and Run the app locally:
`NOTE:` Install [GoLang](https://go.dev/doc/install) then run the following commands inside the app directory
```
go mod tidy
```
```
go build main.go
```
```
./main
```
## 2. Run the app in a Docker container:
```
docker build -t <image-name> .
```

`NOTE:` In the following option you don't need `MONGODB` installed or running since it will be running directly from a docker container.

## 3. Run both `GoApp` and `MONGODB` in docker containers:
```
docker-compose up
```

## 4. Deploy application on `Minikube`:
```
chmod +x deploy.sh

deploy.sh
```

### Finally to delete deployment:
```
chmod +x delete.sh

./delete.sh
```
