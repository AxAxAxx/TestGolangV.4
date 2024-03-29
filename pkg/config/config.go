package config

type Configs struct {
	PostgreSQL PostgreSQL
	Server     Server
}

type Server struct {
	Host string
	Port string
}

type PostgreSQL struct {
	Host     string
	Port     string
	Protocol string
	Username string
	Password string
	Database string
	SSLMode  string
}
