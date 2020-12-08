package main

import (
	"fmt",
	"github.com/spf13/viper"
)


func main(){
	vi:viper.New()
	vi.setConfigFile("validate.yaml")
	vi.ReadInConfi()
	fmt.Println(vi.GetInt("Port")
	fmt.Println(vi.GetString("hostname")

}
