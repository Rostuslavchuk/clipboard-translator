package main

import (
	"encoding/json"
	"strings"
	"github.com/valyala/fasthttp"
)

type Translation struct {
	DetectedSourceLanguage string `json:"detected_source_language"`
	Text                   string `json:"text"`
}
type TranslatedResponse struct {
	Translations []Translation `json:"translations"`
}

func CreateReqBody(map_args map[string]string) *fasthttp.Args {
	args := fasthttp.AcquireArgs()
	for k, v := range map_args {
		args.Set(k, v)
	}
	return args 
}
func TranslateText(text, targetLang, api_key, url string) (string, error) {

	args := map[string]string{
		"auth_key":    api_key,
		"text":        text,
		"target_lang": strings.ToUpper(targetLang),
	}

	argsFast := CreateReqBody(args)
	defer fasthttp.ReleaseArgs(argsFast)

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/x-www-form-urlencoded")
	req.SetBody(argsFast.QueryString())

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	client := &fasthttp.Client{}
	if err := client.Do(req, resp); err != nil {
		return "", err
	}

	var dataResp TranslatedResponse
	if err := json.Unmarshal(resp.Body(), &dataResp); err != nil {
		return "", err
	}

	if len(dataResp.Translations) == 0 {
		return "", nil
	}

	return dataResp.Translations[0].Text, nil
}
