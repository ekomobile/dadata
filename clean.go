package dadata

import "context"

func (c *Client) sendCleanRequest(ctx context.Context, lastURLPart string, source, result interface{}) error {
	return c.sendRequest(ctx, "clean/"+lastURLPart, source, result)
}

// CleanAddresses clean all provided addresses
// Call https://dadata.ru/api/v2/clean/address
func (c *Client) CleanAddresses(sourceAddresses ...string) ([]Address, error) {
	return c.CleanAddressesWithCtx(context.Background(), sourceAddresses...)
}

// CleanAddressesWithCtx clean all provided addresses
// Call https://dadata.ru/api/v2/clean/address
func (c *Client) CleanAddressesWithCtx(ctx context.Context, sourceAddresses ...string) (addresses []Address, err error) {
	if err = c.sendCleanRequest(ctx, "address", &sourceAddresses, &addresses); err != nil {
		addresses = nil
	}
	return
}

// CleanPhones clean all provided phones
// Call https://dadata.ru/api/v2/clean/phone
func (c *Client) CleanPhones(sourcePhones ...string) ([]Phone, error) {
	return c.CleanPhonesWithCtx(context.Background(), sourcePhones...)
}

// CleanPhonesWithCtx clean all provided phones
// Call https://dadata.ru/api/v2/clean/phone
func (c *Client) CleanPhonesWithCtx(ctx context.Context, sourcePhones ...string) (phones []Phone, err error) {
	if err = c.sendCleanRequest(ctx, "phone", &sourcePhones, &phones); err != nil {
		phones = nil
	}
	return
}

// CleanNames clean all provided names
// Call https://dadata.ru/api/v2/clean/name
func (c *Client) CleanNames(sourceNames ...string) ([]Name, error) {
	return c.CleanNamesWithCtx(context.Background(), sourceNames...)
}

// CleanNamesWithCtx clean all provided names
// Call https://dadata.ru/api/v2/clean/name
func (c *Client) CleanNamesWithCtx(ctx context.Context, sourceNames ...string) (names []Name, err error) {
	if err = c.sendCleanRequest(ctx, "name", &sourceNames, &names); err != nil {
		names = nil
	}
	return
}

// CleanEmails clean all provided emails
// Call https://dadata.ru/api/v2/clean/email
func (c *Client) CleanEmails(sourceEmails ...string) ([]Email, error) {
	return c.CleanEmailsWithCtx(context.Background(), sourceEmails...)
}

// CleanEmailsWithCtx clean all provided emails
// Call https://dadata.ru/api/v2/clean/email
func (c *Client) CleanEmailsWithCtx(ctx context.Context, sourceEmails ...string) (emails []Email, err error) {
	if err = c.sendCleanRequest(ctx, "email", &sourceEmails, &emails); err != nil {
		emails = nil
	}
	return
}

// CleanBirthdates clean all provided birthdates
// Call https://dadata.ru/api/v2/clean/birthdate
func (c *Client) CleanBirthdates(sourceBirthdates ...string) ([]Birthdate, error) {
	return c.CleanBirthdatesWithCtx(context.Background(), sourceBirthdates...)
}

// CleanBirthdatesWithCtx clean all provided birthdates
// Call https://dadata.ru/api/v2/clean/birthdate
func (c *Client) CleanBirthdatesWithCtx(ctx context.Context, sourceBirthdates ...string) (birthdates []Birthdate, err error) {
	if err = c.sendCleanRequest(ctx, "birthdate", &sourceBirthdates, &birthdates); err != nil {
		birthdates = nil
	}
	return
}

// CleanVehicles clean all provided vehicles
// Call https://dadata.ru/api/v2/clean/vehicle
func (c *Client) CleanVehicles(sourceVehicles ...string) ([]Vehicle, error) {
	return c.CleanVehiclesWithCtx(context.Background(), sourceVehicles...)
}

// CleanVehiclesWithCtx clean all provided vehicles
// Call https://dadata.ru/api/v2/clean/vehicle
func (c *Client) CleanVehiclesWithCtx(ctx context.Context, sourceVehicles ...string) (vehicles []Vehicle, err error) {
	if err = c.sendCleanRequest(ctx, "vehicle", &sourceVehicles, &vehicles); err != nil {
		vehicles = nil
	}
	return
}

// CleanPassports clean all provided passports
// Call https://dadata.ru/api/v2/clean/passport
func (c *Client) CleanPassports(sourcePassports ...string) ([]Passport, error) {
	return c.CleanPassportsWithCtx(context.Background(), sourcePassports...)
}

// CleanPassportsWithCtx clean all provided passports
// Call https://dadata.ru/api/v2/clean/passport
func (c *Client) CleanPassportsWithCtx(ctx context.Context, sourcePassports ...string) (passports []Passport, err error) {
	if err = c.sendCleanRequest(ctx, "passport", &sourcePassports, &passports); err != nil {
		passports = nil
	}
	return
}
