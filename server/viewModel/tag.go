package viewModel

func NewTag(tag_ string, articles *Articles) *tag {
	return &tag{
		Tag:      tag_,
		Articles: articles,
	}
}

type tag struct {
	Tag      string
	Articles *Articles
}
