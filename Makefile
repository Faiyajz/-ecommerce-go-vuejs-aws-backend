run:
	echo "Run Triggered"
	echo "Golang"
	go run main.go

build:
	env GOOS=linux go build -o bin/api api/main.go
	env GOOS=linux go build -o bin/hello testLambda/main.go

deploy: build
		serverless deploy --aws-profile junk

deploy_dev: build
			serverless deploy --aws-profile junk --param="allowedOrigin=http://localhost:8080" --stage dev