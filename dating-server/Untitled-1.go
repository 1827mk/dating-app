
create golang "identity-service" 
-read config ./conf/app_setting.yaml
-add log config
-add redis config
-add database postgresql
-http use echo
and can up down with load balance on k8s

feature
-register
-login/logout
-delete account
    type User struct {
        ID        uint      `json:"id" gorm:"primaryKey"`
        Email     string    `json:"email" gorm:"unique"`
        Password  string    `json:"-"`
        Name      string    `json:"name"`
        Bio       string    `json:"bio"`
        Age       int       `json:"age"`
        Gender    string    `json:"gender"`
        Location  Location  `json:"location" gorm:"embedded"`
        Photos    []string  `json:"photos" gorm:"type:text[]"`
        CreatedAt time.Time `json:"created_at"`
        UpdatedAt time.Time `json:"updated_at"`
        UpdatedAt time.Time `json:"updated_at"`
    }

    type Location struct {
        Latitude  float64 `json:"latitude"`
        Longitude float64 `json:"longitude"`
    }


Applications feature like tinder app 
-login and auth
-user profile 
-match card 
-near location (ios,android) gps location