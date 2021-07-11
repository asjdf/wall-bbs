package config

type Conf struct {
	DB struct{
		Addr string
		User string
		Pass string
		DB   string
	}
	Redis struct{
		Addr string
		Port string
		Pass string
		DBID int
	}
	Server struct{
		Port string
	}
}
