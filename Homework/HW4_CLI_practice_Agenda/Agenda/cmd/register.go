/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
)

// UserInf is Register User Information
type UserInf struct {
	Username string
	Password string
	Email    string
	Phone    string
}

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "User Register",
	Long:  `To register, please set a username and a password. In addition, you should also provide your e-mail and phone number.`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		phone, _ := cmd.Flags().GetString("phone")

		userInf := UserInf{
			Username: username,
			Password: password,
			Email:    email,
			Phone:    phone,
		}

		userInfJSONData, err := json.MarshalIndent(userInf, "", "    ")
		if err != nil {
			log.Fatal(err)
		}

		userInfFile, err := os.OpenFile(filepath.FromSlash("UserInformation/"+username+".json"), os.O_CREATE|os.O_EXCL|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			if os.IsExist(err) {
				fmt.Println("\nRegister Failed! User already exists!")
			}
			log.Fatal(err)
		}
		defer userInfFile.Close()
		userInfFile.Write(userInfJSONData)

		logFile, err := os.OpenFile(filepath.FromSlash("Agenda.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal(err)

		}
		defer logFile.Close()

		registerLog := log.New(logFile, "[User Register] ", log.LstdFlags)
		registerLog.Println(username + " Registered Successfully!")

		fmt.Println(username + " Registered Successfully!")

	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("user", "u", "Anonymous", "User Name")
	registerCmd.Flags().StringP("password", "p", "", "User Password")
	registerCmd.Flags().StringP("email", "e", "", "User E-mail")
	registerCmd.Flags().StringP("phone", "n", "", "User Phone Number")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
