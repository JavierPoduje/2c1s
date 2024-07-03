export default class MessageHandler {
	width = 0;
	height = 0;

	contructor() {
		this.width = 0;
		this.height = 0;
	}

	/**
	 * Handles a message received from the server. This message is a Uint8Array
	 * with the following format:
	 *
	 * - first byte is the width of the message
	 * - second byte is the heigth of the message
	 * - remaining bytes are the board
	 *
	 * @param {Uint8Array} message
	 */
	handle(message) {
		console.log("message: ", message);
		const messagesDiv = document.getElementById("messages");
		const messageElement = document.createElement("p");
		messageElement.textContent = `Received: ${message}`;
		if (!messagesDiv) {
			console.error("Could not find messagesDiv");
			return;
		}

		messagesDiv.appendChild(messageElement);
	}
}
