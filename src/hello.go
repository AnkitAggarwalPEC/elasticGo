package main

import(
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("Main function")
	viper.SetConfigType("toml")
	viper.AddConfigPath("$GOPATH/config")
	viper.SetConfigName("conf")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error in reading conf", err)
	}else{
		a := viper.GetInt("a")
		fmt.Println(a)
	}

}
