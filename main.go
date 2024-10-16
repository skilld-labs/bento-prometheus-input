package main

import (
	"context"

	_ "bento-prometheus-input/input"

	_ "github.com/warpstreamlabs/bento/public/components/all"
	"github.com/warpstreamlabs/bento/public/service"
)

func main() {
	service.RunCLI(context.Background())
}
