// cmd/app/main.go
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"KIS_HTS/internal/config"
	"KIS_HTS/internal/infrastructure/kis"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	approval, err := kis.RequestApprovalKey(cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("approval:", approval)

	ws := kis.WSClient{}

	// websocket 연결
	err = ws.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// subscribe example
	msg, _ := kis.BuildSubscribeMessage(
		approval,
		"H0STCNT0",
		"005930",
	)

	ws.Send(msg)

	// graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sig
		cancel()
	}()

	// websocket run
	ws.Run(ctx, func(msg []byte) {
		log.Println("recv:", string(msg))
	})

	ws.Close()
}
