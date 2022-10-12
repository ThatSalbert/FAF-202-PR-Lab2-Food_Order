package item

type Order struct {
	Restaurant_Id int     `json:"id"`
	Items         []int   `json:"items"`
	Priority      int     `json:"priority"`
	Max_wait      float32 `json:"max_wait"`
	Created_Time  int64   `json:"created_time"`
}
