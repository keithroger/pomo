
build:
	docker build --tag "getstats" .
run:
	docker run -p 7000:8080 getstats:latest
push:
	docker tag  getstats:latest 118175606273.dkr.ecr.us-west-1.amazonaws.com/pomo-lambda-test:latest
	docker push 118175606273.dkr.ecr.us-west-1.amazonaws.com/pomo-lambda-test:latest
zip:
	zip function.zip *
test:
	go test -cover
vet:
	go vet

lint:
	golangci-lint run --enable-all

fmt:
	gofmt -s -w .
