package testing

type Retriver struct{}

func (Retriver) Get(url string) string {
	return "fake content"
}
