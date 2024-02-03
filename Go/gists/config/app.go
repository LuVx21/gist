package config

type Config struct {
    Server ServerConfig
    Log    LogConfig
    MySQL  MySQL
    Redis  Redis
}

type ServerConfig struct {
    Port  string
    Debug bool
}

type LogConfig struct {
    LogDir   string
    MainLog  string
    ErrorLog string
}

type MySQL struct {
    Host     string
    Username string
    Password string
    Dbname   string
}

type Redis struct {
    Host      string
    Password  string
    Timeout   int
    MaxActive int
    MaxIdle   int
}
