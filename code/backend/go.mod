module backend

go 1.24.1

require (
    sqlite v0.0.0
    github.com/mattn/go-sqlite3 v1.14.24
)

replace sqlite => ./sqlite