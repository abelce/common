package queryType

import (
	"abelce/common/request"

	"github.com/graphql-go/graphql"
	"encoding/json"
)

var Args = graphql.FieldConfigArgument{
			"queryStr": &graphql.ArgumentConfig{
				Type:         graphql.String,
				DefaultValue: "",
			},
			// "page[offset]": &graphql.ArgumentConfig{
			// 	Type:         graphql.Int,
			// 	DefaultValue: 0,
			// },
			// "page[limit]": &graphql.ArgumentConfig{
			// 	Type:         graphql.Int,
			// 	DefaultValue: 20,
			// },
		}

func ListResolver(p graphql.ResolveParams, url string) (interface{}, error) {
	// 请求转发到具体的服务，并获取数据
	// 考虑将服务器上返回的其他信息写入到context上，在ExecutionDidStart中写到exensions结构中，最后在graphql执行完后重新组装起来
	var params string
	if queryStr, ok := p.Args["queryStr"].(string); ok && queryStr != "" {
		params = queryStr
	}
	req := request.Request{
		Url:    url,
		Method: "GET",
		Params: params,
	}
	result, err := req.Do()
	if err != nil {
		return nil, err
	}

	type Result struct {
		Data []interface{} `json:"data"`
	}

	var res Result
	json.Unmarshal(result, &res)
	return res.Data, nil
}

func EntityResolver(p graphql.ResolveParams, url string) (interface{}, error) {
	// 请求转发到具体的服务，并获取数据
	// 考虑将服务器上返回的其他信息写入到context上，在ExecutionDidStart中写到exensions结构中，最后在graphql执行完后重新组装起来
	var params string
	if queryStr, ok := p.Args["queryStr"].(string); ok && queryStr != "" {
		params = queryStr
	}
	req := request.Request{
		Url:    url,
		Method: "GET",
		Params: params,
	}
	result, err := req.Do()
	if err != nil {
		return nil, err
	}

	var res interface{}
	json.Unmarshal(result, &res)
	return res, nil
}

func GetRootQueryType(endpoint string) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			
			"Article": &graphql.Field{
				Type: graphql.NewList(GetArticleType(endpoint)),
				Args: Args,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return ListResolver(p, endpoint + "/v1/articles")
				},
			},
			
			"Product": &graphql.Field{
				Type: graphql.NewList(GetProductType(endpoint)),
				Args: Args,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return ListResolver(p, endpoint + "/v1/products")
				},
			},
			
			"User": &graphql.Field{
				Type: graphql.NewList(GetUserType(endpoint)),
				Args: Args,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return ListResolver(p, endpoint + "/v1/users")
				},
			},
			
		},
	})
}

