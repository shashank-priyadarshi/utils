package ports

type Pool interface {
	Execute(func(...interface{}) []interface{}) error
	Exit()
}
