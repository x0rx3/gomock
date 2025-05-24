package model

type CaseTemplate struct {
	MatchRequestTemplate MatchRequestTemplate `yaml:"MatchRequest" json:"MatchRequest"`
	SetResponseTemplate  SetResponseTemplate  `yaml:"SetResponse" json:"SetResponse"`
}
