package client

import (
	"fmt"
	"strings"
)

type Error struct {
	Message error
}

type Errors struct {
	List []Error
}

func newError(err error) Error {
	clientError := new(Error)
	clientError.Message = err
	return *clientError
}

// Add will add a new error object to the errors list
func (errors *Errors) add(err error) {
	if err != nil {
		errors.List = append(errors.List, newError(err))
	}
}

// Any will return true if there are errors, otherwise false
func (errors Errors) Any() bool {
	return len(errors.List) > 0
}

// Messages return all the messages assocaited with an Errors object as a slice of strings
func (errors Errors) Messages() (messages []string) {
	for _, clientError := range errors.List {
		messages = append(messages, clientError.Message.Error())
	}
	return
}

// JoinMessages will join all the messages associated with an Errors object as a single go error
func (errors Errors) JoinMessages() (err error) {
	if errors.Any() {
		err = fmt.Errorf(strings.Join(errors.Messages(), ", "))
	}
	return
}
