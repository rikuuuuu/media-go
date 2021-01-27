package main

import (
        // "context"
        "log"
        "fmt"
        "net/http"
        "os"
        // "log"

        "mediago/handler"
        pb "mediago/pb"
        
        "github.com/rs/cors"
        "google.golang.org/appengine"
        "mediago/model"
)

const (
        authKey = "Authorization"
)

func twirpAPI() {
        twirpArticleHandler := pb.NewArticleServiceServer(&handler.ArticleService{}, nil)

        twirpAdminUserHandler := pb.NewAdminUserServiceServer(&handler.AdminUserService{}, nil)

        mux := http.NewServeMux()

        mux.Handle(pb.ArticleServicePathPrefix, twirpArticleHandler)
        mux.Handle(pb.AdminUserServicePathPrefix, twirpAdminUserHandler)

        corsWrapper := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST", "GET"},
                AllowedHeaders: []string{
                        "Content-Type", 
                        // "Authorization",
                },
        })
        
        handler := corsWrapper.Handler(mux)

        http.HandleFunc("/", createAppHandler(handler))

        port := os.Getenv("PORT")
        if port == "" {
                port = "8080"
                log.Printf("Defaulting to port %s", port)
        }

        log.Printf("Listening on port %s", port)
        // twirpHandler(handler)
        if err := http.ListenAndServe(":"+port, nil); err != nil {
                log.Fatal(err)
        }

        appengine.Main()
}

func main() {
        twirpAPI();
}

// func unAuthorizeError(w http.ResponseWriter) {
// 	http.Error(w, "unauthorized", 401)
// }

func twirpHandler(h http.Handler) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {

                tokenString := r.Header.Get(authKey)

                var userID model.UserID

                if len(tokenString) >= 7 {
                        userID, _ = model.VerifyAccessToken(tokenString[7:])
                        fmt.Printf("%v", userID)
                        // ctx := r.Context()

                        // newContext, _ := handler.ContextProvider.WithAuthUID(ctx, userID)
                        // fmt.Printf("%v", newContext)
                        // h.ServeHTTP(w, r.WithContext(newContext))
                        h.ServeHTTP(w, r)
                } else {
                        fmt.Printf("%v", "else")
                        h.ServeHTTP(w, r)
                }

                // h.ServeHTTP(w, r)

        }
}

func createAppHandler(h http.Handler) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {
                h.ServeHTTP(w, r)
        }
}