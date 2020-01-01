# weather

Use `weather` to see current weather conditions on the command line.

(todo badges)

# Project Overview

Module `gophers.dev/cmds/weather` provides a command-line utility for reporting
current weather conditions.

The `weather` command can be tweaked to report on a specified location, and the
output customized.

# Getting Started

The `weather` command can be installed by running
```bash
go get gophers.dev/cmds/weather
```

The use of `weather` requires a free token from [openweathermap](https://openweathermap.org/)

# Example Usages

#### get the weather for Austin, TX
```bash
$ cat ~/.weather.json
{
  "token": "<token>",
  "zip": "78701",
  "units": "imperial"
}

$ weather
overcast clouds -- 49.0 °F
```

# Contributing

The `gophers.dev/cmds/weather` module is always improving with new features and
error corrections. For contributing bug fixes and new features please file an
issue.

# License

The `gophers.dev/cmds/weather` module is open source under the [MIT](LICENSE)
