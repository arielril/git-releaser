package gitlab

import (
	"fmt"
	http "git-releaser/internal/integrations"

	"github.com/vicanso/go-axios"
)

type (
	gitlab struct {
		instance *axios.Instance
		Projects []Project
	}

	Gitlab interface {
		GetInstance() *axios.Instance
		AddProject(proj Project)
	}
)

const (
	projectMergeRequestPath = "/projects/:proj_id/merge_requests"
)

func New(baseUrl, pvtToken string) (gl *gitlab) {
	instance := http.New(baseUrl)
	instance.Config.Headers.Set("Authorization", fmt.Sprintf("Bearer %v", pvtToken))

	projs := make([]Project, 0)

	gl = &gitlab{
		instance: instance,
		Projects: projs,
	}

	return
}

func (gl *gitlab) GetInstance() *axios.Instance {
	return gl.instance
}

func (gl *gitlab) AddProject(proj Project) {
	gl.Projects = append(gl.Projects, proj)
}
