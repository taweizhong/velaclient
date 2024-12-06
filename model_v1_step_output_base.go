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

type V1StepOutputBase struct {
	Id     string        `json:"id"`
	Name   string        `json:"name"`
	Phase  string        `json:"phase"`
	Type_  string        `json:"type"`
	Values []V1OutputVar `json:"values"`
}
