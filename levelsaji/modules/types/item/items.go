package item

import (
	"errors"
	"strconv"
)

type Items struct {
	ItemList []Item
	ItemMap  map[string]Item
}

func (items *Items) Initialize() {
	items.ItemList = []Item{}
	items.ItemMap = map[string]Item{}
}

func (items *Items) AddItem(item Item) error {
	if _, exists := items.ItemMap[item.Name]; exists {
		return errors.New(item.Name + "は既に追加済みです")
	} else {
		items.ItemList = append(items.ItemList, item)
		items.ItemMap[item.Name] = item
	}
	return nil
}

func (items *Items) GetItemByIndex(index int) (Item, error) {
	if len(items.ItemList) > index {
		return items.ItemList[index], nil
	} else {
		return Item{}, errors.New("商品番号[" + strconv.Itoa(index) + "]は存在しません")
	}
}

func (items *Items) GetItemByItemName(name string) (Item, error) {
	if item, exists := items.ItemMap[name]; exists {
		return item, nil
	} else {
		return Item{}, errors.New("商品名[" + name + "]は存在しません")
	}
}

func (items *Items) DeleteItemByIndex(index int) (Item, error) {
	if len(items.ItemList) > index {
		item := items.ItemList[index]
		items.ItemList = append(items.ItemList[:index], items.ItemList[index+1:]...)
		return item, nil
	} else {
		return Item{}, errors.New("商品番号[" + strconv.Itoa(index) + "]は存在しません")
	}
}

func (items *Items) DeleteItemByItemName(name string) (Item, error) {
	if item, exists := items.ItemMap[name]; exists {
		delete(items.ItemMap, name)
		return item, nil
	} else {
		return Item{}, errors.New("商品名[" + name + "]は存在しません")
	}
}
