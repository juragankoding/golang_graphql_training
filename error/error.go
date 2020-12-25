package error

import "errors"

var ErrorsNotAuthenticate = errors.New("This request not authenticate by jwt")
