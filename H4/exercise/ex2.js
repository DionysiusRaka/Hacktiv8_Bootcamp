// Exercise 1

compareNumbers(5, 8);
compareNumbers(5, 3);
compareNumbers(4, 4);
compareNumbers(3, 3);
compareNumbers(17, 2);

function compareNumbers(firstNum, SecNum) {
    if (SecNum > firstNum) {
        console.log("true")
    }
    else if (SecNum < firstNum) {
        console.log("false")
    }
    else console.log("-1")
}

// Exercise 2
reverseWord("ABCD");
reverseWord("Asep");


function reverseWord(inputWord) {
    var result = "";

    for (var i = inputWord.length - 1; i >= 0; i--) {
        result += inputWord[i];
    }
    if (result == inputWord) {
        result = "Palindrome";
    }

    console.log(result);
}

// Exercise 3
urutHuruf("aqpwirut");

urutHuruf("qwertyuiopasdfghjklzxcvbnm");
function urutHuruf(inputWord) {
    console.log(inputWord.split('').sort().join(''));
}

// Exercise 4
console.log(isArithmeticProgression([1,3,5,7,9]));
console.log(isArithmeticProgression([2,4,6,8,10]));
console.log(isArithmeticProgression([1, 2, 4, 3, 9, 3]));

function isArithmeticProgression(numbers) {
    const selisih = numbers[1] - numbers[0];
    for (var i = 1; i < numbers.length-1; i++) {
        if (numbers[i + 1] - numbers[i] !== selisih) {
            return false;
        }
    }
    return true
}

// Exercise 5

