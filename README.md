# cfg

cfg is originally a fork from [fig](https://github.com/kkyr/fig). Check out their incredible work in the [fig](https://github.com/kkyr/fig) repository.

cfg is a tiny library for loading an application's config file and its environment into a Go struct. Individual fields can have default values defined or be marked as required.

## Why cfg?

- Define your **configuration**, **validations** and **defaults** in a single location
- Optionally **load from the environment** as well
- Only **3** external dependencies
- Full support for`time.Time`, `time.Duration` & `regexp.Regexp`
- Tiny API
- Decoders for `.yaml`, `.json` and `.toml` files

## Getting Started

`$ go get -d github.com/notnull-co/cfg`

Define your config file:

```yaml
# config.yaml

build: "2020-01-09T12:30:00Z"

server:
    ports:
      - 8080
    cleanup: 1h

logger:
    level: "warn"
    trace: true
```

Define your struct along with _validations_ or _defaults_:

```go
package main

import (
  "fmt"

  "github.com/notnull-co/cfg"
)

type Config struct {
  Build  time.Time `cfg:"build" validate:"required"`
  Server struct {
    Host    string        `cfg:"host" default:"127.0.0.1"`
    Ports   []int         `cfg:"ports" default:"[80,443]"`
    Cleanup time.Duration `cfg:"cleanup" default:"30m"`
  }
  Logger struct {
    Level   string         `cfg:"level" default:"info"`
    Pattern *regexp.Regexp `cfg:"pattern" default:".*"`
    Trace   bool           `cfg:"trace"`
  }
}

func main() {
  var conf Config
  err := cfg.Load(&conf)
  // handle your err
  
  fmt.Printf("%+v\n", conf)
  // Output: {Build:2019-12-25 00:00:00 +0000 UTC Server:{Host:127.0.0.1 Ports:[8080] Cleanup:1h0m0s} Logger:{Level:warn Pattern:.* Trace:true}}
}
```

If a field is not set and is marked as *required* then an error is returned. If a *default* value is defined instead then that value is used to populate the field.

By default Cfg searches for a file named `config.yaml` and `secret.yaml` in the directory it is run from. Add your own files and dirs by passing additional parameters to `Load()`:

```go
cfg.Load(&conf,
  cfg.File("settings.json"),
  // includes the file settings.json to the list of cfg files. 
  cfg.Dirs(".", "/etc/myapp", "/home/user/myapp"),
) // searches for ./settings.json, /etc/myapp/settings.json, /home/user/myapp/settings.json as well as config.yaml and secret.yaml on the same dirs.

```

## Environment

Need to additionally fill fields from the environment? It's as simple as:

```go
cfg.Load(&conf, cfg.UseEnv("MYAPP"))
```

## Usage

See usage [examples](/examples).

## Contributing

PRs are welcome! Please explain your motivation for the change in your PR and ensure your change is properly tested and documented.
