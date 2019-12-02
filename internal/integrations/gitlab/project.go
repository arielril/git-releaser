package gitlab

type (
	project struct {
		Id     int
		Name   string
		Path   string
		WebUrl string `json:"web_url"`
		MRs    []MergeRequest
		gl     Gitlab
	}

	Project interface {
		SubmitMergeRequest(options *MergeRequestOptions) (mrResult Response, err error)
	}
)

func NewProject(id int, name, path, webUrl string, gitlab Gitlab) (p *project) {
	mrs := make([]MergeRequest, 0)
	p = &project{
		Id:     id,
		Name:   name,
		Path:   path,
		WebUrl: webUrl,
		MRs:    mrs,
		gl:     gitlab,
	}

	return
}

func (p *project) SubmitMergeRequest(options *MergeRequestOptions) (mrResult Response, err error) {
	mr := NewMergeRequest(p.Id, options)

	mrResult, err = mr.Submit(p.gl.GetInstance())
	if mrResult != nil {
		mrRes := mrResult.(*mergeRequestResponse)
		mrRes.ProjectName = p.Name
		mrResult = mrRes
	}
	return
}
