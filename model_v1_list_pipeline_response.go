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

type V1ListPipelineResponse struct {
	Pipelines []V1PipelineListItem `json:"pipelines"`
	Total     int32                `json:"total"`
}
