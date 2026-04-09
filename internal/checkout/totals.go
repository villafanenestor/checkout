package checkout

func NewOrder(id string, customer string) Order {
	return Order{
		ID:       id,
		Customer: customer,
		Items:    []Item{},
		Meta:     map[string]string{},
	}
}

func AddItem(o *Order, item Item) {
	o.Items = append(o.Items, item)
}

func RemoveItem(o *Order, sku string) bool {

	for i, item := range o.Items {
		if item.SKU == sku {
			// Se usa append para crear una nueva slice que omita el elemento en la posición i
			// en golang se usan los ... para desempaquetar los elementos de la slice,
			// es como en js ... la diferencia es que en js se usa para spread y en go para desempaquetar
			// o.Items[:i] es la slice desde el inicio hasta la posición i
			// o.Items[i+1:] es la slice desde la posición i+1 hasta el final
			o.Items = append(o.Items[:i], o.Items[i+1:]...)
			return true
		}
	}
	return false
}
