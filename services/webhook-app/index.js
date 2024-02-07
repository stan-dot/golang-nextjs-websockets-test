
const express = require('express');

const app = express();

app.use(express.json());

app.post('/webhook', (req, res) => {
  const webhookPayload = req.body;
  const receivedSignature = req.headers['tally-signature'];

  // Replace 'YOUR_SIGNING_SECRET' with your signing secret
  const yourSigningSecret = 'YOUR_SIGNING_SECRET';

  // Calculate the signature using the signing secret and the payload
  const calculatedSignature = createHmac('sha256', yourSigningSecret)
    .update(JSON.stringify(webhookPayload))
    .digest('base64');

  // Compare the received signature with the calculated signature
  if (receivedSignature === calculatedSignature) {
    // Signature is valid, process the webhook payload
    res.status(200).send('Webhook received and processed successfully.');
  } else {
    // Signature is invalid, reject the webhook request
    res.status(401).send('Invalid signature.');
  }
});

app.get('/', (req, res) => {
  res.status(200).send('hello world');
})

app.listen(3000, () => console.log('Server is running on port 3000'));