package http

import (
	"context"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func JSONMiddleWare(next http.Handler) http.Handler{

    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
        
        //do somthing before
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        next.ServeHTTP(w,r)
        //do something after
    })
}

func LoggingMiddleWare(next http.Handler) http.Handler{

    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
        
      
        log.WithFields(log.Fields{
            "method": r.Method,
            "path": r.URL.Path,
            "IP": r.RemoteAddr,
        }).Info("handled request")
        next.ServeHTTP(w,r)
 
    })
}

func  TimeOutMiddleWare(next http.Handler) http.Handler{

    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
        ctx, cancle := context.WithTimeout(r.Context(), 25 *time.Second)
        defer cancle()
        next.ServeHTTP(w,r.WithContext(ctx))
 
    })
}
