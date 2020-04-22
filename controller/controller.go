package controller

import (
	"golang_side_project_crud_website/models"
	"html/template"
)

type PageContent struct {
	PageTitle string
	PageQuery interface{}
	CsrfTag template.HTML
	IsUser bool
	Collections [] models.Collection
}

func newPageContent() *PageContent {
	var pageContent PageContent
	collections, err := models.FindAllCollections()
	if err != nil {

	}
	pageContent.Collections = collections

	return &pageContent
}
