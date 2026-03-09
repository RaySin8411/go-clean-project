## 專案目標

使用者認證 API 服務

1. 使用者註冊：允許新使用者建立帳號。
2. 使用者登入：驗證使用者身份並核發 JWT (JSON Web Token)。
3. 使用者資訊查詢：一個受保護的端點，使用者必須提供有效的 JWT 才能存取自己的資訊。


## 技術工具箱 (Tech Stack)

* 語言 (Language): Go
* Web 框架 (Framework): Gin - 一個高效能、API 友善的 Web 框架。
* 資料庫互動 (ORM): GORM - 功能強大且廣泛使用的 ORM 函式庫。
* 資料庫遷移 (Migration): golang-migrate - 用於管理資料庫結構的變更。
* 容器化 (Containerization): Docker & Docker Compose - 用於建立一致、可攜的開發與部署環境。
* 設定管理 (Configuration): Viper - 強大的設定檔管理工具，能處理多種格式的設定檔與環境變數。
* API 文件 (API Docs): Swagger/OpenAPI - 自動生成互動式 API 文件，方便前後端協作。
* 程式碼品質 (Linter): golangci-lint - 集合多種 Linter 的工具，強力保障程式碼風格與品質。


## 準備工作 (Prerequisites)

* Go 
* Docker 和 Docker Compose
* VS Code or GoLand(程式碼編輯器)


## 建立專案結構
```cmd
# 建立專案根目錄
mkdir go-clean-project
cd go-clean-project

# 建立核心目錄結構
mkdir -p cmd/api
mkdir -p internal
mkdir -p deployments
mkdir -p docs
mkdir -p pkg

# 初始化 Go Modules
go mod init go-clean-project
```

* cmd/ 

    專案的進入點。如果我們的專案未來有多個執行檔（例如一個 API 服務，一個背景 worker），它們的 main.go 都會放在這裡。

* internal/ 
    
    專案的核心邏輯所在。這是 Go 語言的一個特殊目錄，放在這裡的程式碼只能被同一個專案（go-clean-project）下的程式碼引用，無法被外部專案匯入，這有助於保護你的內部實作。

* deployments/ 
    
    存放與部署相關的檔案，例如資料庫遷移檔案、Dockerfile 等。

* docs/
    
    存放專案文件，以及 Swagger API 文件。

* pkg/
    
    存放可以被外部專案安全引用的程式碼。在我們的專案初期，可能不會用到太多，但建立這個目錄是一個好的實踐。

## 參考
* [Go Clean Architecture API 開發全攻略](https://ithelp.ithome.com.tw/users/20128815/ironman/8486)