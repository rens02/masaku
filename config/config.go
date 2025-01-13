package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type ProgramConfig struct {
	ServerPort    int
	DBPort        int
	DBHost        string
	DBUser        string
	DBPass        string
	DBName        string
	Secret        string
	CloudinaryKey      string
	OpenAiKey	  string
}

func InitConfig() *ProgramConfig {
	var res = new(ProgramConfig)
	res = loadConfig()

	if res == nil {
		logrus.Fatal("Config : Cannot start program, failed to load configuration")
		return nil
	}

	return res
}

func loadConfig() *ProgramConfig {
	var res = new(ProgramConfig)

	godotenv.Load(".env")

	if val, found := os.LookupEnv("SERVER"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("Config : invalid port value,", err.Error())
			return nil
		}
		res.ServerPort = port
	}

	if val, found := os.LookupEnv("DB_PORT"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("Config : invalid db port value,", err.Error())
			return nil
		}
		res.DBPort = port
	}

	if val, found := os.LookupEnv("DB_HOST"); found {
		res.DBHost = val
	}

	if val, found := os.LookupEnv("DB_USER"); found {
		res.DBUser = val
	}

	if val, found := os.LookupEnv("DB_PASS"); found {
		res.DBPass = val
	}

	if val, found := os.LookupEnv("DB_NAME"); found {
		res.DBName = val
	}

	if val, found := os.LookupEnv("SECRET"); found {
		res.Secret = val
	}

	if val, found := os.LookupEnv("LOUDINARY_KEY"); found {
		res.CloudinaryKey = val
	}

	
	if val, found := os.LookupEnv("OPENAI_KEY"); found {
		res.OpenAiKey = val
	}

	return res

}