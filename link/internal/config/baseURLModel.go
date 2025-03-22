package config

type BaseURL struct {
	URL string
}

func (b *BaseURL) GetCreateTaskUrl() string {
	return (b.URL + "/task/create")
}
