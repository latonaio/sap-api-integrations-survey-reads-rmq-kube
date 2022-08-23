package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-survey-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToSurveyCollection(raw []byte, l *logger.Logger) ([]SurveyCollection, error) {
	pm := &responses.SurveyCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to SurveyCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	surveyCollection := make([]SurveyCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		surveyCollection = append(surveyCollection, SurveyCollection{
			ObjectID:                data.ObjectID,
			ETag:                    data.ETag,
			CategoryCode:            data.CategoryCode,
			CategoryCodeText:        data.CategoryCodeText,
			LifeCycleStatusCode:     data.LifeCycleStatusCode,
			LifeCycleStatusCodeText: data.LifeCycleStatusCodeText,
			ValidFromDate:           data.ValidFromDate,
			ValidToDate:             data.ValidToDate,
			ID:                      data.ID,
			Version:                 data.Version,
			Name:                    data.Name,
			LanguageCode:            data.LanguageCode,
			LanguageCodeText:        data.LanguageCodeText,
			EntityLastChangedOn:     data.EntityLastChangedOn,
		})
	}

	return surveyCollection, nil
}
