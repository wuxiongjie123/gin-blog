package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func (tag *Tag)BeforeCreate(scope *gorm.Scope) error  {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (tag *Tag)BeforeUpdate(scope *gorm.Scope) error  {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

func GetTags(pageNum int, pageSize int, maps interface{}) ([]Tag, error) {
	var (
		tags []Tag
		err  error
	)
	if pageSize > 0 && pageNum > 0 {
		err = db.Where(maps).Find(&tags).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = db.Where(maps).Find(&tags).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return tags, nil
}

func GetTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Tag{}).Where(maps).Count(&count).Error; err!=nil{
		return 0, err
	}
	return count, nil
}

func ExistTagByName(name string) (bool,error) {
	var tag Tag
	err:=db.Select("id").Where("name = ? AND deleted_on = ?", name,0).First(&tag).Error
	if err!=nil && err!=gorm.ErrRecordNotFound{
		return false,err
	}
	if tag.ID > 0 {
		return true,nil
	}
	return false,nil
}

func AddTag(name string, state int, createdBy string) error {
	tag := Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}
	if err := db.Create(&tag).Error; err!=nil {
		return err
	}
	return nil
}

func ExistTagById(id int) (bool,error) {
	var tag Tag
	err:=db.Select("id").Where("id = ? AND deleted_on = ?", id,0).First(&tag).Error
	if err!= nil && err!=gorm.ErrRecordNotFound{
		return false,err
	}
	if tag.ID>0 {
		return true,nil
	}
	return false,err
}

func EditTag(id int, data map[string]interface{}) error {
	if err:=db.Where("id = ? AND deleted_on = ? ",id,0).Update(data).Error;err!=nil{
		return err
	}
	return nil
}

func DeleteTag(id int) error {
	if err:=db.Where("id = ?", id).Delete(&Tag{}).Error;err!=nil{
		return err
	}
	return nil
}

func CleanAllTag() bool {
	db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Tag{})
	return true
}