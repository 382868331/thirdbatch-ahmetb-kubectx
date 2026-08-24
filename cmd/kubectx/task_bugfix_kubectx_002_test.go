package main

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixKubectx002SourceContract(t *testing.T) {
    source, err := os.ReadFile("flags.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if argv[0] == \"--readonly\" || argv[0] == \"-r\" {") {
        t.Fatalf("expected source contract is missing")
    }
}
