package item

type Dh_order_post struct {
	Order_id     int     `json:"order_id"`
	Table_id     int     `json:"table_id"`
	Waiter_id    int     `json:"waiter_id"`
	Items        []int   `json:"items"`
	Priority     int     `json:"priority"`
	Max_wait     float32 `json:"max_wait"`
	Pick_up_time int64   `json:"pick_up_time"`
}

type K_order_post struct {
	Order_id        int              `json:"order_id"`
	Table_id        int              `json:"table_id"`
	Waiter_id       int              `json:"waiter_id"`
	Items           []int            `json:"items"`
	Priority        int              `json:"priority"`
	Max_wait        int              `json:"max_wait"`
	Pick_up_time    int64            `json:"pick_up_time"`
	Cooking_time    int              `json:"cooking_time"`
	Cooking_details []Cookingdetails `json:"cooking_details"`
}

type Cookingdetails struct {
	Food_id int `json:"food_id"`
	Cook_id int `json:"cook_id"`
}

func GenKorderpost(order Dh_order_post) *K_order_post {
	var order_post = new(K_order_post)

	order_post.Order_id = order.Order_id
	order_post.Table_id = order.Table_id
	order_post.Waiter_id = order.Waiter_id
	order_post.Items = order.Items
	order_post.Priority = order.Priority
	order_post.Max_wait = int(order.Max_wait)
	order_post.Pick_up_time = order.Pick_up_time

	return order_post
}
