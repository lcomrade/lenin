// Copyright 2021 Leonid Maslakov

// License: GPL-3.0-or-later

// This file is part of Lenpaste.

// Lenpaste is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// Lenpaste is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Lenpaste.  If not, see <https://www.gnu.org/licenses/>.

package lenin

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Remote server errors.
func httpErrorCheck(code int, raw []byte) error {
	//Check HTTP status code
	if code != 200 {
		//Pares error
		var remoteErr remoteError

		//Unmarshal raw
		err := json.Unmarshal(raw, &remoteErr)
		if err != nil {
			return errors.New("remote error: " + err.Error())
		}

		//Return error
		return errors.New("remote error: " + remoteErr.Error)
	}

	return nil
}

// Make GET request.
func getURL(url string) ([]byte, error) {
	var out []byte

	//Get
	resp, err := http.Get(url)
	if err != nil {
		return out, err
	}

	//Read
	out, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return out, err
	}

	//Check to server internal error
	err = httpErrorCheck(resp.StatusCode, out)
	if err != nil {
		return out, err
	}

	//Return
	return out, nil
}

// Make POST request.
func postURL(vals url.Values, url string) ([]byte, error) {
	var out []byte

	//Post
	resp, err := http.PostForm(url, vals)
	if err != nil {
		return out, err
	}

	//Read
	out, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return out, err
	}

	//Check to server internal error
	err = httpErrorCheck(resp.StatusCode, out)
	if err != nil {
		return out, err
	}

	//Return
	return out, nil
}

// Clean base URL. Performs the following actions:
//
//  1. If the URL does not start with "http://" or "https://" adds "http://".
//  2. Deletes the "/" at the end.
func CleanURL(url string) string {
	//Step 1
	if url[0:7] != "http://" && url[0:8] != "https://" {
		url = "http://" + url
	}

	//Step 2
	lastChar := string(url[len(url)-1:])
	if lastChar == "/" {
		url = url[0 : len(url)-1]
	}

	return url
}

// Create new paste.
func New(req NewReq, baseURL string) (NewResp, error) {
	var resp NewResp

	//Create form
	vals := url.Values{}

	vals.Set("title", req.Title)
	vals.Add("text", req.Text)
	vals.Add("expiration", req.Expiration)

	if req.OneUse == true {
		vals.Add("oneUse", "true")

	} else {
		vals.Add("oneUse", "false")
	}

	//Get raw
	raw, err := postURL(vals, baseURL+"/new")
	if err != nil {
		return resp, err
	}

	//Unmarshal raw
	err = json.Unmarshal(raw, &resp)
	if err != nil {
		return resp, err
	}

	//Return
	return resp, nil
}

// Get existed paste by name.
func Get(name string, baseURL string) (GetResp, error) {
	var resp GetResp

	//Get raw
	raw, err := getURL(baseURL + "/get/" + name)
	if err != nil {
		return resp, err
	}

	//Unmarshal raw
	err = json.Unmarshal(raw, &resp)
	if err != nil {
		return resp, err
	}

	//Return
	return resp, nil
}

// Get server about.
func About(baseURL string) (AboutResp, error) {
	var resp AboutResp

	//Get raw
	raw, err := getURL(baseURL + "/about")
	if err != nil {
		return resp, err
	}

	//Unmarshal raw
	err = json.Unmarshal(raw, &resp)
	if err != nil {
		return resp, err
	}

	//Return
	return resp, nil
}

// Get server rules.
func Rules(baseURL string) (RulesResp, error) {
	var resp RulesResp

	//Get raw
	raw, err := getURL(baseURL + "/rules")
	if err != nil {
		return resp, err
	}

	//Unmarshal raw
	err = json.Unmarshal(raw, &resp)
	if err != nil {
		return resp, err
	}

	//Return
	return resp, nil
}

// Get server version information.
func Version(baseURL string) (VersionResp, error) {
	var resp VersionResp

	//Get raw
	raw, err := getURL(baseURL + "/version")
	if err != nil {
		return resp, err
	}

	//Unmarshal raw
	err = json.Unmarshal(raw, &resp)
	if err != nil {
		return resp, err
	}

	//Return
	return resp, nil
}
