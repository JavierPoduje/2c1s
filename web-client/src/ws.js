// TODO: use dotenv instead of hardcoded values
// require("dotenv").config({
//   path: `${__dirname}/../.env`,
// });
import MessageHandler from "./message-handler.js";

/** @type {number} */
// const WEB_PORT = parseInt(process.env.WEB_PORT || "8080");
/** @type {string} */
// const SERVER_HOST = process.env.SERVER_HOST || "127.0.0.1";

/** @type {WebSocket} */
// const ws = new WebSocket(`ws://${SERVER_HOST}:${WEB_PORT}`);
const ws = new WebSocket(`ws://127.0.0.1:8080`);
const messageHandler = new MessageHandler();

ws.onopen = () => {
	console.log("Connected to server");
};

/**
 * Receives an event from the WebSocket connection. This event contains the
 * message sent by the server, which is a Uint8Array with the following format:
 *
 * - first byte is the width of the message
 * - second byte is the heigth of the message
 * - remaining bytes are the board
 *
 * @param {MessageEvent<Blob>} event
 */
ws.onmessage = async (event) => {
	const bytes = await event.data.arrayBuffer();
	const uint8 = new Uint8Array(bytes);
	messageHandler.handle(uint8);
};

/** @param {Event} error */
ws.onerror = (error) => {
	console.error("WebSocket error:", error);
};

ws.onclose = () => {
	console.log("Disconnected from server");
};
