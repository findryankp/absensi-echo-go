package configs

import "absensi/docs"

func InitSwagger() {
	docs.SwaggerInfo.Title = "API E-Absensi"
	docs.SwaggerInfo.Description = "MIG Software Engineer"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "https://eabsensi-findryan.herokuapp.com"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.BasePath = "/"
}
