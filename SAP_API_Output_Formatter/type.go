package sap_api_output_formatter

type Survey struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	APISchema     string `json:"api_schema"`
	SurveyCode    string `json:"survey_code"`
	Deleted       bool   `json:"deleted"`
}

type SurveyCollection struct {
	ObjectID                string `json:"ObjectID"`
	ETag                    string `json:"ETag"`
	CategoryCode            string `json:"CategoryCode"`
	CategoryCodeText        string `json:"CategoryCodeText"`
	LifeCycleStatusCode     string `json:"LifeCycleStatusCode"`
	LifeCycleStatusCodeText string `json:"LifeCycleStatusCodeText"`
	ValidFromDate           string `json:"ValidFromDate"`
	ValidToDate             string `json:"ValidToDate"`
	ID                      string `json:"ID"`
	Version                 string `json:"Version"`
	Name                    string `json:"Name"`
	LanguageCode            string `json:"languageCode"`
	LanguageCodeText        string `json:"languageCodeText"`
	EntityLastChangedOn     string `json:"EntityLastChangedOn"`
}
