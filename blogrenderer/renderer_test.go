package blogrenderer_test

import (
	"bytes"
	approvals "github.com/approvals/go-approval-tests"
	"learn-go-tdd/blogrenderer"
	"testing"
)

func TestRender(t *testing.T) {
	approvals.UseFolder("testdata")
	var (
		aPost = blogrenderer.Post{
			Title:       "Hello, World!",
			Body:        "This is the body.",
			Description: "This is the description.",
			Tags:        []string{"golang", "programming"},
		}
	)

	t.Run("it converts single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blogrenderer.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}
