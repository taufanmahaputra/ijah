package config

type Config struct {
	App      App
	Database Database
}

type App struct {
	Appname string
	Port    string
}

type Database struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
	SSL      string
}