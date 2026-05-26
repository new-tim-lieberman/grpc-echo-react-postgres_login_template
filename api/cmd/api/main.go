package main

import (
	"fmt"

	"github.com/new-timlieberman/gitasy2.0/api/internal/config"
)

func main() {
	cfg := config.Load()

	fmt.Println(cfg.AuthAddr)
}
