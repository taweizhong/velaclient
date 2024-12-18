/*
 * Kubevela api doc
 *
 * Kubevela api doc
 *
 * API version: v1beta1
 * Contact: feedback@mail.kubevela.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package velaclient

type V1UpdateApplicationComponentRequest struct {
	Alias       string             `json:"alias,omitempty"`
	DependsOn   []string           `json:"dependsOn,omitempty"`
	Description string             `json:"description,omitempty"`
	Icon        string             `json:"icon,omitempty"`
	Labels      *map[string]string `json:"labels,omitempty"`
	Properties  string             `json:"properties,omitempty"`
}
