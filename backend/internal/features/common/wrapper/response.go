package wrapper

type Response[T any] struct {
	Data T `json:"data"`
}
