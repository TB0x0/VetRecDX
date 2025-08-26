package db

type DBConfig struct {
	host      string
	port      string
	user      string
	password  string
	dbname    string
	openconns string
	maxidle   string
}