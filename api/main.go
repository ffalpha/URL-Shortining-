pacakge main

import(
	"github.com/gofiber/fiver/v2"
)


func setupRoutes(app * fiber.App){
	app.get("/:url",routes.ResolveURL)
	app.Post("/api/v1",routes.ShortenURL)
}

func main(){
   err:= godotenv.Load()
   app := fiber.New()
}