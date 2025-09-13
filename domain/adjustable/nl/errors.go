package nl

import "fmt"

type ErrNightlightAdapterUnavailable struct {
	Adapter string
}

func (e *ErrNightlightAdapterUnavailable) Error() string {
	return fmt.Sprintf("adapter %s is unavailable", e.Adapter)
}
