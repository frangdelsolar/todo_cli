package currency

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

// The SourceType enum defines the source of the currency exchange rate.
type SourceType string

const (
	Blue     SourceType = "Blue"
	Official SourceType = "Oficial"
)

type Source struct {
	ValueBuy  float64 `json:"buy"`
	ValueSell float64 `json:"sell"`
}

type Rate struct {
	Official Source `json:"oficial"`
	Blue     Source `json:"blue"`
}

type RateMap map[string]Rate

const BASE_URL = "https://api.bluelytics.com.ar/v2/"
const RATES_FILE = "rates.json"

// NewRate creates a new instance of the Rate struct with default values for the Official and Blue sources.
//
// Returns:
// - Rate: a new instance of the Rate struct with the Official and Blue sources initialized to zero values.
func NewRate() Rate {
	return Rate{
		Official: Source{
			ValueBuy:  0,
			ValueSell: 0,
		},
		Blue: Source{
			ValueBuy:  0,
			ValueSell: 0,
		},
	}
}

// GetBlueAverage calculates the average value of the Blue Source in the Rate struct.
//
// It returns the average of the buy and sell values of the Blue Source.
// The return type is float64.
func (r *Rate) GetBlueAverage() float64 {
	return (r.Blue.ValueBuy + r.Blue.ValueSell) / 2
}

// SetOfficial sets the buy and sell values of the Official Source in the Rate struct.
//
// Parameters:
// - buy: a float64 representing the buy value.
// - sell: a float64 representing the sell value.
//
// Return type: None.
func (r *Rate) SetOfficial(buy, sell float64) {
	r.Official.ValueBuy = buy
	r.Official.ValueSell = sell
}

// SetBlue sets the buy and sell values of the Blue Source in the Rate struct.
//
// Parameters:
// - buy: a float64 representing the buy value.
// - sell: a float64 representing the sell value.
//
// Return type: None.
func (r *Rate) SetBlue(buy, sell float64) {
	r.Blue.ValueBuy = buy
	r.Blue.ValueSell = sell
}

// DownloadRates retrieves the currency exchange rates from the specified URL and saves them to a file.
//
// It sends an HTTP GET request to the specified URL, and if the response status code is not 200,
// it returns an error. It then opens the rates file for writing (creating it if it doesn't exist),
// and writes the downloaded data to the file. Finally, it logs the success message and returns nil.
//
// Returns:
// - error: an error if there was an issue downloading or writing the rates.
func DownloadRates() error {
	url := BASE_URL + "evolution.csv"

	log.Info().Msgf("Getting Rates from URL: %s", url)

	resp, err := http.Get(url)
	if err != nil {
		log.Err(err).Msg("Error getting rates")
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Warn().Msgf("Error getting rates. Status code: %d", resp.StatusCode)
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	var ratesMap RateMap = make(map[string]Rate)

	reader := csv.NewReader(resp.Body)

	// Skip header
	_, err = reader.Read()
	if err != nil {
		return fmt.Errorf("error reading rates: %w", err)
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error reading rates: %w", err)
		}

		date := record[0]
		var rateItem Rate

		if _, ok := ratesMap[date]; !ok {
			rateItem = NewRate()
		} else {
			rateItem = ratesMap[date] // Use existing Rate struct
		}

		source := record[1]

		var officialBuy, officialSell, blueBuy, blueSell float64
		if source == string(Official) {
			officialBuy, _ = strconv.ParseFloat(record[2], 64)
			officialSell, _ = strconv.ParseFloat(record[3], 64)
			rateItem.SetOfficial(officialBuy, officialSell)
		} else if source == string(Blue) {
			blueBuy, _ = strconv.ParseFloat(record[2], 64)
			blueSell, _ = strconv.ParseFloat(record[3], 64)
			rateItem.SetBlue(blueBuy, blueSell)
		}

		ratesMap[date] = rateItem
	}

	// Convert ratesMap to JSON
	ratesJson, err := json.MarshalIndent(ratesMap, "", "    ")
	if err != nil {
		return fmt.Errorf("error converting rates to JSON: %w", err)
	}

	// Write rates to file
	err = os.WriteFile(RATES_FILE, ratesJson, 0644)
	if err != nil {
		return fmt.Errorf("error writing rates to file: %w", err)
	}

	log.Info().Interface("file", RATES_FILE).Msg("Downloaded rates and wrote to file")
	return nil
}

// GetRatesByDate retrieves the exchange rate for a given date.
//
// Parameters:
// - date: the date for which the exchange rate is requested (time.Time).
//
// Returns:
// - Rate: the exchange rate for the given date (models.Rate).
// - error: an error if the exchange rate cannot be retrieved (error).
func GetRatesByDate(date time.Time) (Rate, error) {
	refreshFile := false

	// Check file exists
	info, err := os.Stat(RATES_FILE)
	if err != nil {
		log.Warn().Msgf("File does not exist: %s", RATES_FILE)
		refreshFile = true
	} else {
		// Check file is not too old
		if time.Since(info.ModTime()) > 24*time.Hour {
			log.Warn().Msgf("File is older than 24 hours: %s", RATES_FILE)
			refreshFile = true
		}
	}

	// Download rates
	if refreshFile {
		err := DownloadRates()
		if err != nil {
			return Rate{}, err
		}
	}

	// Read rates
	file, err := os.ReadFile(RATES_FILE)
	if err != nil {
		log.Err(err).Msgf("Error reading file %s", RATES_FILE)
		return Rate{}, err
	}

	// Parse rates
	var ratesMap RateMap
	err = json.Unmarshal(file, &ratesMap)
	if err != nil {
		log.Err(err).Msgf("Error parsing rates from file %s", RATES_FILE)
		return Rate{}, err
	}

	// Find rate for date
	minDate := time.Date(2011, 1, 1, 0, 0, 0, 0, time.UTC)

	for {
		if date.Before(minDate) {
			log.Warn().Msgf("No rate found for date: %s", date.Format("2006-01-02"))
			return Rate{}, fmt.Errorf("no rate found for date: %s", date.Format("2006-01-02"))
		}
		if rate, ok := ratesMap[date.Format("2006-01-02")]; ok {
			return rate, nil
		}
		date = date.AddDate(0, 0, -1)
	}
}
