package client

type Promise[T any] struct {
	Res T
	Err error
}
