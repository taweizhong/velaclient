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

type V1ApplicationTemplateVersion struct {
	CreateTime  time.Time `json:"createTime"`
	CreateUser  string    `json:"createUser"`
	Description string    `json:"description"`
	UpdateTime  time.Time `json:"updateTime"`
	Version     string    `json:"version"`
}
