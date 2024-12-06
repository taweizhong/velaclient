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
import (
	"time"
)

type V1ApplicationTriggerBase struct {
	Alias string `json:"alias,omitempty"`
	ComponentName string `json:"componentName,omitempty"`
	CreateTime time.Time `json:"createTime"`
	Description string `json:"description,omitempty"`
	Name string `json:"name"`
	PayloadType string `json:"payloadType"`
	Registry string `json:"registry"`
	Token string `json:"token"`
	Type_ string `json:"type"`
	UpdateTime time.Time `json:"updateTime"`
	WorkflowName string `json:"workflowName"`
}