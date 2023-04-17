package api_paths

import (
	"github.com/tada-team/tdproto"
)

var BillingPaths = []PathSpec{
	{
		Path: "/api/v4/tariffs",
		Get: &OperationSpec{
			Response:    []tdproto.TariffBilling{},
			Description: "Get all tariffs.",
		},
	},
	{
		Path: "/api/v4/workplaces",
		Get: &OperationSpec{
			QueryStruct: tdproto.GetWorkplacesByPersonalAccountRequest{},
			Response:    []tdproto.GetWorkplacesByPersonalAccountResponse{},
			Description: "Get all workplaces.",
		},
		Put: &OperationSpec{
			Request:     tdproto.AddWorkplacesOnPersonalAccountRequest{},
			Response:    tdproto.ActivateWorkplacesOnPersonalAccountResponse{},
			Description: "Add workplaces.",
		},
		Delete: &OperationSpec{
			Request:     tdproto.DeleteUserFromWorkplaceRequest{},
			Response:    tdproto.DeleteUserFromWorkplaceResponse{},
			Description: "Delete workplaces.",
		},
	},
	{
		Path: "/api/v4/bill-teams",
		Get: &OperationSpec{
			Response:    tdproto.GetTeamOnPersonalAccountResponse{},
			Description: "Get teams on personal account.",
		},
	},
}
