package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

type joke struct {
	Categories string `json:"categories"`
	CreatedAt  string `json:"created_at"`
	Icon_url   string `json:"icon_url"`
	ID         string `json:"id"`
	UpdatedAt  string `json:"updated_at"`
	Url        string `json:"url"`
	Value      string `json:"value"`
}

func main() {
	resp, err := http.Get("https://api.chucknorris.io/jokes/random")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	// without a struct
	// ****************
	var dat map[string]interface{}
	json.Unmarshal(bs, &dat)
	fmt.Println(dat["value"].(string))

	// with a known struct
	// **************
	joke1 := joke{}
	json.Unmarshal(bs, &joke1)
	fmt.Println(joke1.Categories)

	return len(bs), nil
}
