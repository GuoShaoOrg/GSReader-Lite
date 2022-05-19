package middlewear

import (
	"context"
	"encoding/json"
	"errors"
	"gs-reader-lite/server/api/controller"
	"gs-reader-lite/server/component"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/crypto/gaes"
	"github.com/gogf/gf/v2/encoding/gbase64"
)

type TokenModel struct {
	UserId         string `json:"userId"`
	UserName       string `json:"username"`
	NickName       string `json:"nickname"`
	Mobile         string `json:"mobile"`
	CreateDate     string `json:"createDate"`
	UpdateDateTime string `json:"updateDateTime"`
	Role           string `json:"role"`
	Token          string `json:"token"`
}

var privateKey = "rsshub-tm01-12-1"

func AuthToken() gin.HandlerFunc {
	return func(request *gin.Context) {
		authorization := request.GetHeader("Authorization")
		token, uid, err := validateToken(request.Request.Context(), authorization)
		if err != nil {
			controller.JsonExitWithStatus(request, http.StatusUnauthorized, 0, "StatusUnauthorized", nil)
			request.Abort()
		}

		if tokenModel, err := ParseToken(token); err != nil {
			component.Logger().Info(request.Request.Context(), "AuthToken invalid")
			controller.JsonExitWithStatus(request, http.StatusUnauthorized, 0, "StatusUnauthorized", nil)
			request.Abort()
		} else {
			if tokenModel.UserId != uid {
				component.Logger().Info(request.Request.Context(), "token invalid tokenModel : ", tokenModel, " ,uid : ", uid)
				controller.JsonExitWithStatus(request, http.StatusUnauthorized, 0, "StatusUnauthorized", nil)
				request.Abort()
			}
		}
	}
}

func CookieToken() gin.HandlerFunc {
	return func(request *gin.Context) {
		authorization, err := request.Cookie("Auth")
		if err != nil {
			request.Params = append(request.Params, gin.Param{Key: "msg", Value: "您还没有登录"}, gin.Param{Key: "title", Value: "Error"})
			request.Redirect(http.StatusTemporaryRedirect, "/view/user/login")
			request.Abort()
		}
		token, uid, err := validateToken(request.Request.Context(), authorization)
		if err != nil {
			request.Params = append(request.Params, gin.Param{Key: "msg", Value: "您还没有登录"}, gin.Param{Key: "title", Value: "Error"})
			request.Redirect(http.StatusTemporaryRedirect, "/view/user/login")
			request.Abort()
		}

		if tokenModel, err := ParseToken(token); err != nil {
			component.Logger().Info(request.Request.Context(), "AuthToken invalid")
			request.Params = append(request.Params, gin.Param{Key: "msg", Value: "身份验证出错"}, gin.Param{Key: "title", Value: "Error"})
			request.Redirect(http.StatusTemporaryRedirect, "/view/user/login")
			request.Abort()
		} else {
			if tokenModel.UserId != uid {
				component.Logger().Info(request.Request.Context(), "token invalid tokenModel : ", tokenModel, " ,uid : ", uid)
				request.Params = append(request.Params, gin.Param{Key: "msg", Value: "身份验证出错"}, gin.Param{Key: "title", Value: "Error"})
				request.Redirect(http.StatusTemporaryRedirect, "/view/user/login")
				request.Abort()
			}
		}
	}
}

func validateToken(cxt context.Context, authString string) (token, uid string, err error) {
	authorizationArray := strings.Split(authString, "@@")
	if len(authorizationArray) < 2 {
		component.Logger().Info(cxt, "Token or uid is null")
		return "", "", errors.New("Token or uid is null")
	}
	token = authorizationArray[0]
	uid = authorizationArray[1]
	if len(token) < 0 || len(uid) < 0 {
		component.Logger().Info(cxt, "AuthToken or uid is null")
		return "", "", errors.New("AuthToken or uid is null")
	}
	return token, uid, nil
}

func ParseToken(tokenString string) (*TokenModel, error) {
	decodeToken, _ := gbase64.Decode([]byte(tokenString))
	decResult, err := gaes.Decrypt(decodeToken, []byte(privateKey))
	ctx := context.Background()
	if err != nil {
		component.Logger().Info(ctx, "token decrypt error : ", tokenString)
		return nil, err
	}
	tokenModel := new(TokenModel)
	if err := json.Unmarshal(decResult, tokenModel); err != nil {
		component.Logger().Info(ctx, "token string decode to json error , token: ", decResult, " ,error : ", err)
		return nil, err
	}
	return tokenModel, nil
}

func CreateToken(tokenData TokenModel) (string, error) {
	ctx := context.Background()
	if jsonToken, err := json.Marshal(tokenData); err != nil {
		component.Logger().Info(ctx, "decode token to json error : ", err)
		return "", err
	} else {
		if token, err := gaes.Encrypt(jsonToken, []byte(privateKey)); err != nil {
			component.Logger().Info(ctx, "aes encrypt string error: ", err)
			return "", err
		} else {
			encodeToken := gbase64.EncodeToString(token)
			return encodeToken, nil
		}
	}
}
