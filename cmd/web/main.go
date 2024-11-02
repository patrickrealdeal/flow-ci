package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/patrickrealdeal/flow-ci/internal/app/web/handlers"
)

func main() {
    app := fiber.New()

    handlers.SetupPipelines(app)

    app.Listen(":3000") 
}
