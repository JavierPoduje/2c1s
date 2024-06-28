const express = require("express");
const http = require("http");
const WebSocket = require("ws");
require("dotenv").config({
	path: `${__dirname}/../.env`,
});

const app = express();
const server = http.createServer(app);
const wss = new WebSocket.Server({ server });

/** @type {number} */
const WEB_PORT = parseInt(process.env.WEB_PORT || "8080");
/** @type {number} */
const SERVER_PORT = parseInt(process.env.SERVER_PORT || "12345");
/** @type {string} */
const SERVER_HOST = process.env.SERVER_HOST || "127.0.0.1";

// Serve static files from the current directory
app.use(express.static(__dirname));

// WebSocket connection handler
wss.on("connection", (ws) => {
	console.log("New WebSocket client connected");

	// Relay messages from TCP server to WebSocket clients
	const net = require("net");
	const tcpClient = new net.Socket();

	tcpClient.connect(SERVER_PORT, SERVER_HOST, () => {
		console.log("Connected to TCP server");
	});

	tcpClient.on("data", (data) => {
		const bytes = new Uint8Array(data);
		console.log("Received from TCP server:", bytes);
		ws.send(data);
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
server.listen(WEB_PORT, () => {
	console.log(`Server is listening on http://${SERVER_HOST}:${WEB_PORT}`);
});
