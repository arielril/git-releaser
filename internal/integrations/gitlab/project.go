package gitlab

import (
	"errors"
	"fmt"

	fp "github.com/novalagung/gubrak"
)

type (
	Project struct {
		Id       int    `json:"id" yaml:"id"`
		Name     string `json:"name" yaml:"name"`
		WebUrl   string `json:"web_url" yaml:"webUrl"`
		Branches struct {
			Develop string `json:"-" yaml:"develop"`
			Master  string `json:"-" yaml:"master"`
		} `json:"-" yaml:"branches"`
		MRs []MergeRequest
		gl  Gitlab
	}

	IProject interface {
		SubmitMergeRequest(options *MergeRequestOptions) (mrResult Response, err error)
		SetGitlab(gl Gitlab)
	}
)

func NewProject(id int, name, webUrl string, gitlab Gitlab) (p *Project) {
	mrs := make([]MergeRequest, 0)
	p = &Project{
		Id:     id,
		Name:   name,
		WebUrl: webUrl,
		MRs:    mrs,
		gl:     gitlab,
	}

	return
}

func (p *Project) SetGitlab(gl Gitlab) {
	p.gl = gl
}

func (p *Project) MergeProject(extra *Project) {
	if p.Id == 0 {
		p.Id = extra.Id
	}

	if p.Name == "" {
		p.Name = extra.Name
	}

	if p.WebUrl == "" {
		p.WebUrl = extra.WebUrl
	}
}

func (p *Project) findProjectByUrl() (proj *Project, err error) {
	projs, err := p.gl.ListOwnedProjects()
	fmt.Printf("Projects: %#v %#v\n\n", err, projs)
	if err != nil {
		return
	}

	finded, err := fp.Find(projs, func(proj *Project, idx int) bool {
		return proj.WebUrl == p.WebUrl
	})

	if finded != nil {
		proj = finded.(*Project)
	}

	return
}

func (p *Project) SubmitMergeRequest(options *MergeRequestOptions) (mrResult Response, err error) {
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

	if options.SourceBranch == "" {
		options.SourceBranch = "develop"
	}

	if options.TargetBranch == "" {
		options.TargetBranch = "master"
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
