package cmd

import (
	"fmt"
	"log"
	"os"

	"com.github/loong/ghelper/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init ghelper",
	Run: func(cmd *cobra.Command, args []string) {
		initDefaultConfig()
	},
}

func initDefaultConfig() {
	// 模拟初始化配置数据
	config := config.Config{
		Author:  "garen mao",
		License: "Apache License 2.0",
		Overite: []struct {
			Source string `yaml:"source"`
			Target string `yaml:"target"`
		}{
			{Source: "/path/to/source1.txt", Target: "/path/to/target1.txt"},
			{Source: "/path/to/source2.txt", Target: "/path/to/target2.txt"},
		},
	}

	// 初始化 Viper
	v := viper.New()
	v.SetConfigType("yaml") // 配置文件类型

	// 将结构体数据加载到 Viper
	v.Set("author", config.Author)
	v.Set("license", config.License)
	v.Set("overite", config.Overite)

	home, _ := os.UserHomeDir()
	// 读取环境变量配置
	var gutil_home string = os.Getenv("GUTIL_HOME")
	// viper.SetConfigFile(gutil_home + "/config.yaml")
	if gutil_home != "" {
		home = gutil_home
	}
	// 指定生成的文件路径
	outputPath := home + "/ghelper.yaml"
	// 写入 YAML 文件
	err := v.WriteConfigAs(outputPath)
	if err != nil {
		log.Fatalf("写入配置文件失败: %v", err)
	}

	fmt.Printf("配置文件已生成: %s\n", outputPath)
}

func init() {
	rootCmd.AddCommand(initCmd)
}
