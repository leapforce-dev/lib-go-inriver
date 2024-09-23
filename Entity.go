package inriver

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	i_types "github.com/leapforce-libraries/go_inriver/types"
	"net/http"
)

type Entity struct {
	Id                    int                    `json:"id"`
	DisplayName           string                 `json:"displayName"`
	DisplayDescription    string                 `json:"displayDescription"`
	Version               string                 `json:"version"`
	LockedBy              string                 `json:"lockedBy"`
	CreatedBy             string                 `json:"createdBy"`
	CreatedDate           i_types.DateTimeString `json:"createdDate"`
	FormattedCreatedDate  string                 `json:"formattedCreatedDate"`
	ModifiedBy            string                 `json:"modifiedBy"`
	ModifiedDate          i_types.DateTimeString `json:"modifiedDate"`
	FormattedModifiedDate string                 `json:"formattedModifiedDate"`
	ResourceUrl           string                 `json:"resourceUrl"`
	EntityTypeId          string                 `json:"entityTypeId"`
	EntityTypeDisplayName string                 `json:"entityTypeDisplayName"`
	Completeness          *int                   `json:"completeness"`
	FieldSetId            string                 `json:"fieldSetId"`
	FieldSetName          string                 `json:"fieldSetName"`
	SegmentId             int                    `json:"segmentId"`
	SegmentName           string                 `json:"segmentName"`
}

func (service *Service) GetEntitySummary(entityId int32) (*Entity, *errortools.Error) {
	var entity Entity

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("entities/%v/summary", entityId)),
		ResponseModel: &entity,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &entity, nil
}
