package main

import (
	"context"

	_ "benthos-prometheus-input/input"

	"github.com/redpanda-data/benthos/v4/public/service"
)

func main() {
	service.RunCLI(context.Background())
}
