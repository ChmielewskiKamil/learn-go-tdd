package blogrenderer

import "io"

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

func Render(w io.Writer, p Post) error {
    _, err := w.Write([]byte("<h1>" + p.Title + "</h1>"))
	return err
}
