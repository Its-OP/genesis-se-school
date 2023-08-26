package domain

type DataConsistencyError struct {
	Message string
}

func (e DataConsistencyError) Error() string {
	return e.Message
}

type EndpointInaccessibleError struct {
	Message string
}

func (e EndpointInaccessibleError) Error() string {
	return e.Message
}

type ArgumentError struct {
	Message string
}

func (e ArgumentError) Error() string {
	return e.Message
}
