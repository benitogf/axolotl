package axolotl

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"github.com/Mparaiso/lodash-go"
)

func getCookie(head string, brand string) (string, error) {
    cookie, err := url.QueryUnescape(head)
	fail := "Could't find the cookie"

	if (err != nil) {
		return "", fmt.Errorf(fail)
	}

	cookieParams := strings.Split(cookie, ";")
    err = lo.Filter(cookieParams, func(element string) bool {
	    return strings.Index(element, brand) != -1
    }, &cookieParams)

    if err != nil {
		return "", fmt.Errorf(fail)
    }

    if (len(cookieParams) == 0) {
		return "", fmt.Errorf(fail)
    }

	return strings.Replace(cookieParams[0], brand, "", 1), nil
}

func EatCookie(r *http.Request, brand string) (string, error) {
	cookie, err := getCookie(r.Header["Cookie"][0], brand)

	if (err != nil) {
		return "", err
	}

	return cookie, nil
}
