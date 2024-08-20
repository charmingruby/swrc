package dto

type CommentOnSnippetTopicInputDTO struct {
	Content         string
	AccountID       string
	SnippetTopicID  string
	ParentCommentID string
}
