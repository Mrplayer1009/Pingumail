module pingumail.com/v0

go 1.21.4

replace server => ./server
replace client => ./client

require github.com/mattn/go-sqlite3 v1.14.22 // indirect
