package main

import (
	"context"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)
	defer cancel()
	init := "Hello, World!"
	ticker := time.NewTicker(time.Second)

mainloop:
	for {
		select {
			case <-ctx.Done():
				break mainloop
			case <-ticker.C:
				init = strings.Repeat(init, 10)
		}
	}
}
