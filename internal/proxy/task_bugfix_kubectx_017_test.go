package proxy

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixKubectx017SourceContract(t *testing.T) {
    source, err := os.ReadFile("readonly.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "debugLog.Printf(\">> %s %s\", r.Method, r.URL.Path)") {
        t.Fatalf("expected source contract is missing")
    }
}
