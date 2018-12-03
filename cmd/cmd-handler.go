package cmd

import (
	"errors"
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type cmdHandler struct {
}

func (h *cmdHandler) login(cmd *cobra.Command, args []string) {
	user, password, err := getUserAndPassword()
	if err != nil {
		log.Fatalln("获取用户名和密码失败！")
	}

	fmt.Printf("%q: %q\n", user, password)
}

func getUserAndPassword() (string, string, error) {
	var err error

	userValidate := func(input string) error {
		if len(input) < 1 {
			return errors.New("用户名不能为空")
		}
		return nil
	}

	userPrompt := promptui.Prompt{
		Label:    "用户名",
		Validate: userValidate,
	}

	user, err := userPrompt.Run()

	if err != nil {
		fmt.Printf("User Prompt failed %v\n", err)
		return "", "", err
	}

	pwValidate := func(input string) error {
		if len(input) < 1 {
			return errors.New("密码不能为空")
		}
		return nil
	}

	passwordPrompt := promptui.Prompt{
		Label:    "密码",
		Mask:     '*',
		Validate: pwValidate,
	}

	password, err := passwordPrompt.Run()

	if err != nil {
		fmt.Printf("Password Prompt failed %v\n", err)
		return "", "", err
	}

	return user, password, err
}
