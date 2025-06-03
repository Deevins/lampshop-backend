package main

import (
	"context"
	"errors"
	"github.com/Deevins/lampshop-backend/order/internal/handler/order"
	"github.com/Deevins/lampshop-backend/order/internal/service"
	"github.com/Deevins/lampshop-backend/order/pkg/logger"
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
	logger.Init(true)

	pool, err := pgxpool.New(ctx, "postgres://postgres:secret@order-db:5432/lampshop_orders?sslmode=disable")
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}
	defer pool.Close()

	orderSvc := service.NewOrderService(pool)

	h := order.NewHandler(orderSvc)
	router := h.InitRoutes()

	srv := &http.Server{
		Addr:    ":8082",
		Handler: router,
	}

	go func() {
		logger.Log.Infof("Server running on %s", ":8082")
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
