package gen_md

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// 分类
type Category struct {
	//id
	Id string `json:"id" valid:"required"`
	//标题
	Title string `json:"title" valid:"required"`
	//是否删除
	IsDeleted bool `json:"isDeleted"`
	//更新时间
	UpdatedTime int64 `json:"updatedTime"`
	//创建时间
	CreatedTime int64 `json:"createdTime"`
	//用户ID
	OperateID string `json:"operateID"`
}

func (entity *Category) Valid() error {
	_, err := govalidator.ValidateStruct(entity)
	if err != nil {
		return err
	}

	return nil
}

func NewCategory(
	id string,
	title string,
	updatedTime int64,
	createdTime int64,
	operateID string,
) (*Category, error) {
	entity := &Category{
		Id:          id,
		Title:       title,
		UpdatedTime: updatedTime,
		CreatedTime: createdTime,
		OperateID:   operateID,
	}

	entity.IsDeleted = false
	entity.CreatedTime = time.Now().Unix()
	entity.UpdatedTime = time.Now().Unix()

	if err := entity.Valid(); err != nil {
		return nil, err
	}

	return entity, nil
}

func (entity *Category) Delete() {
	entity.IsDeleted = false
	entity.UpdatedTime = time.Now().Unix()
}

func (entity *Category) Update(
	title string,
	isDeleted bool,
	updatedTime int64,
	createdTime int64,
	operateID string,
) error {
	entity.Title = title
	entity.IsDeleted = isDeleted
	entity.UpdatedTime = updatedTime
	entity.CreatedTime = createdTime
	entity.OperateID = operateID

	entity.UpdatedTime = time.Now().Unix()

	return entity.Valid()
}
