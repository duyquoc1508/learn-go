package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"

	config "go-demo/configs"
	dto "go-demo/dto"
	model "go-demo/models"
	repoImpl "go-demo/repositories/repoImpl"
	util "go-demo/utils"
)

var jwtKey []byte = []byte(config.JWT_SECRET_KEY)

type Claims struct {
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`
	jwt.StandardClaims
}

func Register(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var payload dto.Register
	err := json.NewDecoder(request.Body).Decode(&payload)
	if err != nil {
		util.ResponseErr(response, http.StatusBadRequest, "Đăng ký thất bại")
		return
	}

	_, err = repoImpl.NewUserRepo().FindUserByEmail(payload.Email)
	if err == nil { // email đã tồn tại
		util.ResponseErr(response, http.StatusConflict, "Email đã được sử dụng")
		return
	}
	hashedPassword, err := hashPassword(payload.Password)
	if err != nil {
		util.ResponseErr(response, http.StatusInternalServerError, err.Error())
		return
	}
	user := model.User{
		Email:       payload.Email,
		Password:    hashedPassword,
		DisplayName: payload.DisplayName,
	}
	err = repoImpl.NewUserRepo().Insert(&user)
	if err != nil {
		util.ResponseErr(response, http.StatusInternalServerError, "Đăng ký thất bại")
		return
	}

	tokenString, err := genToken(user)
	if err != nil {
		util.ResponseErr(response, http.StatusInternalServerError, "Lấy token thất bại")
		return
	}
	responseData := struct {
		AccessToken string `json:"accessToken"`
	}{
		AccessToken: tokenString,
	} // anonymous struct
	util.ResponseOk(response, http.StatusOK, "Đăng ký thành công", responseData)
}

func Login(response http.ResponseWriter, request *http.Request) {
	var payload dto.Login
	err := json.NewDecoder(request.Body).Decode(&payload)
	if err != nil {
		util.ResponseErr(response, http.StatusInternalServerError, "Decode thất bại")
		return
	}
	var user model.User
	user, err = repoImpl.NewUserRepo().FindUserByEmail(payload.Email)
	if err != nil {
		util.ResponseErr(response, http.StatusNotFound, "Không tìm thấy thông tin tài khoản")
		return
	}
	checkPassword := checkPasswordHash(payload.Password, user.Password)
	if !checkPassword {
		fmt.Println("Sai mật khẩu")
		util.ResponseErr(response, http.StatusUnauthorized, "Tên đăng nhập hoặc mật khẩu không đúng")
		return
	}
	tokenString, err := genToken(user)
	if err != nil {
		util.ResponseErr(response, http.StatusInternalServerError, "Đăng nhập thất bại")
		return
	}
	responseData := struct {
		AccessToken string `json:"accessToken"`
	}{
		AccessToken: tokenString,
	}
	util.ResponseOk(response, http.StatusOK, "Đăng nhập thành công", responseData)
}

func GetUser(response http.ResponseWriter, request *http.Request) {
	tokenClaims, ok := request.Context().Value(Claims{}).(*Claims) // get data from request
	if !ok {
		fmt.Println("Get claim data from context failed")
	}
	user, err := repoImpl.NewUserRepo().FindUserByEmail(tokenClaims.Email)
	if err != nil {
		util.ResponseErr(response, http.StatusNotFound, "Không tìm thấy thông tin user")
		return
	}
	util.ResponseOk(response, http.StatusOK, "Lấy thông tin user thành công", user)
}

func genToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24) // 1 day
	claims := &Claims{
		Email:       user.Email,
		DisplayName: user.DisplayName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
