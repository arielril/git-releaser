package gitlab

import (
	"errors"
	"fmt"

	fp "github.com/novalagung/gubrak"
)

type (
	project struct {
		Id     int    `json:"id"`
		Name   string `json:"name"`
		Path   string `json:"path"`
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

func (p *project) MergeProject(extra *project) {
	if p.Id == 0 {
		p.Id = extra.Id
	}

	if p.Name == "" {
		p.Name = extra.Name
	}

	if p.Path == "" {
		p.Path = extra.Path
	}

	if p.WebUrl == "" {
		p.WebUrl = extra.WebUrl
	}
}

func (p *project) findProjectByUrl() (proj *project, err error) {
	projs, err := p.gl.ListOwnedProjects()
	fmt.Printf("Projects: %#v %#v\n\n", err, projs)
	if err != nil {
		return
	}

	finded, err := fp.Find(projs, func(proj *project, idx int) bool {
		return proj.WebUrl == p.WebUrl
	})

	if finded != nil {
		proj = finded.(*project)
	}

	return
}

func (p *project) SubmitMergeRequest(options *MergeRequestOptions) (mrResult Response, err error) {
	if p.Id == 0 && p.WebUrl != "" {
		proj, projErr := p.findProjectByUrl()

		if projErr != nil {
			err = projErr
			return
		} else if proj == nil {
			err = errors.New("failed to find the project with web url")
			return
		}

		p.MergeProject(proj)
	}

	mr := NewMergeRequest(p.Id, options)

	mrResult, err = mr.Submit(p.gl.GetInstance())
	if mrResult != nil {
		mrRes := mrResult.(*mergeRequestResponse)
		mrRes.ProjectName = p.Name
		mrResult = mrRes
	}
	return
}
