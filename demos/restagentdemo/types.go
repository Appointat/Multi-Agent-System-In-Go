package restagentdemo

type Request struct {
	Operator string `json:"op"`
	Args     [2]int `json:"args"`
}

type Response struct {
	Result int `json:"res"`
}
