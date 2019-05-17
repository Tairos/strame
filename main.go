package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Tairos/strame/admin"
	"github.com/Tairos/strame/facebook"
)

func main() {
	messages := make(chan string, 100)
	//ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	c := make(chan os.Signal)

	//CREATE AN ENV VARIABLE FOR THIS RESOURCES PATH
	//CONFIG GILE AS A .cfg SERVICE? AND COMPILE THE HTML? MAYBE TWIG?
	// ex, err := os.Executable()
	// if err != nil {
	// 	panic(err)
	// }
	// ResourcesPath := filepath.Dir(ex)

	//HTML resources loading
	ResourcesPath := os.Getenv("GOPATH") + "/src/github.com/Tairos/strame"
	adminPanel := admin.NewPanel(ResourcesPath)

	//RUN facebook poller
	fb := facebook.NewSource(messages)
	fb.Poll(ctx)

	//RUN admin panel
	go adminPanel.Run(ctx)

	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)
	for {
		select {
		case signal := <-c:
			switch signal {
			case syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1:
				return
			}
		case message := <-messages:
			fmt.Println(message)
		case <-ctx.Done():
			return
		}
	}

	// ws := connection.NewWsServer()
	// ws.Run(ctx, messages)
	//
	// fb := facebook.NewSource(messages)
	// fb.Poll(ctx)

	fmt.Println("STOP")
}
