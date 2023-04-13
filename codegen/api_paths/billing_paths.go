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
}
