package response

import "go-crud-mongo/model"

type AddUserOut struct {
	ResponseCommon
	Data *model.User `json:"data"`
}

type UpdateUserOut struct {
	ResponseCommon
	Data *model.User `json:"data"`
}

type GetUserOut struct {
	ResponseCommon
	Data *model.User `json:"data"`
}

type GetUsersOut struct {
	ResponseCommon
	Data *[]model.User `json:"data"`
}
