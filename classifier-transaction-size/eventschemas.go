package main

type HiFromKnative struct {
	// Msg holds the message from the event
	Msg string `json:"msg,omitempty,string"`
}

type Transaction struct {
	Op string `json:"op"`
	X  struct {
		LockTime int `json:"lock_time"`
		Ver      int `json:"ver"`
		Size     int `json:"size"`
		Inputs   []struct {
			Sequence int64 `json:"sequence"`
			PrevOut  struct {
				Spent   bool   `json:"spent"`
				TxIndex int    `json:"tx_index"`
				Type    int    `json:"type"`
				Addr    string `json:"addr"`
				Value   int    `json:"value"`
				N       int    `json:"n"`
				Script  string `json:"script"`
			} `json:"prev_out"`
			Script string `json:"script"`
		} `json:"inputs"`
		Time      int    `json:"time"`
		TtxIndex  int    `json:"ttx_index"`
		VinSz     int    `json:"vin_sz"`
		Hash      string `json:"hash"`
		VoutSz    int    `json:"vout_sz"`
		RelayedBy string `json:"relayed_by"`
		Out       []struct {
			Spent   bool   `json:"spent"`
			TxIndex int    `json:"tx_index"`
			Type    int    `json:"type"`
			Addr    string `json:"addr"`
			Value   int    `json:"value"`
			N       int    `json:"n"`
			Script  string `json:"script"`
		} `json:"out"`
	} `json:"x"`
}
