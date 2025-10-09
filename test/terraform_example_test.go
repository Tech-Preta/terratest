package test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformExample(t *testing.T) {
	t.Parallel()

	// Configuração das opções do Terraform
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Caminho para o código Terraform
		TerraformDir: "../",

		// Variáveis para passar ao código Terraform
		Vars: map[string]interface{}{
			"content":     "Hello from Terratest! This is a test.",
			"filename":    "test-output.txt",
			"environment": "testing",
		},
	})

	// Limpa os recursos no final do teste
	defer terraform.Destroy(t, terraformOptions)

	// Executa terraform init e terraform apply
	terraform.InitAndApply(t, terraformOptions)

	// Testa os outputs do Terraform
	filePath := terraform.Output(t, terraformOptions, "file_path")
	fileContent := terraform.Output(t, terraformOptions, "file_content")
	environment := terraform.Output(t, terraformOptions, "environment")

	// Validações
	assert.Contains(t, filePath, "test-output.txt")
	assert.Equal(t, "Hello from Terratest! This is a test.", fileContent)
	assert.Equal(t, "testing", environment)

	// Verifica se o arquivo foi criado no sistema de arquivos
	absPath, err := filepath.Abs(filepath.Join("../", filePath))
	assert.NoError(t, err)

	// Verifica se o arquivo existe
	_, err = os.Stat(absPath)
	assert.NoError(t, err, "O arquivo deveria existir")

	// Lê o conteúdo do arquivo e verifica
	actualContent, err := os.ReadFile(absPath)
	assert.NoError(t, err)
	assert.Equal(t, "Hello from Terratest! This is a test.", string(actualContent))
}

func TestTerraformWithDifferentEnvironments(t *testing.T) {
	t.Parallel()

	environments := []string{"dev", "staging", "prod"}

	for _, env := range environments {
		// Captura o valor de env para evitar problemas com closures
		env := env

		t.Run(env, func(t *testing.T) {
			t.Parallel()

			terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
				TerraformDir: "../",
				Vars: map[string]interface{}{
					"content":     fmt.Sprintf("Content for %s environment", env),
					"filename":    fmt.Sprintf("%s-output.txt", env),
					"environment": env,
				},
			})

			defer terraform.Destroy(t, terraformOptions)

			terraform.InitAndApply(t, terraformOptions)

			// Valida o output do ambiente
			outputEnv := terraform.Output(t, terraformOptions, "environment")
			assert.Equal(t, env, outputEnv)

			// Valida o conteúdo
			fileContent := terraform.Output(t, terraformOptions, "file_content")
			assert.True(t, strings.Contains(fileContent, env))
		})
	}
}
