package api

type ApplyJVResp struct {
	
	TraceID    any    `json:"traceId,omitempty"`
	IsSuccess  bool   `json:"isSuccess,omitempty"`
	StatusCode int    `json:"statusCode,omitempty"`
	Message    string `json:"message,omitempty"`
	Data       Data   `json:"data,omitempty"`
}
type Data struct {
	ApplicationCount int `json:"applicationCount,omitempty"`
}