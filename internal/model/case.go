package model

type CaseTemplate struct {
	MatchRequest   MatchRequestTemplate `yaml:"MatchRequest" json:"MatchRequest"`
	SetResponseDto SetResponseTemplate  `yaml:"SetResponse" json:"SetResponse"`
}
