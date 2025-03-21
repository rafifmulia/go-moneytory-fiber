package mock_test

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"net/url"
	"reflect"
	"restfulapi/api"
	"restfulapi/conf"
	"restfulapi/dto"
	"restfulapi/helper"
	"restfulapi/model/domain"
	"restfulapi/repository"
	"restfulapi/router"
	"runtime/debug"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

var (
	fr       *fiber.App
	trxXRepo *repository.TransactionXRepositoryMock // seharusnya trxXRepo di pass via func args trxSvc, cuman kurang sreg, makanya dihilangkan. Sementara waktu kalau mau nyoba-nyoba mock, langsung aja ganti di file trxSvc nya.
)

func initResources() {
	conf.InitDbConnX() // Biar ga nil pointer dereference, soalnya bakal ada inisiasi db di service.
	fr = router.InitRouter()
	trxXRepo = repository.NewTransactionXRepositoryMock(&mock.Mock{})
}

// go test -v -count=1 -failfast -race -cpu=1 -run=^TestCreateTransaction$ ./
func TestCreateTransaction(t *testing.T) {
	initResources()
	defer errorHandler(t)
	dscs := "http mock testing create transaction"
	desc := url.QueryEscape(dscs)
	ds := "2025-03-12T07:24:00+07:00"
	date := url.QueryEscape(ds)
	body := strings.NewReader(fmt.Sprintf("description=%s&amount=1000&date=%s", desc, date))
	req := httptest.NewRequest("POST", "http://localhost:8080/v1/transaction", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-API-Key", "secret")
	req.Header.Set("Accept", "application/json")
	// ctx := req.Context() // Masih integration test.
	// dbx := driver.ExportDbHandleX() // Masih integration test.
	// tx, _ := dbx.Beginx() // Masih integration test.
	v7, err := uuid.NewV7()
	if err != nil {
		t.Fatal(err.Error())
	}
	mtm := time.Date(2025, time.March, 12, 8, 39, 23, 134888000, time.Local)
	trxDate, err := time.Parse(time.RFC3339, ds)
	if err != nil {
		t.Fatal(err.Error())
	}
	trx := &domain.Transaction{
		Uuid:        helper.InlinePointer(v7.String()),
		Description: helper.InlinePointer(dscs),
		Amount:      helper.InlinePointer[float64](1000),
		Ctm:         &mtm,
		Mtm:         &mtm,
		Date:        &sql.NullTime{Valid: true, Time: trxDate},
	}
	// trxXRepo.Mock.On("Save", ctx, tx, trx).Return(trx) // Masih integration test.
	resp, err := fr.Test(req, 200)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	rs := &api.RespDetailTransaction{}
	err = json.Unmarshal(b, rs)
	if err != nil {
		t.Error(err)
	}
	ex := &api.RespDetailTransaction{
		Meta: &api.Meta{
			Code:    201,
			Message: "Success create transaction",
		},
		Data: dto.DomainTrxToApiTrx(trx),
	}
	t.Logf("%s\n", b)
	assert.Equal(t, ex, rs)
}

func errorHandler(t *testing.T) {
	if msg := recover(); msg != nil {
		v := reflect.ValueOf(msg)
		tp := v.Type()
		t.Fatalf("type:%s\nval:%s\nstack trace:%s\n", tp.Name(), msg, debug.Stack())
	}
}
