package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/patrickrealdeal/flow-ci/internal/ci"
)

func SetupPipelines(app *fiber.App) {
    pipelinesGroup := app.Group("/pipelines")

    pipelinesGroup.Post("/check-it-works", postCheckItWorks)
}

type WithRepoUrl struct {
    Url string `json:"url" xml:"url" form:"url"`
}

func postCheckItWorks(c *fiber.Ctx) error {
    body := &WithRepoUrl{}

    if err := c.BodyParser(body); err != nil {
        return err
    }

    var ws ci.Workspace
    ws, err := ci.NewWorkspaceFromGit("./tmp", body.Url, "master")
    if err != nil {
        return err
    }

    return c.SendString(fmt.Sprintf("Cloned repository: %s\nFrom branch: %s\nCommit: %s\nInto Dir: %s\n", body.Url, ws.Branch(), ws.Commit(), ws.Dir()))
}
