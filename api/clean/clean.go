package clean

import (
	"context"

	"github.com/ekomobile/dadata/v2/api/model"
)

type (
	// Requester provides transport level API calls.
	Requester interface {
		// Post makes a POST API call. Assumes sending json-encoded params in a request body.
		Post(ctx context.Context, apiMethod string, params, result interface{}) error
	}

	// Api provides "clean" API.
	Api struct {
		Client Requester
	}
)

// Address clean all provided addresses
// Call https://dadata.ru/api/v2/clean/address
func (a *Api) Address(ctx context.Context, addresses ...string) (result []*model.Address, err error) {
	err = a.Client.Post(ctx, "clean/address", &addresses, &result)
	return
}

// Phone clean all provided phones
// Call https://dadata.ru/api/v2/clean/phone
func (a *Api) Phone(ctx context.Context, phones ...string) (result []*model.Phone, err error) {
	err = a.Client.Post(ctx, "clean/phone", &phones, &result)
	return
}

// Name clean all provided names
// Call https://dadata.ru/api/v2/clean/name
func (a *Api) Name(ctx context.Context, names ...string) (result []*model.Name, err error) {
	err = a.Client.Post(ctx, "clean/name", &names, &result)
	return
}

// Email clean all provided emails
// Call https://dadata.ru/api/v2/clean/email
func (a *Api) Email(ctx context.Context, emails ...string) (result []*model.Email, err error) {
	err = a.Client.Post(ctx, "clean/email", &emails, &result)
	return
}

// Birthday clean all provided birthdates
// Call https://dadata.ru/api/v2/clean/birthdate
func (a *Api) Birthday(ctx context.Context, birthdates ...string) (result []*model.Birthday, err error) {
	err = a.Client.Post(ctx, "clean/birthdate", &birthdates, &result)
	return
}

// Vehicle clean all provided vehicles
// Call https://dadata.ru/api/v2/clean/vehicle
func (a *Api) Vehicle(ctx context.Context, vehicles ...string) (result []*model.Vehicle, err error) {
	err = a.Client.Post(ctx, "clean/vehicle", &vehicles, &result)
	return
}

// Passport clean all provided passports
// Call https://dadata.ru/api/v2/clean/passport
func (a *Api) Passport(ctx context.Context, passports ...string) (result []*model.Passport, err error) {
	err = a.Client.Post(ctx, "clean/passport", &passports, &result)
	return
}
