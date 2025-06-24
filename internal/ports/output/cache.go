package output

import (
	"time"
)

type Cache interface {
	Save(key string, data any, ttl time.Duration) error
	Load(key string) ([]byte, bool)
	Del(key string) error
}
