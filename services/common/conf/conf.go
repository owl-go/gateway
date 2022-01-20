package conf

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var (
	cfg = config{}
	// Global 全局设置
	Global = &cfg.Global
	// Log 日志级别设置
	Log = &cfg.Log
	// Etcd Etcd设置
	Etcd = &cfg.Etcd

	Redis = &cfg.Redis

	Mysql = &cfg.Mysql

	//GCS
	Gcs = &cfg.Gcs

	// Nats 消息中间件设置
	Nats = &cfg.Nats
	//http探针
	Probe = &cfg.Probe
)

func init() {
	if !cfg.parse() {
		showHelp()
		os.Exit(-1)
	}
}

type global struct {
	Pprof   string `mapstructure:"pprof"`
	Version string `mapstructure:"version"`
	Region  string `mapstructure:"region"`
	Zone    string `mapstructure:"zone"`
	Name    string `mapstructure:"name"`
	Nid     string `mapstructure:"nid"`
	Nip     string `mapstructure:"nip"`
	Port    int    `mapstructure:"port"`
	Key     string `mapstructure:"key"`
	Cert    string `mapstructure:"cert"`
}

type log struct {
	Level string `mapstructure:"level"`
}
type redis struct {
	Addrs    []string `mapstructure:"addrs"`
	Password string   `mapstructure:"password"`
	DB       int      `mapstructure:"db"`
}
type mysql struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DB       string `mapstructure:"dbname"`
}

type etcd struct {
	Addrs string `mapstructure:"addrs"`
}

type nats struct {
	URL     string `mapstructure:"url"`
	NatsLog string `mapstructure:"natslog"`
}

type probe struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Cert string `mapstructure:"cert"`
	Key  string `mapstructure:"key"`
}
type config struct {
	Global  global `mapstructure:"global"`
	Log     log    `mapstructure:"log"`
	Etcd    etcd   `mapstructure:"etcd"`
	Nats    nats   `mapstructure:"nats"`
	Redis   redis  `mapstructure:"redis"`
	Mysql   mysql  `mapstructure:"mysql"`
	Gcs     gcs    `mapstructure:"gcs"`
	Probe   probe  `mapstructure:"probe"`
	CfgFile string
}

func showHelp() {
	fmt.Printf("Usage:%s {params}\n", os.Args[0])
	fmt.Println("      -c {config file}")
	fmt.Println("      -h (show help info)")
}

type gcs struct {
	ProjectID    string `mapstructure:"project_id"`
	AvatarBucket string `mapstructure:"avatar_bucket"`
}

func (c *config) load() bool {
	_, err := os.Stat(c.CfgFile)
	if err != nil {
		return false
	}

	viper.SetConfigFile(c.CfgFile)
	viper.SetConfigType("toml")

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("config file %s read failed. %v\n", c.CfgFile, err)
		return false
	}
	err = viper.GetViper().UnmarshalExact(c)
	if err != nil {
		fmt.Printf("config file %s loaded failed. %v\n", c.CfgFile, err)
		return false
	}
	fmt.Printf("config %s load ok!\n", c.CfgFile)
	return true
}

func (c *config) parse() bool {
	flag.StringVar(&c.CfgFile, "c", "configs/configs.toml", "config file")
	help := flag.Bool("h", false, "help info")
	flag.Parse()
	if !c.load() {
		return false
	}

	if *help {
		showHelp()
		return false
	}
	return true
}
