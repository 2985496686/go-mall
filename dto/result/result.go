package result

type Result struct {
	Success      bool        `json:"success"`
	ErrorMessage string      `json:"errorMessage"`
	Data         interface{} `json:"data"`
}

func Ok(data interface{}) *Result {
	return &Result{
		Success:      true,
		ErrorMessage: "",
		Data:         data,
	}
}

func Fail(errorMessage string) *Result {
	return &Result{
		Success:      false,
		ErrorMessage: errorMessage,
		Data:         nil,
	}
}
