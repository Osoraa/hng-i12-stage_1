# hng-i12-stage_1 - Number Classification API

Creates an API which takes a number and returns math properties and a fun fact

## Tools

- golang ~= 1.23.5
- go gin web framework

## Installation & Run

```bash
# Clone & install deps
git clone https://github.com/Osoraa/hng-i12-stage_1.git
cd hng-i12-stage_1
go get .
go build .
sudo ./hng-i12-stage_1

# Or just download this project
go get github.com/Osoraa/hng-i12-stage_1
```

### Number struct used

```go
type nomba struct {
    Number     int      `json:"number"`
    Prime      bool     `json:"is_prime"`
    Perfect    bool     `json:"is_perfect"`
    Properties []string `json:"properties"`
    Sum        int      `json:"digit_sum"`
    Fact       string   `json:"fun_fact"`
}
```

```bash
# Build and Run
cd go-todo-rest-api-example
go build
./go-todo-rest-api-example

# API Endpoint : http://127.0.0.1[:80]
```

## API

### /api/clasify-number

- `GET` : Get math properties of randomly generated numbers and their fun fact

### /api/clasify-number?number=`num`

- `GET` : Get math properties of _num_ and a fact

  - Returns a `200` response when num is a valid integer with the example JSON below

  ```json
  {
    "number": 371,
    "is_prime": false,
    "is_perfect": false,
    "properties": ["armstrong", "odd"],
    "digit_sum": 11,
    "fun_fact": "371 is an Armstrong number because 3^3 + 7^3 + 1^3 = 371"
  }
  ```
  
  - Returns a `400` response if num is not a valid integer with the example JSON below
  
  ```json
  {
    "number": "num_passed",
    "error": "description"
  }  
  ```

## Tasks

- [x] Support the GET protocol
- [x] Enable and handle CORS
- [x] Handle default input
- [x] Handle locally computable properties
- [x] Get fun fact from [NumbersAPI](http://numbersapi.com)
- [x] Organize the code with packages
- [x] Document README.md
- [x] Build and test
