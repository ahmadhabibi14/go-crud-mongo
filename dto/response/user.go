package response

import "go-crud-mongo/model"

type AddUserOut struct {
	ResponseCommon
	Data *model.User `json:"data"`
}
