package model

type ServerOptions struct {
	DBOptions DBOptions
	Port      *int
}

type DBOptions struct {
	User     string
	Password string
	Connect  string
	DBName   string
}
