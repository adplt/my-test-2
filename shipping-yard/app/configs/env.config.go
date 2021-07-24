package configs

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	variables "project/app/variables"

	"runtime"

	"github.com/joho/godotenv"
)

// Config is
var (
	Env = GetENV()
	OS  = runtime.GOOS
)

// ENV function
func ENV(args ...string) string {
	var key, defaultValue string
	if len(args) < 1 {
		return ""
	}

	key = args[0]
	if len(args) > 1 {
		defaultValue = args[1]
	}

	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	err := LoadFile()
	if err != nil {
		return ""
	}

	// recheck again after load file
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultValue
}

// LoadFile function
func LoadFile(dirEnv ...string) error {
	rootDir := variables.REPO_NAME

	if len(dirEnv) > 0 {
		rootDir = dirEnv[0]
	}

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	env := os.Getenv("ENV")
	if strings.ToUpper(env) == "TEST" {
		dir = filepath.Clean(filepath.Join(dir, "../.."))
	}

	re := regexp.MustCompile(`^(.*` + rootDir + `)`)
	rootPath := string(re.Find([]byte(dir)))
	for _, v := range []string{"/", `\`} {
		dirSplit := strings.Split(dir, v)
		rootPathLength := len(strings.Split(rootPath, v))
		filePathENV := strings.Join(dirSplit[:rootPathLength], v)
		if OS == "windows" {
			checkPath := filePathENV[len(filePathENV)-1:]
			if !(checkPath == "/" || checkPath == "\\") && checkPath != ":" {
				filePathENV += "/"
			}
		} else {
			filePathENV += "/"
		}
		filePathENV += `.env`

		matches, err := filepath.Glob(filePathENV)
		if err != nil {
			return err
		}
		if len(matches) != 0 {
			err = godotenv.Load(filePathENV)
			if err != nil {
				return err
			}
			break
		}
	}
	return nil
}

// CreateHash function
func CreateHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (e *Environtment) setENV() {
	env := os.Getenv("ENV")
	switch env {
	case variables.DEVELOPMENT:
		e.AppHost = ENV("DEV__APP__HOST")
		e.AppPort = ENV("DEV__APP__PORT")
		e.AppURL = ENV("DEV__APP__URL")
		e.DBConn1 = ENV("DEV__DB_CONNECTION_1")
		e.DBConn2 = ENV("DEV__DB_CONNECTION_2")
		e.DbDialect = ENV("DEV__DB__DIALECT")
		e.DbUserName = ENV("DEV__DB__USER")
		e.DbPassword = ENV("DEV__DB__PASS")
		e.DbName1 = ENV("DEV__DB__DB_NAME_1")
		e.DbName2 = ENV("DEV__DB__DB_NAME_2")
		e.DbPort = ENV("DEV__DB__PORT")
		e.DbHost = ENV("DEV__DB__HOST")
		e.Timeout = ENV("DEV__API__TIMEOUT")
		e.RedisHost = ENV("DEV__REDIS__HOST")
		e.RedisPort = ENV("DEV__REDIS__PORT")
		e.Logger = ENV("DEV__LOGGER__ALLOW")
		e.JwtSecretKey = ENV("DEV__JWT__SECRET_KEY")
		e.Sha256Salt = ENV("DEV__SHA256__SALT")
		e.SwaggerAllow = ENV("DEV__SWAGGER__ALLOW")
		e.SwaggerURL = ENV("DEV__SWAGGER__URL")
	case variables.TESTING:
		e.AppHost = ENV("TEST__APP__HOST")
		e.AppPort = ENV("TEST__APP__PORT")
		e.AppURL = ENV("TEST__APP__URL")
		e.DBConn1 = ENV("TEST__DB_CONNECTION_1")
		e.DBConn2 = ENV("TEST__DB_CONNECTION_2")
		e.DbDialect = ENV("TEST__DB__DIALECT")
		e.DbUserName = ENV("TEST__DB__USER")
		e.DbPassword = ENV("TEST__DB__PASS")
		e.DbName1 = ENV("TEST__DB__DB_NAME_1")
		e.DbName2 = ENV("TEST__DB__DB_NAME_2")
		e.DbPort = ENV("TEST__DB__PORT")
		e.DbHost = ENV("TEST__DB__HOST")
		e.Timeout = ENV("TEST__API__TIMEOUT")
		e.RedisHost = ENV("TEST__REDIS__HOST")
		e.RedisPort = ENV("TEST__REDIS__PORT")
		e.Logger = ENV("TEST__LOGGER__ALLOW")
		e.JwtSecretKey = ENV("TEST__JWT__SECRET_KEY")
		e.Sha256Salt = ENV("TEST__SHA256__SALT")
		e.SwaggerAllow = ENV("TEST__SWAGGER__ALLOW")
		e.SwaggerURL = ENV("TEST__SWAGGER__URL")
	case variables.PRODUCTION:
		e.AppHost = ENV("PROD__APP__HOST")
		e.AppPort = ENV("PROD__APP__PORT")
		e.AppURL = ENV("PROD__APP__URL")
		e.DBConn1 = ENV("PROD__DB_CONNECTION_1")
		e.DBConn2 = ENV("PROD__DB_CONNECTION_2")
		e.DbDialect = ENV("PROD__DB__DIALECT")
		e.DbUserName = ENV("PROD__DB__USER")
		e.DbPassword = ENV("PROD__DB__PASS")
		e.DbName1 = ENV("PROD__DB__DB_NAME_1")
		e.DbName2 = ENV("PROD__DB__DB_NAME_2")
		e.DbPort = ENV("PROD__DB__PORT")
		e.DbHost = ENV("PROD__DB__HOST")
		e.Timeout = ENV("PROD__API__TIMEOUT")
		e.RedisHost = ENV("PROD__REDIS__HOST")
		e.RedisPort = ENV("PROD__REDIS__PORT")
		e.Logger = ENV("PROD__LOGGER__ALLOW")
		e.JwtSecretKey = ENV("PROD__JWT__SECRET_KEY")
		e.Sha256Salt = ENV("PROD__SHA256__SALT")
		e.SwaggerAllow = ENV("PROD__SWAGGER__ALLOW")
		e.SwaggerURL = ENV("PROD__SWAGGER__URL")
	}
}

// GetENV is
func GetENV() *Environtment {
	E := &Environtment{}
	E.setENV()
	return E
}
