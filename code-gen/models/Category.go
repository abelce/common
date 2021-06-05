package gen_md

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// 分类
type Category struct {
	  //id
  Id string `json:"id" valid:"required"`
  //名称
  Title string `json:"title" valid:"required"`
  //是否删除
  IsDeleted bool `json:"isDeleted"`
  //更新时间
  UpdateTime int64 `json:"updateTime"`
  //创建时间
  CreateTime int64 `json:"createTime"`
  //用户ID
  OperateID string `json:"operateID"`
  //对应的业务对象
  ObjectTypeId string `json:"objectTypeId"`
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
  updateTime int64,
  createTime int64,
  operateID string,
  objectTypeId string,
	) (*Category, error) {
	entity := &Category {
		   Id: id,
  Title: title,
  UpdateTime: updateTime,
  CreateTime: createTime,
  OperateID: operateID,
  ObjectTypeId: objectTypeId, 
		}

    entity.IsDeleted = false
    entity.CreateTime = time.Now().Unix()
	entity.UpdateTime = time.Now().Unix()
	
    if err := entity.Valid(); err != nil {
		return nil, err
	}

	return entity, nil
}

func (entity *Category) Delete() {
	entity.IsDeleted = false
	entity.UpdateTime = time.Now().Unix()
}

func (entity *Category) Update(
	  title string,
  isDeleted bool,
  updateTime int64,
  createTime int64,
  operateID string,
  objectTypeId string,
) error {
	  entity.Title=title
  entity.IsDeleted=isDeleted
  entity.UpdateTime=updateTime
  entity.CreateTime=createTime
  entity.OperateID=operateID
  entity.ObjectTypeId=objectTypeId

	entity.UpdateTime = time.Now().Unix()

	return entity.Valid()
}

