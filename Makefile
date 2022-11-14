build: 
	go build -o bin/
	cp -r static bin/static
clean:
	rm -rf bin/
run: build
	./bin/learn-go-web