package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redis_rate/v9"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"rate-limiting-gin-go/controllers"
	_ "rate-limiting-gin-go/docs"
	"rate-limiting-gin-go/middlewares"
	"time"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "DbWV0cfe",
	})
	_ = rdb.FlushDB(ctx).Err()

	limiter := redis_rate.NewLimiter(rdb)

	userController, err := controllers.NewUserController()
	if err != nil {
		panic(err)
	}
	rateMiddleware, err := middlewares.NewRateMiddleware(limiter)
	if err != nil {
		panic(err)
	}
	engine := gin.Default()
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	userRouter := engine.Group("user")
	userRouter.Use(rateMiddleware.Limit(10, time.Minute)).GET("/:id", userController.FindUser)
	err = engine.Run()
	if err != nil {
		panic(err)
	}
}
