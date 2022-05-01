package integration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
)

type testCase struct {
	Raw []map[string]interface{} `json:"expected"`
}

func newTestCase(t *testing.T, file string) testCase {
	jsonFile, err := os.Open(path.Join("testdata", file))
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var tc testCase
	require.NoError(t, json.Unmarshal([]byte(byteValue), &tc))

	return tc
}

func (tc testCase) unmarshalModels(models interface{}) error {
	bytes, err := json.Marshal(tc.Raw)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(bytes), models)
}
