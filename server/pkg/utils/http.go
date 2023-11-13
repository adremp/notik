package utils

import (
	"net/http"
	"notik/internal/users/users_repo"
	"notik/pkg/httpErrors"
	"notik/pkg/logger"

	"github.com/labstack/echo/v4"
	"github.com/microcosm-cc/bluemonday"
)

var sanitizer *bluemonday.Policy

func init() {
	sanitizer = bluemonday.UGCPolicy()
}

func CreateCookie(token string, age int) *http.Cookie {
	return &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		MaxAge:   age,
		HttpOnly: true,
		Secure:   true,
	}
}

func SanitizeRequest[T any](c echo.Context, request *T) error {
	if err := c.Bind(request); err != nil {
		return err
	}

	sanitize(request)

	return Validate.StructCtx(c.Request().Context(), request)
}

func sanitize(data interface{}) {
	switch d := data.(type) {
	case map[string]interface{}:
		for k, v := range d {
			switch tv := v.(type) {
			case string:
				d[k] = sanitizer.Sanitize(tv)
			case map[string]interface{}:
				sanitize(tv)
			case []interface{}:
				sanitize(tv)
			case nil:
				delete(d, k)
			}
		}
	case []interface{}:
		if len(d) > 0 {
			switch d[0].(type) {
			case string:
				for i, s := range d {
					d[i] = sanitizer.Sanitize(s.(string))
				}
			case map[string]interface{}:
				for _, t := range d {
					sanitize(t)
				}
			case []interface{}:
				for _, t := range d {
					sanitize(t)
				}
			}
		}
	}
}

func ErrResponseWithLog(c echo.Context, log logger.Logger, err error) error {
	status, e := httpErrors.RequestError(err)
	log.Errorf("Error: %s; IP: %s; Causes: %+v", err, c.Request().RemoteAddr, e.ErrCauses)
	return c.JSON(status, e)
}

func GetUserFromCtx(c echo.Context) (*users_repo.User, error) {
	user, ok := c.Get("user").(users_repo.User)
	if !ok {
		return &user, httpErrors.ErrUnauthorized
	}
	return &user, nil
}
