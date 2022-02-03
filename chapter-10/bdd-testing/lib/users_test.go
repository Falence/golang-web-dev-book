package lib_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "web-dev-with-golang-book-by-shiju/chapter-10/bdd-testing/lib"
)

type FakeUserRepository struct {
	DataStore []User
}

func (repo *FakeUserRepository) GetAll() []User {
	return repo.DataStore
}

func (repo *FakeUserRepository) Create(user User) error {
	err := repo.Validate(user)
	if err!= nil {
		return err
	}
	repo.DataStore = append(repo.DataStore, user)
	return nil
}

func (repo *FakeUserRepository) Validate(user User) error {
	for _, u := range repo.DataStore {
		if u.Email == user.Email {
			return errors.New("the email already exists")
		}
	}
	return nil
}

func NewFakeUserRepo() *FakeUserRepository {
	return &FakeUserRepository{
		DataStore: []User{
			User{"Falence", "Lemungoh", "falence@lemungoh.com"},
			User{"Precious", "Zemoh", "precious@zemoh.com"},
			User{"Fiemina", "Chiafie", "fiemina@chiafie.com"},
		},
	}
}


var _ = Describe("Users", func() {
	BeforeEach(func() {

	})

	Describe("Get Users", func() {
		Context("Get all Users", func() {
			It("should get list of Users", func() {

			})
		})
	})

	Describe("Post a new User", func() {
		Context("Provide valid User data", func() {
			It("should create a new User and get HTTP status: 201", func() {})
		})
		Context("Provide User data that contains duplicate email id", func() {
			It("should get HTTP status: 400", func() {})
		})
	})
})
