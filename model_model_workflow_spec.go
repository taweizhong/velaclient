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

type ModelWorkflowSpec struct {
	Mode  *V1alpha1WorkflowExecuteMode `json:"mode,omitempty"`
	Steps []ModelWorkflowStep          `json:"steps,omitempty"`
}
