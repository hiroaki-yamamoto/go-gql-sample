package auth

import (
  "context"
  "errors"
  "log"
  "net/http"
  "os"
  "time"
)

import (
  "github.com/gbrlsnchs/jwt"
  "github.com/jinzhu/gorm"
)

import (
  "github.com/hiroaki-yamamoto/go-gql-sample/backend/models"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
  name string
}

func extractToken(token string) (*jwt.JWT, error) {
  secret := os.Getenv("SECRET_KEY")
  if len(secret) < 1 {
    return nil, errors.New("No secret key specified")
  }
  signer := jwt.NewHS512(secret)
  now := time.Now().UTC()
  payload, sig, err := jwt.Parse(token)
  if err != nil {
    return nil, err
  }
  var jot jwt.JWT
  if err = jwt.Unmarshal(payload, &jot); err != nil {
    return nil, err
  }
  if err = signer.Verify(payload, sig); err != nil {
    return nil, err
  }
  err = jot.Validate(
    jwt.IssuedAtValidator(now),
    jwt.ExpirationTimeValidator(now),
    jwt.AudienceValidator("go-gql-example"),
  )
  if err != nil {
    return nil, err
  }
  return &jot, nil
}

// AuthenticationMiddleware is Middleware for authentication.
func AuthenticationMiddleware(db *gorm.DB) func(http.Handler) http.Handler {
  return func(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      c, err := r.Cookie("session")
      if err != nil || c == nil {
        next.ServeHTTP(w, r)
        log.Print(err)
        return
      }
      token, error := extractToken(c.Value)
      if error != nil {
        next.ServeHTTP(w, r)
        log.Print(error)
        return
      }
      var user models.User
      err = db.Where(&models.User{Username: token.ID}).First(&user).Error
      if err != nil {
        next.ServeHTTP(w, r)
        log.Print(error)
        return
      }
      ctx := context.WithValue(r.Context(), userCtxKey, user)
      r = r.WithContext(ctx)
      next.ServeHTTP(w, r)
    })
  }
}

// GetUser tries to retrieve user instance from current context.
func GetUser(ctx context.Context) *models.User {
  raw, _ := ctx.Value(userCtxKey).(*models.User)
  return raw
}
