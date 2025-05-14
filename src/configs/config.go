package configs

import (
	"fmt"
	"os"
)

type DbConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"dbname"`
	Charset  string `mapstructure:"charset"`
	Timeout  string `mapstructure:"timeout"`
}

func (c *DbConfig) GetDSN() string {
	c.Host = os.Getenv("DB_HOST")
	c.Port = os.Getenv("DB_PORT")
	c.Username = os.Getenv("DB_USERNAME")
	c.Password = os.Getenv("DB_PASSWORD")
	c.DbName = os.Getenv("DB_NAME")
	c.Charset = os.Getenv("DB_CHARSET")
	c.Timeout = os.Getenv("DB_TIMEOUT")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.Username, c.Password, c.Host, c.Port, c.DbName)
}
