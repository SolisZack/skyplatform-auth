package response

import "skyplatform-auth/request"

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
