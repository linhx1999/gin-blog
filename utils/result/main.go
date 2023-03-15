package result

type result struct {
	Message string `json:"message"`
	Data    []any  `json:"data"`
}

func New(msg string, data ...any) *result {
	if len(data) == 0 {
		return &result{
			Message: msg,
			Data:    []any{},
		}
	} else {
		return &result{
			Message: msg,
			Data:    data,
		}
	}
}
