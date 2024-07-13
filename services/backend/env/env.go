package env

import (
	"fmt"
	"os"
)

type EnvManager struct {
	dbHost 		string;
	dbUser		string;
	dbName		string;
	dbPassword	string;
	sslmode		string;
	DbSource	string;
	DbDriver 	string;
}

func NewEnvManager() *EnvManager {
	return &EnvManager{}
}

func (env *EnvManager)dataSourceName() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
	 env.dbHost, env.dbUser, env.dbPassword, env.dbName, env.sslmode)
}

func SetupEnv(file string) *EnvManager {
	env := NewEnvManager()
	env.loadEnv()
	env.DbSource = env.dataSourceName()
	return env
}

func (e *EnvManager) loadEnv() {
	
/* 	fd, err := os.Open(file)
	if err != nil {
		fmt.Printf("Error opening file, file not found")
	}
	defer fd.Close() */

	e.dbHost = os.Getenv("DBHOST")
	e.dbUser = os.Getenv("DBUSER")
	e.dbName = os.Getenv("DBNAME")
	e.dbPassword = os.Getenv("DBPASS")
	e.sslmode = os.Getenv("SSLMODE")
	e.DbDriver = os.Getenv("DBDRIVER")
}
