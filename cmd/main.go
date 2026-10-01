package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/fajarworks/backend-eventhub/internal/config"
	"github.com/fajarworks/backend-eventhub/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title           			Eventhub
// @version         			1.0
// @description     			Aplication for eventhub API
// @host      					localhost:8081

// @BasePath  					/

// @securityDefinitions.apikey 	BearerToken
// @in 							header
// @name 						Authorization
// @description 				Bearer token used as identiity for accessing backend

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println(err.Error())
	}

	pdb := config.NewPsqlDb(os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	pool, err := pdb.Conn()
	if err != nil {
		log.Println("can't connect to database :", err.Error())
	}
	defer pool.Close()

	if err := pool.Ping(context.Background()); err != nil {
		log.Println("database is not ready\n:", err.Error())
	}
	rc := config.NewRedis(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_USERNAME"), os.Getenv("REDIS_PASSWORD"), os.Getenv("REDIS_PORT"))
	rdb := rc.Connect()
	defer rdb.Close()
	status := rdb.Ping(context.Background())

	res, err := status.Result()
	if err != nil {
		fmt.Println("gagal connect redis : ", err)

	}
	fmt.Println(res)
	fmt.Println("berhasil connect redis")

	r := gin.Default()

	router.MainRouter(r, pool, rdb)
	fmt.Println("connected to data base")
	r.Run(fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("PORT")))
}
