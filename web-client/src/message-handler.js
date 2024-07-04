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
		const board = document.querySelector(".board");

		if (!board) {
			console.error("couldn't find the board element");
			return;
		}

		const width = message[0];
		const height = message[1];

		this.buildBoard(width, height);

		const boardItems = message.slice(2);

		for (let y = 0; y < height; y++) {
			for (let x = 0; x < width; x++) {
				const index = y * width + x;

				const cell = this.getCell(x, y);

				if (!cell) {
					console.error("couldn't find the cell element");
					return;
				}

				if (boardItems[index] === 1) {
					cell.className = "cell-alive";
				} else {
					cell.className = "cell-dead";
				}
			}
		}
	}

	/**
	 * Returns the cell at the given coordinates.
	 *
	 * @param {number} x
	 * @param {number} y
	 * @returns {Element | null}
	 */
	getCell(x, y) {
		return document.querySelector(`[x="${x}"][y="${y}"]`);
	}

	/**
	 * Builds the board with the given width and height.
	 *
	 * @param {number} width
	 * @param {number} height
	 */
	buildBoard(width, height) {
		const board = document.querySelector(".board");

		if (!board) {
			console.error("couldn't find the board element");
			return;
		}

    // clear board content
		while (board.firstChild) {
			board.removeChild(board.firstChild);
		}

		for (let y = 0; y < height; y++) {
			for (let x = 0; x < width; x++) {
				const cell = document.createElement("div");

				cell.className = "cell";

				cell.setAttribute("x", x.toString());
				cell.setAttribute("y", y.toString());

				board.appendChild(cell);
			}
		}
	}
}
