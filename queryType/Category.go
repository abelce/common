package queryType

import (
	// gen_md "github.com/abelce/common/code-gen/models"

	"github.com/graphql-go/graphql"
)

var singleCategoryType *graphql.Object // 使用单例模式
func GetCategoryType(endpoint string) *graphql.Object {
	if singleCategoryType != nil {
		return singleCategoryType
	}
	singleCategoryType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Category",
		Fields: graphql.Fields{

			"id": &graphql.Field{
				Type:        graphql.String,
				Description: "id",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if p.Source == nil {
						return nil, nil
					}
					if val, exist := p.Source.(map[string]interface{})["id"]; exist {
						return val, nil
					}
					return nil, nil
				},
			},

			"title": &graphql.Field{
				Type:        graphql.String,
				Description: "标题",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if p.Source == nil {
						return nil, nil
					}
					if val, exist := p.Source.(map[string]interface{})["title"]; exist {
						return val, nil
					}
					return nil, nil
				},
			},

			"isDeleted": &graphql.Field{
				Type:        graphql.Boolean,
				Description: "是否删除",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if p.Source == nil {
						return nil, nil
					}
					if val, exist := p.Source.(map[string]interface{})["isDeleted"]; exist {
						return val, nil
					}
					return nil, nil
				},
			},

			"updatedTime": &graphql.Field{
				Type:        graphql.Int,
				Description: "更新时间",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if p.Source == nil {
						return nil, nil
					}
					if val, exist := p.Source.(map[string]interface{})["updatedTime"]; exist {
						return val, nil
					}
					return nil, nil
				},
			},

			"createdTime": &graphql.Field{
				Type:        graphql.Int,
				Description: "创建时间",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if p.Source == nil {
						return nil, nil
					}
					if val, exist := p.Source.(map[string]interface{})["createdTime"]; exist {
						return val, nil
					}
					return nil, nil
				},
			},

			"operateID": &graphql.Field{
				Type:        graphql.String,
				Description: "用户ID",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if p.Source == nil {
						return nil, nil
					}
					if val, exist := p.Source.(map[string]interface{})["operateID"]; exist {
						return val, nil
					}
					return nil, nil
				},
			},
		},
	})
	return singleCategoryType
}
