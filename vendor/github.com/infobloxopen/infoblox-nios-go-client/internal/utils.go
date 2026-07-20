package internal

type MappedNullable interface {
	ToMap() (map[string]interface{}, error)
}
