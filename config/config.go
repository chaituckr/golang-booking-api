package config

import (
  "log"
  "io/ioutil"
  "encoding/json"
)

type Config struct {
  Server   string
  Database string
}

func (c *Config) Read() {
  bytes, _ := ioutil.ReadFile("config.json")
  if err := json.Unmarshal(bytes, &c); err != nil {
    log.Fatal(err)
  }
}