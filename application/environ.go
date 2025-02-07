package application

import (
	"errors"
	"os"
	"strings"
)

// Environ provides an interface to access environment variables
type Environ map[string]string

func NewEnviron() *Environ {
	return NewEnvironFromStrings(os.Environ())
}

func NewEnvironFromStrings(s []string) *Environ {
	env := make(Environ)
	for _, keyValue := range s {
		parts := strings.SplitN(keyValue, "=", 2)
		if len(parts) == 2 {
			env[parts[0]] = parts[1]
		}
	}
	return &env
}

func (e *Environ) ToStrings() []string {
	idx := 0
	keyValue := make([]string, len(*e))
	for key, value := range *e {
		keyValue[idx] = key + "=" + value
		idx++
	}
	return keyValue
}

func (e *Environ) HasAll(keys ...string) bool {
	for _, key := range keys {
		if _, exists := (*e)[key]; !exists {
			return false
		}
	}
	return true
}

func (e *Environ) HasAny(keys ...string) bool {
	for _, key := range keys {
		if _, exists := (*e)[key]; exists {
			return true
		}
	}
	return false
}

func (e *Environ) GetAny(keys ...string) (string, error) {
	for _, key := range keys {
		if value, exists := (*e)[key]; exists {
			return value, nil
		}
	}
	return "", errors.New("Any keys are not found: " + strings.Join(keys, " or "))
}

func (e *Environ) Filter(filter func(k string, v string) bool) QueryMap[string, string] {
	return NewQueryMap(*e).Filter(filter)
}

func (e *Environ) Replace(replacer func(k string, v string) (string, string)) QueryMap[string, string] {
	return NewQueryMap(*e).Replace(replacer)
}
