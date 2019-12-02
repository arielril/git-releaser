package gitlab

import (
	"strconv"

	"github.com/vicanso/go-axios"
)

type (
	MergeRequestOptions struct {
		SourceBranch string `json:"source_branch"`
		TargetBranch string `json:"target_branch"`
		Title        string `json:"title"`
	}

	mergeRequest struct {
		ProjectId int
		Options   *MergeRequestOptions
	}

	MergeRequest interface {
		Submit(mr *axios.Instance) (mrResult Response, err error)
	}
)

func NewMergeRequest(projId int, opts *MergeRequestOptions) (mr MergeRequest) {
	mr = &mergeRequest{
		ProjectId: projId,
		Options:   opts,
	}
	return
}

func (mr *mergeRequest) Submit(instance *axios.Instance) (mrResult Response, err error) {
	reqConfig := &axios.Config{
		Method: "POST",
		URL:    projectMergeRequestPath,
		Params: map[string]string{
			"proj_id": strconv.Itoa(mr.ProjectId),
		},
		Body: mr.Options,
	}

	res, err := instance.Request(reqConfig)

	if err == nil {
		var mrRes *mergeRequestResponse
		_ = res.JSON(&mrRes)
		mrResult = mrRes
	}

	return
}
