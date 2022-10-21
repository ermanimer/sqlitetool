package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/alecthomas/kong"
)

func main() {
	ctx, cancelFunc := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancelFunc()

	cli := kong.Parse(&CLI{})

	if err := cli.Run(ContextWrapper{ctx}); err != nil {
		fmt.Println(err.Error())
		return
	}
}
