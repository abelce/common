package queryType

import (
	gen_md "abelce/common/code-gen/models"
	"abelce/common/request"

	"github.com/graphql-go/graphql"
	"encoding/json"
)

func GetRootQueryType(endpoint string) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			
			"Article": &graphql.Field{
				Type: GetArticleType(endpoint),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type:         graphql.String,
						DefaultValue: "",
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// 请求转发到具体的服务，并获取数据
					if id, ok := p.Args["id"].(string); ok && id != "" {
						req := request.Request{
							Url: endpoint + "/v1/articles/" + id,
							Method: "GET",
						}
						result, err := req.Do()
						if err != nil {
							return nil, err
						}
						var entity gen_md.Product
						err = json.Unmarshal(result, &entity)
						if err != nil {
							return nil, nil
						}
						
						return entity, nil	
					}

					return nil, nil
				},
			},
			
			"Category": &graphql.Field{
				Type: GetCategoryType(endpoint),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type:         graphql.String,
						DefaultValue: "",
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// 请求转发到具体的服务，并获取数据
					if id, ok := p.Args["id"].(string); ok && id != "" {
						req := request.Request{
							Url: endpoint + "/v1/categorys/" + id,
							Method: "GET",
						}
						result, err := req.Do()
						if err != nil {
							return nil, err
						}
						var entity gen_md.Product
						err = json.Unmarshal(result, &entity)
						if err != nil {
							return nil, nil
						}
						
						return entity, nil	
					}

					return nil, nil
				},
			},
			
			"Product": &graphql.Field{
				Type: GetProductType(endpoint),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type:         graphql.String,
						DefaultValue: "",
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// 请求转发到具体的服务，并获取数据
					if id, ok := p.Args["id"].(string); ok && id != "" {
						req := request.Request{
							Url: endpoint + "/v1/products/" + id,
							Method: "GET",
						}
						result, err := req.Do()
						if err != nil {
							return nil, err
						}
						var entity gen_md.Product
						err = json.Unmarshal(result, &entity)
						if err != nil {
							return nil, nil
						}
						
						return entity, nil	
					}

					return nil, nil
				},
			},
			
			"User": &graphql.Field{
				Type: GetUserType(endpoint),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type:         graphql.String,
						DefaultValue: "",
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// 请求转发到具体的服务，并获取数据
					if id, ok := p.Args["id"].(string); ok && id != "" {
						req := request.Request{
							Url: endpoint + "/v1/users/" + id,
							Method: "GET",
						}
						result, err := req.Do()
						if err != nil {
							return nil, err
						}
						var entity gen_md.Product
						err = json.Unmarshal(result, &entity)
						if err != nil {
							return nil, nil
						}
						
						return entity, nil	
					}

					return nil, nil
				},
			},
			
		},
	})
}

