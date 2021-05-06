package gen_md

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// 文章
type Article struct {
	  //文章id
  Id string `json:"id" valid:"required"`
  //文章标题
  Name string `json:"name" valid:"required"`
  //文章类型，不同的类型对应不用的泪容格式
  CategoryId string `json:"categoryId" valid:"required"`
  //是否删除
  IsDeleted bool `json:"isDeleted"`
  //更新时间
  UpdatedTime int64 `json:"updatedTime"`
  //创建时间
  CreatedTime int64 `json:"createdTime"`
  //用户ID
  OperateID string `json:"operateID"`
  //内容
  Content string `json:"content"`
}

func (entity *Article) Valid() error {
	_, err := govalidator.ValidateStruct(entity)
	if err != nil {
		return err
	}

	return nil
}

func NewArticle(
	  id string,
  name string,
  categoryId string,
  updatedTime int64,
  createdTime int64,
  operateID string,
  content string,
	) (*Article, error) {
	entity := &Article {
		   Id: id,
  Name: name,
  CategoryId: categoryId,
  UpdatedTime: updatedTime,
  CreatedTime: createdTime,
  OperateID: operateID,
  Content: content, 
		}

    entity.IsDeleted = false
    entity.CreatedTime = time.Now().Unix()
	entity.UpdatedTime = time.Now().Unix()
	
    if err := entity.Valid(); err != nil {
		return nil, err
	}

	return entity, nil
}

func (entity *Article) Delete() {
	entity.IsDeleted = false
	entity.UpdatedTime = time.Now().Unix()
}

func (entity *Article) Update(
	  name string,
  categoryId string,
  isDeleted bool,
  updatedTime int64,
  createdTime int64,
  operateID string,
  content string,
) error {
	  entity.Name=name
  entity.CategoryId=categoryId
  entity.IsDeleted=isDeleted
  entity.UpdatedTime=updatedTime
  entity.CreatedTime=createdTime
  entity.OperateID=operateID
  entity.Content=content

	entity.UpdatedTime = time.Now().Unix()

	return entity.Valid()
}

