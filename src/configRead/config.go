package configRead

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type Config struct {
	Name string
}

//var cfg = pflag.StringP("config", "c", "", "apiserver config file path.")

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	if err := c.initConfig(); err != nil {
		return err
	}
	c.watchConfig()
	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name)
	} else {
		viper.AddConfigPath("src/configRead") // 如果没有指定配置文件，则解析默认的配置文件
		viper.SetConfigName("conf")
	}
	viper.SetConfigType("yaml") // 设置读取文件格式
	viper.AutomaticEnv()        // 读取匹配的环境变量
	viper.SetEnvPrefix("viperDemo")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}

	return nil
}

func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
	})
}

//import (
//"fmt"
//"github.com/spf13/viper"
//"goExample/src/configRead"
//)
//import "github.com/spf13/pflag"
//
//var cfg = pflag.StringP("config", "c", "", "config file path.")
//func main() {
//	pflag.Parse()
//	if err := configRead.Init(*cfg); err != nil {
//		panic(err)
//	}
//
//	fmt.Println(viper.GetString("common.database.passwd"))
//}
