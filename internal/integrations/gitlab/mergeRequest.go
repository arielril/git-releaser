package gitlab

import (
	"encoding/json"
	"fmt"
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
		Submit(mr *axios.Instance)
	}
)

func NewMergeRequest(projId int, opts *MergeRequestOptions) (mr MergeRequest) {
	mr = &mergeRequest{
		ProjectId: projId,
		Options:   opts,
	}
	return
}

func (mr *mergeRequest) Submit(instance *axios.Instance) {
	projId := strconv.Itoa(mr.ProjectId)
	reqConfig := &axios.Config{
		Method: "POST",
		URL:    projectMergeRequestPath,
		Params: map[string]string{
			"proj_id": projId,
		},
		Body: mr.Options,
	}

	res, err := instance.Request(reqConfig)

	if err != nil {
		fmt.Println("Error requesting gitlab:", err)
		return
	}

	var parsedRes interface{}
	_ = json.Unmarshal(res.Data, parsedRes)
	fmt.Printf("gitlab response: %#v\n", parsedRes)
}
