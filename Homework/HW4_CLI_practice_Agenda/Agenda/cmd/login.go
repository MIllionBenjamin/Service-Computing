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

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "User Login",
	Long:  `Use username and password to login.`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("loginUser")
		password, _ := cmd.Flags().GetString("loginPassword")

		userInfFile, err := os.OpenFile(filepath.FromSlash("UserInformation/"+username+".json"), os.O_RDONLY, 0666)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println("\nLogin Failed! User doesn't exist!")
			}
			log.Fatal(err)
		}
		defer userInfFile.Close()
		data := make([]byte, 4096)
		dataLength, err := userInfFile.Read(data)
		if err != nil {
			log.Fatal(err)
		}

		var userInf UserInf
		JSONerr := json.Unmarshal(data[:dataLength], &userInf)
		if JSONerr != nil {
			log.Fatal(JSONerr)
		}

		if userInf.Password != password {
			log.Fatal("Login Failed! Password is incorrect!")
		}

		currentUserFile, err := os.OpenFile("curUser.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			log.Fatal(err)
		}
		currentUserFile.Write([]byte(userInf.Username))

		logFile, err := os.OpenFile("Agenda.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal(err)

		}
		defer logFile.Close()

		registerLog := log.New(logFile, "[User Login] ", log.LstdFlags)
		registerLog.Println(username + " Login Successfully!")

		fmt.Println(username + " Login Successfully!")
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("loginUser", "U", "Anonymous", "login User Name")
	loginCmd.Flags().StringP("loginPassword", "P", "", "login User Password")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
