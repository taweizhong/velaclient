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

import (
	"time"
)

type V1EnvBindingBase struct {
	Alias              string               `json:"alias,omitempty"`
	AppDeployName      string               `json:"appDeployName"`
	AppDeployNamespace string               `json:"appDeployNamespace"`
	ComponentSelector  *V1ComponentSelector `json:"componentSelector,omitempty"`
	CreateTime         time.Time            `json:"createTime"`
	Description        string               `json:"description,omitempty"`
	Name               string               `json:"name"`
	TargetNames        []string             `json:"targetNames"`
	Targets            []V1EnvBindingTarget `json:"targets,omitempty"`
	UpdateTime         time.Time            `json:"updateTime"`
	Workflow           *V1NameAlias         `json:"workflow"`
}
