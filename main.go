package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func main() {
	godotenv.Load()
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading.env file")
	}

	ctx := context.Background()

	// Access API Key as environment variable
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Gemini model version
	model := client.GenerativeModel("gemini-pro")

	response, err := model.GenerateContent(ctx, genai.Text("Write a story about a magic backpack."))
	if err != nil {
		log.Fatal(err)
	}

	// Display's information from the generative AI back to the user
	fmt.Println(response)
}
