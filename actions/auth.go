package actions

import (
	"fmt"
	"os"

	"github.com/gobuffalo/buffalo"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/twitter"
	"github.com/as27/setenv"
	"net/http"
)

func init() {
	gothic.Store = App().SessionStore

	setenv.File(".env")

	goth.UseProviders(
		github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), fmt.Sprintf("%s%s", App().Host, "/auth/github/callback")),
		twitter.New(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), fmt.Sprintf("%s%s", App().Host, "/auth/twitter/callback")),
	)
}

func AuthCallback(c buffalo.Context) error {
	user, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return c.Error(401, err)
	}
	c.Session().Set("token", user.AccessToken)

	rTo := c.Session().GetOnce("login_redirect_to")
	rToStr, ok := rTo.(string)
	if !ok || rToStr == "" {
		rToStr = "/jobs"
	}
	// Do something with the user, maybe register them/sign them in
	return c.Redirect(http.StatusFound, rToStr)
}

func SetCurrentUser() buffalo.MiddlewareFunc {
	return func(h buffalo.Handler) buffalo.Handler {
		return func(c buffalo.Context) error {
			t := c.Session().Get("token")
			if t == nil {
				// set the redirect URL
				c.Session().Set("login_redirect_to", c.Request().URL.String())
				return c.Redirect(http.StatusFound, "/")
			}
			return h(c)
		}
	}
}

func AuthDestroy(c buffalo.Context) error {
	c.Session().Delete("token")
	return c.Redirect(http.StatusFound, "/")
}