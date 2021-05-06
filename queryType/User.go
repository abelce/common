package queryType

import (
	gen_md "abelce/common/code-gen/models"

	"github.com/graphql-go/graphql"
)


var singleUserType *graphql.Object // 使用单例模式
func GetUserType(endpoint string) *graphql.Object{
	if singleUserType != nil {
		return singleUserType
	}
	singleUserType = graphql.NewObject(graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			
			"id": &graphql.Field {
				Type: graphql.String,
				Description: "用户ID",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if p.Source == nil {
						return nil, nil
					}
					entity := p.Source.(gen_md.User)
					return entity.Id, nil
				},
			},
			
			
			"name": &graphql.Field {
				Type: graphql.String,
				Description: "姓名",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if p.Source == nil {
						return nil, nil
					}
					entity := p.Source.(gen_md.User)
					return entity.Name, nil
				},
			},
			
			
			"sex": &graphql.Field {
				Type: graphql.String,
				Description: "性别",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if p.Source == nil {
						return nil, nil
					}
					entity := p.Source.(gen_md.User)
					return entity.Sex, nil
				},
			},
			
			
			"isDeleted": &graphql.Field {
				Type: graphql.Boolean,
				Description: "是否删除",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if p.Source == nil {
						return nil, nil
					}
					entity := p.Source.(gen_md.User)
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
					entity := p.Source.(gen_md.User)
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
					entity := p.Source.(gen_md.User)
					return entity.CreatedTime, nil
				},
			},
			
			
		},
	})
	return singleUserType
}
