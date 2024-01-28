package middlewares

import (
	"os"
)

func Read(path string) string {
	data, err := os.ReadFile(path)
	FailOnError(err, "Failed to read file")
	return string(data)
}
