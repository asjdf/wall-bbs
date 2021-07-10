package config

type Conf struct {
	DB struct{
		Addr string
		User string
		Pass string
		DB   string
	}
	Server struct{
		Port string
	}
}
