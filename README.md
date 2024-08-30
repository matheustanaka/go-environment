# Desafio: Configuracão de Ambiente com Docker Compose

Realizacão do desafio de criacão de ambiente de com Docker compose

# Estrutura do projeto

```sh
├── docker-compose.yaml # Criacão de multiplos containers ao mesmo tempo
├── Dockerfile # Criacão de imagem customizada
├── Dockerfile.test # Meu primeiro dockerfile em golang, foi apenas um teste, deixei aqui para comparar no futuro
├── go.mod # Arquivo que contém a versão de dependências
├── go.sum # Armazena checksums para todas as versões de dependências
├── main.go # Código principal
```

# Como executar o código em seu ambiente local

```sh
git clone

cd project

docker-compose -up --build -d
```
