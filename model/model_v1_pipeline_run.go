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

type V1PipelineRun struct {
	ContextName string `json:"contextName"`
	ContextValues []ModelValue `json:"contextValues"`
	PipelineName string `json:"pipelineName"`
	PipelineRunName string `json:"pipelineRunName"`
	Project *V1NameAlias `json:"project"`
	Record int64 `json:"record"`
	Spec *V1alpha1WorkflowRunSpec `json:"spec"`
	Status *V1alpha1WorkflowRunStatus `json:"status"`
}