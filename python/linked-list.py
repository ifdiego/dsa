class Node:
    def __init__(self, value):
        self.value = value
        self.next = None

class LinkedList:
    def __init__(self):
        self.head = None

    def append(self, value):
        new = Node(value)
        if not self.head:
            self.head = new
            return
        last = self.head
        while last.next:
            last = last.next
        last.next = new

    def remove(self, value):
        current = self.head
        if current and current.value == value:
            self.head = current.next
            current = None
            return
        previous = None
        while current and current.value != value:
            previous = current
            current = current.next
        if not current:
            print("value not found")
            return
        previous.next = current.next
        current = None

    def print(self):
        current = self.head
        while current:
            print(current.value, end=" -> ")
            current = current.next
        print("None")

if __name__ == "__main__":
    list = LinkedList()
    for num in [24, 83, 11, 56, 2, 71, 39, 48, 65, 90]:
        list.append(num)
    list.print()
    for num in [83, 56, 39]:
        list.remove(num)
    list.print()
