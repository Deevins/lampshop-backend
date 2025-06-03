package main

import (
	"context"
	"errors"
	"github.com/Deevins/lampshop-backend/product/internal/handler/product"
	"github.com/Deevins/lampshop-backend/product/internal/service"
	"github.com/Deevins/lampshop-backend/product/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var (
		ctx = context.Background()
	)
	logger.Init(true) // true = dev mode (читаемый вывод)

	//cfg, err := config.LoadConfig("../internal/config")
	//if err != nil {
	//	logger.Log.Fatalf("cannot load config: %s", err)
	//}
	//
	//if err := initConfig(); err != nil {
	//	log.Fatalf("can not read config file %s", err.Error())
	//}

	pool, err := pgxpool.New(ctx, "postgres://postgres:secret@product-db:5432/lampshop_products?sslmode=disable")
	//pool, err := pgxpool.New(ctx, "postgres://postgres:secret@localhost:5432/lampshop_products?sslmode=disable")
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}
	defer pool.Close()

	productSvc := service.NewProductService(pool)

	h := product.NewHandler(productSvc)
	router := h.InitRoutes()

	srv := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	go func() {
		logger.Log.Infof("Server running on %s", ":8081")
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Log.Fatalf("listen: %s", err)
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
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
