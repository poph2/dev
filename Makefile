
init-nodejs-h:
	go run main.go init nodejs -h

build:
	go build

install:
	go install

hive-build:
	go run main.go build -c "/Users/pop/PopProjects/sd-service"