package integration_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"net/url"
	"reflect"
	"restfulapi/conf"
	"restfulapi/router"
	"runtime/debug"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

var (
	fr *fiber.App
)

func initResources() {
	conf.InitDbConnX()
	fr = router.InitRouter()
}

// go test -v -count=1 -failfast -race -cpu=1 -run='^TestCreateTransaction$' ./
func TestCreateTransaction(t *testing.T) {
	initResources()
	defer errorHandler(t)
	desc := url.QueryEscape("http integration testing create transaction")
	date := url.QueryEscape("2025-03-12T07:24:00+07:00")
	body := strings.NewReader(fmt.Sprintf("description=%s&amount=1000&date=%s", desc, date))
	r := httptest.NewRequest("POST", "http://localhost:8080/v1/transaction", body)
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Set("X-API-Key", "secret")
	r.Header.Set("Accept", "application/json")
	// adaptor.FiberHandler(trxH)
	// fr.ServeHTTP(w, r)
	resp, err := fr.Test(r, 200)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	d := make(map[string]map[string]interface{})
	err = json.Unmarshal(b, &d)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, float64(201), (d["meta"]["code"]).(float64))
	t.Log(d["meta"])
}

func errorHandler(t *testing.T) {
	if msg := recover(); msg != nil {
		v := reflect.ValueOf(msg)
		tp := v.Type()
		t.Fatalf("type:%s\nval:%s\nstack trace:%s\n", tp.Name(), msg, debug.Stack())
	}
}
