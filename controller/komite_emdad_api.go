package controller

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func CheckUserLoginKomiteEmdad(phoneNumber string, requestVerificationToken string) (bool, string, error) {
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

func CheckVerificationCode(enteredCode string, verificationToken string) (bool, string, error) {
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
	if contentType == "application/json; charset=utf-8" {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return false, "", fmt.Errorf("error reading response body, error: %v", err)
		}
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
		return true, "", nil
	}
}

func Identity() (bool, string, string, error) {
	req, err := http.NewRequest(http.MethodGet, "https://ekram.emdad.ir/Identity", nil)
	if err != nil {
		return false, "", "", fmt.Errorf("error in set request,error: %v", err)
	}
	cli := &http.Client{}
	res, err := cli.Do(req)
	if err != nil {
		return false, "", "", fmt.Errorf("error in send request,error: %v", err)
	}
	defer res.Body.Close()

	jar, err := cookiejar.New(nil)
	if err != nil {
		return false, "", "", fmt.Errorf("cookie jar,error: %v", err)
	}
	u, err := url.Parse("https://ekram.emdad.ir/Identity")
	if err != nil {
		return false, "", "", fmt.Errorf("url parse error,error: %v", err)
	}
	jar.SetCookies(u, res.Cookies())
	ck := jar.Cookies(u)
	var cookieStr []string
	for _, c := range ck {
		cookieStr = append(cookieStr, c.String())
	}
	siteCookie := strings.Join(cookieStr, "; ")

	if res.Body != nil {
		tokenizer := html.NewTokenizer(res.Body)
		for {
			tokenType := tokenizer.Next()
			switch tokenType {
			case html.ErrorToken:
				// End of document
				return false, "", "", fmt.Errorf("error: %v", err)
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
						return true, value, siteCookie, nil
					}
				}
			}
		}
	}
	return false, "", siteCookie, nil
}

func Registry() (bool, string, string, error) {
	req, err := http.NewRequest(http.MethodGet, "https://ekram.emdad.ir/Home/Register", nil)
	if err != nil {
		return false, "", "", fmt.Errorf("error in set request,error: %v", err)
	}
	cli := &http.Client{}
	res, err := cli.Do(req)
	if err != nil {
		return false, "", "", fmt.Errorf("error in send request,error: %v", err)
	}
	defer res.Body.Close()

	jar, err := cookiejar.New(nil)
	if err != nil {
		return false, "", "", fmt.Errorf("cookie jar,error: %v", err)
	}
	u, err := url.Parse("https://ekram.emdad.ir/Home/Register")
	if err != nil {
		return false, "", "", fmt.Errorf("url parse error,error: %v", err)
	}
	jar.SetCookies(u, res.Cookies())
	ck := jar.Cookies(u)
	var cookieStr []string
	for _, c := range ck {
		cookieStr = append(cookieStr, c.String())
	}
	siteCookie := strings.Join(cookieStr, "; ")

	if res.Body != nil {
		tokenizer := html.NewTokenizer(res.Body)
		for {
			tokenType := tokenizer.Next()
			switch tokenType {
			case html.ErrorToken:
				// End of document
				return false, "", "", fmt.Errorf("error: %v", err)
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
						return true, value, siteCookie, nil
					}
				}
			}
		}
	}
	return false, "", siteCookie, nil
}

func SendUserInformation(firstName string, lastName string, codeMeli string, birthDate string, phoneNumber string, enteredCode string, verificationToken string, jobId string) (string, string, error) {
	var birthDay, birthMonth, birthYear int
	birthInfo := strings.Split(birthDate, "-")
	for k, v := range birthInfo {
		value, err := strconv.Atoi(v)
		if err != nil {
			return "", "", fmt.Errorf("error in conver birth date from string to integer, error: %v", err)
		}

		if k == 0 {
			birthYear = value
		} else if k == 1 {
			birthMonth = value
		} else if k == 2 {
			birthDay = value
		}
	}

	password := "qwertyTest123@"
	body := map[string]interface{}{
		"PersonType":                      "ش",
		"FirstName":                       firstName,
		"LastName":                        lastName,
		"CodeMeli":                        codeMeli,
		"DateOfBirthDay":                  birthDay,
		"DateOfBirthMonth":                birthMonth,
		"DateOfBirthYear":                 birthYear,
		"JobMainCat":                      jobId,
		"Password":                        password,
		"ConfirmPassword":                 password,
		"verifyNumberVM.PhoneNumber":      phoneNumber,
		"verifyNumberVM.VerificationCode": enteredCode,
		"__RequestVerificationToken":      verificationToken,
		"HamkariHozeId":                   0,
		"JobCat":                          "",
		"verifyNumberVM.JobId":            "",
		"X-Requested-With":                "XMLHttpRequest",
		"OrgName":                         "",
	}

	formData := url.Values{}
	for key, value := range body {
		formData.Add(key, fmt.Sprintf("%v", value))
	}

	res, err := http.PostForm("https://ekram.emdad.ir/Home/Register", formData)
	if err != nil {
		return "", "", fmt.Errorf("error in send the http request, error: %v", err)
	}
	defer res.Body.Close()

	contentType := res.Header.Get("Content-Type")
	if contentType == "application/json; charset=utf-8" {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "", "", fmt.Errorf("error reading response body, error: %v", err)
		}

		var responseBody map[string]interface{}
		err = json.Unmarshal(body, &responseBody)
		if err != nil {
			return "", "", fmt.Errorf("error in unmarshalling the response body, error: %v", err.Error())
		}

		fmt.Printf("User with phone number: %v get this message: %v\n", phoneNumber, responseBody["message"].(string))
		return responseBody["message"].(string), password, nil
	} else {
		fmt.Printf("User with phone number: %v give a html code", phoneNumber)
		return "", password, nil
	}
}

func GetCaptcha(siteCookie string) ([]byte, string, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://ekram.emdad.ir/get-captcha-image?%d", time.Now().UnixMilli()), nil)
	if err != nil {
		return nil, "", err
	}
	req.Header.Set("Cookie", siteCookie)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, "", fmt.Errorf("error in send get the token request, error: %v", err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error in read the response body, error: %v", err)
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, "", fmt.Errorf("cookie jar,error: %v", err)
	}
	u, err := url.Parse("https://ekram.emdad.ir/get-captcha-image")
	if err != nil {
		return nil, "", fmt.Errorf("url parse error,error: %v", err)
	}
	jar.SetCookies(u, req.Cookies())
	jar.SetCookies(u, res.Cookies())
	ck := jar.Cookies(u)
	var cookieStr []string
	for _, c := range ck {
		cookieStr = append(cookieStr, c.String())
	}
	siteCookie = strings.Join(cookieStr, "; ")

	return body, siteCookie, nil
}

func GetIdentity(NationalCode string, password string, CaptchaCode string, verificationCode string, cookie string) ([]byte, string, error) {
	fmt.Printf("\npassword: %v \n", password)
	fmt.Printf("\nCaptchaCode: %v \n", CaptchaCode)
	fmt.Printf("\nverificationCode: %v \n", verificationCode)
	fmt.Printf("\ncookie: %v \n", cookie)
	fmt.Printf("\nNationalCode: %v \n", NationalCode)

	encryptPassword, err := EncryptAES(password)
	if err != nil {
		return nil, "", err
	}

	data := url.Values{
		"Username":                   []string{NationalCode},
		"Password":                   []string{encryptPassword},
		"RememberMe":                 []string{"true", "false"},
		"CaptchaCode":                []string{CaptchaCode},
		"HashedPW":                   []string{"0"},
		"__RequestVerificationToken": []string{verificationCode},
	}

	//reqBody := data.Encode()
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://ekram.emdad.ir/Identity"), nil)
	req.PostForm = data
	//req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost:5000/Identity"), bytes.NewBufferString(reqBody))
	if err != nil {
		return nil, "", err
	}
	req.Header.Set("Cookie", cookie)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", "curl/7.81.0")

	//client := http.Client{
	//	Transport: &http.Transport{
	//		DisableCompression: true,
	//	},
	//}

	//res, err := client.Do(req)
	//if err != nil {
	//	return nil, "", fmt.Errorf("Error in send get the token request, error: %v", err)
	//}
	//resCookie := res.Header.Get("Set-Cookie")
	//body, err := io.ReadAll(res.Body)
	//if err != nil {
	//	return nil, "", fmt.Errorf("Error in read the response body, error: %v", err)
	//}

	args := []string{"-v", req.URL.String()}
	for k, varr := range req.Header {
		for _, v := range varr {
			args = append(args, "-H")
			args = append(args, fmt.Sprintf("%s: %s", k, v))
		}
	}
	args = append(args, "--data-raw", data.Encode())

	cmd := exec.Command("curl", args...)
	buf := bytes.Buffer{}
	errBuf := bytes.Buffer{}
	cmd.Stdout = &buf
	cmd.Stderr = &errBuf
	err = cmd.Run()
	if err != nil {
		return nil, "", err
	}

	outErr := errBuf.String()

	startReading := false
	statusCode := ""
	resHeaders := http.Header{}
	for _, line := range strings.Split(outErr, "\n") {
		if strings.HasPrefix(line, "< HTTP/1.1") {
			startReading = true
			temp := strings.TrimSpace(strings.TrimPrefix(line, "< HTTP/1.1"))
			idx := strings.Index(temp, " ")
			statusCode = temp[:idx]
			if statusCode != "302" {
				return nil, "", fmt.Errorf("invalid status code: %s", statusCode)
			}
			continue
		}

		if startReading {
			if strings.HasPrefix(line, "<") {
				l := strings.TrimSpace(strings.TrimPrefix(line, "<"))
				if len(l) > 0 {
					arr := strings.Split(l, ":")
					key := strings.TrimSpace(arr[0])
					value := strings.TrimSpace(arr[1])

					resHeaders.Add(key, value)
				}
			}
		}
	}

	res := http.Response{
		Header: resHeaders,
	}
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, "", fmt.Errorf("cookie jar,error: %v", err)
	}
	u, err := url.Parse("https://ekram.emdad.ir/get-captcha-image")
	if err != nil {
		return nil, "", fmt.Errorf("url parse error,error: %v", err)
	}
	jar.SetCookies(u, req.Cookies())
	jar.SetCookies(u, res.Cookies())
	ck := jar.Cookies(u)
	var cookieStr []string
	for _, c := range ck {
		cookieStr = append(cookieStr, c.String())
	}
	siteCookie := strings.Join(cookieStr, "; ")
	return nil, siteCookie, nil
	//return body, resCookie, nil
}

func EncryptAES(password string) (string, error) {
	// Example configuration
	key := []byte("8080808080808080")
	iv := []byte("8080808080808080")
	fmt.Printf("password: %v", password)
	// Create a new AES cipher block using the specified key size
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("Error creating AES cipher block,Error: %v", err)
	}

	// Create a new CBC mode block cipher using the IV
	mode := cipher.NewCBCEncrypter(block, iv)

	// Pad the plaintext to a multiple of the block size (PKCS#7 padding)
	paddedPlaintext := padPlaintext([]byte(password), aes.BlockSize)

	// Encrypt the padded plaintext
	ciphertext := make([]byte, len(paddedPlaintext))
	mode.CryptBlocks(ciphertext, paddedPlaintext)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// padPlaintext pads the plaintext to a multiple of blockSize using PKCS#7 padding
func padPlaintext(plaintext []byte, blockSize int) []byte {
	padding := blockSize - len(plaintext)%blockSize
	pad := make([]byte, padding)
	for i := range pad {
		pad[i] = byte(padding)
	}
	return append(plaintext, pad...)
}

func SendResetPasswordRequest(siteCookie string, verificationToken string, nationalCode string, phoneNumber string, captchaCode string, jobId string) error {
	body := url.Values{
		"CodeMelli":                  []string{nationalCode},
		"PhoneNumber":                []string{phoneNumber},
		"PhoneOrMail":                []string{"1"},
		"CaptchaCode":                []string{captchaCode},
		"JobId":                      []string{jobId},
		"__RequestVerificationToken": []string{verificationToken},
	}
	req, err := http.NewRequest(http.MethodPost, "https://ekram.emdad.ir/Identity/ForgotPassword", bytes.NewBufferString(body.Encode()))
	if err != nil {
		return fmt.Errorf("error in set request, error: %v", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", siteCookie)
	client := http.Client{
		Transport: &http.Transport{DisableCompression: true},
	}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error in send request, error: %v", err)
	}
	if res.StatusCode == http.StatusFound || res.StatusCode == http.StatusOK {
		return nil
	} else {
		return fmt.Errorf("error somthing went wrong")
	}
}

func GetJobId() (string, error) {
	req, err := http.NewRequest(http.MethodGet, "https://ekram.emdad.ir/Identity/Login", nil)
	if err != nil {
		return "", fmt.Errorf("error in set request, error: %v", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error in send request, error: %v", err)
	}
	parsedDoc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var jobIdValue string
	parsedDoc.Find("a[href^='/Identity/ForgotPassword']").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			// Parse the query parameters from the href
			params, err := url.ParseQuery(strings.Split(href, "?")[1])
			if err == nil {
				jobIdValue = params.Get("JobId")
			}
		}
	})
	return jobIdValue, nil
}

// ************************** Get Token *********************************

func GetToken() (string, error) {
	mapBody := map[string]string{
		"grant_type": "password",
		"username":   "bale_user",
		"password":   "2024315_UserbalE",
	}
	bodyMapped, err := json.Marshal(mapBody)
	if err != nil {
		return "", fmt.Errorf("error in marshaling the body of request, error: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, "https://apim.emdad.ir:9443/oauth2/token", bytes.NewBuffer(bodyMapped))
	if err != nil {
		return "", fmt.Errorf("error in set get the token request, error: %v", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error in send get the token request, error: %v", err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error in read the response body, error: %v", err)
	}
	return string(body), nil
}

// ==========================  chose children ============================

type cityStr struct {
	CityId       int
	CityName     string
	ProviderId   int
	ProviderName string
}

func GetCityOfState(stateId string, city string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, "https://ekram.emdad.ir/CIty/GetByProvinceId", nil)
	if err != nil {
		return "", fmt.Errorf("error in set the request get the city of the stat, error: %v", err)
	}
	q := req.URL.Query()
	q.Add("id", stateId)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error in send the request get the city of the stat, error: %v", err)
	}

	if res.Body != nil {
		var response []cityStr
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "", fmt.Errorf("error in parse the body, error: %v", err)
		}
		err = json.Unmarshal(body, &response)
		if err != nil {
			return "", fmt.Errorf("error in unmarshal the body, error: %v", err)
		}
		cityId, err := FindCityID(city, response)
		if err != nil {
			return "", nil
		}
		cityIdString := strconv.Itoa(int(cityId))
		return cityIdString, nil
	}
	return "", nil
}

func FindCityID(cityName string, cities []cityStr) (int64, error) {
	for _, v := range cities {
		if v.CityName == cityName {
			return int64(v.CityId), nil
		}
	}
	return 0, fmt.Errorf("your city entered is wrong")
}

func GetChildren(cookie string, verificationToken string) (map[string]string, error) {
	body := url.Values{
		"mohseninOrphanSearchVM.ProvinceId":                              []string{"17"},
		"mohseninOrphanSearchVM.SenAz":                                   []string{""},
		"mohseninOrphanSearchVM.SenTa":                                   []string{""},
		"mohseninOrphanSearchVM.BirthDate":                               []string{""},
		"mohseninOrphanSearchVM.OrphanSearchForRelationSearchFirstName":  []string{""},
		"mohseninOrphanSearchVM.MemberShipCode":                          []string{""},
		"mohseninOrphanSearchVM.CityId":                                  []string{"882"},
		"mohseninOrphanSearchVM.OrphanSearchForRelationSearchBranchCode": []string{"-1"},
		"mohseninOrphanSearchVM.OrphanSearchForRelationGender":           []string{"-1"},
		"mohseninOrphanSearchVM.Action":                                  []string{"AddToList"},
		"mohseninOrphanSearchVM.OrphanSearchForRelationHealthStatus":     []string{"-1"},
		"mohseninOrphanSearchVM.IsSeyed":                                 []string{"-1"},
		"mohseninOrphanSearchVM.PlanType":                                []string{"-1"},
		"mohseninOrphanSearchVM.NoeSokonat":                              []string{"-1"},
		"__RequestVerificationToken":                                     []string{verificationToken},
		"X-Requested-With":                                               []string{"XMLHttpRequests"},
	}
	//req, err := http.NewRequest(http.MethodPost, "http://localhost:5001/Orphan", bytes.NewBufferString(body.Encode()))
	req, err := http.NewRequest(http.MethodPost, "https://ekram.emdad.ir/Orphan", bytes.NewBufferString(body.Encode()))
	if err != nil {
		return nil, fmt.Errorf("error in set the http request, error: %v", err)
	}
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Cookie", cookie)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error in send the http request, error: %v", err)
	}

	//defer res.Body.Close()
	//resbody, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	return "", err
	//}

	if res.Body != nil {
		parsedDoc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		//fullName, err := ParseChildrenInfo(parsedDoc)
		//if err != nil {
		//	fmt.Printf("ERrror:%v", err.Error())
		//	return "", err
		//}
		childInfo := make(map[string]string)
		parsedDoc.Find("table#tblOrphansList").Find("tbody").Find("tr").Each(func(i int, row *goquery.Selection) {
			checkbox := row.Find("input[type=checkbox]")
			id := strings.TrimSpace(checkbox.AttrOr("value", ""))

			fullName := row.Find("td.fullName").Text()
			fullName = strings.Replace(fullName, "\n", " ", -1)
			fullName = strings.TrimSpace(fullName)
			childInfo[fullName] = id
		})

		//fullNameString := strings.Join(fullNames, "\n")
		//fullNames := strings.Split(fullName, "\n")
		if childInfo != nil {
			return childInfo, nil
		}
	}
	return nil, nil
}

func ParseChildrenInfo(doc *goquery.Document) (string, error) {
	//doc, err := html.Parse(htmlCode)
	//if err != nil {
	//	return "", fmt.Errorf("Error: %v", err)
	//}

	// Find the <td> element with class "fullName"
	//fullName := doc.Find("table#tblOrphansList").Find("tbody").Find("tr").Children().Each(func(i int, s *goquery.Selection) {
	//	s.Find("input[type=checkbox]").Each(func(i int, selection *goquery.Selection) {
	//		id := strings.TrimSpace(selection.AttrOr("value", ""))
	//	})
	//	fullName := s.Find("td.fullName").Text()
	//
	//})

	return "", nil
	//var fullNames []string
	//
	//var findChild func(*html.Node)
	//findChild = func(node *html.Node) {
	//	if node.Type == html.ElementNode && node.Data == "tr" {
	//		for _, attr := range node.Attr {
	//			if attr.Key == "class" && attr.Val == "fullName" {
	//				fullNames = append(fullNames, node.FirstChild.Data)
	//			}
	//		}
	//	}
	//	for c := node.FirstChild; c != nil; c = c.NextSibling {
	//		findChild(c)
	//	}
	//}
	//
	//findChild(doc)
	//if fullNames != nil {
	//	return fmt.Sprintf("Full Name: %s\n", fullNames), nil
	//} else {
	//	return fmt.Sprintln("Full Name not found."), nil
	//}
}

func GetListOfMyChildren(cookie string, verificationToken string) error {
	req, err := http.NewRequest(http.MethodGet, "https://ekram.emdad.ir/Orphan/MyList", nil)
	if err != nil {
		return fmt.Errorf("error in set the http request, error: %v", err)
	}
	//req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Cookie", cookie)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error in send the http request, error: %v", err)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(body)

	doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		return fmt.Errorf("Error: %v\n", err)
	}

	var fullNames []string
	var listChildren func(*html.Node)
	listChildren = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "td" {
			for _, attr := range node.Attr {
				if attr.Key == "class" && attr.Val == "fullName" {
					if node.FirstChild != nil {
						fullNames = append(fullNames, node.FirstChild.Data)
					}
				}
			}
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			listChildren(c)
		}
	}

	listChildren(doc)
	if fullNames != nil {
		fmt.Printf("fullName equal to: %v \n", fullNames)
		return nil
	} else {
		return fmt.Errorf("error in find the full name class")
	}
}
