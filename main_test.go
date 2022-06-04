package main

import (
	"github.com/michelonfelipe/url-shortener/controllers"
	"github.com/michelonfelipe/url-shortener/database"
	"github.com/michelonfelipe/url-shortener/models"

	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestGetExistingUrl(t *testing.T) {
	url := models.Url{Shortened: "shortened", Original: "http://original.com"}

	database.DB.Create(&url)
	defer database.DB.Unscoped().Delete(&url)

	ts := httptest.NewServer(controllers.SetupRouter())
	defer ts.Close()

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}}

	req, _ := client.Get(fmt.Sprintf("%s/%s", ts.URL, url.Shortened))

	assert.Equal(t, 301, req.StatusCode)
	assert.Equal(t, req.Header.Get("Location"), url.Original)
}

func TestGetNonExistingUrl(t *testing.T) {
	ts := httptest.NewServer(controllers.SetupRouter())
	defer ts.Close()

	req, _ := http.Get(fmt.Sprintf("%s/shortened", ts.URL))
	body, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()

	assert.Equal(t, 404, req.StatusCode)
	assert.Equal(t, "Url not found", string(body))
}

func TestCreateUrl(t *testing.T) {
	ts := httptest.NewServer(controllers.SetupRouter())
	defer ts.Close()

	values := map[string]string{"original": "https://original.com"}
	jsonValue, _ := json.Marshal(values)
	defer database.DB.Where("original = ?", values["original"]).Unscoped().Delete(&models.Url{})

	req, _ := http.Post(
		fmt.Sprintf("%s/urls", ts.URL),
		"application/json",
		bytes.NewBuffer(jsonValue),
	)
	body, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()

	assert.Equal(t, 201, req.StatusCode)
	assert.True(t, strings.Contains(string(body), "ID"))
}

func TestCreateInvalidUrl(t *testing.T) {
	ts := httptest.NewServer(controllers.SetupRouter())
	defer ts.Close()

	values := map[string]int{"original": 2}
	jsonValue, _ := json.Marshal(values)

	req, _ := http.Post(
		fmt.Sprintf("%s/urls", ts.URL),
		"application/json",
		bytes.NewBuffer(jsonValue),
	)

	assert.Equal(t, 400, req.StatusCode)
}

func TestHome(t *testing.T) {
	ts := httptest.NewServer(controllers.SetupRouter())
	defer ts.Close()

	req, _ := http.Get(ts.URL)
	body, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()

	assert.Equal(t, 200, req.StatusCode)
	assert.True(t, strings.Contains(string(body), "Url shortener"))
}

func TestMain(m *testing.M) {
	_ = godotenv.Load(".env")
	database.SetupDB()
	models.Migrate()

	exitVal := m.Run()

	os.Exit(exitVal)
}
