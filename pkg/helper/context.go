package helper

import (
	"context"
	"time"
	"vote-system/internal/entity"
	"vote-system/pkg/constant"

	"github.com/gofiber/fiber/v2"
)

func CreateContextWithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

func CreateContextWithCustomTimeout(timeout int) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
}

func SetValueToContext(ctx context.Context, c *fiber.Ctx) context.Context {
	var userId, email string

	userId, ok := c.Locals("user-id").(string)
	if !ok {
		userId = "0"
	}
	email, ok = c.Locals("email").(string)
	if !ok {
		email = "0"
	}

	ctx = context.WithValue(ctx, constant.HeaderContext, entity.ValueContext{
		UserId: userId,
		Email:  email,
	})

	return context.WithValue(ctx, constant.FiberContext, c)
}

func GetValueContext(ctx context.Context) (valueCtx entity.ValueContext) {
	return ctx.Value(constant.HeaderContext).(entity.ValueContext)
}
