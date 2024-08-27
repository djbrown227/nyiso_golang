package config

type Config struct {
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
    APIPort    string
}

func LoadConfig() Config {
    return Config{
        DBHost:     "127.0.0.1",
        DBPort:     "3306",
        DBUser:     "root",
        DBPassword: "Pavilion227",
        DBName:     "nyiso_database",
        APIPort:    ":8080",
    }
}
