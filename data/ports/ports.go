package ports

type Data interface {
	Generate(map[string]interface{}) (string, error)
	Validate(string, map[string]interface{}) error
	Refresh(string, map[string]interface{}) (string, error)
}
