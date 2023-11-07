.PHONY: build clean run

clean:
	rm ${PWD}/build/example_db
	rm ${PWD}/data/example.db

build:
	go build -ldflags -s -o ${PWD}/build/example_db "${PWD}"/src/main.go

run: build
	${PWD}/build/example_db
