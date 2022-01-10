# rate-limiting-gin-go
## Table of Contents

 * [專案描述](#專案描述)
 * [執行專案](#執行專案)

## 專案描述
實作rate limiting middleware

## 執行專案

#### 執行應用程式

```bash
#到專案目錄下
$ cd path_to_dir/rate-limiting-gin-go

# 下載第三方套件
$ go mod download

# 生成swagger文檔
$ swag init 

# 啟動redis
$ docker-compose up 

# 編譯專案(輸出到當前目錄下,檔案名為main)
$ go build -o main . 

# 執行應用程式
$ ./main 
```

# swagger 頁面
```bash
$ http://localhost:8080/swagger/index.html
```
