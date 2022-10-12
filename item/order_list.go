package item

var OrderList Queue

type Queue struct {
	Elements []Cl_order_post
}

func (ol *Queue) Enqueue(order Cl_order_post) {
	ol.Elements = append(ol.Elements, order)
}

func (ol *Queue) Dequeue() *Cl_order_post {
	if ol.IsEmpty() {
		return nil
	}
	order := ol.Elements[0]
	if ol.GetSize() == 1 {
		ol.Elements = nil
		return &order
	}
	ol.Elements = ol.Elements[1:]
	return &order
}

func (ol *Queue) GetSize() int {
	return len(ol.Elements)
}

func (ol *Queue) IsEmpty() bool {
	return len(ol.Elements) == 0
}
