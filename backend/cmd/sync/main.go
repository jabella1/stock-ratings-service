package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/jabella1/stock-ratings-service/internal/features/common/utils"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/infra/db/postgres"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/infra/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/joho/godotenv"
)

type StockRating struct {
	Ticker     string      `json:"ticker"`
	Company    string      `json:"company"`
	Brokerage  pgtype.Text `json:"brokerage"`
	Action     pgtype.Text `json:"action"`
	RatingFrom pgtype.Text `json:"rating_from"`
	RatingTo   pgtype.Text `json:"rating_to"`
	TargetFrom string      `json:"target_from"`
	TargetTo   string      `json:"target_to"`
}

type Response struct {
	Items    []StockRating `json:"items"`
	NextPage *string       `json:"next_page"`
}

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("No .env file found")
	}

	context := context.Background()
	connectionString := os.Getenv("CONNECTION_STRING")
	if connectionString == "" {
		log.Fatal("CONNECTION_STRING is not set")
		return
	}
	connection, err := postgres.NewConnection(context, connectionString)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer connection.Close(context)
	queries := postgres.NewQueries(connection)
	var next_page *string
	baseUrlEnv := "CONNECTION_KARENAI_BASEURL"
	baseUrl := validateEmptyString(os.Getenv(baseUrlEnv), baseUrlEnv)
	endpointSwechallengeEnv := "CONNECTION_KARENAI_ENDPOINTS_SWECHALLENGE"
	endpointsSwechallenge := validateEmptyString(os.Getenv(endpointSwechallengeEnv), endpointSwechallengeEnv)
	url := baseUrl + endpointsSwechallenge
	for {
		var requestURL string
		if next_page != nil {
			requestURL = url + "?next_page=" + *next_page
		} else {
			requestURL = url
		}

		request, err := http.NewRequest("GET", requestURL, nil)
		if err != nil {
			log.Fatal("Failed to create request:", err)
		}
		request.Header.Set("Authorization", "Bearer "+validateEmptyString(os.Getenv("CONNECTION_KARENAI_TOKEN"), "CONNECTION_KARENAI_TOKEN"))
		response, err := http.DefaultClient.Do(request)
		if err != nil {
			log.Fatal("Failed to perform request:", err)
		}
		defer response.Body.Close()
		var apiResponse Response
		if err := json.NewDecoder(response.Body).Decode(&apiResponse); err != nil {
			log.Fatal("Failed to decode response:", err)
		}

		for _, item := range apiResponse.Items {
			err := queries.UpsertStockRating(context, sqlc.UpsertStockRatingParams{
				Ticker:     item.Ticker,
				Company:    item.Company,
				Brokerage:  item.Brokerage,
				Action:     item.Action,
				RatingFrom: item.RatingFrom,
				RatingTo:   item.RatingTo,
				TargetFrom: utils.NumericFromString(item.TargetFrom),
				TargetTo:   utils.NumericFromString(item.TargetTo),
			})

			if err != nil {
				log.Printf("Error con %s: %v", item.Ticker, err)
			}
		}

		if apiResponse.NextPage == nil || *apiResponse.NextPage == "" {
			break
		}
		next_page = apiResponse.NextPage
	}
}

func validateEmptyString(stringToValidate string, fieldName string) string {
	if stringToValidate == "" {
		log.Fatal(fieldName + " is required")
	}
	return stringToValidate
}
