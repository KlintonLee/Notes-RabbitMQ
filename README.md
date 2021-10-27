# RabbitMQ com Golang

A POC ainda está incompleta, porém, já existe implementação da ferramenta.

## Ferramentas necessárias para testa-lo
- Docker & docker-compose - Se for subir o container do RabbitMQ pelo docker-compose.yaml do repositório.
- Golang instalado

## 👨🏻‍💻 Como utilizar?
- Primeiro é necessário que o rabbitmq esteja executando, com o usuário e senha `admin`;
  - Se for utilizar o docker-compose.yaml basta executar: `docker-compose up -d`
- Baixar as dependências utilizas com o comando `go mod vendor`
- Compilar e executar o arquivo de entrada: `go run main.go`

## Fluxo da aplicação
- A aplicação tentará se conectar ao RabbitMQ;
- Tentará criar um canal de comunicação;
- Na inexistência da fila `TestQueue`, será criado;
- Publicará a mensagem `Hello World`
- E será exibido no terminal a mensagem consumida `Hello World`

Obs: O Consumer está observando todas mensagens que caem na fila e exibindo no terminal sempre que uma nova mensagem cai na fila.

Para efetuar um teste, é possível acessar o RabbitMQ pelo `http://localhost:15672` e publicar uma nova mensagem, ou, acessar o container e publicar uma mensagem manualmente pela linha de comando:
```
docker exec -it rabbitmq bash
rabbitmqadmin -u admin -p admin publish routing_key="TestQueue" payload="Ola Mundo!"
```