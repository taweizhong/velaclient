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

type V1PipelineRunBriefing struct {
	ContextName     string       `json:"contextName"`
	ContextValues   []ModelValue `json:"contextValues"`
	EndTime         string       `json:"endTime"`
	Finished        bool         `json:"finished"`
	Message         string       `json:"message"`
	Phase           string       `json:"phase"`
	PipelineRunName string       `json:"pipelineRunName"`
	StartTime       string       `json:"startTime"`
}
