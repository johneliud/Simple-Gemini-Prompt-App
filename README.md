# Simple-Gemini-Prompt-App

This is a simple tool build using Go programming language that integrates with Gemini Pro API to provide responses from user text prompts.

The project's inspiration stems from GDG Kisumu Build With AI event.

## Features

Features and functionality of the the project below include:

- API integration with Gemini Pro AI

The tool takes text prompts/input from the user and is able to provide more information or context with the help of Gemini AI.

## Getting Started

To get started with the program, the user needs to have Go programming language installed on their laptop. Use the link below to download and install Go:

https://go.dev/doc/install

The other necessity you should have is a valid Gemini Application Programming Interface (API) key from Google. To purchase an API key, use the link below:

https://ai.google.dev/aistudio

## Project Installation

Once done with the steps above, you can now clone the repository to your local environment.
- Clone the repository using the command
```
git clone https://github.com/johneliud/Simple-Gemini-Prompt-App.git
```

- Change path to the correct directory
```
cd Simple-Gemini-Prompt-App
```

Having completed the steps above, the user will be ready to run and interact with the program

## Usage

Before running the project, the user must paste their purchased API key for the program to make requests from using the steps below:
- Navigate to the .env file
- Paste the API key on the the GEMINI_API_KEY variable. The key should be enclosed within the quotation marks e.g
```
GEMINI_API_KEY = "Paste API Key here"
```

With the above steps, the user can now successfully run the program from the terminal using the command:
```
go run main.go
```

or

```
go run .
```

If successful, the program should display information based on the users text prompt.

## Contact
Incase of further enquiries, please reach out via the email address johneliud4@gmail.com