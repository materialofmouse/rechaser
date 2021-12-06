package connect

type RequestData struct {
	Data struct {
		Session string `json:"session"`
		Command struct {
			Action string `json:"action"`
			Direction int `json:"direction"`
		} `json:"command"`
	} `json:"data"`
}



