package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jabella1/stock-ratings-service/internal/features/common/utils"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/app/command"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/app/command/handler"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/infra/db/cockroach"
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
	connection, err := cockroach.NewConnection(context, connectionString)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer connection.Close(context)
	queries := cockroach.NewQueries(connection)
	var next_page *string
	baseUrlEnv := "CONNECTION_KARENAI_BASEURL"
	baseUrl, err := utils.ValidateEmptyString(os.Getenv(baseUrlEnv), baseUrlEnv)
	if err != nil {
		log.Fatalf("Error de configuración: %v", err)
	}
	endpointSwechallengeEnv := "CONNECTION_KARENAI_ENDPOINTS_SWECHALLENGE"
	endpointsSwechallenge, err := utils.ValidateEmptyString(os.Getenv(endpointSwechallengeEnv), endpointSwechallengeEnv)
	if err != nil {
		log.Fatalf("Error de configuración: %v", err)
	}
	url := baseUrl + endpointsSwechallenge
	var stockRatingRepository = cockroach.CreateSqlcStockRatingRepository(queries)
	var unitOfWork = cockroach.CreateUnitOfWork(connection)
	var saveStockRatingCommandHandler = handler.CreateSaveStockRatingCommandHandler(stockRatingRepository, unitOfWork)
	connectionKarenToken, err := utils.ValidateEmptyString(os.Getenv("CONNECTION_KARENAI_TOKEN"), "CONNECTION_KARENAI_TOKEN")
	if err != nil {
		log.Fatalf("Error de configuración: %v", err)
	}
	yahooFinanceBaseURL, err := utils.ValidateEmptyString(os.Getenv("CONNECTION_YAHOO_FINANCE_BASEURL"), "YAHOO_FINANCE_BASEURL")
	if err != nil {
		log.Fatalf("Error de configuración: %v", err)
	}
	interval, err := utils.ValidateEmptyString(os.Getenv("CONNECTION_YAHOO_FINANCE_INTERVAL"), "YAHOO_FINANCE_INTERVAL")
	if err != nil {
		log.Fatalf("Error de configuración: %v", err)
	}
	rangeVal, err := utils.ValidateEmptyString(os.Getenv("CONNECTION_YAHOO_FINANCE_RANGE"), "YAHOO_FINANCE_RANGE")
	if err != nil {
		log.Fatalf("Error de configuración: %v", err)
	}
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
		request.Header.Set("Authorization", "Bearer "+connectionKarenToken)
		response, err := http.DefaultClient.Do(request)
		if err != nil {
			log.Fatal("Failed to perform request:", err)
		}

		var apiResponse Response
		if err := json.NewDecoder(response.Body).Decode(&apiResponse); err != nil {
			log.Fatal("Failed to decode response:", err)
		}
		response.Body.Close()

		for _, item := range apiResponse.Items {
			targetFromNumeric := utils.NumericFromString(item.TargetFrom)
			targetToNumeric := utils.NumericFromString(item.TargetTo)
			targetFromMapped := utils.Float64FromNumeric(targetFromNumeric)
			targetToMapped := utils.Float64FromNumeric(targetToNumeric)

			currentPrice := GetOrFetchCurrentPrice(item.Ticker, yahooFinanceBaseURL, interval, rangeVal)
			saveStockRatingCommandHandler.SaveStockRating(context, &command.SaveStockRatingCommand{
				Ticker:       item.Ticker,
				Company:      item.Company,
				Brokerage:    &item.Brokerage.String,
				Action:       item.Action.String,
				RatingFrom:   item.RatingFrom.String,
				RatingTo:     item.RatingTo.String,
				TargetFrom:   *targetFromMapped,
				TargetTo:     *targetToMapped,
				CurrentPrice: currentPrice,
			})

		}

		if apiResponse.NextPage == nil || *apiResponse.NextPage == "" {
			break
		}
		next_page = apiResponse.NextPage
	}
}

func GetOrFetchCurrentPrice(ticker string, yahooFinanceBaseURL, interval, rangeVal string) float64 {
	price, err := getCurrentPrice(ticker, yahooFinanceBaseURL, interval, rangeVal)
	if err != nil {
		log.Printf("Warning: using fallback price for %s", ticker)
		return 0
	} else {
		log.Printf("Current price for %s: %.4f", ticker, price)
	}
	return price
}

func getCurrentPrice(ticker string, yahooFinanceBaseURL, interval, rangeVal string) (float64, error) {
	url := fmt.Sprintf(
		"%s/%s?interval=%s&range=%s",
		yahooFinanceBaseURL,
		ticker,
		interval,
		rangeVal,
	)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	var result struct {
		Chart struct {
			Result []struct {
				Meta struct {
					RegularMarketPrice float64 `json:"regularMarketPrice"`
				} `json:"meta"`
			} `json:"result"`
		} `json:"chart"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, fmt.Errorf("error decoding yahoo response: %w", err)
	}

	if len(result.Chart.Result) == 0 {
		return 0, fmt.Errorf("no hay datos para %s", ticker)
	}

	return result.Chart.Result[0].Meta.RegularMarketPrice, nil
}
