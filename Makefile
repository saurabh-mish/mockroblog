default: exectests

exectests:
	go test -v ./tests/*

coverage:
	go test -v ./tests/* -coverprofile=coverage.out -coverpkg ./...

report:
	go tool cover -html coverage.out -o coverage.html
