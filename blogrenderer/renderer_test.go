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

		got := buf.String()
		want := `<h1>Hello, World!</h1><p>This is the description.</p>Tags:<ul><li>golang</li><li>programming</li></ul>`

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
