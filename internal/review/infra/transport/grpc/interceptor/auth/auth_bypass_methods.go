package auth

var (
	ReviewMethodsToBypass = map[string]bool{
		"/proto.ReviewService/FetchSnippetTopics": true,
		"/proto.ReviewService/FetchSnippets":      true,
	}
)
