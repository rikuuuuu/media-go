package main

import (
        // "context"
        "log"
        "net/http"
        "os"

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

        // adminHandler := handler.NewAdminUserService()
        twirpAdminUserHandler := pb.NewAdminUserServiceServer(&handler.AdminUserService{}, nil)

        mux := http.NewServeMux()

        mux.Handle(pb.ArticleServicePathPrefix, twirpArticleHandler)
        mux.Handle(pb.AdminUserServicePathPrefix, twirpAdminUserHandler)

        corsWrapper := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST", "GET"},
                AllowedHeaders: []string{
                        "Content-Type", 
                        "Authorization",
                },
        })
        
        h := corsWrapper.Handler(mux)

        http.HandleFunc("/", createAppHandler(h))

        port := os.Getenv("PORT")
        if port == "" {
                port = "8080"
                log.Printf("Defaulting to port %s", port)
        }

        log.Printf("Listening on port %s", port)
        if err := http.ListenAndServe(":"+port, twirpHandler(h, handler.NewContextProvider())); err != nil {
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

func twirpHandler(h http.Handler, contextProvider handler.ContextProvider) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {

                tokenString := r.Header.Get(authKey)

                var userID model.UserID

                if len(tokenString) >= 7 {
                        userID, _ = model.VerifyAccessToken(tokenString[7:])
                        ctx := r.Context()

                        handler.NewAdminUserService()
                        newContext, _ := contextProvider.WithAuthUID(ctx, userID)
                        h.ServeHTTP(w, r.WithContext(newContext))
                } else {
                        h.ServeHTTP(w, r)
                }
        }
}

func createAppHandler(h http.Handler) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {
                h.ServeHTTP(w, r)
        }
}