
build:
	docker build --tag "getstats" .
run:
	docker run -p 7000:8080 getstats:latest
test:
	curl -X POST "http://localhost:7000/2015-03-31/functions/function/invocations"
push:
	docker tag  getstats:latest 118175606273.dkr.ecr.us-west-1.amazonaws.com/pomo-lambda-test:latest
	docker push 118175606273.dkr.ecr.us-west-1.amazonaws.com/pomo-lambda-test:latest
zip:
	zip function.zip *
