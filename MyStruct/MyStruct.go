package MyStruct

type ExportStruct struct {
	Value int
}

func (es *ExportStruct) PrintValue() int {
	return es.Value
}

func PrintVal() int {
	return 123
}
