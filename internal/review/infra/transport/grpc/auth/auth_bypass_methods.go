package auth

var (
	ReviewMethodsToBypass = map[string]bool{
		"/proto.ReviewService/Register": true,
	}
)
