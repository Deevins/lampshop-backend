package main

import (
	"context"
	"errors"
	"github.com/Deevins/lampshop-backend/product/internal/handler/order"
	"github.com/Deevins/lampshop-backend/product/pkg/logger"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger.Init(true) // true = dev mode (читаемый вывод)

	//cfg, err := config.LoadConfig("../internal/config")
	//if err != nil {
	//	logger.Log.Fatalf("cannot load config: %s", err)
	//}

	if err := initConfig(); err != nil {
		log.Fatalf("can not read config file %s", err.Error())
	}

	h := order.NewHandler()
	router := h.InitRoutes()

	srv := &http.Server{
		Addr:    viper.GetString("http_server.port"),
		Handler: router,
	}

	go func() {
		logger.Log.Infof("Server running on %s", viper.GetString("http_server.port"))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Log.Fatalf("listen: %s", err)
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Fatalf("Server forced to shutdown: %s", err)
	}

	logger.Log.Info("Server exited")

	// TODO: if have time - implement GRPC
	//srv := server.NewGRPC("8081")
	//productAPI := grpc2.NewProductAPI()
	//if err := srv.Start(
	//	func(s *grpc.Server) {
	//		productv1.RegisterProductServiceServer(s, productAPI)
	//	}); err != nil {
	//	log.Fatalf("gRPC server failed: %v", err)
	//}
}

func initConfig() error {
	viper.SetConfigType("yml")
	viper.AddConfigPath("internal/config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
