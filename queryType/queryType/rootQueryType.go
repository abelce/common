package queryType

import (
	gen_md "abelce/common/code-gen/models"
	"abelce/common/request"

	"encoding/json"

	"github.com/graphql-go/graphql"
)

func GetRootQueryType(endpoint string) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{

			// 	"Product": &graphql.Field{
			// 		Type: ProductType,
			// 		Args: graphql.FieldConfigArgument{
			// 			"id": &graphql.ArgumentConfig{
			// 				Type:         graphql.String,
			// 				DefaultValue: "",
			// 			},
			// 		},
			// 		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 			// 请求转发到具体的服务，并获取数据
			// 			if id, ok := p.Args["id"].(string); ok && id != "" {
			// 				req := request.Request{
			// 					Url: endpoint + "/v1/products/" + id,
			// 					Method: "GET",
			// 				}
			// 				result, err := req.Do()
			// 				if err != nil {
			// 					return nil, err
			// 				}
			// 				var entity gen_md.Product
			// 				err = json.Unmarshal(result, &entity)
			// 				if err != nil {
			// 					return nil, nil
			// 				}

			// 				return entity, nil
			// 			}

			// 			return nil, nil
			// 		},
			// 	},
			"Product": &graphql.Field{
				Type: graphql.NewList(ProductType),
				Args: graphql.FieldConfigArgument{
					"queryStr": &graphql.ArgumentConfig{
						Type:         graphql.String,
						DefaultValue: "",
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// 请求转发到具体的服务，并获取数据
					// 考虑将服务器上返回的其他信息写入到context上，在ExecutionDidStart中写到exensions结构中，最后在graphql执行完后重新组装起来
					var params string
					if queryStr, ok := p.Args["queryStr"].(string); ok && queryStr != "" {
						params = queryStr
					}
					req := request.Request{
						Url:    endpoint + "/v1/products",
						Method: "GET",
						Params: params,
					}
					result, err := req.Do()
					if err != nil {
						return nil, err
					}

					type Result struct {
						Data []gen_md.Product `json:"data"`
					}

					var res Result
					json.Unmarshal(result, &res)
					return res.Data, nil
				},
			},

			// 	"User": &graphql.Field{
			// 		Type: UserType,
			// 		Args: graphql.FieldConfigArgument{
			// 			"id": &graphql.ArgumentConfig{
			// 				Type:         graphql.String,
			// 				DefaultValue: "",
			// 			},
			// 		},
			// 		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 			// 请求转发到具体的服务，并获取数据
			// 			if id, ok := p.Args["id"].(string); ok && id != "" {
			// 				req := request.Request{
			// 					Url: endpoint + "/v1/users/" + id,
			// 					Method: "GET",
			// 				}
			// 				result, err := req.Do()
			// 				if err != nil {
			// 					return nil, err
			// 				}
			// 				var entity gen_md.Product
			// 				err = json.Unmarshal(result, &entity)
			// 				if err != nil {
			// 					return nil, nil
			// 				}

			// 				return entity, nil
			// 			}

			// 			return nil, nil
			// 		},
			// 	},
			"User": &graphql.Field{
				Type: graphql.NewList(UserType),
				Args: graphql.FieldConfigArgument{
					"queryStr": &graphql.ArgumentConfig{
						Type:         graphql.String,
						DefaultValue: "",
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// 请求转发到具体的服务，并获取数据
					// 考虑将服务器上返回的其他信息写入到context上，在ExecutionDidStart中写到exensions结构中，最后在graphql执行完后重新组装起来
					var params string
					if queryStr, ok := p.Args["queryStr"].(string); ok && queryStr != "" {
						params = queryStr
					}
					req := request.Request{
						Url:    endpoint + "/v1/users",
						Method: "GET",
						Params: params,
					}
					result, err := req.Do()
					if err != nil {
						return nil, err
					}

					type Result struct {
						Data []gen_md.Product `json:"data"`
					}

					var res Result
					json.Unmarshal(result, &res)
					return res.Data, nil
				},
			},
		},
	})
}
