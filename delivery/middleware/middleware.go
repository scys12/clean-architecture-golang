package middleware

import (
	"errors"
	"net/http"

	"github.com/scys12/clean-architecture-golang/pkg/payload/response"

	"github.com/labstack/echo/v4"
	"github.com/scys12/clean-architecture-golang/pkg/session"
)

func SessionMiddleware(s session.SessionStore, role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			cookie, err := ctx.Cookie("sessionID")
			if err != nil {
				response.Error(ctx, http.StatusUnauthorized, err)
				return err
			}
			sessionID := cookie.Value
			sess, err := s.Get(sessionID)
			if err != nil {
				response.Error(ctx, http.StatusInternalServerError, err)
				return err
			}
			if role != sess.UserRole {
				response.Error(ctx, http.StatusUnauthorized, errors.New("Wrong authorization"))
				return nil
			}
			ctx.Set("sessionID", sessionID)
			ctx.Set("userID", sess.UserID.String())
			return next(ctx)
		}
	}
}

func CORS() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
			c.Response().Header().Set("Accept", "application/json")
			return next(c)
		}
	}
}
