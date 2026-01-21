package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/jabella1/stock-ratings-service/internal/features/stock/app/query/handler"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/infra/db/postgres"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/interface/api/rest"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
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
	getListStockRatingQueryHandler := handler.CreateGetListStockRatingQueryHandler(context, stockRatingRepository)
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: getAllowedOrigins(),
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.Recover())

	rest.CreateStockRatingController(e, stockRatingByTickerQueryHandler, getListStockRatingQueryHandler)

	if err := e.Start(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func getAllowedOrigins() []string {
	var origins string
	if origins = os.Getenv("ALLOWED_ORIGINS"); origins == "" {
		log.Fatal("ALLOWED_ORIGINS is not set")
	}

	formattedOrigins := strings.Split(origins, ",")

	for i, origin := range formattedOrigins {
		formattedOrigins[i] = strings.TrimSpace(origin)
	}

	return formattedOrigins
}
