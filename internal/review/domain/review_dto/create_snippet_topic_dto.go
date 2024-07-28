package review_dto

type CreateSnippetTopicInputDTO struct {
	Title       string
	Description string
	AccountID   string
}

type CreateSnippetTopicOutputDTO struct {
	ID string
}
