package storage

type Storage interface {
	User() UserRepository
	Order() OrderRepository
	Good() GoodRepository
}
