package initialize

import (
	"flag"
	"fmt"
	"os"
)

func LoadEnv() string {
	env := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage:: server -e {mode}")
		os.Exit(1)
	}

	flag.Parse()

	return GetEnv(*env)
}

func GetEnv(env string) string {
	if len(env) == 0 {
		return "default"
	}

	return env
}
