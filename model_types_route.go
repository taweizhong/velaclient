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

type TypesRoute struct {
	Headers     []TypesHeader     `json:"headers,omitempty"`
	Method      string            `json:"method"`
	Path        string            `json:"path"`
	Permission  *TypesPermission  `json:"permission,omitempty"`
	ResourceMap map[string]string `json:"resourceMap,omitempty"`
}
