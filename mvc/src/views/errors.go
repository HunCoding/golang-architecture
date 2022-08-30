package views

type ErrorView struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"errorCode"`
}

func ParseError(
	message string,
	errorCode int,
) *ErrorView {
	return &ErrorView{
		message, errorCode,
	}
}
