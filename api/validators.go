package api

import (
	"encoding/json"
	"strconv"
)

type ValidatorsResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  []struct {
		PubKey      string `json:"pub_key"`
		VotingPower string `json:"voting_power"`
	} `json:"result"`
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error"`
}

func (a *Api) Validators(height int) (*ValidatorsResponse, error) {

	params := make(map[string]string)
	if height > 0 {
		params["height"] = strconv.Itoa(height)
	}

	res, err := a.client.R().SetQueryParams(params).Get("/validators")
	if err != nil {
		return nil, err
	}

	result := new(ValidatorsResponse)
	err = json.Unmarshal(res.Body(), result)
	if err != nil {
		return nil, err
	}

	return result, nil
}