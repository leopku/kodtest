package store

type IStore interface {
	Get(string) any
	Set(string, any) error
}
