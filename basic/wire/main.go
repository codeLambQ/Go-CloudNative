package main

import "fmt"

type Config struct {
	RedisIp   string
	RedisPort string
	MysqlIp   string
	MysqlPort string
}

type MysqlClient struct{}
type RedisClient struct{}

type App struct{}

func NewConfig() Config {
	return Config{}
}
func NewMysqlClient(config Config) MysqlClient {
	return MysqlClient{}
}
func NewRedisClient(config Config) RedisClient {
	return RedisClient{}
}

func NewApp(m MysqlClient, r RedisClient) App {
	return App{}
}

func main() {
	app := WireApp()
	fmt.Println(app)
}
