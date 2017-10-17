package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/pop"
	"github.com/connelevalsam/BuffaloProjects/job-position/models"
	"github.com/pkg/errors"
)

// PositionHandler is a default handler to serve up
// the position page.
func PositionHandler(c buffalo.Context) error {
	trans := c.Value("tx").(*pop.Connection)

	listJob, err := catJobs(trans)
	if err != nil {
		return errors.Wrap(err, "unable to load data")
	}
	c.Set("listJob", listJob)

	return c.Render(200, r.HTML("position.html"))
}

//categorize the jobs
func catJobs(trans *pop.Connection) (map[string][]models.Job, error) {
	var jobPos []models.Job
	listJob := make(map[string][]models.Job)
	if err := trans.All(&jobPos); err != nil {
		return nil, err
	}

	for _, j := range jobPos {
		category := j.Category.String
		if category == "" {
			category = "Others"
		}
		listJob[category] = append(listJob[category], j)
	}
	return listJob, nil
}