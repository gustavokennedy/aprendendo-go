FROM golang:latest

# Define o diretório atual dentro do container
WORKDIR /app

RUN GO111MODULE=on

# Copia o go mod e o aquivo sum 
COPY go.mod go.sum ./

# Baixa todas as dependências
RUN go mod download

COPY . . 

# Builda a aplicação
RUN go build -o main .

# Expoe a porta 5000
EXPOSE 5000

# Comando para rodar o executável
CMD ["./main"]