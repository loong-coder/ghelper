package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "ghelper is a developer tools",
	Long:  "ghelper is a developer tools, help soft developer professional work",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	rootCmd.ExecuteC()
}

var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)
	// 持久化标识 当前命令及其所有下级命令都可以使用
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/ghelper.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		// 读取环境变量配置
		var gutil_home string = os.Getenv("GUTIL_HOME")
		// viper.SetConfigFile(gutil_home + "/config.yaml")
		if gutil_home != "" {
			home = gutil_home
		}

		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName("ghelper")
	}

	// 检查环境变量，将配置的键值加载到viper中
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("use config file:", viper.ConfigFileUsed())

}
