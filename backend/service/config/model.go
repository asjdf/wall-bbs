package config

type Config struct {
	MySQL
	Redis
	HttpEngine
}
type MySQL struct {
	Addr string
	User string
	Pass string
	DB   string
	ConnectTimeout uint
}
type Redis struct {
	Addr string
	Port string
	Pass string
	DBID int
}
type HttpEngine struct {
	Port string
}
