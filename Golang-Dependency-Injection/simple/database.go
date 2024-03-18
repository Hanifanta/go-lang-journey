package simple

type Database struct {
	Name string
}

type DBPostgres Database
type DBMySql Database

func NewDBPostgres() *DBPostgres {
	return (*DBPostgres)(&Database{Name: "Postgres"})
}

func NewDBMySql() *DBMySql {
	return (*DBMySql)(&Database{Name: "MySql"})
}

type DatabaseRepository struct {
	DBPostgres *DBPostgres
	DBMySql    *DBMySql
}

func NewDatabaseRepository(postgres *DBPostgres, mySql *DBMySql) *DatabaseRepository {
	return &DatabaseRepository{DBPostgres: postgres, DBMySql: mySql}
}
