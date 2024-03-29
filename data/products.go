package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	SlackUsername   string    `json:"slack_name"`
	Day             string    `json:"current_day"`
	Time            time.Time `json:"utc_time"`
	Track           string    `json:"track"`
	Github_file_url string    `json:"github_file_url"`
	Github_repo_url string    `json:"github_repo_url"`
	Status_code     int       `json:"status_code"`
}

type Products Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func GetProducts() Products {
	return productList()
}

func productList() Products {
	currentTime := time.Now()
	formattedTimeStr := currentTime.Format("2006-01-02T15:04:05Z")
	layout := "2006-01-02T15:04:05Z"
	parsedTime, _ := time.Parse(layout, formattedTimeStr)

	return Products{
		SlackUsername:   "uchedingba",
		Day:             "Monday",
		Time:            parsedTime,
		Track:           "backend",
		Github_file_url: "https://github.com/Blaqollar/hng/blob/main/main.go",
		Github_repo_url: "https://github.com/Blaqollar/hng",
		Status_code:     200,
	}
}
