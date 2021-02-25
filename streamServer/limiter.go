package streamServer

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/**
limit connect number
 */
type ConnLimiter struct {
	maxinum int
	token chan int
}

func GetToken(c *gin.Context){
	if limiter == nil {
		InitializeConnLimiter(1000)
	}
	if !limiter.getToken() {
		c.String(http.StatusBadRequest, "server busy")
		fmt.Println("server busy")
		c.Abort()
		return
	}
	c.Next()
	limiter.releaseToken()
}

func InitializeConnLimiter(limitNum int) {
	limiter = &ConnLimiter{
		maxinum: limitNum,
		token: make(chan int, limitNum),
	}
}

var limiter *ConnLimiter

func (cl *ConnLimiter) getToken() bool {
	if len(cl.token) >= cl.maxinum {
		return false
	}
	cl.token <- 1
	return true
}

func (cl *ConnLimiter) releaseToken() {
	<-cl.token
}


func TestLimiter(c *gin.Context){
	time.Sleep(time.Second * 10)
	c.String(http.StatusOK, "ok")
}

 