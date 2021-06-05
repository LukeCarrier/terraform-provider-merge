package marshal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestYamlMarshal(t *testing.T) {
	input := NewUnmarshalResult()
	input["name"] = "Luke"
	marshaller := YamlMarshaller{}
	result, err := marshaller.Marshal(input)

	assert.Nil(t, err)
	assert.Equal(t, "name: Luke\n", *result)
}

func TestYamlUnmarshal(t *testing.T) {
	input := "name: Luke\n"
	marshaller := YamlMarshaller{}
	result := NewUnmarshalResult()
	err := marshaller.Unmarshal(input, result)

	assert.Nil(t, err)
	assert.Equal(t, "Luke", result["name"])
}

func TestYamlUnmarshalWithMalformedInput(t *testing.T) {
	input := "{"
	marshaller := YamlMarshaller{}
	result := NewUnmarshalResult()
	err := marshaller.Unmarshal(input, result)

	assert.NotNil(t, err)
	assert.Empty(t, result)
}
