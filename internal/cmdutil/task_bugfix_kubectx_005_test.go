package cmdutil

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixKubectx005SourceContract(t *testing.T) {
    source, err := os.ReadFile("deprecated.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if key == `KUBECTX_CURRENT_FGCOLOR` || key == `KUBECTX_CURRENT_BGCOLOR` {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if key != `KUBECTX_CURRENT_FGCOLOR` || key == `KUBECTX_CURRENT_BGCOLOR` {") {
        t.Fatalf("mutated source contract is still present")
    }
}
