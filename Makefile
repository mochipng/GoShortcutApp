# build program
build:
	go build -o shortcutApp .

# run
run:
	go run .

# build and run during development
dev: build
	./shortcutApp

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