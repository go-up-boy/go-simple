package captcha

import (
	"github.com/mojocn/base64Captcha"
	"go-simple/app"
	"go-simple/pkg/config"
	"go-simple/pkg/redis"
	"sync"
)

type Captcha struct {
	Base64Captcha *base64Captcha.Captcha
}

var once sync.Once

var internalCaptcha *Captcha

func NewCaptcha() *Captcha {
	once.Do(func() {
		internalCaptcha = &Captcha{}

		store := RedisStore{
			RedisClient: redis.Redis,
			keyPrefix: config.GetString("app.name") + ":captcha",
		}

		driver := base64Captcha.NewDriverDigit(
			config.GetInt("captcha.height"),
			config.GetInt("captcha.width"),
			config.GetInt("captcha.length"),
			config.GetFloat64("captcha.maxskew"),
			config.GetInt("captcha.dotcount"),
			)
		internalCaptcha.Base64Captcha = base64Captcha.NewCaptcha(driver, &store)
	})
	return internalCaptcha
}

func (c *Captcha) GenerateCaptcha() (id string, b64s string, err error) {
	return c.Base64Captcha.Generate()
}

func (c *Captcha) VerifyCaptcha(id string, answer string, clear ...bool) (match bool) {
	if !app.IsProduction() && id == config.GetString("captcha.testing_key") {
		return true
	}
	var _clear bool = false
	if len(clear) > 0 && clear[0] == true {
		_clear = true
	}
	return c.Base64Captcha.Verify(id, answer, _clear)
}
