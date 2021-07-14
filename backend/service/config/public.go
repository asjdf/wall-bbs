package config

func GetMySQLConfig() MySQL {
	return configService.MySQL
}

func GetRedisConfig() Redis {
	return configService.Redis
}

func GetHttpEngineConfig() HttpEngine {
	return configService.HttpEngine
}
