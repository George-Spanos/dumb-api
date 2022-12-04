.PHONE: deps build clean run
deps:
	go mod tidy
build: deps
	go build -o bin/
	cp -r static bin/static
launch: build 
	./bin/learn-go-web
clean:
	rm -rf bin/
run: build
	go run main.go