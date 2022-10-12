package item

type FS_order_post struct {
	Items        []int   `json:"items"`
	Priority     int     `json:"priority"`
	Max_wait     float32 `json:"max_wait"`
	Created_Time int64   `json:"created_time"`
}

type Cl_order_post struct {
	Client_Id int   `json:client_id`
	Orders    Order `json:orders`
}

type Register_post struct {
	Restaurant_Id int    `json:restaurant_id`
	Name          string `json:name`
	Address       string `json:address`
	Menu_Items    int    `json:menu_items`
	Menu          []Food `json:menu`
}

func GenFS_order_post(order Cl_order_post) *FS_order_post {
	var fs_order = new(FS_order_post)
	fs_order.Items = order.Orders.Items
	fs_order.Priority = order.Orders.Priority
	fs_order.Max_wait = order.Orders.Max_wait
	fs_order.Created_Time = order.Orders.Created_Time
	return fs_order
}
