package main

import (
	"github.com/HaitaoYue/myweb/utils"
	"github.com/gin-gonic/gin"
	"log"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	r := gin.Default()

	SetupRouter(r)

	return r
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	r := setupRouter()
	r.Use(GinBodyLogMiddleware)

	log.Printf("%s Start service %s", utils.Green, utils.Reset)

	r.Run(":8080")
}

//func loadTemplate() (*template.Template, error) {
//	t := template.New("")
//	for name, file := range Assets.Files {
//		if file.IsDir() || !strings.HasSuffix(name, ".tmpl") {
//			continue
//		}
//		h, err := ioutil.ReadAll(file)
//		if err != nil {
//			return nil, err
//		}
//		t, err = t.New(name).Parse(string(h))
//		if err != nil {
//			return nil, err
//		}
//	}
//	return t, nil
//}

