package chaincode

type fieldType string

var (
	String fieldType = "String"
	Integer fieldType = "Integer"
	Boolean fieldType = "Boolean"
	Long fieldType = "Long"
	Double fieldType = "Double"
	Float fieldType = "Float"
	UnknownFieldType fieldType = "Unknown"
)


func GetFieldType(name string) fieldType {
	switch name {
	case "String":
		return String
	case "Integer":
		return Integer
	case "Boolean":
		return Boolean
	case "Long":
		return Long
	case "Double":
		return Double
	case "Float":
		return Float
	default:
		return UnknownFieldType
	}
}