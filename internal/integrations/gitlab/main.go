package gitlab

import (
	"fmt"
	http "git-releaser/internal/integrations"
	"net/url"

	"github.com/vicanso/go-axios"
)

type (
	gitlab struct {
		instance *axios.Instance
		Projects []IProject
	}

	Gitlab interface {
		GetInstance() *axios.Instance
		AddProject(proj IProject)
		ListOwnedProjects() (projs []*Project, err error)
	}
)

const (
	projectMergeRequestPath = "/projects/:proj_id/merge_requests"
)

func New(baseUrl, pvtToken string) (gl *gitlab) {
	instance := http.New(baseUrl)
	instance.Config.Headers.Set("Authorization", fmt.Sprintf("Bearer %v", pvtToken))
	instance.Config.ResponseInterceptors = append(instance.Config.ResponseInterceptors, ParseErrorResponse)

	projs := make([]IProject, 0)

	gl = &gitlab{
		instance: instance,
		Projects: projs,
	}

	return
}

func (gl *gitlab) GetInstance() *axios.Instance {
	return gl.instance
}

func (gl *gitlab) AddProject(proj IProject) {
	gl.Projects = append(gl.Projects, proj)
}

func (gl *gitlab) ListOwnedProjects() (projs []*Project, err error) {
	reqConfig := &axios.Config{
		Method: "GET",
		URL:    "/projects",
		Query: url.Values{
			"owned":  []string{"true"},
			"simple": []string{"true"},
		},
	}

	res, err := gl.instance.Request(reqConfig)

	if err == nil {
		_ = res.JSON(&projs)
	}
	return
}
