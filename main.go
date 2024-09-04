package main

import (
	"context"

	_ "benthos-prometheus-input/input"

	_ "github.com/redpanda-data/connect/v4/public/components/all"
	"github.com/redpanda-data/benthos/v4/public/service"
)

func main() {
	service.RunCLI(context.Background())
}
