package source

import (
  "context"
)

type Source interface {
	Poll(ctx context.Context)
}
