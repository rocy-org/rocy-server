package resource

import "github.com/rocy-org/rocy-server/utils/db"

var DB = db.DB

func (r *Resource) GetAllResource() ([]Resource, error) {
	var data []Resource
	err := DB.Find(&data).Error
	return data, err
}

func (r *Resource) Create() error {
	return DB.Create(r).Error
}
