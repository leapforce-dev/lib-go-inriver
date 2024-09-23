package inriver

// ErrorResponse stores general Inriver API error response
type ErrorResponse struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int    `json:"status"`
	Detail   int    `json:"detail"`
	Instance int    `json:"instance"`
	TraceId  string `json:"traceId"`
}
