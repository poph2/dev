
init-nodejs-h:
	go run main.go init nodejs basic-nodejs-app -c examples

build:
	go build

install:
	go install

hive-build:
	go run main.go build -c "/Users/pop/PopProjects/sd-service"

hive-clean:
	go run main.go clean -c "/Users/pop/PopProjects/sd-service"
