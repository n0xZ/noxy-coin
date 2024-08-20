package config

import "os"

func Get(k string) string {
	return os.Getenv(k)
}
