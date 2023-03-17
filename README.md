pixGo-CLI é uma CLI para gerenciar e armazenar chaves PIX em um banco de dados local SQLite.

## Dependências

O projeto utiliza as seguintes bibliotecas:

- [Cobra](https://github.com/spf13/cobra) 
- [GORM](https://gorm.io/) 



## Build

```bash
go build -o pixGo-cli
````

## Usage

##### Adicionar uma chave PIX:
```bash
./pixGo-cli add -k 12345678901 -a exemplo
````

##### Listar todas as chaves PIX::
```bash
./pixGo-cli list
````