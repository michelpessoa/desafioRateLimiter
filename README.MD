# Go Expert - Desfio Rate Limiter

Go | REST | Wire | Docker | Redis


### Para executar o projeto sega os seguintes passos:


1. `git clone https://github.com/michelpessoa/desafioRateLimiter`
2. `go mod tidy` para instalar todas as dependências
3. `docker-compose up` tenho o docker instalado, para iniciar os serviços redis
4. `go run main.go wire_gen.go` execute o comando para iniciar os serviços dentro da pasta cmd/ratelimiter


### Para testar o rate limiter por IP

1. `abra um terminal e execute o seguinte comando abaixo: `

    for i in \`seq 1 60 \` ; do 
curl --request GET --url http://localhost:8080/; done

- este comando irá executar 60 chamadas simultaneamente ao servidor e deverá receber 10 repostas iniciais OK e o restante o bloqueio, aguarde 5 segundos que repita o processo, onde ser receberá o mesmo resultado 


### Para testar o rate limiter por Token

1. `abra um terminal e execute o seguinte comando abaixo: `

    for i in \`seq 1 60 \` ; do 
curl --request GET --url http://localhost:8080/ --header 'API_KEY: token123'; done

- este comando irá executar 60 chamadas simultaneamente ao servidor e deverá receber 10 repostas iniciais OK e o restante o bloqueio, aguarde 5 segundos que repita o processo, onde ser receberá o mesmo resultado


### Para visualizar o registro bloqueados no REDIS

1. `abra um novo terminal enquanto o bloqueio é executado e execute o seguinte comando abaixo: `

docker exec -it redis bash

2. `redis-cli` para acessar o redis

3. `keys *` para listar os registros no redis

- será apresentado um bloqueio por IP ou por token

