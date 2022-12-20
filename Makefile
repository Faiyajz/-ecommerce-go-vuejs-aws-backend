run:
	echo "Run Triggered"
	echo "Golang"
	go run main.go

build:
	env GOOS=linux go build -o bin/hello testLambda/main.go

deploy: build
		serverless deploy --aws-profile junk