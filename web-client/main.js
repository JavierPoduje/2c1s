const express = require("express");
const http = require("http");
const WebSocket = require("ws");

const app = express();
const server = http.createServer(app);
const wss = new WebSocket.Server({ server });

const port = 8080;
const wsPort = 12345;

// Serve static files from the current directory
app.use(express.static(__dirname));

// WebSocket connection handler
wss.on("connection", (ws) => {
	console.log("New WebSocket client connected");

	// Relay messages from TCP server to WebSocket clients
	const net = require("net");
	const tcpClient = new net.Socket();

	tcpClient.connect(wsPort, "127.0.0.1", () => {
		console.log("Connected to TCP server");
	});

	tcpClient.on("data", (data) => {
		console.log("Received from TCP server:", data.toString());
		ws.send(data.toString());
	});

	tcpClient.on("close", () => {
		console.log("Connection to TCP server closed");
		ws.close();
	});

	tcpClient.on("error", (err) => {
		console.error("TCP error:", err);
		ws.close();
	});

	ws.on("close", () => {
		console.log("WebSocket client disconnected");
		tcpClient.end();
	});

	ws.on("error", (err) => {
		console.error("WebSocket error:", err);
		tcpClient.end();
	});
});

// Start the server
server.listen(port, () => {
	console.log(`Server is listening on http://localhost:${port}`);
});
