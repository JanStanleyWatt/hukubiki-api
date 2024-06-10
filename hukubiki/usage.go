package hukubiki

type JsonUsage struct {
	Description string `json:"description"`
}

func Usage() JsonUsage {
	usage := JsonUsage{
		Description: "Usage: hukubiki <json>",
	}

	return usage
}
