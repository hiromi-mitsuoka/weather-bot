package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Weather struct {
	Area     string `json:"targetArea"`
	HeadLine string `json:"headlineText"`
	Body     string `json:"text"`
}

// To test on console with package main
// func main() {
// 	fmt.Print(GetWeather())
// }

func GetWeather() (str string, err error) {
	body, err := httpGetBody("https://www.jma.go.jp/bosai/forecast/data/overview_forecast/130000.json")
	if err != nil {
		return str, err
	}

	weather, err := formatWeather(body)
	if err != nil {
		return str, err
	}

	result := weather.ToS()

	return result, nil
}

func httpGetBody(url string) ([]byte, error) {
	// NOTE: Post the HTTP request and receive a response (https://pkg.go.dev/net/http)
	resp, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf("Get Http Error: %s", err)
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("IO Read Error: %s", err)
		return nil, err
	}

	// NOTE: The client must close the response body when finished with it: (https://qiita.com/stk0724/items/dc400dccd29a4b3d6471)
	defer resp.Body.Close()

	return body, nil
}

// NOTE: Return pointer
func formatWeather(body []byte) (*Weather, error) {
	// NOTE: Create pointer
	weather := new(Weather)
	if err := json.Unmarshal(body, weather); err != nil {
		err = fmt.Errorf("JSON Unmarshal error: %s", err)
		return nil, err
	}

	return weather, nil
}

// Method Definition
func (w *Weather) ToS() string {
	area := fmt.Sprintf("%sの天気です。\n", w.Area)
	head := fmt.Sprintf("%s\n", w.HeadLine)
	body := fmt.Sprintf("%s\n", w.Body)
	result := area + head + body

	return result
}
