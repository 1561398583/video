package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"time"
	"yx.com/videos/streamServer"
	"yx.com/videos/utils"
)


func main()  {
	r := gin.New()

	logf, err := os.Create(LOG_DIR + "gin.log")
	if err != nil {
		log.Fatal(err)
	}

	pdw := utils.NewPdWriter(LOG_DIR + "pdlog/")

	//log write to 2 place
	logWriter := io.MultiWriter(logf, pdw)

	//gin packed all properties in LogFormatterParams,
	//then give you LogFormatterParams,
	//then you  can decide how to organize these properties to a string,
	//return the string
	logFormatterFunc := func(params gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC1123),
			params.Method,
			params.Path,
			params.Request.Proto,
			params.StatusCode,
			params.Latency,
			params.Request.UserAgent(),
			params.ErrorMessage,
		)
	}
	loggerConfig := gin.LoggerConfig{
		Formatter: logFormatterFunc,
		Output: logWriter,
	}
	logger := gin.LoggerWithConfig(loggerConfig)
	r.Use(logger)

	frc, err := os.Create(LOG_DIR + "recover.log")
	if err != nil {
		log.Fatal(err)
	}
	r.Use(gin.RecoveryWithWriter(frc))

	streamServer.InitializeConnLimiter(1)
	r.Use(streamServer.GetToken)
	r.GET("/testLimiter", streamServer.TestLimiter)

	//test connect
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message" : "pong",
		})
	})

	r.GET("/panic", func(c *gin.Context) {
		// panic with a string -- the custom middleware could save this to a database or report it to the user
		panic("foo")
	})

	/**
	upload files
	 */
	r.MaxMultipartMemory = 8 << 24  //128M
	//set directory of HTML
	r.LoadHTMLGlob("/home/yx/Videos/web/html/*")
	r.GET("/uploadFilePage", func(context *gin.Context) {
		context.HTML(http.StatusOK, "uploadFilePage.tmpl", gin.H{
			"title": "Main website",
		})
	})
	r.POST("/uploadFile", gin.HandlerFunc(utils.UploadFile))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

