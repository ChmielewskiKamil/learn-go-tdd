package blogrenderer_test

import (
	"bytes"
	"io"
	"learn-go-tdd/blogrenderer"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
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

	postRenderer, err := blogrenderer.NewPostRenderer()

	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := postRenderer.Render(&buf, aPost); err != nil {
			approvals.VerifyString(t, buf.String())
		}
	})

	t.Run("it renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}

		posts := []blogrenderer.Post{{Title: "First Post"}, {Title: "Second Post"}}

		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := blogrenderer.NewPostRenderer()

	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, aPost)
	}
}
