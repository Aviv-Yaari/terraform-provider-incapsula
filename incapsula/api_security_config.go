package incapsula

type ViolationActions struct {
	InvalidUrlViolationAction        string `json:"invalidUrlViolationAction"`
	InvalidMethodViolationAction     string `json:"invalidMethodViolationAction"`
	MissingParamViolationAction      string `json:"missingParamViolationAction"`
	InvalidParamNameViolationAction  string `json:"invalidParamNameViolationAction"`
	InvalidParamValueViolationAction string `json:"invalidParamValueViolationAction"`
}

type UserViolationActions struct {
	//InvalidUrlViolationAction        string `json:"invalidUrlViolationAction"`
	MissingParamViolationAction      string `json:"missingParamViolationAction"`
	InvalidParamNameViolationAction  string `json:"invalidParamNameViolationAction"`
	InvalidParamValueViolationAction string `json:"invalidParamValueViolationAction"`
}

type EndpointResponse struct {
	Id     int    `json:"id"`
	Path   string `json:"path"`
	Method string `json:"method"`
	//PiiList                      []string             `json:"piiList"`
	ViolationActions             UserViolationActions `json:"violationActions"`
	SpecificationViolationAction string               `json:"specificationViolationAction"`
}

