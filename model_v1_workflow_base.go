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

type V1WorkflowBase struct {
	Alias       string           `json:"alias"`
	CreateTime  time.Time        `json:"createTime"`
	Default_    bool             `json:"default"`
	Description string           `json:"description"`
	Enable      bool             `json:"enable"`
	EnvName     string           `json:"envName"`
	Mode        string           `json:"mode"`
	Name        string           `json:"name"`
	Steps       []V1WorkflowStep `json:"steps,omitempty"`
	SubMode     string           `json:"subMode"`
	UpdateTime  time.Time        `json:"updateTime"`
}
