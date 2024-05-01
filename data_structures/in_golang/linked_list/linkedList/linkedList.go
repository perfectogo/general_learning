package main

import (
	"errors"
	"fmt"
)

type OrederBy string

const (
	Desc OrederBy = "desc"
	Asc  OrederBy = "asc"
)

// LinkedList ning node bir vaqtda o'ziga tegishli bo'lgan qiymatni va o'zidan keyingi nodening adressinig saqlaydi.
type Node struct {
	Data     int
	NextNode *Node
}

/*
LinkedList elementi nodelardan tashkil topgan chiziqli toplam hisoblanadi.
quyidagi LinkedList type o'zida nodelarning eng birinchisini va shu nodega ulangan barcha nodelarning sonini saqalash uchun fieldalri mavjud
*/
type LinkedList struct {
	HeadNode *Node
	Length   int
}

func (l *LinkedList) PrintLength() {
	if l != nil {
		fmt.Println(l.Length)
	}
}

func NewLinkedList(datas ...int) LinkedList {

	var linkedList = LinkedList{}

	for i := len(datas) - 1; i >= 0; i-- {
		newNode := &Node{Data: datas[i], NextNode: linkedList.HeadNode}
		linkedList.HeadNode = newNode
		linkedList.Length++
	}

	return linkedList
}

/*
LinkedListni boshidan bitta yangi node qo'shish nuchun method.
bu method data nomli parametrga ega bo'lib, bu parametrdan yangi node yaratish uchun foydalaniladi
*/
func (l *LinkedList) Prepend(data int) {

	node := &Node{Data: data}

	if l.HeadNode == nil {
		l.HeadNode = node
	} else {
		node.NextNode = l.HeadNode
		l.HeadNode = node
	}

	l.Length++
}

/*
LinkedListni boshidan bir-qancha yangi nodlar qo'shish nuchun method.
*/
func (l *LinkedList) PrependMultiple(datas ...int) {
	for i := len(datas) - 1; i >= 0; i-- {

		node := &Node{Data: datas[i]}

		if l == nil {
			l = &LinkedList{HeadNode: node}
		} else {
			node.NextNode = l.HeadNode
			l.HeadNode = node
		}

		l.Length++
	}
}

// LinkedList oxiridan yangi node qo'shish uchun method
func (l *LinkedList) Append(data int) {

	if l == nil {
		l = &LinkedList{HeadNode: &Node{Data: data}}
		l.Length++
		return
	}

	if l.HeadNode == nil {
		l.HeadNode = &Node{Data: data}
		return
	}

	lastNode := l.GetLastNode()
	lastNode.NextNode = &Node{Data: data}
	l.Length++
}

/*
LinkedListni oxirgi nodeini olish uchun method.
Bu method LinkedListning oxirgi nodeini qaytaradi va bu node dan LinkedList oxiridan yangi node qo'shish uchun foydalaniiladi
*/
func (l *LinkedList) GetLastNode() (node *Node) {

	node = l.HeadNode

	if node == nil {
		return
	}

	for node.NextNode != nil {
		node = node.NextNode
	}

	return
}

/*
LikedListning istalgan pozitsiyasidan yangi e'element qo'shish uchun method! bu methodning position va data nomli parametrli mavjud bo'lib
position dan kerakli positionni topish uchun va datadan yangi node yasash uchun foydalaniladi!
*/
func (l *LinkedList) Insert(position, data int) error {

	if position < 0 || position > l.Length+1 {
		return errors.New("insert: Index out of bounds")
	}

	newNode := &Node{Data: data, NextNode: nil}
	var prevNode, currentNode *Node
	currentNode = l.HeadNode

	for position > 0 {
		prevNode = currentNode
		currentNode = currentNode.NextNode
		position--
	}

	if prevNode != nil {
		prevNode.NextNode = newNode
		newNode.NextNode = currentNode
	} else {
		newNode.NextNode = currentNode
		l.HeadNode = newNode
	}

	l.Length++
	return nil
}

// LinkedList elementlarini ko'rsatuchi method
func (l *LinkedList) Print() {

	if l == nil {
		fmt.Println("[]")
		return
	}

	fmt.Print("[")
	for node := l.HeadNode; node != nil; node = node.NextNode {
		fmt.Print(node.Data)
		if node.NextNode != nil {
			fmt.Print(" ")
		}
	}
	fmt.Print("]\n")
}

func (l *LinkedList) DeleteFirst() error {
	if l == nil {
		return errors.New("DeleteFirst: List is nil")
	}

	if l.HeadNode == nil {
		return errors.New("DeleteFirst: List is empty")
	}

	l.HeadNode = l.HeadNode.NextNode
	l.Length--

	return nil
}

func (l *LinkedList) DeleteEnd() error {
	if l == nil {
		return errors.New("DeleteEnd: List is nil")
	}

	if l.HeadNode == nil {
		return errors.New("DeleteEnd: List is empty")
	}

	if l.HeadNode.NextNode == nil {
		l.HeadNode = nil
		l.Length--
		return nil
	}

	prev := l.HeadNode
	currentNode := l.HeadNode.NextNode

	for currentNode.NextNode != nil {
		prev = currentNode
		currentNode = currentNode.NextNode
	}

	prev.NextNode = nil
	l.Length--

	return nil
}

func (l *LinkedList) DeleteByPosition(position int) error {
	if l == nil {
		return errors.New("DeleteByPosition: List is nil")
	}

	if l.HeadNode == nil {
		return errors.New("DeleteByPosition: List is empty")
	}

	if position < 0 || position >= l.Length {
		return errors.New("DeleteByPosition: Index out of bounds")
	}

	if position == 0 {
		l.HeadNode = l.HeadNode.NextNode
		l.Length--
		return nil
	}

	current := l.HeadNode
	prev := &Node{}

	for position > 0 {
		prev = current
		current = current.NextNode
		position--
	}

	if current != nil {
		prev.NextNode = current.NextNode
		l.Length--
	}

	return nil
}

func (l *LinkedList) Clear() {
	if l != nil {
		l.HeadNode = nil
	}
}

/*
LinkedListning istalgan positsiyasidagi nodeining valuesni olish uchun metod
*/
func (l *LinkedList) GetDataByPosition(position int) (int, error) {

	if position < 0 || position > l.Length-1 {
		return 0, errors.New("GetDataByPosition: Index out of bounds")
	}

	currentNode := l.HeadNode

	for range position {
		currentNode = currentNode.NextNode
	}

	return currentNode.Data, nil
}

func (l *LinkedList) Traverse() {
	current := l.HeadNode

	for current != nil {
		// Access the data stored in the current node
		fmt.Println(current.Data)

		// Move to the next node
		current = current.NextNode
	}
}

func (l *LinkedList) GetMax() (int, error) {

	if l == nil {
		return 0, errors.New("GetMax: List is nil")
	}

	if l.HeadNode == nil {
		return 0, errors.New("GetMax: List is empty")
	}

	max := l.HeadNode.Data

	for currentNode := l.HeadNode.NextNode; currentNode != nil; currentNode = currentNode.NextNode {
		if currentNode.Data > max {
			max = currentNode.Data
		}
	}

	return max, nil
}

func (l *LinkedList) GetMin() (int, error) {
	if l == nil {
		return 0, errors.New("GetMax: List is nil")
	}

	if l.HeadNode == nil {
		return 0, errors.New("GetMax: List is empty")
	}

	min := l.HeadNode.Data

	for currentNode := l.HeadNode.NextNode; currentNode != nil; currentNode = currentNode.NextNode {
		if currentNode.Data < min {
			min = currentNode.Data
		}
	}

	return min, nil
}

func (l *LinkedList) Sort(ord OrederBy) {

}

func main() {
	//var list LinkedList = NewLinkedList(1, 2, 3)
	// list.Print()
	// fmt.Println(list.Length)

	// prepend methodini tekshirish
	// list.Prepend(4)
	// list.Print()
	// fmt.Println(list.Length)

	// append methodini tekshirish
	// list.Append(77)
	// list.Print()
	// fmt.Println(list.Length)

	// insert methodini tekshirish
	// if err := list.Insert(5, 7); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// list.Print()
	// fmt.Println(list.Length)

	// data, _ := list.GetDataByPosition(6)
	// fmt.Println(data)
	// fmt.Println(list.Length)

	// for i := range list.Length {
	// 	data, _ := list.GetDataByPosition(i)
	// 	fmt.Println(i, "data", data)
	// }

	// deleteFirst methodini tekshirish
	// for range list.Length {
	// 	list.DeleteFirst()
	// 	list.Print()
	// 	fmt.Println(list.Length)
	// }

	// list.Print()
	// list.DeleteEnd()
	// list.Print()

	// for range list.Length {
	// 	list.DeleteEnd()
	// 	list.Print()
	// 	list.PrintLength()
	// }
	// list.Print()

	// list.Print()
	// list.DeleteByPosition(0)
	// list.Print()

	var s LinkedList = NewLinkedList(1, 3, 4, 5, 6, 1, 3, 4, 5)
	fmt.Println(s.GetMax())
	s.MergeSort()
	s.Print()

}

func mergeSort(head *Node) *Node {
	// Base case: If the list is empty or has only one node, it's already sorted
	if head == nil || head.NextNode == nil {
		return head
	}

	// Find the middle of the list using slow and fast pointers
	var prev *Node
	slow, fast := head, head
	for fast != nil && fast.NextNode != nil {
		prev = slow
		slow = slow.NextNode
		fast = fast.NextNode.NextNode
	}

	// Split the list into two halves
	prev.NextNode = nil // Split the list by disconnecting the halves
	left := mergeSort(head)
	right := mergeSort(slow)

	// Merge the sorted halves
	return merge(left, right)
}

func merge(left, right *Node) *Node {
	dummy := &Node{} // Dummy node to simplify merging logic
	current := dummy

	// Merge the two sorted lists
	for left != nil && right != nil {
		if left.Data >= right.Data {
			current.NextNode = left
			left = left.NextNode
		} else {
			current.NextNode = right
			right = right.NextNode
		}
		current = current.NextNode
	}

	// Append any remaining nodes from the left or right list
	if left != nil {
		current.NextNode = left
	}
	if right != nil {
		current.NextNode = right
	}

	return dummy.NextNode // Return the merged list (skip the dummy node)
}

// MergeSort sorts a linked list in ascending order using Merge Sort
func (l *LinkedList) MergeSort() {
	l.HeadNode = mergeSort(l.HeadNode)
}
