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

type V1UpdateApplicationRequest struct {
	Alias string `json:"alias,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
	Description string `json:"description,omitempty"`
	Icon string `json:"icon,omitempty"`
	Labels map[string]string `json:"labels,omitempty"`
}