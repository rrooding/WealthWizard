package database

import (
  "fmt"
)

type Config struct {
  Host string
  Database string
  User string
  Password string
}

func (c *Config) isValid() bool {
  return c.Host != "" && c.Database != "" && c.User != "" && c.Password != ""
}

func (c *Config) GetDSN() (string, error) {
  if !c.isValid() {
    return "", fmt.Errorf("invalid database configuration")
  }

  return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=require TimeZone=Europe/Amsterdam", c.Host, c.User, c.Password, c.Database), nil
}
