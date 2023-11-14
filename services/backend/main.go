package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"log"
	"net/http"
	"time"
)

func setupRouter() *gin.Engine {

	r := gin.Default()
	config := cors.DefaultConfig()
	log.Println(config)
	config.AllowOrigins = []string{"http://localhost:3000"}
	r.Use(cors.New(config))

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/handler-initial-data", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"text": "initial"})
	})

	r.GET("/handler", func(c *gin.Context) {
		// todo make sure CORS is ok
		c.JSON(http.StatusOK, gin.H{"text": "updated"})
	})

	r.GET("/socket", func(c *gin.Context) {
		conn, _, _, err := ws.UpgradeHTTP(c.Request, c.Writer)
		if err != nil {
			log.Println("error with WebSOcket: ", err)
			c.Writer.WriteHeader(http.StatusMethodNotAllowed)
			// todo not sure is ok
			return
		}
		go func() {
			defer conn.Close()
			time.Sleep(time.Second)
			err = wsutil.WriteServerMessage(conn, ws.OpText, []byte(`{"text": "from-websocket-in-gin"}`))
			if err != nil {
				log.Println("error writing WebSocket data: ", err)
				return
			}
		}()
		return
	})

	// that for custom 404
	// r.NoRoute()
	log.Println("Server is available at http://localhost:8000")
	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8000
	r.Run(":8000")
}
