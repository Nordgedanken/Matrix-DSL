package js

type pkgJson struct {
	Name         string        `json:"name,omitempty"`
	Version      string        `json:"version,omitempty"`
	Description  string        `json:"description,omitempty"`
	Main         string        `json:"main,omitempty"`
	Dependencies []interface{} `json:"dependencies,omitempty"`
}
