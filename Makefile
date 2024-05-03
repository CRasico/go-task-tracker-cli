run: build 
	@./bin/go-task-tracker-cli	
	
debug: build 
	@./bin/go-task-tracker-cli -debug 

build:
	@go build -C cmd -o ../bin/go-task-tracker-cli
