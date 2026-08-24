package kubeconfig

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixKubectx015SourceContract(t *testing.T) {
    source, err := os.ReadFile("contexts.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if contexts == nil {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if false && contexts == nil {") {
        t.Fatalf("mutated source contract is still present")
    }
}
