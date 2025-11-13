# build program
build:
	go build -o shortcutApp main.go

# run
run:
	go run main.go

# build and run during development
dev: build
	shortcutApp.exe

# remove binary
clean:
	rm -f shortcutApp

# help section
help:
	@echo "Usage: make <target>"
	@echo "---------------------------------------------"
	@echo "  build - build the program"
	@echo "  run - run the program"
	@echo "  dev - build and run during development"
	@echo "  clean - remove existing binary"
	@echo "  help - display this help message"