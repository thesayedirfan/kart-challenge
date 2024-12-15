# oolio backend challege

### Installation
1. have latest version of go installed
    *  verify the installation of go

    ```bash
    go version
    ```
2. Clone the repository 
    ```bash
    git clone https://github.com/thesayedirfan/kart-challenge

    cd kart-challenge
    ```
3. Start The Project
    ```bash
    make run

    ```

## Project Structure

```bash

├── Makefile
├── README.md
├── api
│   ├── openapi.yaml
│   ├── order_api.go
│   └── product_api.go
├── backend-challenge
│   └── README.md
├── cmd
│   └── main.go
├── design.fig
├── domain
│   ├── order.go
│   └── product.go
├── errors
│   └── errors.go
├── go.mod
├── go.sum
├── middleware
│   └── api_key.go
├── repository
│   ├── order_repository.go
│   ├── order_repository_test.go
│   ├── product_repository.go
│   └── product_repository_test.go
├── usecase
│   ├── order_usecase.go
│   └── product_usecase.go
└── utils
    ├── coupons.go
    └── discount.go

```


> It seems the repository has exceeded its data quota for Git LFS (Large File Storage), and this is preventing me from pushing these files containing coupon codes


## Approach Used For Reading Coupon Files
file location
```bash
utils
    ├── coupons.go
```

- we read all the files from the directory called coupons
- we create goroutines to read each file and also creating a local map of unique codes
- we merge all the local maps code count into a global map we use Mutexes to deal with concurrency
- we can also used sync.map but it is optmised for Read-Heavy Workloads with Occasional Writes



