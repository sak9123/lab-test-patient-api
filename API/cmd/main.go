package main

import (
	"context"
	"hospitalApi/cmd/config"
	"hospitalApi/pkg/api"
	"hospitalApi/pkg/entity"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	log.Println("starting...")
	initTimeZone()

	dbSqlServer := initDB()

	apiMux := api.APIMux(api.APIConfig{
		DB: dbSqlServer,
	})

	sqlDb, _ := dbSqlServer.DB()
	defer sqlDb.Close()

	srv := &http.Server{
		Addr:         "0.0.0.0:5000",
		WriteTimeout: time.Minute * 10,
		ReadTimeout:  time.Minute * 10,
		IdleTimeout:  time.Minute * 10,
		Handler:      apiMux,
	}

	serverError := make(chan error, 1)

	go func() {
		serverError <- srv.ListenAndServe()
	}()

	log.Println("The service is ready to listen and serve.")

	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverError:
		log.Fatal(err)
	case sig := <-gracefulStop:
		log.Println("shutdown", "status", "shutdown started", "signal", sig)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		_ = srv.Shutdown(ctx)
		defer cancel()
	}

	log.Println("The service is shutting down...")

	log.Println("terminated...")

	os.Exit(0)
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}

func setGormUTCTime(input *gorm.DB) *gorm.DB {
	return input.Session(&gorm.Session{NowFunc: func() time.Time { return time.Now().UTC() }})
}

func initDB() *gorm.DB {
	configuration := config.New()
	dbSqlServer, err := gorm.Open(postgres.Open(configuration.ConnectionString), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
	})
	if err != nil {
		log.Println("can't connect to db..")
		return nil
	}

	sqlDb, err := dbSqlServer.DB()
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	sqlDb.SetConnMaxLifetime(15 * time.Minute)
	log.Println("db connected successfully..")

	dbSqlServer = setGormUTCTime(dbSqlServer)
	entity.Migration(dbSqlServer)

	return dbSqlServer
}
