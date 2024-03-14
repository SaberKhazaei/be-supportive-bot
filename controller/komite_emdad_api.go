package controller

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"strings"
)

func CheckUserLoginKomiteEmdad(phoneNumber int64, requestVerificationToken string) (bool, string, error) {
	body := map[string]interface{}{
		"PhoneNumber":                phoneNumber,
		"__RequestVerificationToken": requestVerificationToken,
		"X-Requested-With":           "XMLHttpRequest",
	}

	formData := url.Values{}
	for key, value := range body {
		formData.Add(key, fmt.Sprintf("%v", value))
	}
	res, err := http.PostForm("https://ekram.emdad.ir/Home/SendSms", formData)
	if err != nil {
		return false, "", fmt.Errorf("error in send the http request, error: %v", err)
	}
	defer res.Body.Close()

	if res.Body != nil {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return false, "", fmt.Errorf("error reading response body, error: %v", err)
		}

		var responseBody map[string]interface{}
		err = json.Unmarshal(body, &responseBody)
		if err != nil {
			return false, "", fmt.Errorf("error in unmarshalling the response body, error: %v", err.Error())
		}
		if responseBody["message"].(string) == "پیامک ارسال شد" {
			return true, responseBody["message"].(string), nil
		} else {
			return false, responseBody["message"].(string), nil
		}
	}
	return false, "", nil
}

func CheckVerificationCode(enteredCode int64, verificationToken string) (bool, string, error) {
	reqBody := map[string]interface{}{
		"VerificationCode":           enteredCode,
		"__RequestVerificationToken": verificationToken,
		"X-Requested-With":           "XMLHttpRequest",
	}

	formData := url.Values{}
	for key, value := range reqBody {
		formData.Add(key, fmt.Sprintf("%v", value))
	}
	res, err := http.PostForm("https://ekram.emdad.ir/Home/NewUser", formData)
	if err != nil {
		return false, "", fmt.Errorf("error in send the http request, error: %v", err)
	}
	defer res.Body.Close()
	contentType := res.Header.Get("Content-Type")
	fmt.Printf("content type : %v\n", contentType)
	if contentType == "application/json; charset=utf-8" {
		fmt.Println("TEST Application json")
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return false, "", fmt.Errorf("error reading response body, error: %v", err)
		}
		fmt.Printf("body: %v \n ", string(body))
		var responseBody map[string]interface{}
		err = json.Unmarshal(body, &responseBody)
		if err != nil {
			return false, "", fmt.Errorf("error in unmarshalling the response body, error: %v", err.Error())
		}

		fmt.Printf("success value: %v \n", responseBody["success"].(bool))
		if responseBody["success"].(bool) {
			return true, responseBody["message"].(string), nil
		} else {
			return false, responseBody["message"].(string), nil
		}
	} else {
		fmt.Println("TEST not Application json")
		return true, "", nil
	}
}

func Registry() (bool, string, error) {
	req, err := http.NewRequest(http.MethodGet, "https://ekram.emdad.ir/Home/Register", nil)
	if err != nil {
		return false, "", fmt.Errorf("error in set request,error: %v", err)
	}
	cli := &http.Client{}
	res, err := cli.Do(req)
	if err != nil {
		return false, "", fmt.Errorf("error in send request,error: %v", err)
	}
	defer res.Body.Close()

	if res.Body != nil {
		tokenizer := html.NewTokenizer(res.Body)
		for {
			tokenType := tokenizer.Next()
			switch tokenType {
			case html.ErrorToken:
				// End of document
				return false, "", fmt.Errorf("error: %v", err)
			case html.SelfClosingTagToken, html.StartTagToken:
				token := tokenizer.Token()
				if token.Data == "input" {
					isTargetInput := false
					value := ""
					for _, attr := range token.Attr {
						if attr.Key == "name" && strings.TrimSpace(attr.Val) == "__RequestVerificationToken" {
							isTargetInput = true
						}
						if attr.Key == "value" {
							value = strings.TrimSpace(attr.Val)
						}
					}
					if isTargetInput {
						return true, value, nil
					}
				}
			}
		}
	}
	return false, "", nil
}

func SendUserInformation(firstName string, lastName string, codeMeli int64, birthDate time.Time, phoneNumber int64, enteredCode int64, verificationToken string) (string, string, error) {
	password := "qwertyTest123@"
	body := map[string]interface{}{
		"PersonType":                      "ش",
		"FirstName":                       firstName,
		"LastName":                        lastName,
		"CodeMeli":                        codeMeli,
		"DateOfBirthDay":                  birthDate.Day(),
		"DateOfBirthMonth":                birthDate.Month(),
		"DateOfBirthYear":                 birthDate.Year(),
		"JobMainCat":                      5027,
		"Password":                        password,
		"ConfirmPassword":                 password,
		"verifyNumberVM.PhoneNumber":      phoneNumber,
		"verifyNumberVM.VerificationCode": enteredCode,
		"__RequestVerificationToken":      verificationToken,
		"X-Requested-With":                "XMLHttpRequest",
	}

	formData := url.Values{}
	for key, value := range body {
		formData.Add(key, fmt.Sprintf("%v", value))
	}

	res, err := http.PostForm("https://ekram.emdad.ir/Home/NewUser", formData)
	if err != nil {
		return "", "", fmt.Errorf("error in send the http request, error: %v", err)
	}
	defer res.Body.Close()

	contentType := res.Header.Get("Content-Type")
	if contentType == "application/json; charset=utf-8" {
		fmt.Println("TEST 1")
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "", "", fmt.Errorf("error reading response body, error: %v", err)
		}
		fmt.Println("TEST 2")

		var responseBody map[string]interface{}
		err = json.Unmarshal(body, &responseBody)
		if err != nil {
			return "", "", fmt.Errorf("error in unmarshalling the response body, error: %v", err.Error())
		}
		fmt.Println("TEST 3")
		fmt.Printf("TEST MESSAGE: %v \n", responseBody["message"].(string))

		if responseBody["success"].(bool) {
			return responseBody["message"].(string), password, nil
		} else {
			return responseBody["message"].(string), "", nil
		}
	} else {
		fmt.Println("TEST 4")
		responseBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "", "", fmt.Errorf("error reading response body, error: %v", err)
		}
		return "", password, nil
	}
}
