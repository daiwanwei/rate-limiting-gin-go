package middlewares

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v9"
	"time"
)

type RateMiddleware interface {
	Limit(rate int, period time.Duration) gin.HandlerFunc
}

type rateMiddleware struct {
	limiter *redis_rate.Limiter
}

func NewRateMiddleware(limiter *redis_rate.Limiter) (middleware RateMiddleware, err error) {
	return &rateMiddleware{limiter: limiter}, nil
}

func (middleware *rateMiddleware) Limit(rate int, period time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.FullPath()
		limit := redis_rate.Limit{
			Rate:   rate,
			Burst:  rate,
			Period: period,
		}
		res, err := middleware.limiter.Allow(context.Background(), key, limit)
		if err != nil {
			c.AbortWithStatus(500)
			fmt.Println(err)
			return
		}
		fmt.Printf("path(%s):allowed(%d),remaining(%d)\n", key, res.Allowed, res.Remaining)
		if res.Allowed == 0 {
			c.AbortWithStatus(401)
			fmt.Printf("path(%s):over rate limit!\n", key)
			return
		}
		c.Next()
	}
}
