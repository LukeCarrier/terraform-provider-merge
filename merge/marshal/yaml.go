package marshal

import "gopkg.in/yaml.v3"

type YamlMarshaller struct{}

func (m *YamlMarshaller) Marshal(data UnmarshalledData) (*string, error) {
	bytes, err := yaml.Marshal(data)
	result := string(bytes)
	return &result, err
}

func (m *YamlMarshaller) Unmarshal(data string, v UnmarshalledData) error {
	bytes := []byte(data)
	return yaml.Unmarshal(bytes, &v)
}
