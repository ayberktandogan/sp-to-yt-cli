package utils

import (
	"log"
	"net/url"
)

func AppendQueries(ou string, queries map[string]string) url.URL {
	nu, err := url.Parse(ou)
	if err != nil {
		log.Fatal(err)
	}

	values := nu.Query()

	for k, v := range queries {
		values.Add(k, v)
	}

	nu.RawQuery = values.Encode()

	return *nu
}
