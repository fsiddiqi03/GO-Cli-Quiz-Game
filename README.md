# Go Cli Quiz Game

A simple command-line quiz game built with Go that reads questions and answers from a CSV file and quizzes the user interactively.


## Description

**GO-Cli-Quiz-Game** is a command-line application that quizzes users on various topics. It reads a CSV file containing questions and answers, presents each question to the user, and checks the user's answers against the correct ones. It’s a fun way to test your knowledge on any subject you choose!

## Features

- Read questions from a CSV file
- Display questions interactively
- Validate user answers
- Display the number of correct answers at the end
- Easily extendable with different sets of questions

## Installation

To run this project locally, you'll need Go installed on your machine. If Go is not installed, follow the instructions [here](https://golang.org/doc/install).

### Steps:

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/GO-Cli-Quiz-Game.git
    cd GO-Cli-Quiz-Game
    ```

2. Build the project:

    ```bash
    go build .
    ```

## Usage

1. Ensure you have a CSV file with your quiz questions (example: `problems.csv`).

2. Run the quiz game:

    ```bash
    ./GO-Cli-Quiz-Game -csv=path/to/your/problems.csv
    ```

   For example:

    ```bash
    ./GO-Cli-Quiz-Game -csv=problems.csv
    ```

3. Answer the questions as they are presented. Your score will be displayed at the end.

### Options

- `-csv`: Path to the CSV file containing questions and answers.

## CSV Format

The CSV file should have questions and answers in the following format:

```csv
5+5,10
7+3,10
1+1,2
8-3,5
```
