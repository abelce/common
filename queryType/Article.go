package queryType

import (
	// gen_md "abelce/common/code-gen/models"

	"github.com/graphql-go/graphql"
)

var singleArticleType *graphql.Object // 使用单例模式
func GetArticleType(endpoint string) *graphql.Object {
	if singleArticleType != nil {
		return singleArticleType
	}
	singleArticleType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Article",
		Fields: graphql.Fields{

			"id": &graphql.Field{
				Type:        graphql.String,
				Description: "文章id",
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

			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "文章标题",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if p.Source == nil {
						return nil, nil
					}
					if val, exist := p.Source.(map[string]interface{})["name"]; exist {
						return val, nil
					}
					return nil, nil
				},
			},

			"categoryId": &graphql.Field{
				Type:        graphql.String,
				Description: "文章类型，不同的类型对应不用的泪容格式",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if p.Source == nil {
						return nil, nil
					}
					if val, exist := p.Source.(map[string]interface{})["categoryId"]; exist {
						return val, nil
					}
					return nil, nil
				},
			},

			"category": &graphql.Field{
				Type: GetCategoryType(endpoint),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return EntityResolver(p, endpoint+"/v1/categorys/")
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

			"content": &graphql.Field{
				Type:        graphql.String,
				Description: "内容",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if p.Source == nil {
						return nil, nil
					}
					if val, exist := p.Source.(map[string]interface{})["content"]; exist {
						return val, nil
					}
					return nil, nil
				},
			},
		},
	})
	return singleArticleType
}
