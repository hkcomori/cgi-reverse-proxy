package application

import (
	"errors"
	"os"
	"strconv"
)

type Config struct {
	UpstreamAddress string
	UpstreamTimeout int
}

type castFunc[T any] func(s string) (T, error)

func NewConfig() (*Config, error) {
	config := &Config{}

	paramsMapping := []parameters{
		// Required
		newRequired(&config.UpstreamAddress, "GROXY_UPSTREAM_ADDRESS", castString),
		// Optional
		newOptional(&config.UpstreamTimeout, "GROXY_UPSTREAM_TIMEOUT", 180, castInt),
	}
	for _, p := range paramsMapping {
		if err := p.load(); err != nil {
			return nil, err
		}
	}

	return config, nil
}

type parameters interface {
	load() error
}

type required[T any] struct {
	field *T
	key string
	castFunc castFunc[T]
}

func newRequired[T any](field *T, key string, castFunc castFunc[T]) *required[T] {
	return &required[T]{
		field: field,
		key: key,
		castFunc: castFunc,
	}
}

func (p *required[T]) load() error {
	env := os.Getenv(p.key)
	if env == "" {
		return errors.New(p.key + " is required but not got")
	}
	castedValue, err := p.castFunc(env)
	*p.field = castedValue
	return err
}

type optional[T any] struct {
	field *T
	key string
	defaultValue T
	castFunc castFunc[T]
}

func newOptional[T any](field *T, key string, defaultValue T, castFunc castFunc[T]) *optional[T] {
	return &optional[T]{
		field: field,
		key: key,
		defaultValue: defaultValue,
		castFunc: castFunc,
	}
}

func (p *optional[T]) load() error {
	env := os.Getenv(p.key)
	if env == "" {
		*p.field = p.defaultValue
		return nil
	}
	castedValue, err := p.castFunc(env)
	*p.field = castedValue
	return err
}

func castString(s string) (string, error) {
	return s, nil
}

func castInt(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return n, nil
}
