package marshal

import "fmt"

type UnmarshalledData = map[string]interface{}

type Marshaller interface {
	Marshal(data UnmarshalledData) (*string, error)
	Unmarshal(data string, v UnmarshalledData) error
}

func NewUnmarshalResult() UnmarshalledData {
	return make(UnmarshalledData, 0)
}

func NewMarshaller(format string) (Marshaller, error) {
	switch format {
	case "json":
		return &JsonMarshaller{}, nil
	case "yaml":
		return &YamlMarshaller{}, nil
	default:
		return nil, fmt.Errorf("unknown format %s", format)
	}
}
