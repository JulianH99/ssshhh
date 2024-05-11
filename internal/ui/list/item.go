package list

import "github.com/charmbracelet/bubbles/list"

type listItem struct {
	title string
	desc  string
}

func (i listItem) FilterValue() string {
	return i.title
}

func (i listItem) Title() string {
	return i.title
}

func (i listItem) Description() string {
	return i.desc
}

func fromBubbleItem(i list.Item) listItem {
	return i.(listItem)
}

func fromBubbleArray(items []list.Item) []listItem {
	listItems := make([]listItem, len(items))

	for i, item := range items {
		listItems[i] = fromBubbleItem(item)
	}

	return listItems
}

func NewItem(title, desc string) listItem {
	return listItem{title: title, desc: desc}
}
