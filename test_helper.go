package simc_config

import (
	"fmt"
	"testing"
)

func assertEqualString(t *testing.T, a string, b string) {
  if a != b {
		t.Error(fmt.Sprintf("'%s' should be '%s'", a, b))
	}
}

func assertEqualInt(t *testing.T, a int, b int) {
	if a != b {
		t.Error(fmt.Sprintf("'%d' should be '%d'", a, b))
	}
}

func assertEqualStringSlice(t *testing.T, a []string, b []string) {
	if len(a) != len(b) {
		t.Error(fmt.Sprintf("'%v' should be '%v'", a, b))
	}
	for i, _ := range b {
		if a[i] != b[i] {
			t.Error(fmt.Sprintf("'%v' should be '%v'", a, b))
		}
	}

}
