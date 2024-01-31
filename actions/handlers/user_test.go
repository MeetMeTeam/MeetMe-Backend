package handlers

//
//import (
//	"meetme/be/actions/repositories/interfaces"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//
//	"github.com/labstack/echo/v4"
//	"github.com/stretchr/testify/assert"
//)
//
//var (
//	mockDB = map[string]*interfaces.User{
//		"jon@labstack.com": &User{"Jon Snow", "jon@labstack.com"},
//	}
//	userJSON = `{
//		"data": [
//			{
//				"id": "65b8eed7a2652927683f7136",
//				"username": "String",
//				"displayName": "String",
//				"birthday": "Date",
//				"email": "String",
//				"image": ""
//			}
//		],
//		"message": "Get users success.[test automated deploy]"
//	}`
//)
//
//func TestGetAllUser(t *testing.T) {
//	// Setup
//	e := echo.New()
//	req := httptest.NewRequest(http.MethodGet, "/api", nil)
//	rec := httptest.NewRecorder()
//	c := e.NewContext(req, rec)
//	c.SetPath("/users")
//	//c.SetParamNames("email")
//	//c.SetParamValues("jon@labstack.com")
//	//h := &handler{mockDB}
//	//
//	//// Assertions
//	if assert.NoError(t, h.getUser(c)) {
//		assert.Equal(t, http.StatusOK, rec.Code)
//		assert.Equal(t, userJSON, rec.Body.String())
//	}
//}
