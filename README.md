# Follow below steps to run this app in your local machine

## Environment variables to be added

export MYSQL*CONNECTION_URL="<User>:<Password>@/<Database>?parseTime=true"
export PORT= \_Mention the PORT where you want to run the application by default it will be 8080*

## Command to run download all dependencies

go get -d ./...

## Command to run the application

go run main.go
