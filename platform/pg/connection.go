package pg

type Connection interface {
	Load(query string) interface{}
}
