package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/praveen001/go-db-migration/cmd"
)

func main() {
	cmd.Execute()
}
