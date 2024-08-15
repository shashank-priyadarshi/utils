package constants

type Type string

type Database Type

const (
	MONGODB Database = "mongodb"
	REDIS   Database = "redis"
	MYSQLDB Database = "mysql"
)

type RDBMS Database

type Driver RDBMS

const (
	MYSQL       Driver = "mysql"
	POSTGRES    Driver = "postgres"
	SQLITE      Driver = "sqlite"
	VITESS      Driver = "vitess"
	COCKROACHDB Driver = "cockroachdb"
)

type ORM RDBMS

const (
	GORM ORM = "gorm"
)
