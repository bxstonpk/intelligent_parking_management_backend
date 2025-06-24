require('dotenv').config();
const express = require('express');
const http = require('http');
const WebSocket = require('ws');

const PORT = process.env.PORT || 3000;
const WS_PATH = process.env.WS_PATH || '/ws';

const app = express();
const server = http.createServer(app);
const wss = new WebSocket.Server({ server, path: WS_PATH });

let clients = new Set();

wss.on('connection', (ws) => {
  console.log('Client connected');
  clients.add(ws);

  ws.on('message', (data, isBinary) => {
    // ส่งต่อเฉพาะ binary (ภาพ)
    if (isBinary) {
      for (const client of clients) {
        if (client !== ws && client.readyState === WebSocket.OPEN) {
          client.send(data, { binary: true });
        }
      }
    }
  });

  ws.on('close', () => {
    console.log('Client disconnected');
    clients.delete(ws);
  });
});

app.get('/', (req, res) => {
  res.send('Stream Server Running');
});

server.listen(PORT, () => {
  console.log(`Server running at http://localhost:${PORT}`);
  console.log(`WebSocket path: ws://localhost:${PORT}${WS_PATH}`);
});
