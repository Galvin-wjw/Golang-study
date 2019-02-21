package main

import (
	"fmt"
	"reflect"
)

// 定义节点
type ListNode struct {
	next *ListNode
	value interface{}
}

//单链表
type LinkedList struct {
	head *ListNode
}

func (this *LinkedList) Print()  {
	cur := this.head
	format := ""
	for nil != cur{
		format += fmt.Sprintf("%+v",cur.value)
		cur = cur.next
		if nil != cur{
			format += "->"
		}

	}
	fmt.Println(format)
	//fmt.Println(reflect.TypeOf(format))
}

//在节点 a b之间插入节点x O(1)
func (this *LinkedList) insert(a *ListNode, b *ListNode,x *ListNode){
	x.next = b
	a.next = x
}

//删除节点x的后续节点  O(1)
func (this *LinkedList) delete(x *ListNode)  {
	x.next = x.next.next
}

// 单链表反转 O(n)
func (this *LinkedList) Reverse()  {
	if nil == this.head || nil == this.head.next {
		return
	}

	var pre *ListNode = nil
	cur := this.head.next
	for nil != cur{
		tmp := cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}

	this.head.next = pre
}


func main() {
	var l *LinkedList

	n5 := &ListNode{value:5,next:nil}
	n4 := &ListNode{value:4,next:n5}
	n3 := &ListNode{value:3,next:n4}
	n2 := &ListNode{value:2,next:n3}
	n1 := &ListNode{value:1,next:n2}
	fmt.Println(reflect.TypeOf(n1))
	l = &LinkedList{head:&ListNode{next:n1}}
	l.Print()

	n6 := &ListNode{value:0}
	l.insert(n3,n4,n6)
	l.Print()

	l.delete(n4)
	l.Print()

	l.Reverse()
	l.Print()

}
