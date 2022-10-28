package adress

import (
	"kang-sayur-backend/model/domain/address"
	request_body "kang-sayur-backend/model/web/request_body/address"

	"go.mongodb.org/mongo-driver/mongo"
)

type addressRepository struct {
	collection mongo.Collection
}

// Create implements address.AddressRepository
func (*addressRepository) Create(data *request_body.CreateAddress) (*address.Address, error) {
	panic("unimplemented")
}

// Delelete implements address.AddressRepository
func (*addressRepository) Delelete(data *request_body.DeleteOrReadAddress) error {
	panic("unimplemented")
}

// Read implements address.AddressRepository
func (*addressRepository) Read(data *request_body.DeleteOrReadAddress) ([]address.Address, error) {
	panic("unimplemented")
}

// ReadOne implements address.AddressRepository
func (*addressRepository) ReadOne(data *request_body.ReadOne) (*address.Address, error) {
	panic("unimplemented")
}

// Update implements address.AddressRepository
func (*addressRepository) Update(data *request_body.UpdateAddress) (*address.Address, error) {
	panic("unimplemented")
}

func AddressRepository(db *mongo.Database) address.AddressRepository {
	return &addressRepository{collection: *db.Collection("address")}
}
