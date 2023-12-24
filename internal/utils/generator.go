package utils

import (
	"github.com/oklog/ulid/v2"
)

func StateGenerator() string {
	return ulid.Make().String()
}
