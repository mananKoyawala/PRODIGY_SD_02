# Guess the Number Game

## Description

This is a program that challenges the user to guess a randomly generated number between 1 and 100. The user is prompted to input their guess, and the program provides feedback on whether the guess is too high or too low. The game continues until the user correctly guesses the number, and it displays the number of attempts it took to win the game.

## Features

- Random number generation between 1 and 100.
- User input validation to ensure proper numerical input.
- Feedback provided if the guess is too high, too low, or close to the correct number.
- Displays the total number of attempts made once the correct number is guessed.

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/mananKoyawala/PRODIGY_SD_02.git
   cd PRODIGY_SD_02
   ```

2. **Run the application:**
   ```bash
   go run main.go
   ```

## Usage

1. Run the application.
2. You will be presented with a menu that introduces the game and explains the rules.
3. Enter your guesses when prompted.
4. The application will provide feedback if your guess is too high, too low, or close to the correct number.
5. Continue guessing until you find the correct number.
   The application will display the total number of attempts made once the correct number is guessed.

## Example

```bash
Guess the Number
Guess the number between 1-100

Enter integer number
50

guessed number is too high then generate number

Let's guess again

Enter integer number
25

guessed number is too low then generate number

Let's guess again

Enter integer number
37

you guessed it correct.
Numbers of attempts took to guess number is 3
```

## Contributing

Feel free to fork this project and submit pull requests. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT License](LICENSE)
