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

type V1ApplicationTrait struct {
	Alias       string           `json:"alias,omitempty"`
	CreateTime  time.Time        `json:"createTime"`
	Description string           `json:"description,omitempty"`
	Properties  *ModelJsonStruct `json:"properties"`
	Type_       string           `json:"type"`
	UpdateTime  time.Time        `json:"updateTime"`
}
