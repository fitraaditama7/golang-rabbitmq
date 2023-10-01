package middleware

import (
	"account-service/constants"
	"account-service/utils/logger"
	"account-service/utils/responses"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func Config() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		rid := c.Get(fiber.HeaderXRequestID)

		if rid == "" {
			rid = uuid.New().String()
			c.Set(fiber.HeaderXRequestID, rid)
		}

		header := make(map[string]string)
		c.Request().Header.VisitAll(func(key, val []byte) {
			k := bytes.NewBuffer(key).String()
			header[k] = bytes.NewBuffer(val).String()
		})

		headerByte, _ := json.Marshal(header)
		fields := logger.Fields{
			Id:       rid,
			RemoteIp: c.IP(),
			Method:   c.Method(),
			Host:     c.Hostname(),
			Protocol: c.Protocol(),
			StartAt:  time.Now(),
		}

		defer func() {
			var err error
			rvr := recover()
			if rvr != nil {
				var ok bool
				err, ok = rvr.(error)
				if !ok {
					err = fmt.Errorf("%v", rvr)
				}

				resp := responses.Error(constants.ErrorInternalServer)
				c.Status(http.StatusInternalServerError).JSON(resp)
			}

			logger.LogMiddleware(c.Path(),
				c.Context(),
				headerByte,
				c.Request().URI().QueryString(),
				c.Body(),
				c.Response().Body(),
				c.Response().StatusCode(),
				fields,
				err)
		}()

		return c.Next()
	}
}
