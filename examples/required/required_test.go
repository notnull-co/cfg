package required

import (
	"fmt"
	"time"

	"github.com/notnull-co/cfg"
)

type Config struct {
	Cache struct {
		Size            int           `conf:"size" default:"10000"`
		CleanupInterval time.Duration `conf:"cleanup_interval" validate:"required"`
	} `conf:"cache"`
	Tags []string `conf:"tags" validate:"required"`
}

func ExampleLoad() {
	var conf Config
	err := cfg.Load(&conf, cfg.File("config.json"), cfg.Tag("conf"))
	fmt.Println(err)

	// Output:
	// cache.cleanup_interval: required validation failed, tags: required validation failed
}
