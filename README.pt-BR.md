# Citassauro 🦕

Uma API simples e robusta para escrita e gerenciamento de citações, desenvolvida com Golang.

## 📦 Sobre

Citassauro é um projeto open source focado em fornecer uma API leve e eficiente para manipulação de citações — ideal para testes, aprendizado ou como base para projetos maiores.  
Construído com propósito educacional e para compor um portfólio profissional.

## 🚀 Tecnologias

- [Go (Golang)](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin) — framework HTTP
- [GORM](https://gorm.io/) — ORM para banco de dados
- SQLite (por padrão, mas adaptável para outros bancos)
- Docker (opcional para produção)

## 🔧 Funcionalidades

- 🔹 Criar novas citações
- 🔹 Listar todas as citações
- 🔹 Buscar por ID
- 🔹 Deletar citações
- 🔹 Atualizar citações (em breve)

## ⚙️ Como rodar

```bash
# Clone o repositório
git clone https://github.com/walacysilvam/citassauro.git
cd citassauro

# Instale as dependências
go mod tidy

# Rode a aplicação localmente
go run main.go