package msg

import "fmt"

var ErrorNotFound = fmt.Errorf("Not Found!")

var ErrorAlreadyExists = fmt.Errorf("Already Exists!")

var ErrorUnAuthorized = fmt.Errorf("UnAuthorized!")

var ErrorServerSide = fmt.Errorf("Server Error!")
