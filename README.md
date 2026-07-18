# Go Calculator

A simple command-line calculator built with Go that performs basic arithmetic operations. This project was created to strengthen my understanding of Go fundamentals, user input handling, conditional statements, and functions.

## Features

* Addition (`+`)
* Subtraction (`-`)
* Multiplication (`*`)
* Division (`/`)
* Prevents division by zero with an appropriate error message
* Interactive command-line interface

## Technologies Used

* Go (Golang)

## Getting Started

### Clone the repository

```bash
git clone https://github.com/olaiyadigitech/calculator.git
```

### Navigate to the project

```bash
cd calculator
```

### Run the application

```bash
go run .
```

## Example

### Valid Calculation

```text
Enter first number: 20
Enter operation (+, -, *, /): /
Enter second number: 4

Result: 5
```

### Division by Zero

```text
Enter first number: 20
Enter operation (+, -, *, /): /
Enter second number: 0

Error: Division by zero is not allowed.
```

## Concepts Practiced

* Variables
* User input (`fmt.Scan`)
* Conditional statements (`switch`)
* Arithmetic operations
* Error handling
* Basic Go project structure

## Future Improvements

* Support advanced mathematical operations
* Add command history
* Improve input validation
* Build a graphical user interface (GUI)
* Expose the calculator as a REST API

## Author

**Ajayi Oluwafemi Olaiya**

* GitHub: https://github.com/olaiyadigitech
