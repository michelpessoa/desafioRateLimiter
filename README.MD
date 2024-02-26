# Go Expert - Desfio Rate Limiter

Go | REST | Wire | Docker | Redis


### Overview do projeto 

A ferramenta foi projetada para controlar a taxa de solicitações recebidas em um servidor web.
Funciona analisando tokens e/ou endereços IP e aplicando políticas de limitação de taxa com base na configuração fornecida.

O limitador realiza a análise com base no token fornecido no cabeçalho da solicitação com a chave API_KEY.
Se for encontrada uma configuração para esse token, os limites e o tempo de bloqueio especificados serão aplicados.

Se nenhum token for passado ou o token diferente do que foi configuração, os valores padrão serão aplicados com base no endereço IP da solicitação.

### Alteração da configuração padrão

Caso queira alterar a configuração padrão, deve-se alterar os valores no arquivo .env na pasta cmd/ratelimite

- Define o número máximo de requisiçõe liberadas: 
MAX_REQUESTS=10  
- Token definido com o número máximo de requisições para requisições desse tipo:
TOKENS=[{"name": "token123", "max_requests": 20, "cooldown_seconds": 5}]

- Quantidade de tempo da janela de limite:
TTL_SECONDS=10

- Quantidade de tempo de bloqueio
COOLDOWN_SECONDS=5


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

### Para visualizar o resultados dos teste

1. `abra um terminal e execute os seguintes comandos`

    docker build -t go-docker-test .
    
    docker run -d --name testes go-docker-test

    docker logs testes