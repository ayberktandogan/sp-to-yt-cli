package utils

import (
	"log"
	"net/url"
)

func AppendQueries(ou string, queries map[string]string) string {
	nu, err := url.Parse(ou)
	if err != nil {
		log.Fatal(err)
	}

	values := nu.Query()

	for k, v := range queries {
		if v == "" || v == "0" {
			continue
		}
		values.Add(k, v)
	}

	nu.RawQuery = values.Encode()

	return nu.String()
}
