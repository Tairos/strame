package connection

import (
  "fmt"
  "context"
)

type WsServer struct {}

func NewWsServer() Server {
  return &WsServer{}
}

func (ws *WsServer) Run(ctx context.Context, ch chan string) {
  go func (ctx context.Context, ch chan string) {
    for {
      select {
      case <-ctx.Done():
        return
      case msg := <-ch:
        fmt.Println(msg)
      }
    }
  }(ctx, ch)
}
