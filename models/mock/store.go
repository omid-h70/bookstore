package mock

type MockStore struct {
	userMockDB UserMockDB
}

func NewUserMockStore() MockStore {
	return MockStore{
		userMockDB: NewUserMockDB(),
	}
}
