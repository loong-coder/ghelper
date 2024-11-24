package cmd

import (
	"fmt"
	"io"
	"os"

	"com.github/loong/ghelper/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var overiteCmd = &cobra.Command{
	Use: "overite",
	Run: func(cmd *cobra.Command, args []string) {
		var config config.Config
		if err := viper.Unmarshal(&config); err != nil {
			fmt.Println(err)
		}
		overiteFile(config)
		fmt.Println("overite success")
	},
}

// copyFile 复制文件，若目标文件存在则覆盖
func copyFile(source, target string) error {
	// 打开源文件
	srcFile, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("打开源文件失败: %v", err)
	}
	defer srcFile.Close()

	// 创建或覆盖目标文件
	tgtFile, err := os.Create(target)
	if err != nil {
		return fmt.Errorf("创建目标文件失败: %v", err)
	}
	defer tgtFile.Close()

	// 复制文件内容
	_, err = io.Copy(tgtFile, srcFile)
	if err != nil {
		return fmt.Errorf("复制文件内容失败: %v", err)
	}

	return nil
}

func overiteFile(config config.Config) {
	overite := config.Overite
	for _, v := range overite {
		err := copyFile(v.Source, v.Target)
		if err != nil {
			fmt.Printf("复制文件发生错误 sourceFile: %s, targetFile: %s, error: %v\n", v.Source, v.Target, err)
		}
	}
}

func init() {
	rootCmd.AddCommand(overiteCmd)
}
