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

type V1LoginRequest struct {
	Code     string `json:"code,omitempty"`
	Password string `json:"password,omitempty"`
	Username string `json:"username,omitempty"`
}
