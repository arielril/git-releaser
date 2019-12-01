package telegram

import (
	"fmt"

	"github.com/vicanso/go-axios"
)

type (
	errorResponse struct {
		Description string `json:"description"`
		Ok          string `json:"ok,omitempty"`
	}
)

func (er *errorResponse) String() string {
	return er.Description
}

func ParseErrorResponse(res *axios.Response) (err error) {
	if res.Status >= 400 {
		var errResponse errorResponse
		parseErr := res.JSON(&errResponse)

		if parseErr != nil {
			err = fmt.Errorf("failed to parse response. error=%s", parseErr.Error())
		} else {
			err = fmt.Errorf("%s", errResponse)
		}
	}
	return
}
