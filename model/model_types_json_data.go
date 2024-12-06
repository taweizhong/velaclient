/*
 * Kubevela api doc
 *
 * Kubevela api doc
 *
 * API version: v1beta1
 * Contact: feedback@mail.kubevela.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type TypesJsonData struct {
	AuthSecret *TypesKubernetesSecret `json:"authSecret,omitempty"`
	AuthType string `json:"authType,omitempty"`
	Backend bool `json:"backend"`
	BackendService *TypesKubernetesService `json:"backendService"`
	BackendType string `json:"backendType"`
	Category string `json:"category"`
	Id string `json:"id"`
	Includes []TypesIncludes `json:"includes"`
	Info *TypesInfo `json:"info"`
	KubePermissions []V1PolicyRule `json:"kubePermissions,omitempty"`
	Name string `json:"name"`
	Proxy bool `json:"proxy"`
	Requirement *TypesRequirement `json:"requirement,omitempty"`
	Routes []TypesRoute `json:"routes,omitempty"`
	SubType string `json:"subType"`
	Type_ string `json:"type"`
}