package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var _ = godotenv.Load(".env") // Cargar del archivo llamado ".env"

var (
	ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", //"<user>:<password>@tcp(127.0.0.1:3306)/<dbname>"
		os.Getenv("USER_DB"),
		os.Getenv("PASSWORD_DB"),
		os.Getenv("HOST_DB"),
		os.Getenv("PORT_DB"),
		os.Getenv("NAME_DB"))
	SecretPassword = fmt.Sprint(os.Getenv("SECRET_JWT"))
)
