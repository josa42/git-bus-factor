package github

import (
	"github.com/AlecAivazis/survey"
	keychain "github.com/lunixbochs/go-keychain"
)

const keychainService = "github.com/josa42/git-bus-factor"

// Login :
func Login() {

	if getToken() != "" {
		replace := false
		prompt := &survey.Confirm{
			Message: "Replace current token?",
		}
		survey.AskOne(prompt, &replace, nil)
		if !replace {
			return
		}
	}

	token := ""
	prompt := &survey.Password{
		Message: "Token:",
		Help:    "Create a GitHub access token at https://github.com/settings/tokens",
	}
	survey.AskOne(prompt, &token, nil)

	if token == "" {
		// fmt.Println("Error: ")
		return
	}

	setToken(token)
}

// # Token
//	tokenPassword = Security::GenericPassword.find(service: KEYCHAIN_SERVICE)
//	token = nil
//	unless tokenPassword
//		puts "Please create a GitHub access token at https://github.com/settings/tokens"
//		while token.nil? || token.empty? do
//			token = ask("Token: ")
//			Security::GenericPassword.add(KEYCHAIN_SERVICE, '', token) unless token.nil? || token.empty?
//		end
//	else
//		token = tokenPassword.password
//	end
func getToken() string {
	token, error := keychain.Find(keychainService, "token")
	if error == nil && token != "" {
		return token
	}

	return ""
}

func setToken(token string) bool {
	error := keychain.Add(keychainService, "token", token)
	return error == nil
}

func removeToken() bool {
	error := keychain.Remove(keychainService, "token")
	return error == nil
}
