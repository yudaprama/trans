function decrementChar(c) {
	return String.fromCharCode(c.charCodeAt(0) - 1);
}

function incrementChar(c) {
	return String.fromCharCode(c.charCodeAt(0) + 1);
}

function makePalindrome(input) {
	let steps = 0;
	for (let i = 0; i < input.length / 2; i++) {
		let firstLetter = input[i];
		let lastLetter = input[input.length - i - 1];
		while (firstLetter !== lastLetter) {
			steps++;
			if (firstLetter.localeCompare(lastLetter) === -1) {
				lastLetter = decrementChar(lastLetter);
			} else {
				lastLetter = incrementChar(lastLetter);
			}
		}
	}
	return steps;
}

console.log(`abc > ${makePalindrome("abc")} steps`);
console.log(`abcde > ${makePalindrome("abcde")} steps`);
console.log(`aaa > ${makePalindrome("aaa")} steps`);

// Run with `node palindrome.js`