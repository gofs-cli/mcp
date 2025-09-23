package utils

type ReturnError struct {
	Message string
}

func (m *ReturnError) Error() string {
	return m.Message
}
