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

type V1ListPipelineRunResponse struct {
	Runs  []V1PipelineRunBriefing `json:"runs"`
	Total int64                   `json:"total"`
}
