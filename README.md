# 🐰 RabbitMQ

## 📬  O que é o RabbitMQ?
- O RabbitMQ é um Message Broker open-source, em outras palavras, um software de mensageria;
- Ele foi desenvolvido em cima do Erlang, o que traz sua fama de ser extremamente rápido e capaz de suportar alta carga;
- As mensagens caem na memória, ou seja, ainda mais rápido.
- Ele tem como principal protocolo de comunicação o AMQP(_Advanced Message Queueing Protocol_)

## 🏭 Como o RabbitMQ funciona?
Ele abre uma única conexão persistente, e dentro dele encontramos sub-conexões. Seu papel é agir como um intermediador de serviços, em sua maioria tornando o fluxo assíncrono e evitando um alto delay de resposta.

Seu funcionamento básico é composto por:
- Publisher - Serviço que publica a mensagem. Porém, ele não manda a mensagem diretamente ao consumidor e nem mesmo para a fila, do contrário não poderiamos enviar para diversas filas;
- Consumer - Serviço que vai receber/consumir a mensagem;
- Queue - Buffer que irá armazenar as mensagens, ou seja, a mensagem do publisher cai na fila, e conforme forem consumidas saem da fila;
- Exchange - Recebe a mensagem do Publisher, processa e descobre para qual fila deve enviar.

![RabbitMQ flow](/assets/rabbitmq-flow.png)

## 📩 Tipos de Exchange
- **Direct** - Quando eu mando uma mensagem, ela é enviada para um fila específica baseada na `routing key`.
  - Em resumo, a mensagem vai com um `binding key` que dará match com o `routing key` da mensagem.

![Direct exchange](/assets/direct-exchange.png)

- **Fanout** - Quando mando uma mensagem, ela envia esta mensagem para `todas` as filas relacionadas a esta exchange.

![Fanout exchange](/assets/fanout-exchange.png)

- **Topic** - Nesta exchange são definidas algumas regras, e baseando na routing key da mensagem, é encaminhada para fila desejada.

![Topic exchange](/assets/topic-exchange.png)

## Principais propriedades das Queues (filas)
Nota: Segue o conceito de FIFO - _First In, First Out_;

- **Durable**: Se ela deve ser salva mesmo depois do restart do Broker. Por padrão usamos filas duráveis;
- **Auto-delete**: Removida automaticametne quanto o consumer se desconecta;
- **Expiry**: Define o tempo que não há mais mensagens ou cliente consumindo, por exemplo, se dentro de 3 horas ninguém consumiu, a fila é deletada;
- **Message TTL (time to live)**: Tempo de vida da mensagem, se não houve consumo dentro do tempo de vida da mensagem, ela é removida da fila;
- **Overflow**: quando a fila transborda:
    - Drop head (remove a mensagem mais antiga)
    - Reject publish - quando a fila está lotada, o publicador não consegue mais publicar e recebe a mensagem que a fila rejeitou;
- **Exclusive**: Somente channel que criou pode acessar;
- **Max length ou bytes**: Quantidade de mensagens ou o tamanho em bytes máximo permitido;

### 📨 Dead letter queues
- Algumas mensagens não conseguem ser entregues por qualquer motivo;
- São encaminhadas para uma Exchange específica que roteia as mensagens para um dead letter queue;
- Tais mensagens podem ser consumidas e averiguadas posteriormente, por exemplo, suponha que a mensagem seja publicada com expiração de um dia. No dia seguinte se ninguém consumir, ela cairá nessa exchange alternativa que tem uma fila. Então eu posso ter um sistema para tratar essas mensagens que por algum motivo não foi interpretada;

### 🐢 Lazy Queues
- Mensagens são armazenadas em disco, as vezes o fluxo de mensagem é tão grande que os consumers não dão conta de ler essas mensagens e o consumo de memória começa a subir.
- As lazy queues garantem o ritmo dessas leituras e essas mensagens não são perdidas;
- Exige alto consumo de Input/Output, tornando tudo mais custoso.
- Pense sempre se realmente é necessário a utilização das lazy queues;

## Docker
Para subir o ambiente a partir do `docker-compose.yaml` deste serviço, disponibilize as portas `15672` e `5672`. Em seguida execute dentro deste repositório (caminho do docker-compose) o seguinte comando:
```
docker-compose up -d
```
Abra no browser `http://localhost:15672` e preencha o usuário e senha como `admin`

### Publicando mensagem manualmente pelo Docker
Nota: Certifique-se de ter uma fila com o nome `TestQueue`, o exemplo é baseado no docker-compose deste repositório.
```
docker exec -it rabbitmq bash
rabbitmqadmin -u docker -p docker publish routing_key="TestQueue" payload="Ola Mundo!"
```