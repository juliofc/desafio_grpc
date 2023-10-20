# Desafio gRPC

Projeto criado para o desafio de gRPC

## Como executar:

**Passo 1:** Faça o clone do projeto no diretório de sua preferência.

```shell
git clone github.com/juliofc/desafio_grpc
```

**Passo 2:** Acesse o diretório do projeto e suba o servidor

```shell
go run cmd/main.go
```

**Passo 3:** Execute o comando abaixo:

```shell
evans -r repl
```

**Passo 5:** Para criar um novo produto digite o comando abaixo

```shell
call CreateProduct
```

Preencha os campos solicitados

**Passo 6:** Para consultar os produtos criados, digite o comando abaixo

```shell
call FindProducts
```

**Observação:** Antes de executar as chamadas dos métodos certifique de estar no pacote github.com.codeedu.codepix e no serviço ProductService.

Caso não esteja, execute os comandos abaixo:

```shell
package github.com.codeedu.codepix
```

```shell
service ProductService
```
