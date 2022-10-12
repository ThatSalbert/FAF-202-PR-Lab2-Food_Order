package item

var RestaurantList Restaurant_Queue

type Restaurant_Queue struct {
	Elements []Register_post
}

func (rl *Restaurant_Queue) Enqueue(order Register_post) {
	rl.Elements = append(rl.Elements, order)
}

func (rl *Restaurant_Queue) SearchFor(Restaurant_Id int) *Register_post {
	for _, restaurant := range rl.Elements {
		if Restaurant_Id == restaurant.Restaurant_Id {
			return &restaurant
		}
	}
	return nil
}

func (rl *Restaurant_Queue) IsEmpty() bool {
	return len(rl.Elements) == 0
}
