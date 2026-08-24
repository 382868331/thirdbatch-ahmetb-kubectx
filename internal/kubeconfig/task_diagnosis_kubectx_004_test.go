package kubeconfig

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisKubectx004SourceContract(t *testing.T) {
    source, err := os.ReadFile("kubeconfig.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if rn.YNode().Kind != yaml.MappingNode {") {
        t.Fatalf("expected source contract is missing")
    }
}
