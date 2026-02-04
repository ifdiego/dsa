class Stack:
    def __init__(self):
        self.data = []

    def is_empty(self):
        return len(self.data) == 0

    def push(self, item):
        self.data.append(item)

    def pop(self):
        if self.is_empty():
            print("stack is empty!")
        else:
            del self.data[-1]

    def print(self):
        print(self.data)

if __name__ == "__main__":
    stack = Stack()
    nums = [12, 6, 18, 19, 21, 11, 3, 5, 4, 24, 18]
    for num in nums:
        stack.push(num)
    stack.print()
    stack.pop()
    stack.print()
