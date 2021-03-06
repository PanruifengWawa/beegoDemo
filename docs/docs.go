package docs

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/swagger"
)

const (
    Rootinfo string = `{"apiVersion":"1.0.0","swaggerVersion":"1.2","apis":[{"path":"/t_token","description":"oprations for TToken\n"},{"path":"/t_user","description":"oprations for TUser\n"}],"info":{"title":"beego Test API","description":"beego has a very cool tools to autogenerate documents for your API","contact":"astaxie@gmail.com","termsOfServiceUrl":"http://beego.me/","license":"Url http://www.apache.org/licenses/LICENSE-2.0.html"}}`
    Subapi string = `{"/t_token":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/t_token","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/","description":"","operations":[{"httpMethod":"POST","nickname":"Post","type":"","summary":"create TToken","parameters":[{"paramType":"body","name":"body","description":"\"body for TToken content\"","dataType":"TToken","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":201,"message":"models.TToken","responseModel":""},{"code":403,"message":"body is empty","responseModel":""}]}]},{"path":"/:id","description":"","operations":[{"httpMethod":"GET","nickname":"Get","type":"","summary":"get TToken by id","parameters":[{"paramType":"path","name":"id","description":"\"The key for staticblock\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.TToken","responseModel":"TToken"},{"code":403,"message":":id is empty","responseModel":""}]}]},{"path":"/","description":"","operations":[{"httpMethod":"GET","nickname":"Get All","type":"","summary":"get TToken","parameters":[{"paramType":"query","name":"query","description":"\"Filter. e.g. col1:v1,col2:v2 ...\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"fields","description":"\"Fields returned. e.g. col1,col2 ...\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"sortby","description":"\"Sorted-by fields. e.g. col1,col2 ...\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"order","description":"\"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ...\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"limit","description":"\"Limit the size of result set. Must be an integer\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"offset","description":"\"Start position of result set. Must be an integer\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.TToken","responseModel":"TToken"},{"code":403,"message":"","responseModel":""}]}]},{"path":"/:id","description":"","operations":[{"httpMethod":"PUT","nickname":"Update","type":"","summary":"update the TToken","parameters":[{"paramType":"path","name":"id","description":"\"The id you want to update\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"body","name":"body","description":"\"body for TToken content\"","dataType":"TToken","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.TToken","responseModel":"TToken"},{"code":403,"message":":id is not int","responseModel":""}]}]},{"path":"/:id","description":"","operations":[{"httpMethod":"DELETE","nickname":"Delete","type":"","summary":"delete the TToken","parameters":[{"paramType":"path","name":"id","description":"\"The id you want to delete\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"delete success!","responseModel":""},{"code":403,"message":"id is empty","responseModel":""}]}]}],"models":{"TToken":{"id":"TToken","properties":{"Id":{"type":"int","description":"","format":""},"Token":{"type":"string","description":"","format":""}}}}},"/t_user":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/t_user","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/","description":"","operations":[{"httpMethod":"POST","nickname":"Post","type":"","summary":"create TUser","parameters":[{"paramType":"body","name":"body","description":"\"body for TUser content\"","dataType":"TUser","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":201,"message":"models.TUser","responseModel":""},{"code":403,"message":"body is empty","responseModel":""}]}]},{"path":"/:id","description":"","operations":[{"httpMethod":"GET","nickname":"Get","type":"","summary":"get TUser by id","parameters":[{"paramType":"path","name":"id","description":"\"The key for staticblock\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.TUser","responseModel":"TUser"},{"code":403,"message":":id is empty","responseModel":""}]}]},{"path":"/","description":"","operations":[{"httpMethod":"GET","nickname":"Get All","type":"","summary":"get TUser","parameters":[{"paramType":"query","name":"query","description":"\"Filter. e.g. col1:v1,col2:v2 ...\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"fields","description":"\"Fields returned. e.g. col1,col2 ...\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"sortby","description":"\"Sorted-by fields. e.g. col1,col2 ...\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"order","description":"\"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ...\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"limit","description":"\"Limit the size of result set. Must be an integer\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"offset","description":"\"Start position of result set. Must be an integer\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.TUser","responseModel":"TUser"},{"code":403,"message":"","responseModel":""}]}]},{"path":"/:id","description":"","operations":[{"httpMethod":"PUT","nickname":"Update","type":"","summary":"update the TUser","parameters":[{"paramType":"path","name":"id","description":"\"The id you want to update\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"body","name":"body","description":"\"body for TUser content\"","dataType":"TUser","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.TUser","responseModel":"TUser"},{"code":403,"message":":id is not int","responseModel":""}]}]},{"path":"/:id","description":"","operations":[{"httpMethod":"DELETE","nickname":"Delete","type":"","summary":"delete the TUser","parameters":[{"paramType":"path","name":"id","description":"\"The id you want to delete\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"delete success!","responseModel":""},{"code":403,"message":"id is empty","responseModel":""}]}]}],"models":{"TUser":{"id":"TUser","properties":{"Id":{"type":"int","description":"","format":""},"Name":{"type":"string","description":"","format":""}}}}}}`
    BasePath string= "/v1"
)

var rootapi swagger.ResourceListing
var apilist map[string]*swagger.APIDeclaration

func init() {
	if beego.BConfig.WebConfig.EnableDocs {
		err := json.Unmarshal([]byte(Rootinfo), &rootapi)
		if err != nil {
			beego.Error(err)
		}
		err = json.Unmarshal([]byte(Subapi), &apilist)
		if err != nil {
			beego.Error(err)
		}
		beego.GlobalDocAPI["Root"] = rootapi
		for k, v := range apilist {
			for i, a := range v.APIs {
				a.Path = urlReplace(k + a.Path)
				v.APIs[i] = a
			}
			v.BasePath = BasePath
			beego.GlobalDocAPI[strings.Trim(k, "/")] = v
		}
	}
}


func urlReplace(src string) string {
	pt := strings.Split(src, "/")
	for i, p := range pt {
		if len(p) > 0 {
			if p[0] == ':' {
				pt[i] = "{" + p[1:] + "}"
			} else if p[0] == '?' && p[1] == ':' {
				pt[i] = "{" + p[2:] + "}"
			}
		}
	}
	return strings.Join(pt, "/")
}
