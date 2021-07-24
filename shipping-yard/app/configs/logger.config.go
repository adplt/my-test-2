package configs

import "errors"

// FancyHandleError function
func FancyHandleError(errCode error) (b bool) {
	if errCode != nil {

	}
	return
}

// FancyHandleResponse function
func FancyHandleResponse(dataPrint string) (b bool) {
	if dataPrint == "" {
		FancyHandleError(errors.New("Message is empty"))
		return
	} else {
	}
	return
}
