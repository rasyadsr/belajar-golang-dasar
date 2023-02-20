package database

var connection string

/* akan di eksekusi ketika package di panggil */
func init() {
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
