package services

import (
	"os"
	"testing"
	"github.com/golang/mock/gomock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"messageapp/mocks"
	"messageapp/models"
)

var userService UserService

func TestMain(m *testing.M) {
	
	println("before test")
	ret := m.Run()
	println("after test")
	os.Exit(ret)
}

func TestCreateUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repo := mocks.NewMockUserRepository(mockCtrl)
	service := &UserServiceImpl{Ur: repo}
	user := models.User{
		Name: "testuser",
	}

	objId, err := primitive.ObjectIDFromHex("6319073602de3f59eb3b9853")
	if err != nil {
		t.Fail()
	}

	newUser := models.User{
		Name: "testuser",
		Id: objId,
	}

	repo.EXPECT().CreateUser(user).Return(&newUser, nil).Times(1)

	_, err = service.CreateUser(user)
	if err != nil {
		t.Fail()
	}
}

func TestDeleteUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repo := mocks.NewMockUserRepository(mockCtrl)
	service := &UserServiceImpl{Ur: repo}

	repo.EXPECT().DeleteUser("123").Return(nil).Times(1)

	err := service.DeleteUser("123")
	if err != nil {
		t.Fail()
	}
}

func TestUpdateUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repo := mocks.NewMockUserRepository(mockCtrl)
	service := &UserServiceImpl{Ur: repo}

	user := models.User{
		Name: "updated user",
	}

	updatedUser := models.User{
		Name: "updated user",
	}

	repo.EXPECT().UpdateUser("123", user).Return(&updatedUser, nil).Times(1)

	ret, err := service.UpdateUser("123", user)
	if err != nil {
		t.Fail()
	}

	if ret.Name != user.Name {
		t.Fail()
	}
}

func TestFindUserById(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repo := mocks.NewMockUserRepository(mockCtrl)
	service := &UserServiceImpl{Ur: repo}

	objId, _ := primitive.ObjectIDFromHex("6319073602de3f59eb3b9853")

	user := models.User{
		Name: "test user",
		Id: objId,
	}

	repo.EXPECT().FindUserById("123").Return(&user, nil).Times(1)

	foundUser, err := service.FindUserById("123")
	if err != nil {
		t.Fail()
	}

	if foundUser.Name != user.Name {
		t.Fail()
	}
}

func TestFindUserByName(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repo := mocks.NewMockUserRepository(mockCtrl)
	service := &UserServiceImpl{Ur: repo}

	objId, _ := primitive.ObjectIDFromHex("6319073602de3f59eb3b9853")
	user := models.User{
		Name: "testuser",
		Id: objId,
	}

	repo.EXPECT().FindUserByName("testuser").Return(&user, nil).Times(1)

	foundUser, err := service.FindUserByName("testuser")
	if err != nil {
		t.Fail()
	}
	if foundUser.Id != user.Id {
		t.Fail()
	}
}

func TestGetAllUsers(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repo := mocks.NewMockUserRepository(mockCtrl)
	service := &UserServiceImpl{Ur: repo}

	objId, _ := primitive.ObjectIDFromHex("6319073602de3f59eb3b9853")
	var users = []models.User{
		models.User{
			Name: "testuser1",
			Id: objId,
		},
	}

	repo.EXPECT().GetAllUsers().Return(users, nil).Times(1)

	allUsers, err := service.GetAllUsers()
	if err != nil {
		t.Fail()
	}

	if len(allUsers) != 1 {
		t.Fail()
	}

	if allUsers[0].Id != objId {
		t.Fail()
	}
}