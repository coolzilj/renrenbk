package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewRenrenbkCmd() *cobra.Command {
	hdl := cmdHandler{}

	loginCmd := &cobra.Command{
		Use:   "login",
		Short: "登录",
		Run:   hdl.login,
	}

	rootCmd := &cobra.Command{
		Use:   "renrenbk",
		Short: "人人网备份工具",
	}

	rootCmd.AddCommand(loginCmd)
	return rootCmd
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.AddConfigPath(home)
	viper.SetConfigName(".renrenbk")
}
