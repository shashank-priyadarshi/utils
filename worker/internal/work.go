package internal

type Work struct {
	ID   string
	Work func(...interface{}) []interface{}
}
