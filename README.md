# Terratest - Demonstração Completa

[![Terratest CI](https://github.com/nataliagranato/terratest/actions/workflows/terratest.yml/badge.svg)](https://github.com/nataliagranato/terratest/actions/workflows/terratest.yml)

Este repositório é uma demonstração completa de como utilizar o **Terratest** para testar infraestrutura como código (IaC) escrita em Terraform, incluindo sua integração com GitHub Actions para testes automatizados em CI/CD.

## 📚 O que é Terratest?

**Terratest** é uma biblioteca Go que fornece padrões e funções auxiliares para testar infraestrutura, com suporte de primeira classe para Terraform, Packer, Docker, Kubernetes e muito mais.

Desenvolvido pela [Gruntwork](https://gruntwork.io/), o Terratest permite escrever testes automatizados para validar que sua infraestrutura funciona conforme esperado antes de implantá-la em produção.

### Características principais:

- 🧪 **Testes automatizados** para código Terraform
- 🔄 **Execução de ciclo completo**: init, plan, apply, destroy
- ✅ **Validação de outputs** e recursos criados
- 🔁 **Retry automático** para operações que podem falhar temporariamente
- 🎯 **Suporte a múltiplos provedores** cloud (AWS, Azure, GCP, etc.)
- 📊 **Integração com frameworks de teste** Go (testing, testify)

## 🎯 Por que Terratest é importante?

### 1. **Confiabilidade**
- Detecta erros antes que o código chegue à produção
- Valida que os recursos são criados corretamente
- Testa diferentes cenários e configurações

### 2. **Qualidade**
- Garante que mudanças não quebrem funcionalidades existentes
- Documenta o comportamento esperado da infraestrutura
- Facilita refatorações com segurança

### 3. **Automação**
- Integra-se facilmente com pipelines CI/CD
- Executa testes automaticamente a cada mudança
- Reduz testes manuais e erros humanos

### 4. **Economia**
- Identifica problemas cedo no ciclo de desenvolvimento
- Reduz custos de correção de bugs em produção
- Otimiza o uso de recursos cloud

## 📋 Pré-requisitos

Para executar este projeto, você precisará ter instalado:

- **Go** 1.21 ou superior ([Download](https://golang.org/dl/))
- **Terraform** 1.0 ou superior ([Download](https://www.terraform.io/downloads))
- **Git** para clonar o repositório

### Instalação do Go

```bash
# Linux/macOS
wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

# Verificar instalação
go version
```

### Instalação do Terraform

```bash
# Linux/macOS
wget https://releases.hashicorp.com/terraform/1.6.0/terraform_1.6.0_linux_amd64.zip
unzip terraform_1.6.0_linux_amd64.zip
sudo mv terraform /usr/local/bin/

# Verificar instalação
terraform version
```

## 🚀 Como usar este repositório

### 1. Clone o repositório

```bash
git clone https://github.com/nataliagranato/terratest.git
cd terratest
```

### 2. Instale as dependências Go

```bash
go mod download
```

### 3. Execute os testes

```bash
cd test
go test -v -timeout 30m
```

## 📁 Estrutura do Projeto

```
terratest/
├── .github/
│   └── workflows/
│       └── terratest.yml       # GitHub Actions workflow
├── test/
│   └── terraform_example_test.go  # Testes Terratest
├── main.tf                     # Código Terraform principal
├── variables.tf                # Variáveis do Terraform
├── outputs.tf                  # Outputs do Terraform
├── go.mod                      # Dependências Go
└── README.md                   # Esta documentação
```

## 🧪 Exemplos de Testes

### Teste Básico

O arquivo `test/terraform_example_test.go` contém dois exemplos de testes:

#### 1. Teste Simples
```go
func TestTerraformExample(t *testing.T) {
    terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
        TerraformDir: "../",
        Vars: map[string]interface{}{
            "content":     "Hello from Terratest!",
            "filename":    "test-output.txt",
            "environment": "testing",
        },
    })

    defer terraform.Destroy(t, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)

    // Valida outputs
    filePath := terraform.Output(t, terraformOptions, "file_path")
    assert.Contains(t, filePath, "test-output.txt")
}
```

#### 2. Teste com Múltiplos Ambientes
```go
func TestTerraformWithDifferentEnvironments(t *testing.T) {
    environments := []string{"dev", "staging", "prod"}
    
    for _, env := range environments {
        t.Run(env, func(t *testing.T) {
            // Testa cada ambiente em paralelo
        })
    }
}
```

## 🔄 Integração com GitHub Actions

Este repositório inclui um workflow completo de CI/CD que:

1. ✅ Instala Go e Terraform
2. 📦 Faz cache das dependências
3. 🧪 Executa todos os testes Terratest
4. 🧹 Limpa recursos automaticamente

### Workflow Example

```yaml
name: Terratest CI

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  terratest:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v4
      - uses: hashicorp/setup-terraform@v3
      - run: go mod download
      - run: cd test && go test -v -timeout 30m
```

## 🔧 Código Terraform de Exemplo

Este repositório usa um exemplo simples com o provider `local` para demonstrar conceitos sem necessidade de credenciais cloud:

- **main.tf**: Cria arquivos locais com conteúdo configurável
- **variables.tf**: Define variáveis para customização
- **outputs.tf**: Exporta valores para validação nos testes

### Executar manualmente

```bash
# Inicializar Terraform
terraform init

# Ver o plano
terraform plan

# Aplicar as mudanças
terraform apply -auto-approve

# Destruir recursos
terraform destroy -auto-approve
```

## 📊 Boas Práticas

### 1. **Isolamento de Testes**
```go
t.Parallel() // Executa testes em paralelo
```

### 2. **Cleanup Automático**
```go
defer terraform.Destroy(t, terraformOptions)
```

### 3. **Tratamento de Erros Temporários**
```go
terraform.WithDefaultRetryableErrors(t, options)
```

### 4. **Timeouts Adequados**
```go
go test -v -timeout 30m
```

### 5. **Variáveis Parametrizadas**
```go
Vars: map[string]interface{}{
    "environment": "testing",
}
```

## 🌟 Casos de Uso Reais

Embora este exemplo use o provider `local`, o Terratest é amplamente usado para testar:

- **AWS**: EC2, S3, RDS, Lambda, ECS, etc.
- **Azure**: VMs, Storage, AKS, Functions, etc.
- **GCP**: Compute Engine, GCS, GKE, Cloud Functions, etc.
- **Kubernetes**: Deployments, Services, ConfigMaps, etc.
- **Docker**: Containers, Networks, Volumes, etc.

## 📚 Recursos Adicionais

- [Documentação Oficial Terratest](https://terratest.gruntwork.io/)
- [Repositório GitHub Terratest](https://github.com/gruntwork-io/terratest)
- [Exemplos Terratest](https://github.com/gruntwork-io/terratest/tree/master/examples)
- [Blog Gruntwork](https://blog.gruntwork.io/)
- [Terraform Testing Best Practices](https://www.terraform.io/docs/language/modules/testing.html)

## 🤝 Contribuindo

Contribuições são bem-vindas! Sinta-se à vontade para:

1. Fazer fork do projeto
2. Criar uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abrir um Pull Request

## 📝 Licença

Este projeto é open source e está disponível para fins educacionais.

## 👤 Autor

**Natalia Granato**

- GitHub: [@nataliagranato](https://github.com/nataliagranato)

## ⭐ Mostre seu apoio

Se este projeto foi útil para você, considere dar uma ⭐️!

---

**Nota**: Este é um projeto de demonstração educacional. Para uso em produção, adapte os exemplos às suas necessidades específicas e sempre siga as melhores práticas de segurança.