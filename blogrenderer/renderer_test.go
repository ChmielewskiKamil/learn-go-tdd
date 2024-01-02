package blogrenderer_test

import (
	"bytes"
	"learn-go-tdd/blogrenderer"
	"testing"
)

func TestBlogRenderer(t *testing.T) {
	var (
		aPost = blogrenderer.Post{
			Title:       "Hello, World!",
			Body:        "This is the body of my blog post.",
			Description: "This is the description of my blog post.",
			Tags:        []string{"golang", "programming", "tutorial"},
		}
	)

	t.Run("it renders the HTML post title", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blogrenderer.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>Hello, World!</h1>`

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
