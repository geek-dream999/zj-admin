package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/warehouse"

type ExaCustomerResponse struct {
	Customer warehouse.ExaCustomer `json:"customer"`
}
