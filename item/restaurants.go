package item

var Menu_List Restaurants

type Restaurants struct {
	NRestaurants     int               `json:restaurants`
	Restaurants_Data []Restaurant_Data `json:restaurants_data`
}

type Restaurant_Data struct {
	Name       string `json:name`
	Menu_Items int    `json:menu_items`
	Menu       []Food `json:menu`
}

func GenRestaurantsMenu() *Restaurants {
	var Rpayload = new(Restaurants)
	Rpayload.NRestaurants = len(RestaurantList.Elements)
	var RD []Restaurant_Data
	RD = GetListOfRestaurants()
	Rpayload.Restaurants_Data = RD
	return Rpayload
}

func GetListOfRestaurants() []Restaurant_Data {
	var RD []Restaurant_Data
	for _, restaurant := range RestaurantList.Elements {
		var Restaurant_Data = new(Restaurant_Data)
		Restaurant_Data.Name = restaurant.Name
		Restaurant_Data.Menu_Items = restaurant.Menu_Items
		Restaurant_Data.Menu = restaurant.Menu
		RD = append(RD, *Restaurant_Data)
	}
	return RD
}
