package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/pop"
	"github.com/connelevalsam/BuffaloProjects/job-position/models"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	trans := c.Value("tx").(*pop.Connection)
	runs := trans.Last(&models.Job{})
	c.Set("aJob", runs)
	return c.Render(200, r.HTML("index.html"))
}