package jsonutil_test

import (
	"encoding/json"
	"github.com/rayspock/go-graphql-impostor/pkg/jsonutil"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestJSONUtilFixture(t *testing.T) {
	suite.Run(t, new(JSONUtilFixture))
}

type JSONUtilFixture struct {
	suite.Suite
}

func (s *JSONUtilFixture) SetupTest() {
}

func (s *JSONUtilFixture) TestMarshalVariableMap() {
	// Given
	jsonStr := `{"test":1,"test2":"test","test3":1.1,"test4":{"test":1}
	,"test5":[1,2,3],"test6":[{"test":1},{"test":false}],"test7":true, "test8":null}`
	var v jsonutil.VariableMap
	err := json.Unmarshal([]byte(jsonStr), &v)
	s.NoError(err)

	// When
	b, err := json.Marshal(v)

	// Then
	s.NoError(err)
	s.NotEmpty(b)
	v2 := jsonutil.VariableMap{}
	err = json.Unmarshal(b, &v2)
	s.NoError(err)
	s.Equal(v, v2)
}
