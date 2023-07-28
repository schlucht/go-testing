
testcov:
	go test -coverprofile=coverage.out && go tool cover -html=coverage.out

test:
	go test .

run: 
	go run webapp/cmd/web/*.go