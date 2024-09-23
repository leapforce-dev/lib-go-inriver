package inriver

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
)

const (
	apiName string = "Inriver"
)

type Service struct {
	apiUrl      string
	apiKey      string
	httpService *go_http.Service
}

type ServiceConfig struct {
	ApiUrl string
	ApiKey string
}

func NewService(config *ServiceConfig) (*Service, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("ServiceConfig must not be a nil pointer")
	}

	if config.ApiUrl == "" {
		return nil, errortools.ErrorMessage("Service Api Url not provided")
	}

	if config.ApiKey == "" {
		return nil, errortools.ErrorMessage("Service Api Key not provided")
	}

	httpService, e := go_http.NewService(&go_http.ServiceConfig{})
	if e != nil {
		return nil, e
	}

	return &Service{
		apiUrl:      config.ApiUrl,
		apiKey:      config.ApiKey,
		httpService: httpService,
	}, nil
}

func (service *Service) httpRequest(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	// add api key header
	header := http.Header{}
	header.Set("X-inRiver-APIKey", service.apiKey)
	(*requestConfig).NonDefaultHeaders = &header

	// add error model
	errorResponse := ErrorResponse{}
	(*requestConfig).ErrorModel = &errorResponse

	request, response, e := service.httpService.HttpRequest(requestConfig)
	if errorResponse.Title != "" {
		e.SetMessage(errorResponse.Title)
	}

	return request, response, e
}

func (service *Service) url(path string) string {
	return fmt.Sprintf("%s/%s", service.apiUrl, path)
}

func (service *Service) ApiName() string {
	return apiName
}

func (service *Service) ApiKey() string {
	return service.apiKey
}

func (service *Service) ApiCallCount() int64 {
	return service.httpService.RequestCount()
}

func (service *Service) ApiReset() {
	service.httpService.ResetRequestCount()
}
