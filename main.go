package main

import (
	"bank/handler"
	"bank/logs"
	"bank/repository"
	"bank/service"
	"fmt"
	"net/http"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func main() {
	// salt := make([]byte, 8)
	// http://www.ietf.org/rfc/rfc2898.txt
	// Salt.
	// c, _ := rand.Read()
	// fmt.Println(salt)
	// rand.Read(salt)
	// fmt.Println(salt)
	// df := pbkdf2.Key([]byte("some password"), salt, 4096, 32, sha1.New)
	// df2 := pbkdf2.Key([]byte("some password"), salt, 4096, 32, sha1.New)
	// fmt.Println(df)
	// fmt.Println(df2)
	// fmt.Println(bytes.Compare(df, df2))
	initTimeZone()
	initConfig()
	db := initDatabase()
	//customer
	customerRepositoryDB := repository.NewCustomerRepositoryDB(db)
	customerService := service.NewCustomerService(customerRepositoryDB)
	customerHandler := handler.NewCustomerHandler(customerService)

	//account
	accountRespositoryDB := repository.NewAccountRepositoryDB(db)
	accountService := service.NewAccountService(accountRespositoryDB)
	accountHandler := handler.NewAccountHandler(accountService)

	router := mux.NewRouter()
	//customer
	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)

	//account
	router.HandleFunc("/customers/{id:[0-9]+}/accounts", accountHandler.GetAccount).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}/accounts", accountHandler.NewAccount).Methods(http.MethodPost)
	logs.Info("Banking Service Start")
	http.ListenAndServe(fmt.Sprintf(":%v", viper.GetInt("db.port")), router)

}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = ict
}

func initDatabase() *sqlx.DB {
	dsn := fmt.Sprintf("%v:%v@/%v",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.database"),
	)

	db, err := sqlx.Open(viper.GetString("db.driver"), dsn)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(3 * time.Minute) // time out
	db.SetMaxOpenConns(10)                 //เปิดได้กี่เครื่อง
	db.SetMaxIdleConns(10)
	return db
}
