package main

import (
    "fmt"
    "net/http"
    "os"
    "github.com/gin-gonic/gin"
    "encoding/json"
    "io/ioutil"
)

type Stock struct {
    Symbol string  `json:"symbol"`
    Price  float64 `json:"price"`
    Change float64 `json:"change"`
}

var stocks []Stock

// Fetch Stock Data (This is a mock function. You can replace it with an actual API call)
func fetchStockData() []Stock {
    // Mock stock data
    return []Stock{
        {Symbol: "NVDA", Price: 134.54, Change: 1.24},
        {Symbol: "TSLA", Price: 243.20, Change: -0.53},
        {Symbol: "GME", Price: 20.63, Change: -0.39},
        {Symbol: "GERN", Price: 4.27, Change: 0.00},
    }
}

// Fetch stock data from a real API (Example using Alpha Vantage)
func fetchStockDataFromAPI(symbol string) (Stock, error) {
    apiKey := os.Getenv("ALPHA_VANTAGE_API_KEY")
    url := fmt.Sprintf("https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=%s&apikey=%s", symbol, apiKey)
    
    resp, err := http.Get(url)
    if err != nil {
        return Stock{}, err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return Stock{}, err
    }

    var result map[string]interface{}
    json.Unmarshal(body, &result)

    stockData := result["Global Quote"].(map[string]interface{})
    price, _ := strconv.ParseFloat(stockData["05. price"].(string), 64)
    change, _ := strconv.ParseFloat(stockData["10. change percent"].(string), 64)

    stock := Stock{
        Symbol: symbol,
        Price:  price,
        Change: change,
    }
    return stock, nil
}

func main() {
    router := gin.Default()

    // API endpoint for fetching stock data
    router.GET("/stocks", func(c *gin.Context) {
        stocks = fetchStockData() // Replace this with actual API call
        c.JSON(http.StatusOK, stocks)
    })

    // Example of a specific stock query from the API
    router.GET("/stocks/:symbol", func(c *gin.Context) {
        symbol := c.Param("symbol")
        stock, err := fetchStockDataFromAPI(symbol)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, stock)
    })

    // Start the server
    router.Run(":8080")
}
