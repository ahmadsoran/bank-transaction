package constant

import "os"

var (
	DSN, DBName, Port, SecretKey, HashKey string
	AUTO_MIGRATE                          bool
)

func InitENVS() {
	DSN = os.Getenv("CORE_DATABASE_DATA_DSN")
	Port = os.Getenv("PORT")
	SecretKey = os.Getenv("SECRET_KEY")
	HashKey = os.Getenv("HASH_KEY")
	AUTO_MIGRATE = os.Getenv("CORE_AUTO_MIGRATE") == "true"
}
