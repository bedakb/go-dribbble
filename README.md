# Go Dribbble
Unofficial Go library for interacting with Dribbble API v2 (WIP).
Please refer to official [Dribbble API v2 Docs](http://developer.dribbble.com/v2/) for more information about API itself.
 
## Install
```
go get github.com/bedakb/go-dribbble
```

## Usage
```
import "github.com/bedakb/go-dribbble"

cfg := dribbble.NewConfig(accessToken)
d, _ := dribbble.New(cfg)

// Get currently logged in user
user, _ := d.User.GetUser()
fmt.Printf("%v", user)
```
