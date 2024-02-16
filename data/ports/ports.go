package ports

type Data interface {
	Generate() error
	Validate() error
	Refresh() error
}
