package queryType

import (
	gen_md "abelce/common/code-gen/models"

	"github.com/graphql-go/graphql"
)


var singleProductType *graphql.Object // 使用单例模式
func GetProductType(endpoint string) *graphql.Object{
	if singleProductType != nil {
		return singleProductType
	}
	singleProductType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			
			"id": &graphql.Field {
				Type: graphql.String,
				Description: "产品ID",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if p.Source == nil {
						return nil, nil
					}
					entity := p.Source.(gen_md.Product)
					return entity.Id, nil
				},
			},
			
			
			"name": &graphql.Field {
				Type: graphql.String,
				Description: "产品名称",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if p.Source == nil {
						return nil, nil
					}
					entity := p.Source.(gen_md.Product)
					return entity.Name, nil
				},
			},
			
			
			"categoryId": &graphql.Field {
				Type: graphql.String,
				Description: "产品类型",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if p.Source == nil {
						return nil, nil
					}
					entity := p.Source.(gen_md.Product)
					return entity.CategoryId, nil
				},
			},
			
			"category": &graphql.Field {
				Type: GetCategoryType(endpoint),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return EntityResolver(p, endpoint + "/v1/categorys/")
				},
			},
			
			
			"isDeleted": &graphql.Field {
				Type: graphql.Boolean,
				Description: "是否删除",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if p.Source == nil {
						return nil, nil
					}
					entity := p.Source.(gen_md.Product)
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
					entity := p.Source.(gen_md.Product)
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
					entity := p.Source.(gen_md.Product)
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
					entity := p.Source.(gen_md.Product)
					return entity.OperateID, nil
				},
			},
			
			
			"overview": &graphql.Field {
				Type: graphql.String,
				Description: "概览",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if p.Source == nil {
						return nil, nil
					}
					entity := p.Source.(gen_md.Product)
					return entity.Overview, nil
				},
			},
			
			
			"logo": &graphql.Field {
				Type: graphql.String,
				Description: "概览",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if p.Source == nil {
						return nil, nil
					}
					entity := p.Source.(gen_md.Product)
					return entity.Logo, nil
				},
			},
			
			
		},
	})
	return singleProductType
}
