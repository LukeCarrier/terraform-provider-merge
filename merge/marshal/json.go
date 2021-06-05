package marshal

import "encoding/json"

type JsonMarshaller struct{}

func (m *JsonMarshaller) Marshal(data UnmarshalledData) (*string, error) {
	bytes, err := json.Marshal(data)
	result := string(bytes)
	return &result, err
}

func (m *JsonMarshaller) Unmarshal(data string, v UnmarshalledData) error {
	bytes := []byte(data)
	return json.Unmarshal(bytes, &v)
}
