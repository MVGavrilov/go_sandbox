package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"mysrvr/app/common/database"
	"mysrvr/app/common/server"
	"mysrvr/app/router"
)

type Config struct {
  Database database.PostgresConfig `json:"database"`
  Server server.ServerConfig `json:"server"`
}

func LoadConfig(path string) Config {
  var config Config
  configFile, err := os.Open(path)
  defer configFile.Close()
  if err != nil {
    fmt.Println(err.Error())
  }
  jsonParser := json.NewDecoder(configFile)
  jsonParser.Decode(&config)
  return config
}

func main() {
  configPath := os.Getenv("CONFIG_PATH")
  if configPath == "" {
    configPath = "config/config.json"
  }

  log.Default().Println("Loading config from", configPath)
  appConfig:=LoadConfig(configPath)

  log.Default().Println("Connecting to database")
  database.Connect(appConfig.Database)
  defer database.Close()

  log.Default().Println("Starting server")
  server.Start(appConfig.Server, router.Init())
}