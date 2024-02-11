# Linky
Fully complete Linked list library. This a Go port of C library [lists](https://github.com/joegasewicz/lists).
### API

#### ListInit
ListInit should be called before any operation can be used on with
the linked list
```
l := LinkedListInit(data)
```

#### ListDestroy
ListDestroy destroys the linked list. No operations are permitted to
be called on the linked list after ListDestroy is called
```
err = LinkedListDestroy(l.LinkedList)
```

#### InsertNext
InsertNext inserts a node after the current linked list element.
If the element is nil then the new element gets placed at the head
of the list & then return the tail.
```
err, tailNode := l.LinkedList.InsertNext(l.tailNode, data)
```

#### RemoveNext
RemoveNext removes the next element after node & returns an error
if one exists & the removed node's data. If the node is nil then the
head of the list is removed & the head's data returned.
```
err, node := l.LinkedList.RemoveNext(nil)
```
#### ListSize
ListSize returns the total number of nodes in the list.
```
size := l.LinkedList.ListSize()
```
#### ListHead
ListHead returns the head of the list
```
head := l.LinkedList.ListHead()
```

### ListTail
ListTail returns the tail node of the list.
```
tail := l.Tail()
```

#### IsHead
IsHead determines if element is the head of the list.
```
isHead := l.IsHead(node)
```

#### IsTail
IsTail determines if element is the tail of the list.
```go
isTail := l.IsTail(node)
```
#### Evaluate
Evaluate returns the data of element's Node

```go
err, result = l.Evaluate(node)
```

#### NextNode
NextNode returns the next Node in the list.
```go
err, nextNode := l.NextNode(node)
```
