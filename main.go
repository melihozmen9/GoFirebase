package main

import (
	"context"
	firebase "firebase.google.com/go"
	"log"
	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/melihozmen9/GoFirebase/handlers"
)

func main() {

	_, err := createClient()
	if err != nil {
		log.Printf("error creating the client %v", err)
		return
	}

	r := gin.Default()

	r.GET("/api/health", handlers.HealtCheckHandler())

	r.POST(
		"/api/health",
		func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"ok":true})
		}
	)

	r.Run("127.0.0.1:9090")
}

func createClient() (*firestore.Client, error) {
// test the local connection
ctx := context.Background() // context oluşturduk.
conf := &firebase.Config{ProjectID: "myfirstgoproject-6e761"} // configurasyon oluşturduk.
app, err := firebase.NewApp(ctx, conf) 

if err != nil { // connectionla ilgili hata var mı onu kontrol ediyoruz.
	log.Printf("error initializing app %v", err)
	return nil, err
}

client, err := app.Firestore(ctx) // client'i alıyoruz.
if err != nil {
	 log.Print("error when creating client")
	 log.Fatalln(err) // işlemi sonlandırır.
}
return client, nil
}
// eğer her şey yolunda ise reference oluşutuyoruz.
// ref := client.Collection("todos").NewDoc() // client'dan yeni bir collection oluşturduk.  ve newDoc ile yeni biritem oluştruduk.
// result, err := ref.Set(ctx, map[string]interface{}{
//	"title": "23 ocak akşamı işi",
//	"description": "23 ocak akşamı işinin detayı",
// })
// if err != nil {
// 	log.Printf("error happened when creating a todo: %v", err)
// 	return
// }
// log.Printf("Result is [%v]", result )
