package ds

import (
	"log"
	"os"
	"os/signal"
	"time"
	"context"
	"net/http"
	"gin-lnp-api/app"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	r *gin.Engine
}

func NewHttpServer(r *gin.Engine) *HttpServer {
	return &HttpServer{
		r: r,
	}
}

func (httpSrv *HttpServer) Start() {
	srv := &http.Server{
		Addr:    app.Config.HttpServerPort,
		Handler: httpSrv.r,
	}

	go func() {
		// service connections
		log.Println("Start Http Server ...")

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("http listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Http Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Http Server Shutdown:", err)
	}

	log.Println("Http Server exiting")
}