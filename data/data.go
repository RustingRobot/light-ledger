package data

type Data struct {
	Expenses expenses `json:"expenses"`
}

type expenses struct {
	Description []string `json:"desc"`
	Cost        []string `json:"cost"`
}
