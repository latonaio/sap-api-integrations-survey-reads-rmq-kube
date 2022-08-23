package responses

type SurveyCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
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
		} `json:"results"`
	} `json:"d"`
}
