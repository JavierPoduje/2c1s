require("dotenv").config({
	path: `${__dirname}/../.env`,
});

/** @type {number} */
const WEB_PORT = parseInt(process.env.WEB_PORT || "8080");
/** @type {string} */
const SERVER_HOST = process.env.SERVER_HOST || "127.0.0.1";

/** @type {WebSocket} */
const ws = new WebSocket(`ws://${SERVER_HOST}:${WEB_PORT}`);

ws.onopen = () => {
	console.log("Connected to server");
};

// TODO: define the structue of the MessageEvent object
/** @param {MessageEvent} event */
ws.onmessage = (event) => {
	console.log("Received:", event.data);
	const messagesDiv = document.getElementById("messages");
	const messageElement = document.createElement("p");
	messageElement.textContent = `Received: ${event.data}`;

	if (!messagesDiv) {
		console.error("Could not find messagesDiv");
		return;
	}

	messagesDiv.appendChild(messageElement);
};

/** @param {Event} error */
ws.onerror = (error) => {
	console.error("WebSocket error:", error);
};

ws.onclose = () => {
	console.log("Disconnected from server");
};
