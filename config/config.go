package config

type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"db-name"`
	Config   string `mapstructure:"config"`
}

type RedisConfig struct {
	RedisDb     string `mapstructure:"redis-db"`
	RedisAddr   string `mapstructure:"redis-addr"`
	RedisPw     string `mapstructure:"redis-pw"`
	RedisDbName string `mapstructure:"redis-db-name"`
}

func (m *MySQLConfig) Dsn() string {
	return m.Username + ":" + m.Password + "@(" + m.Host + ":" + m.Port + ")/" + m.DBName + "?" + m.Config
}

type SystemConfig struct {
	Port string `mapstructure:"port"`
}

type ServerConfig struct {
	MySqlConfig  MySQLConfig  `mapstructure:"mysql"`
	SystemConfig SystemConfig `mapstructure:"system"`
	RedisConfig  RedisConfig  `mapstructure:"redis"`
}
