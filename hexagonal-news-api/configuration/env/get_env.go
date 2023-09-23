package env

import "os"

func GetNewsTokenAPI() string {
	return os.Getenv("NEWS_API_KEY")
}
