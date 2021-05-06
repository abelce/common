package queryType

import (
	gen_md "abelce/common/code-gen/models"
	
	"github.com/graphql-go/graphql"
)


var GetCategoryType = func(endpoint string) *graphql.Object{
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Category",
		Fields: graphql.Fields{
			
			"id": &graphql.Field {
				Type: graphql.String,
				Description: "id",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if p.Source == nil {
						return nil, nil
					}
					entity := p.Source.(gen_md.Category)
					return entity.Id, nil
				},
			},
			
			
			"title": &graphql.Field {
				Type: graphql.String,
				Description: "标题",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if p.Source == nil {
						return nil, nil
					}
					entity := p.Source.(gen_md.Category)
					return entity.Title, nil
				},
			},
			
			
			"isDeleted": &graphql.Field {
				Type: graphql.Boolean,
				Description: "是否删除",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if p.Source == nil {
						return nil, nil
					}
					entity := p.Source.(gen_md.Category)
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
					entity := p.Source.(gen_md.Category)
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
					entity := p.Source.(gen_md.Category)
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
					entity := p.Source.(gen_md.Category)
					return entity.OperateID, nil
				},
			},
			
			
		},
	})
}
