package utils

/*RestErr Make an error struct for all of our microservices*/
type RestErr struct {
	StatusCode int
	Message    string
	Error      string
}

func ErrResponse(statusCode int, err error) RestErr {
	return RestErr{StatusCode: statusCode, Message: err.Error()}
}
