# Go Algoritmos

Repositório destinado ao catálogo de altoritmos utilizados no dia a dia.

## Download, instalação e configuração

-   **Download oficial:** https://go.dev/dl/

```bash
# Remove o diretorio se existir, e extrai o arquivo diretório /usr/local
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.20.6.linux-amd64.tar.gz

# Cria a pasta onde irá ficar os pacotes
mkdir $HOME/go

# Edita o zshrc
nano ~/.zshrc

# Insere
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin

# Atualiza o zshrc
source ~/.zshrc

# Verifica a versão
go version
```

## Comandos

```bash
# Exibe a versão do go
go version

# Exibe as variáveis do go
go env

# Inicializa um módulo
go mod init <nome_do_modulo>

# Baixa as dependências de um projeto
go mod tidy

# Cria o build da aplicação de acordo com a variável GOOS do go env
go build

# Cria um build para windows
GOOS=windows go build
```

## Pacotes

```bash
# UUID
go get github.com/google/uuid
```

## Links

-   [ViaCEP](https://viacep.com.br/)
-   [Converte JSON to struct](https://mholt.github.io/json-to-go/)
