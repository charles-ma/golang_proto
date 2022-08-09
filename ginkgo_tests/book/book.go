package book

type Book struct {
	Title  string
	Author string
	Pages  int
}

const (
	CategoryNoval      = "Noval"
	CatetoryShortStory = "ShortStory"
)

func (b Book) Category() string {
	if b.Pages > 300 {
		return CategoryNoval
	}
	return CatetoryShortStory
}
