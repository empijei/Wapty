package apis

import (
	"fmt"
)

// Err returns a command that contains the error message as argument and ERR as action
func Err(m interface{}) *Command {
	message := fmt.Sprint(m)
	return &Command{
		Action: ERR,
		Args:   map[ArgName]string{ERR: message},
	}
}
