package main

import (
	"context"
	"log"
	"os"

	"github.com/jabella1/stock-ratings-service/internal/features/stock/app/query/handler"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/infra/db/postgres"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/interface/api/rest"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	connectionString := os.Getenv("CONNECTION_STRING")
	if connectionString == "" {
		log.Fatal("CONNECTION_STRING is not set")
	}
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = ":8080"
	}
	context := context.Background()
	connection, err := postgres.NewConnection(context, connectionString)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer connection.Close(context)
	queries := postgres.NewQueries(connection)
	stockRatingRepository := postgres.CreateSqlcStockRatingRepository(queries)
	stockRatingByTickerQueryHandler := handler.CreateGetStockRatingByTickerQueryHandler(stockRatingRepository)

	e := echo.New()
	rest.CreateStockRatingController(e, stockRatingByTickerQueryHandler)

	if err := e.Start(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
