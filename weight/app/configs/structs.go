package configs

// Environtment struct
type Environtment struct {
	AppHost      string
	AppPort      string
	AppURL       string
	DBConn1      string
	DBConn2      string
	DbDialect    string
	DbUserName   string
	DbPassword   string
	DbName1      string
	DbName2      string
	DbPort       string
	DbHost       string
	Timeout      string
	RedisHost    string
	RedisPort    string
	Logger       string
	JwtSecretKey string
	Sha256Salt   string
	SwaggerURL   string
	SwaggerAllow string
}
