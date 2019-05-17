package connection

import (
  "context"
)

type Server interface {
	Run(ctx context.Context, ch chan string)
}
