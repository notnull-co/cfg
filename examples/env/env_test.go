package env

import (
	"fmt"
	"os"

	"github.com/notnull-co/cfg"
)

type Config struct {
	Database struct {
		Host     string `validate:"required"`
		Port     int    `validate:"required"`
		Database string `validate:"required" cfg:"db"`
		Username string `validate:"required"`
		Password string `validate:"required"`
	}
	Container struct {
		Args []string `default:"[/bin/sh]"`
	}
}

func ExampleLoad() {
	os.Clearenv()
	check(os.Setenv("APP_DATABASE_HOST", "pg.internal.corp"))
	check(os.Setenv("APP_DATABASE_USERNAME", "mickey"))
	check(os.Setenv("APP_DATABASE_PASSWORD", "mouse"))
	check(os.Setenv("APP_CONTAINER_ARGS", "[-p,5050:5050]"))

	var conf Config
	err := cfg.Load(&conf, cfg.UseEnv("app"))
	check(err)

	fmt.Println(conf.Database.Host)
	fmt.Println(conf.Database.Port)
	fmt.Println(conf.Database.Database)
	fmt.Println(conf.Database.Username)
	fmt.Println(conf.Database.Password)
	fmt.Println(conf.Container.Args)

	// Output:
	// pg.internal.corp
	// 5432
	// users
	// mickey
	// mouse
	// [-p 5050:5050]
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
