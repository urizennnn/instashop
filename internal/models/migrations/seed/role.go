package seed

import (
	"github.com/urizennnn/instashop/internal/models"
	"github.com/urizennnn/instashop/utility"
	"gorm.io/gorm"
)

func SeedRoles(logger *utility.Logger, db *gorm.DB) {
	roles := []models.Role{
		{
			ID:   utility.GenerateUUID(),
			Name: "administrator",
			Desc: "Full access",
		},
		{
			ID:   utility.GenerateUUID(),
			Name: "user",
			Desc: "Limited access",
		},
	}
	for _, role := range roles {
		err := role.CreateRole(db)
		logger.Info("Seeding role: ", role.ID)
		if err != nil {
			logger.Error("Error while seeding roles", err)
		}
	}
}
