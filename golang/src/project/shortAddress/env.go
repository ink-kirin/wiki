package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
	S Storage
}

func getEnv() *Env {
	// Load will read your env files 加载env文件
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	addr := os.Getenv("APP_REDIS_ADDR")
	if addr == "" {
		addr = "localhost"
	}
	port := os.Getenv("APP_REDIS_PORT")
	if port == "" {
		port = "6379"
	}
	passwd := os.Getenv("APP_REDIS_PASSWD")
	if passwd == "" {
		passwd = ""
	}
	dbS := os.Getenv("APP_REDIS_DB")
	if dbS == "" {
		dbS = "0"
	}
	// dbS 为字符串，将字符串转换为整型int
	db, err := strconv.Atoi(dbS)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("connect to redis (add: %s port: %s password: %s db: %d)", addr, port, passwd, db)
	r := NewRedisCli(addr+":"+port, passwd, db)
	return &Env{S: r}
}
