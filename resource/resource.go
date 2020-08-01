package resource

import "strings"

const (
	DefaultResourceRootPath = "/usr/local/resources"
)

type resource struct {
	Group string
	Name string
	RootPath string
}

func NewResource(group string, name string) *resource {

	return &resource{Group: group, Name: name, RootPath: DefaultResourceRootPath}
}

func (r *resource) GetResourcePath(path string) string {

	return strings.Join([]string{r.RootPath, r.Group, r.Name, path}, "/")
}

func GetResourceRootPath() string {

	return DefaultResourceRootPath
}