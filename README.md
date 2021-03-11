# Follow below steps to run this app in your local machine

## Environment variables to be added

export MYSQL_CONNECTION_URL="Username:Password@/Database?parseTime=true"

export PORT= *Mention the PORT where you want to run the application by default it will be 8080*

## Command to download all dependencies

go get -d ./...

## Command to run the application

go run main.go
