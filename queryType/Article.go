package queryType

import (
	gen_md "abelce/common/code-gen/models"
	
	"github.com/graphql-go/graphql"
)


var GetArticleType = func(endpoint string) *graphql.Object{
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Article",
		Fields: graphql.Fields{
			
			"id": &graphql.Field {
				Type: graphql.String,
				Description: "文章id",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if p.Source == nil {
						return nil, nil
					}
					entity := p.Source.(gen_md.Article)
					return entity.Id, nil
				},
			},
			
			
			"name": &graphql.Field {
				Type: graphql.String,
				Description: "文章标题",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if p.Source == nil {
						return nil, nil
					}
					entity := p.Source.(gen_md.Article)
					return entity.Name, nil
				},
			},
			
			
			"categoryId": &graphql.Field {
				Type: graphql.String,
				Description: "文章类型，不同的类型对应不用的泪容格式",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if p.Source == nil {
						return nil, nil
					}
					entity := p.Source.(gen_md.Article)
					return entity.CategoryId, nil
				},
			},
			
			
			"isDeleted": &graphql.Field {
				Type: graphql.Boolean,
				Description: "是否删除",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if p.Source == nil {
						return nil, nil
					}
					entity := p.Source.(gen_md.Article)
					return entity.IsDeleted, nil
				},
			},
			
			
			"updatedTime": &graphql.Field {
				Type: graphql.Int,
				Description: "更新时间",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if p.Source == nil {
						return nil, nil
					}
					entity := p.Source.(gen_md.Article)
					return entity.UpdatedTime, nil
				},
			},
			
			
			"createdTime": &graphql.Field {
				Type: graphql.Int,
				Description: "创建时间",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if p.Source == nil {
						return nil, nil
					}
					entity := p.Source.(gen_md.Article)
					return entity.CreatedTime, nil
				},
			},
			
			
			"operateID": &graphql.Field {
				Type: graphql.String,
				Description: "用户ID",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if p.Source == nil {
						return nil, nil
					}
					entity := p.Source.(gen_md.Article)
					return entity.OperateID, nil
				},
			},
			
			
			"content": &graphql.Field {
				Type: graphql.String,
				Description: "内容",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if p.Source == nil {
						return nil, nil
					}
					entity := p.Source.(gen_md.Article)
					return entity.Content, nil
				},
			},
			
			
		},
	})
}
