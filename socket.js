const WebSocket = require('ws');
const express = require('express');
const http = require('http');

const app = express();
const server = http.createServer(app);
const wss = new WebSocket.Server({ server });

wss.on('connection', function connection(ws) {
    console.log('Client connected');

    ws.on('message', (data) => {
        // ส่งข้อมูลวิดีโอไปยัง Client อื่นที่กำลังเชื่อมต่อ
        wss.clients.forEach((client) => {
            if (client !== ws && client.readyState === WebSocket.OPEN) {
                client.send(data);
                console.log(data);
            }
        });
    });

    ws.on('close', () => {
        console.log('Client disconnected');
    });
});

server.listen(3000, () => {
    console.log('WebSocket server running on ws://localhost:3000');
});