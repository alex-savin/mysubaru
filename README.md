# MySubaru!GO
Is a simple API client to interact with My Subaru service (https://www.mysubaru.com/) via HTTP, written in Go

## News
  * v0.0.0-dev First public release on May 9, 2020

## Features
  * Simple and chainable methods for settings and request

## Installation
```bash
# Go Modules
require github.com/alex-savin/mysubaru
```

## Usage
The following samples will assist you to become as comfortable as possible with mysubaru library.
```go
// Import hassky into your code and refer it as `mysubaru`.
import "github.com/alex-savin/mysubaru"
```

#### Get Location request
```go
// Create a MySubaru Client
car, _ := mysubaru.New()

car.R().GetLocation()
```

#### Start/Stop Engine request
```go
car.R().StartEngine()

car.R().StopEngine()

```
