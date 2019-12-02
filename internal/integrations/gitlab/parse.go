package gitlab

import (
	"fmt"
	"strings"

	"github.com/vicanso/go-axios"
)

type (
	errorResponse struct {
		Messages    []string `json:"message,omitempty"`
		Error       string   `json:"error,omitempty"`
		Description string   `json:"error_description,omitempty"`
	}

	mergeRequestResponse struct {
		ID          int `json:"iid"`
		ProjectId   int `json:"project_id"`
		ProjectName string
		MergeStatus string `json:"merge_status"`
		MergeError  string `json:"merge_error,omitempty"`
		Sha         string `json:"sha,omitempty"`
		WebUrl      string `json:"web_url"`
		Changes     int    `json:"changes_count,omitempty,string"`
		CreatedAt   string `json:"created_at"`
	}

	Response interface {
		String() string
	}
)

func (er errorResponse) String() string {
	if er.Messages != nil && len(er.Messages) > 0 {
		return er.Messages[0]
	}
	return fmt.Sprintf("%s. %s", er.Error, er.Description)
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

func (mrRes *mergeRequestResponse) String() string {
	return fmt.Sprintf(`
		Project: %s (%d)
			Merge Request (%d): [available here](%s)
			Changes: %d
			Status: %s
			Sha: %s
			Date: %s
	`,
		mrRes.ProjectName, mrRes.ProjectId,
		mrRes.ID, mrRes.WebUrl,
		mrRes.Changes,
		strings.ReplaceAll(mrRes.MergeStatus, "_", "\\_"),
		mrRes.Sha,
		mrRes.CreatedAt,
	)
}
