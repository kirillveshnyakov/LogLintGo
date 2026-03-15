package rules

import (
	"strings"
)

func findKeyWord(phrase string) string {
	phrase = strings.ToLower(phrase)
	switch {
	case strings.Contains(phrase, "password"):
		return "password"
	case strings.Contains(phrase, "passwd"):
		return "passwd"
	case strings.Contains(phrase, "pwd"):
		return "pwd"
	case strings.Contains(phrase, "passphrase"):
		return "passphrase"

	case strings.Contains(phrase, "token"):
		return "token"
	case strings.Contains(phrase, "api_key"):
		return "api_key"
	case strings.Contains(phrase, "apikey"):
		return "apikey"
	case strings.Contains(phrase, "api"):
		return "api"
	case strings.Contains(phrase, "key"):
		return "key"
	case strings.Contains(phrase, "access_key"):
		return "access_key"
	case strings.Contains(phrase, "secret_key"):
		return "secret_key"
	case strings.Contains(phrase, "private_key"):
		return "private_key"
	case strings.Contains(phrase, "signing_key"):
		return "signing_key"
	case strings.Contains(phrase, "secret"):
		return "secret"

	case strings.Contains(phrase, "auth_token"):
		return "auth_token"
	case strings.Contains(phrase, "access_token"):
		return "access_token"
	case strings.Contains(phrase, "refresh_token"):
		return "refresh_token"
	case strings.Contains(phrase, "credentials"):
		return "credentials"
	case strings.Contains(phrase, "auth"):
		return "auth"

	case strings.Contains(phrase, "session_id"):
		return "session_id"
	case strings.Contains(phrase, "session"):
		return "session"
	case strings.Contains(phrase, "cookie"):
		return "cookie"

	case strings.Contains(phrase, "credit_card"):
		return "credit_card"
	case strings.Contains(phrase, "card_number"):
		return "card_number"
	case strings.Contains(phrase, "cvv"):
		return "cvv"
	case strings.Contains(phrase, "ssn"):
		return "ssn"
	case strings.Contains(phrase, "social_security"):
		return "social_security"
	case strings.Contains(phrase, "passport"):
		return "passport"
	case strings.Contains(phrase, "taxpayer"):
		return "taxpayer"

	case strings.Contains(phrase, "private"):
		return "private"
	case strings.Contains(phrase, "jwt"):
		return "jwt"
	case strings.Contains(phrase, "bearer"):
		return "bearer"
	case strings.Contains(phrase, "cipher"):
		return "cipher"
	case strings.Contains(phrase, "encryption"):
		return "encryption"
	case strings.Contains(phrase, "database_url"):
		return "database_url"
	case strings.Contains(phrase, "db_password"):
		return "db_password"
	case strings.Contains(phrase, "connection_string"):
		return "connection_string"

	default:
		return ""
	}
}

func CheckSensitiveWords(msg string) (string, bool) {
	if keyWord := findKeyWord(msg); keyWord != "" {
		return keyWord, true
	}
	return "", false
}
