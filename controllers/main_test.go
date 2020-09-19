package controllers

import (
	"github.com/felipe-michelon/url-shortener/database"
	"github.com/felipe-michelon/url-shortener/models"

	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	ts := httptest.NewServer(SetupRouter())
	defer ts.Close()

	req, _ := http.Get(ts.URL)
	body, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()

	assert.Equal(t, 200, req.StatusCode)
	assert.Equal(t, "Hello there", string(body))
}

func TestMain(m *testing.M) {
	_ = godotenv.Load("../.env")
	database.SetupDB()
	models.Migrate()

	exitVal := m.Run()

	os.Exit(exitVal)
}
