package checkout

import "fmt"

func RunDemo() {
	PrintHeader("Hola Checkout Engine")
	order := NewOrder("ORDER-001", "NESTOR")
	AddItem(&order, Item{
		SKU:   "A1",
		Name:  "Producto A",
		Price: 10,
		Qty:   1,
	})
	AddItem(&order, Item{
		SKU:   "B2",
		Name:  "Producto B",
		Price: 20,
		Qty:   2,
	})

	status := RemoveItem(&order, "A1")
	if status {
		PrintKV("REMOVE_STATUS", "SUCCESS")
	} else {
		PrintKV("REMOVE_STATUS", "FAIL")
	}

	PrintKV("ORDER_ID", order.ID)
	PrintKV("CUSTOMER", order.Customer)
	PrintKV("TOTAL_ITEMS", fmt.Sprintf("%d", len(order.Items)))
	PrintDivider()
}
