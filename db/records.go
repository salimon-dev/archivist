package db

import (
	"salimon/archivist/types"

	"gorm.io/gorm"
)

func RecordsModel() *gorm.DB {
	return DB.Model(types.Record{})
}

func FindRecord(query interface{}, args ...interface{}) (*types.Record, error) {
	var user types.Record

	result := DB.Model(types.Record{}).Where(query, args...).Find(&user)

	if result.RowsAffected == 0 {
		return nil, result.Error
	} else {
		return &user, nil
	}
}

func InsertRecord(user *types.Record) error {
	result := DB.Model(types.Record{}).Create(user)
	return result.Error
}
func UpdateRecord(user *types.Record) error {
	result := DB.Model(types.Record{}).Where("id = ?", user.Id).Updates(user)
	return result.Error
}

func FindRecords(query interface{}, offset int, limit int, args ...interface{}) ([]types.Record, error) {
	var users []types.Record
	result := DB.Model(types.Record{}).Select("*").Where(query, args...).Offset(offset).Limit(limit).Find(users)
	return users, result.Error
}
