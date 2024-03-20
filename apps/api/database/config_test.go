package database

import (
  "testing"
)

func Test_Config_GetDSN(t *testing.T) {
  tests := []struct {
    name string
    config *Config
    expectedDSN string
    wantErr bool
  }{
    {
      "Valid data",
      &Config{Host: "localhost", Database: "testdb", User: "testuser", Password: "testpw"},
      "host=localhost user=testuser password=testpw dbname=testdb port=5432 sslmode=require TimeZone=Europe/Amsterdam",
      false,
    },
    {
      "Missing data",
      &Config{Host: "localhost", Database: "testdb", Password: "testpw"},
      "",
      true,
    },
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      dsn, err := tt.config.GetDSN()

      if (err != nil) != tt.wantErr {
        t.Errorf("GetDSN() error %v, wantErr %v", err, tt.wantErr)
      }

      if dsn != tt.expectedDSN {
        t.Errorf("GetDSN() = %v, want %v", dsn, tt.expectedDSN)
      }
    })
  }
}
