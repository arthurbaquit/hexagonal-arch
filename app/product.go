package app

import "errors"

type ProductInterface interface {
	isValid() (bool, error)
	Enable() error
	Disable() error
	GetId() string
	GetStatus() string
	GetName() string
	GetPrice() int64 // in decimals => 1000 = 10.00
}

const (
	ProductStatusEnabled  = "enabled"
	ProductStatusDisabled = "disabled"
)

type Product struct {
	Id     string `json:"id"`
	Status string `json:"status"`
	Name   string `json:"name"`
	Price  int64  `json:"price"`
}

func (p *Product) IsValid() (bool, error) {
	if p.Name == "" {
		return false, errors.New("Product name is required")
	}
	if p.Price <= 0 {
		return false, errors.New("Product price must be greater than zero")
	}
	if p.Status != ProductStatusEnabled && p.Status != ProductStatusDisabled {
		return false, errors.New("Product status should be enable or disabled")
	}
	return true, nil
}

func (p *Product) Enable() error {
	if p.Status == ProductStatusEnabled {
		return errors.New("Product is already enabled")
	}
	if p.Price <= 0 {
		return errors.New("Product price must be greater than zero to be enabled")
	}
	p.Status = ProductStatusEnabled
	return nil
}

func (p *Product) Disable() error {
	if p.Status == ProductStatusDisabled {
		return errors.New("Product is already disabled")
	}
	p.Status = ProductStatusDisabled
	return nil
}

func (p *Product) GetId() string {
	return p.Id
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetPrice() int64 {
	return p.Price
}
