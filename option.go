package cfg

// Option configures how cfg loads the configuration.
type Option func(f *cfg)

// File returns an option that configures the filename that cfg
// looks for to provide the config values.
//
// The name must include the extension of the file. Supported
// file types are `yaml`, `yml`, `json` and `toml`.
//
//	cfg.Load(&cfg, cfg.File("config.toml"))
//
// If this option is not used then cfg looks for a file with name `config.yaml`.
func File(name string) Option {
	return func(f *cfg) {
		if len(f.filename) == 0 {
			f.filename = []string{}
		}
		f.filename = append(f.filename, name)
	}
}

// IgnoreFile returns an option which disables any file lookup.
//
// This option effectively renders any `File` and `Dir` options useless. This option
// is most useful in conjunction with the `UseEnv` option when you want to provide
// config values only via environment variables.
//
//	cfg.Load(&cfg, cfg.IgnoreFile(), cfg.UseEnv("my_app"))
func IgnoreFile() Option {
	return func(f *cfg) {
		f.ignoreFile = true
	}
}

// Dirs returns an option that configures the directories that cfg searches
// to find the configuration file.
//
// Directories are searched sequentially and the first one with a matching config file is used.
//
// This is useful when you don't know where exactly your configuration will be during run-time:
//
//	cfg.Load(&cfg, cfg.Dirs(".", "/etc/myapp", "/home/user/myapp"))
//
// If this option is not used then cfg looks in the directory it is run from.
func Dirs(dirs ...string) Option {
	return func(f *cfg) {
		f.dirs = dirs
	}
}

// Tag returns an option that configures the tag key that cfg uses
// when for the alt name struct tag key in fields.
//
//	cfg.Load(&cfg, cfg.Tag("config"))
//
// If this option is not used then cfg uses the tag `cfg`.
func Tag(tag string) Option {
	return func(f *cfg) {
		f.tag = tag
	}
}

// TimeLayout returns an option that conmfigures the time layout that cfg uses when
// parsing a time in a config file or in the default tag for time.Time fields.
//
//	cfg.Load(&cfg, cfg.TimeLayout("2006-01-02"))
//
// If this option is not used then cfg parses times using `time.RFC3339` layout.
func TimeLayout(layout string) Option {
	return func(f *cfg) {
		f.timeLayout = layout
	}
}

// UseEnv returns an option that configures cfg to additionally load values
// from the environment, after it has loaded values from a config file.
//
//	cfg.Load(&cfg, cfg.UseEnv("my_app"))
//
// Values loaded from the environment overwrite values loaded by the config file (if any).
//
// Cfg looks for environment variables in the format PREFIX_FIELD_PATH or
// FIELD_PATH if prefix is empty. Prefix is capitalised regardless of what
// is provided. The field's path is formed by prepending its name with the
// names of all surrounding fields up to the root struct. If a field has
// an alternative name defined inside a struct tag then that name is
// preferred.
//
//	type Config struct {
//	  Build    time.Time
//	  LogLevel string `cfg:"log_level"`
//	  Server   struct {
//	    Host string
//	  }
//	}
//
// With the struct above and UseEnv("myapp") cfg would search for the following
// environment variables:
//
//	MYAPP_BUILD
//	MYAPP_LOG_LEVEL
//	MYAPP_SERVER_HOST
func UseEnv(prefix string) Option {
	return func(f *cfg) {
		f.useEnv = true
		f.envPrefix = prefix
	}
}

// UseStrict returns an option that configures cfg to return an error if
// there exists additional fields in the config file that are not defined
// in the config struct.
//
//	cfg.Load(&cfg, cfg.UseStrict())
//
// If this option is not used then cfg ignores any additional fields in the config file.
func UseStrict() Option {
	return func(f *cfg) {
		f.useStrict = true
	}
}
