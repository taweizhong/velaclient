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

type V1UserProjectBase struct {
	Alias       string        `json:"alias"`
	Description string        `json:"description"`
	JoinTime    time.Time     `json:"joinTime"`
	Name        string        `json:"name"`
	Owner       *V1NameAlias  `json:"owner,omitempty"`
	Roles       []V1NameAlias `json:"roles"`
}
