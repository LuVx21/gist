package config

type Redis struct {
    Host      string
    Password  string
    Timeout   int
    MaxActive int
    MaxIdle   int
}
