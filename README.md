# README

This project is for learning purposes. Main goal is to experiment with C
interoperability with CGO. The example used is a reference on how to wrap the
SQLITE C header files in CGO.


## USAGE

- `cd cgosqlite/`
- `make run`
- `sqlite3 data/example.db`
- `SELECT * FROM employee;`


## Note To Self

Faced `killed: 9` issue when running program after indirectly compiling C
program with CGO on MacOS, used the following flags as a fix:

`go build -ldflags -s -o ${PWD}/build/example_db "${PWD}"/src/main.go`

It worked without these flags before I upgraded either (haven't done a follow-up):
 - X-code GUI
 - MacOS Sonoma (Version. 14.1)

Works if I compile C with clang separately on MacOS. On Linux the Go build
compiles without passing `-ldflags -s`.


## Useful links:

- C header: https://www.sqlite.org/2023/sqlite-amalgamation-3440000.zip
- Go sqlite library: https://github.com/mattn/go-sqlite3
- Style guide to header only libs (in deprecated dir): https://github.com/nothings/stb
