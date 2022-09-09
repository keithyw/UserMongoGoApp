package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gorilla/mux"
	"github.com/golang/mock/gomock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"messageapp/mocks"
	"messageapp/models"
)

func TestUserPostHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	svc := mocks.NewMockUserService(mockCtrl)
	controller := &UserController{Svc: svc}

	var u models.User
	var postJson = []byte(`{"name":"testuser"}`)

	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(postJson))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	user := models.User{
		Name: "testuser",
	}

	objId, _ := primitive.ObjectIDFromHex("6319073602de3f59eb3b9853")
	newUser := models.User{
		Name: "testuser",
		Id: objId,
	}

	svc.EXPECT().FindUserByName(user.Name).Return(nil, nil).Times(1)
	svc.EXPECT().CreateUser(user).Return(&newUser, nil).Times(1)
	
	controller.UserPostHandler(w, req)

	response := w.Result()
	defer response.Body.Close()
	jsonStr, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fail()
	}

	if response.StatusCode != http.StatusOK {
		t.Fail()
	}

	err = json.Unmarshal(jsonStr, &u)
	if err != nil {
		t.Fail()
	}

	if u.Name != newUser.Name {
		t.Fail()
	}
}

func TestUserUpdateHandler(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	svc := mocks.NewMockUserService(mockCtrl)
	controller := &UserController{Svc: svc}
	var postJson = []byte(`{"name":"testuser2"}`)
	var u models.User

	objId, _ := primitive.ObjectIDFromHex("6319073602de3f59eb3b9853")

	req := httptest.NewRequest(http.MethodPut, "/users/" + objId.Hex(), bytes.NewBuffer(postJson))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	updateUser := models.User{
		Name: "testuser2",
	}

	foundUser := models.User{
		Name: "testuser2",
		Id: objId,
	}

	params := map[string]string{
		"id": objId.Hex(),
	}
	req = mux.SetURLVars(req, params)

	svc.EXPECT().UpdateUser(objId.Hex(), updateUser).Return(&foundUser, nil).Times(1)

	controller.UserUpdateHandler(w, req)
	response := w.Result()
	defer response.Body.Close()
	jsonStr, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fail()
	}

	if response.StatusCode != http.StatusOK {
		t.Fail()
	}

	err = json.Unmarshal(jsonStr, &u)
	if err != nil {
		t.Fail()
	}
	if u.Name != updateUser.Name {
		t.Fail()
	}
}

func TestUserDeleteHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	svc := mocks.NewMockUserService(mockCtrl)
	controller := &UserController{Svc: svc}

	objId, _ := primitive.ObjectIDFromHex("6319073602de3f59eb3b9853")

	req := httptest.NewRequest(http.MethodDelete, "/users/" + objId.Hex(), nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	params := map[string]string{
		"id": objId.Hex(),
	}
	req = mux.SetURLVars(req, params)

	svc.EXPECT().DeleteUser(objId.Hex()).Return(nil).Times(1)

	controller.UserDeleteHandler(w, req)

}