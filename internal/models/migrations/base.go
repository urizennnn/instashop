package migrations

import (
	"fmt"

	"github.com/urizennnn/instashop/pkg/repository/storage"
	"gorm.io/gorm"
)

func RunAllMigrations(db *storage.Database) {

	// verification migration
	MigrateModels(db.Postgresql, AuthMigrationModels(), AlterColumnModels())

}

func MigrateModels(db *gorm.DB, models []interface{}, AlterColums []AlterColumn) {
	_ = db.AutoMigrate(models...)

	for _, d := range AlterColums {
		err := d.UpdateColumnType(db)
		if err != nil {
			fmt.Println("error migrating ", d.TableName, "for column", d.Column, ": ", err)
		}

	}

	// d := AlterColumn{
	// 	Model:     mdl.OrgUserManagement{},
	// 	TableName: "org_user_managements",
	// 	Column:    "role_id",
	// }

	// d.DropColumn(db)

	// d = AlterColumn{
	// 	Model:     mdl.OrgUserManagement{},
	// 	TableName: "org_user_managements",
	// 	Column:    "role_id",
	// 	Type:      "uuid",
	// }

	// d.AddColumn(db)

}
