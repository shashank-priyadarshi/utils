package algo_test

import (
	_ "github.com/onsi/ginkgo/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.ssnk.in/utils/algo"
	"testing"
)

// go test -count=1 -json -v .

type TrieTestSuite struct {
	suite.Suite
	trie  *algo.Trie[string, string]
	tests []test
}

type test struct {
	name     string
	input    input
	expected expected
}

type input struct {
	keys  []string
	value string
}

type expected struct {
	output string
	exists bool
	err    error
}

func TestTrie(t *testing.T) {
	suite.Run(t, new(TrieTestSuite))
}

func (t *TrieTestSuite) SetupAllSuite() {
	t.trie = algo.NewTrie[string, string]()
}

func (t *TrieTestSuite) BeforeTest() {
	t.tests = []test{{
		name: "happy path",
		input: input{
			keys:  []string{"hello", "world", "my", "name", "is"},
			value: "jack",
		},
		expected: expected{
			output: "jack",
			exists: true,
			err:    nil,
		},
	}}
}

func (t *TrieTestSuite) TestTrie_InsertAndSearch() {
	for _, tt := range t.tests {

		t.trie.Insert(tt.input.keys, tt.input.value)

		v, ok := t.trie.Search(tt.input.keys)
		// Debug logs
		suite.TestingSuite(t).T().Logf("Testing InsertAndSearch: %s", tt.name)
		suite.TestingSuite(t).T().Logf("Expected value: %s, Got: %s", tt.expected.output, v)
		suite.TestingSuite(t).T().Logf("Expected exists: %v, Got: %v", tt.expected.exists, ok)

		assert.Equal(suite.TestingSuite(t).T(), tt.expected.output, v, "Failed for test case: %s", tt.name)
		assert.Equal(suite.TestingSuite(t).T(), tt.expected.exists, ok, "Failed for test case: %s", tt.name)
	}
}

func (t *TrieTestSuite) TestTrie_DeleteAndSearch() {

	for _, tt := range t.tests {

		err := t.trie.Delete(tt.input.keys)
		// Debug logs
		suite.TestingSuite(t).T().Logf("Testing Delete: %s", tt.name)
		suite.TestingSuite(t).T().Logf("Expected error: %v, Got: %v", nil, err)

		assert.NoError(suite.TestingSuite(t).T(), err, "Delete operation failed for test case: %s", tt.name)

		v, ok := t.trie.Search(tt.input.keys)

		expected := expected{
			output: "",
			exists: false,
			err:    nil,
		}

		// Debug logs
		suite.TestingSuite(t).T().Logf("Expected value: %s, Got: %s", expected.output, v)
		suite.TestingSuite(t).T().Logf("Expected exists: %v, Got: %v", expected.exists, ok)

		assert.Equal(suite.TestingSuite(t).T(), expected.err, err, "Failed for test case: %s", tt.name)
		assert.Equal(suite.TestingSuite(t).T(), expected.output, v, "Failed for test case: %s", tt.name)
		assert.Equal(suite.TestingSuite(t).T(), expected.exists, ok, "Failed for test case: %s", tt.name)
	}
}
