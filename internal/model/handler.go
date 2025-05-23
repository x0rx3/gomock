package model

type HandlerTamplate struct {
	Name  string                    `yaml:"Name" json:"Name"`
	Path  string                    `yaml:"Path" json:"Path"`
	Cases map[string][]CaseTemplate `yaml:"Cases" json:"Cases"`
}
