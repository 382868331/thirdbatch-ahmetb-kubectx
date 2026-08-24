package kubeconfig

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixKubectx013SourceContract(t *testing.T) {
    source, err := os.ReadFile("currentcontext.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if len(k.files) == 0 {") {
        t.Fatalf("expected source contract is missing")
    }
}
