package token

import (
	"time"
)

// TtyParameter kubectl tty param
type TtyParameter struct {
	Id        int
	Title     string
	Arg       string
	Sign      string
	UserAgent string
}

// Cache interface that defines token cache behavior
type Cache interface {
	Get(token string) *TtyParameter
	Delete(token string) error
	Add(token string, param *TtyParameter, d time.Duration) error
	GetValue(token string) string
	AddKeyValue(token string, value any, d time.Duration) error
}
