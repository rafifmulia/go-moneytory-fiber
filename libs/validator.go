package libs

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
	mu       *sync.Mutex = &sync.Mutex{}
)

func init() {
	mu.Lock()
	defer mu.Unlock()
	validate = validator.New(validator.WithRequiredStructEnabled())
}

func ExportValidator() *validator.Validate {
	mu.Lock()
	defer mu.Unlock()
	return validate
}
