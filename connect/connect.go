package connect

type RequestData struct {
	Data struct {
		Session string `json:"session"`
		Command struct {
			Action string `json:"action"`
			Direction string `json:"direction"`
		} `json:"command"`
	} `json:"data"`
}



