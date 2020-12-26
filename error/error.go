package error

import "errors"

var ErrorsNotAuthenticate = errors.New("This request not authenticate by jwt")

var ErrorsUpdateNotEffect = errors.New("Nothing row been updated")

var NotActionOnThisFunction = errors.New("Not action this function")
