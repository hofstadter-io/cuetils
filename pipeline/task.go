package pipeline

import (
	"cuelang.org/go/tools/flow"
)

type Task interface {
  Run(t *flow.Task, err error) error
}
