const globals = require("globals");

module.exports = [
	{
		rules: {
			"no-unused-vars": "error",
			"no-undef": "warn",
		},
		files: ["**/*.js"],
		languageOptions: {
			globals: {
				...globals.browser,
				...globals.jquery,
				...globals.node,
			},
		},
	},
];
