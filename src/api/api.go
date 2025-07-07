package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/omidganji1380/golang-clean-web-api/api/routes"
	"github.com/omidganji1380/golang-clean-web-api/api/validations"
	"github.com/omidganji1380/golang-clean-web-api/config"
	"net"
)

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "unknown"
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			return ipnet.IP.String()
		}
	}
	return "unknown"
}
func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := val.RegisterValidation("mobile", validations.IranianMobileNumberValidator, true)
		if err != nil {
			return
		}
	}
	r.Use(gin.Logger(), gin.Recovery() /*, middlewares.TestMiddleware()*/)

	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		test_router := v1.Group("/test")
		routes.Health(health)
		routes.TestRouter(test_router)
	}
	fmt.Printf("%sLocal IP: http://localhost:%s%s\n", "\033[0;32m", cfg.Server.Port, "\033[0m")
	ip := getLocalIP()
	fmt.Printf("%sNetwork: http://%s:%s%s\n", "\033[0;32m", ip, cfg.Server.Port, "\033[0m")

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}
