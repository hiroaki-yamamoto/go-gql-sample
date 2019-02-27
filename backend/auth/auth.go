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

import "github.com/hiroaki-yamamoto/go-gql-sample/backend/models"

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

func composeToken(username string) ([]byte, error) {
  secret := os.Getenv("SECRET_KEY")
  if len(secret) < 1 {
    return nil, errors.New("No secret key specified")
  }
  signer := jwt.NewHS512(secret)
  now := time.Now().UTC()
  jot := &jwt.JWT{
    Issuer: "go-gql-sample",
    Subject: "identity",
    Audience: "go-gql-example",
    ExpirationTime: now.Add(2 * time.Hour).Unix(),
    NotBefore: now.Unix(),
    IssuedAt: now.Unix(),
    ID: username,
  }
  jot.SetAlgorithm(signer)
  jot.SetKeyID("identity")
  payload, err := jwt.Marshal(jot)
  if err != nil {
    return nil, err
  }
  return signer.Sign(payload)
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
      if len(token.ID) < 1 {
        next.ServeHTTP(w, r)
        log.Print("Not authenticated user")
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

// Login sets user session to cookie named "session"
func Login(w *http.ResponseWriter, user *models.User) {
  tok, err := composeToken(user.Username)
  if err != nil {
    log.Print(err)
    return
  }
  http.SetCookie(*w, &http.Cookie{
    Name: "session",
    Value: string(tok),
  })
}

// Logout sets empty username to cookie named "session"
func Logout(w *http.ResponseWriter, user *models.User) {
  tok, err := composeToken("")
  if err != nil {
    log.Print(err)
    return
  }
  http.SetCookie(*w, &http.Cookie{
    Name: "session",
    Value: string(tok),
  })
}
