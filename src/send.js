const amqp = require('amqplib');

const run = async () => {
  const amqpServer = 'amqp://admin:admin@localhost:5672';
  const connection = await amqp.connect(amqpServer);
  const channel = await connection.createChannel();
  await channel.assertQueue('TestQueue', { durable: true });

  const testData = JSON.stringify({ hello: "world" });
  channel.sendToQueue('TestQueue', Buffer.from(testData));
  await channel.close()
  await connection.close()
}

run()