package dto

type CompareSnippetVersionsModificationsInputDTO struct {
	VersionsToCompare []string
}

type CompareSnippetVersionsModificationsOutputDTO struct {
	Modifications []string
}
