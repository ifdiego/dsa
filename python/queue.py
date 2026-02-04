class Queue:
    def __init__(self):
        self.data = []

    def is_empty(self):
        return len(self.data) == 0

    def enqueue(self, item):
        self.data.append(item)

    def dequeue(self):
        if self.is_empty():
            print("queue is empty!")
        else:
            del self.data[0]

    def print(self):
        print(self.data)

if __name__ == "__main__":
    queue = Queue()
    nums = [12, 6, 18, 19, 21, 11, 3, 5, 4, 24, 18]
    for num in nums:
        queue.enqueue(num)
    queue.print()
    queue.dequeue()
    queue.print()
