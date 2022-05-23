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

	// ApiImp provides "clean" API.
	ApiImp struct {
		Client Requester
	}

	Api interface {
		Address(context.Context, ...string) ([]*model.Address, error)
		Phone(context.Context, ...string) ([]*model.Phone, error)
		Name(context.Context, ...string) ([]*model.Name, error)
		Email(context.Context, ...string) ([]*model.Email, error)
		Birthday(context.Context, ...string) ([]*model.Birthday, error)
		Vehicle(context.Context, ...string) ([]*model.Vehicle, error)
		Passport(context.Context, ...string) ([]*model.Passport, error)
	}
)

// Address clean all provided addresses
// Call https://dadata.ru/api/v2/clean/address
func (a *ApiImp) Address(ctx context.Context, addresses ...string) (result []*model.Address, err error) {
	err = a.Client.Post(ctx, "clean/address", &addresses, &result)
	return
}

// Phone clean all provided phones
// Call https://dadata.ru/api/v2/clean/phone
func (a *ApiImp) Phone(ctx context.Context, phones ...string) (result []*model.Phone, err error) {
	err = a.Client.Post(ctx, "clean/phone", &phones, &result)
	return
}

// Name clean all provided names
// Call https://dadata.ru/api/v2/clean/name
func (a *ApiImp) Name(ctx context.Context, names ...string) (result []*model.Name, err error) {
	err = a.Client.Post(ctx, "clean/name", &names, &result)
	return
}

// Email clean all provided emails
// Call https://dadata.ru/api/v2/clean/email
func (a *ApiImp) Email(ctx context.Context, emails ...string) (result []*model.Email, err error) {
	err = a.Client.Post(ctx, "clean/email", &emails, &result)
	return
}

// Birthday clean all provided birthdates
// Call https://dadata.ru/api/v2/clean/birthdate
func (a *ApiImp) Birthday(ctx context.Context, birthdates ...string) (result []*model.Birthday, err error) {
	err = a.Client.Post(ctx, "clean/birthdate", &birthdates, &result)
	return
}

// Vehicle clean all provided vehicles
// Call https://dadata.ru/api/v2/clean/vehicle
func (a *ApiImp) Vehicle(ctx context.Context, vehicles ...string) (result []*model.Vehicle, err error) {
	err = a.Client.Post(ctx, "clean/vehicle", &vehicles, &result)
	return
}

// Passport clean all provided passports
// Call https://dadata.ru/api/v2/clean/passport
func (a *ApiImp) Passport(ctx context.Context, passports ...string) (result []*model.Passport, err error) {
	err = a.Client.Post(ctx, "clean/passport", &passports, &result)
	return
}
