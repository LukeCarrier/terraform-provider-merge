package marshal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonMarshal(t *testing.T) {
	input := NewUnmarshalResult()
	input["name"] = "Luke"
	marshaller := JsonMarshaller{}
	result, err := marshaller.Marshal(input)

	assert.Nil(t, err)
	assert.Equal(t, "{\"name\":\"Luke\"}", *result)
}

func TestJsonUnmarshal(t *testing.T) {
	input := "{\"name\":\"Luke\"}"
	marshaller := JsonMarshaller{}
	result := NewUnmarshalResult()
	err := marshaller.Unmarshal(input, result)

	assert.Nil(t, err)
	assert.Equal(t, "Luke", result["name"])
}

func TestJsonUnmarshalWithMalformedInput(t *testing.T) {
	input := "{"
	marshaller := JsonMarshaller{}
	result := NewUnmarshalResult()
	err := marshaller.Unmarshal(input, result)

	assert.NotNil(t, err)
	assert.Empty(t, result)
}
