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
	"io/ioutil"
	"net/http"
	"net/url"
)

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

	//Return
	return out, nil
}

// Create new paste.
func New(req NewReq, baseURL string) (NewResp, error) {
	var resp NewResp

	//Create form
	vals := url.Values{}
	vals.Set("title", req.Title)
	vals.Add("text", req.Text)
	vals.Add("expiration", req.Expiration)

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
func Get(req GetReq, baseURL string) (GetResp, error) {
	var resp GetResp

	//Get raw
	raw, err := getURL(baseURL + "/get/" + req.Name)
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
