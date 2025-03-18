package msg

import "fmt"

var ErrorNotFound = fmt.Errorf("Not Found!")

var ErrorBadRequest = fmt.Errorf("Bad Request, Please Try Again...")

var ErrorAlreadyExists = fmt.Errorf("Already Exists!")

var ErrorUnAuthorized = fmt.Errorf("UnAuthorized!")

var ErrorServerSide = fmt.Errorf("Server Error!")

var ErrorConflict = fmt.Errorf("Conflict Occured")
