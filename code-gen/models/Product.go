package gen_md

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// 产品
type Product struct {
	  //产品ID
  Id string `json:"id" valid:"required"`
  //产品名称
  Name string `json:"name" valid:"required"`
  //产品类型
  CategoryId string `json:"categoryId" valid:"required"`
  //是否删除
  IsDeleted bool `json:"isDeleted"`
  //更新时间
  UpdateTime int64 `json:"updateTime"`
  //创建时间
  CreateTime int64 `json:"createTime"`
  //用户ID
  OperateID string `json:"operateID"`
  //概览
  Overview string `json:"overview"`
  //概览
  Logo string `json:"logo"`
}

func (entity *Product) Valid() error {
	_, err := govalidator.ValidateStruct(entity)
	if err != nil {
		return err
	}

	return nil
}

func NewProduct(
	  id string,
  name string,
  categoryId string,
  updateTime int64,
  createTime int64,
  operateID string,
  overview string,
  logo string,
	) (*Product, error) {
	entity := &Product {
		   Id: id,
  Name: name,
  CategoryId: categoryId,
  UpdateTime: updateTime,
  CreateTime: createTime,
  OperateID: operateID,
  Overview: overview,
  Logo: logo, 
		}

    entity.IsDeleted = false
    entity.CreateTime = time.Now().Unix()
	entity.UpdateTime = time.Now().Unix()
	
    if err := entity.Valid(); err != nil {
		return nil, err
	}

	return entity, nil
}

func (entity *Product) Delete() {
	entity.IsDeleted = false
	entity.UpdateTime = time.Now().Unix()
}

func (entity *Product) Update(
	  name string,
  categoryId string,
  isDeleted bool,
  updateTime int64,
  createTime int64,
  operateID string,
  overview string,
  logo string,
) error {
	  entity.Name=name
  entity.CategoryId=categoryId
  entity.IsDeleted=isDeleted
  entity.UpdateTime=updateTime
  entity.CreateTime=createTime
  entity.OperateID=operateID
  entity.Overview=overview
  entity.Logo=logo

	entity.UpdateTime = time.Now().Unix()

	return entity.Valid()
}

