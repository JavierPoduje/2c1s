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
		const width = message[0];
		const height = message[1];
		const boardFromMessage = message.slice(2);

		this.buildBoard(width, height);
		this.colorizeBoard(boardFromMessage, width, height);
	}

	/**
	 * Colorizes the board based on the given board items.
	 *
	 * @param {Uint8Array} boardItems
	 * @param {number} width
	 * @param {number} height
	 */
	colorizeBoard(boardItems, width, height) {
		for (let y = 0; y < height; y++) {
			for (let x = 0; x < width; x++) {
				const index = y * width + x;

				const cell = this.getCell(x, y);
				if (!cell) {
					console.error("couldn't find the cell element");
					return;
				}

				const isAlive = boardItems[index] === 1;
				cell.className = `cell-${isAlive ? "alive" : "dead"}`;
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
		const board = this.getBoard();
		if (!board) {
			console.error("couldn't find the board element");
			return;
		}

		this.clearBoard();
		this.fillBoard(width, height);
	}

	/**
	 * Get the board element
	 *
	 * @returns {Element | null}
	 */
	getBoard() {
		return document.querySelector(".board");
	}

	/**
	 * Fill the board with "neutral" cells (they aren't alive nor dead just yet)
	 *
	 * @param {number} width
	 * @param {number} height
	 */
	fillBoard(width, height) {
		const board = document.querySelector(".board");
		if (!board) {
			console.error("couldn't find the board element");
			return;
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

	/**
	 * Remove all the children of the board
	 */
	clearBoard() {
		const board = document.querySelector(".board");
		if (!board) {
			console.error("couldn't find the board element");
			return;
		}

		// clear board content
		while (board.firstChild) {
			board.removeChild(board.firstChild);
		}
	}
}
