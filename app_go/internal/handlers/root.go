package handlers

import (
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

// NewRoot creates a handler for the root page returning random number fact.
func NewRoot() func(*fiber.Ctx) error {
	file := "./static/index.html"
	content, err := os.ReadFile(file)
	if err != nil {
		panic("can not open static content")
	}

	template := string(content)
	return func(c *fiber.Ctx) error {
		return c.Type("html").SendString(fmt.Sprintf(template, getRandomFact()))
	}
}

// getRandomFact goes to API and pulls a random fact.
func getRandomFact() string {
	const maxRandomNumber = 2024
	const api = "http://numbersapi.com/%d"
	const noAnswer = "Unfortunately, no fact for today :("

	num := rand.IntN(maxRandomNumber)
	resp, err := http.Get(fmt.Sprintf(api, num))

	if err != nil {
		return noAnswer
	}

	fact, err := io.ReadAll(resp.Body)
	if err != nil {
		return noAnswer
	}
	return string(fact)

}
