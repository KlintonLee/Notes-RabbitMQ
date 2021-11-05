const amqp = require('amqplib');

const run = async () => {
  const amqpServer = 'amqp://admin:admin@localhost:5672';
  const connection = await amqp.connect(amqpServer);
  const channel = await connection.createChannel();
  await channel.assertQueue('TestQueue');
  console.log(' [*] Waiting for messages in %s. To exit press CTRL+C', 'TestQueue');
  channel.consume('TestQueue', data => {
    console.log(`Received data: ${Buffer.from(data.content)}`);
    channel.ack(data);
  })
}

run();