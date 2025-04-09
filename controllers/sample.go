package controllers

import (
	"github.com/verbeux-ai/crm-integration/interfaces"
)

type sampleController struct {
	//sampleController interfaces2.Repository[models.CRMContact, *repositories.CRMContact]
	//crmIntegration   interfaces2.Repository[models.CRMIntegration, *repositories.CRMIntegration]
}

func NewCRMContact() interfaces.CRMContact {
	return &sampleController{
		//sampleController: repositories.NewPostgresRepository[models.CRMContact](repositories.NewCRMContact()),
		//crmIntegration:   repositories.NewPostgresRepository[models.CRMIntegration](repositories.NewCRMIntegration()),
	}
}
