package dbconfig

import "fmt"

type users struct {
	id int
	name string
	age int
	modalidade int
	donated bool
}

const PostgresDriver = "postgres"
const User = "postgres"
const Host = "localhost"
const Port = "5432"
const Password = "postgres"
const DbName = "users"
const TableName = "users"

var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+
    "password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)