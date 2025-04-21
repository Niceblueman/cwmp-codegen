package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/Niceblueman/cwmp-codegen/internal/models"
)

// createBenchmarkModel creates a data model for benchmarking with the specified number of objects
func createBenchmarkModel(numObjects, paramsPerObject int) *models.DataModel {
	model := &models.DataModel{
		Name:        "BenchmarkModel",
		Description: "Model for benchmarking",
		Version:     "1.0",
		Objects:     make([]models.Object, numObjects),
	}

	for i := 0; i < numObjects; i++ {
		obj := models.Object{
			Name:        fmt.Sprintf("Object%d", i),
			Description: fmt.Sprintf("Benchmark object %d", i),
			Parameters:  make([]models.Parameter, paramsPerObject),
		}

		for j := 0; j < paramsPerObject; j++ {
			paramType := "string"
			if j%3 == 0 {
				paramType = "int"
			} else if j%3 == 1 {
				paramType = "boolean"
			}

			obj.Parameters[j] = models.Parameter{
				Name:        fmt.Sprintf("Param%d_%d", i, j),
				Description: fmt.Sprintf("Parameter %d of object %d", j, i),
				Type:        paramType,
			}
		}

		model.Objects[i] = obj
	}

	return model
}

func BenchmarkGenerateGolang(b *testing.B) {
	tmpDir := b.TempDir()
	model := createBenchmarkModel(10, 10) // 10 objects with 10 parameters each

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := GenerateGolang(model, tmpDir)
		if err != nil {
			b.Fatalf("Error in GenerateGolang: %v", err)
		}
	}
}

func BenchmarkGenerateTypeScript(b *testing.B) {
	tmpDir := b.TempDir()
	model := createBenchmarkModel(10, 10) // 10 objects with 10 parameters each

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := GenerateTypeScript(model, tmpDir)
		if err != nil {
			b.Fatalf("Error in GenerateTypeScript: %v", err)
		}
	}
}

func BenchmarkGenerateCHeader(b *testing.B) {
	tmpDir := b.TempDir()
	model := createBenchmarkModel(10, 10) // 10 objects with 10 parameters each

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := GenerateCHeader(model, tmpDir)
		if err != nil {
			b.Fatalf("Error in GenerateCHeader: %v", err)
		}
	}
}

func BenchmarkGenerateAll(b *testing.B) {
	tmpDir := b.TempDir()
	goDir := filepath.Join(tmpDir, "go")
	tsDir := filepath.Join(tmpDir, "ts")
	cDir := filepath.Join(tmpDir, "c")

	for _, dir := range []string{goDir, tsDir, cDir} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			b.Fatalf("Failed to create directory %s: %v", dir, err)
		}
	}

	model := createBenchmarkModel(10, 10) // 10 objects with 10 parameters each

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Generate all languages in sequence
		_, err := GenerateGolang(model, goDir)
		if err != nil {
			b.Fatalf("Error in GenerateGolang: %v", err)
		}

		_, err = GenerateTypeScript(model, tsDir)
		if err != nil {
			b.Fatalf("Error in GenerateTypeScript: %v", err)
		}

		_, err = GenerateCHeader(model, cDir)
		if err != nil {
			b.Fatalf("Error in GenerateCHeader: %v", err)
		}
	}
}
