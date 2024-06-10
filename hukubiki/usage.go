package hukubiki

type jsonUsage struct {
	Description string `json:"description"`
}

func Usage() jsonUsage {
	usage := jsonUsage{
		Description: "`hukubiki` is a random number generator.",
	}
	return usage
}
