package review_dto

type CompareSnippetVersionsModificationsInputDTO struct {
	VersionsToCompare []string
}

type CompareSnippetVersionsModificationsOutputDTO struct {
	Modifications []string
}
