package model

import "fmt"

type ErrNotFound struct{}

func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("Task not found.")
}
