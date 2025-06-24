package service

import (
	"encoding/json"
	"go-hex-temp/internal/infrastructure/logx"
	"go-hex-temp/internal/ports/output"
	"time"
)

type callbackOne[T any] func() (*T, error)
type callbackMany[T any] func() ([]T, error)

type typedCache[T any] struct {
	inner output.Cache
}

func (tc *typedCache[T]) ReadOne(key string, ttl time.Duration, cb callbackOne[T]) (*T, error) {
	if raw, ok := tc.inner.Load(key); ok {
		var val T
		if err := json.Unmarshal(raw, &val); err == nil {
			logx.Debug("read from cache")
			return &val, nil
		}
	}

	val, err := cb()
	if err != nil {
		return nil, err
	}

	if err := tc.inner.Save(key, val, ttl); err != nil {
		return nil, err
	}
	logx.Debug("read from callback")
	return val, nil
}

func (tc *typedCache[T]) ReadMany(key string, ttl time.Duration, cb callbackMany[T]) ([]T, error) {
	if raw, ok := tc.inner.Load(key); ok {
		var val []T
		if err := json.Unmarshal(raw, &val); err == nil {
			return val, nil
		}
		// Optionally: log or handle unmarshal error here
	}

	val, err := cb()
	if err != nil {
		return nil, err
	}

	if err := tc.inner.Save(key, val, ttl); err != nil {
		return nil, err
	}
	return val, nil
}
