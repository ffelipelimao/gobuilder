package mockbuilder

type CustomerMockBuilder struct {
	mock entities.Customer
}

func Customer() *CustomerMockBuilder {
	return &CustomerMockBuilder{
		mock: MockCustomer,
	}
}

func (b *CustomerMockBuilder) Build() entities.Customer {
	return b.mock
}

func (b *CustomerMockBuilder) WithAddress(value string) entities.Customer {
	b.mock.Address = value
	return b.mock
}

func (b *CustomerMockBuilder) WithName(value string) entities.Customer {
	b.mock.Name = value
	return b.mock
}

func (b *CustomerMockBuilder) WithPhone(value int64) entities.Customer {
	b.mock.Phone = value
	return b.mock
}
