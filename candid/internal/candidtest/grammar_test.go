package candidtest_test

import (
	"os"
	"testing"

	"github.com/aviate-labs/agent-go/candid/internal/candidtest"
)

func TestData(t *testing.T) {
	rawDid, err := os.ReadFile("../../idl/testdata/prim.test.did")
	if err != nil {
		t.Fatal(err)
	}
	p, err := candidtest.NewParser([]rune(string(rawDid)))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := p.ParseEOF(candidtest.TestData); err != nil {
		t.Fatal(err)
	}
}
