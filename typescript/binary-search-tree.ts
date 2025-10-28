class BinarySearchTreeNode<T> {
    data: T;
    left: BinarySearchTreeNode<T> | null;
    right: BinarySearchTreeNode<T> | null;

    constructor(data: T) {
        this.data = data;
        this.left = null;
        this.right = null;
    }
}

class BinarySearchTree<T> {
    root: BinarySearchTreeNode<T> | null;

    constructor() {
        this.root = null;
    }

    insert(value: T): void {
        const newNode = new BinarySearchTreeNode(value);
        if (!this.root) {
            this.root = newNode;
        } else {
            this.insertNode(this.root, newNode);
        }
    }

    private insertNode(current: BinarySearchTreeNode<T>, newNode: BinarySearchTreeNode<T>): void {
        if (current.data > newNode.data) {
            if (!current.right) {
                current.right = newNode;
            } else {
                this.insertNode(current.right, newNode);
            }
        } else {
            if (!current.left) {
                current.left = newNode;
            } else {
                this.insertNode(current.left, newNode);
            }
        }
    }

    search(value: T): boolean {
        return this.searchNode(this.root, value);
    }

    private searchNode(current: BinarySearchTreeNode<T> | null, value: T): boolean {
        if (!current) return false;
        if (value < current.data) {
            return this.searchNode(current.left, value);
        } else if (value > current.data) {
            return this.searchNode(current.right, value);
        } else {
            return true;
        }
    }

    delete(value: T): void {
        this.root = this.deleteNode(this.root, value);
    }

    private deleteNode(current: BinarySearchTreeNode<T> | null, value: T): BinarySearchTreeNode<T> | null {
        if (current === null) {
            console.log("current null:", current);
            return null;
        }
        if (value < current.data) {
            console.log("current left:", current);
            current.left = this.deleteNode(current.left, value);
        } else if (value > current.data) {
            console.log("current right:", current);
            current.right = this.deleteNode(current.right, value);
        } else {
            console.log("current delete:", current);
            if (current.left === null && current.right === null) return null;
            if (current.left === null) return current.right;
            if (current.right == null) return current.left;
            const min = this.minValue(current.right);
            current.data = min;
            current.right = this.deleteNode(current.right, min);
        }
        return current;
    }

    minValue(node: BinarySearchTreeNode<T>): T {
        let current = node;
        while (current.left)
            current = current.left;
        return current.data;
    }

    inOrderTraverse(current: BinarySearchTreeNode<T> | null): void {
        if (current) {
            this.inOrderTraverse(current.left);
            console.log(current.data);
            this.inOrderTraverse(current.right);
        }
    }

    preOrderTraverse(current: BinarySearchTreeNode<T> | null): void {
        if (current) {
            console.log(current.data);
            this.inOrderTraverse(current.left);
            this.inOrderTraverse(current.right);
        }
    }

    postOrderTraverse(current: BinarySearchTreeNode<T> | null): void {
        if (current) {
            this.inOrderTraverse(current.left);
            this.inOrderTraverse(current.right);
            console.log(current.data);
        }
    }
}

const tree = new BinarySearchTree<number>();
tree.insert(42);
tree.insert(8);
tree.insert(99);
tree.insert(31);
console.log("inorder traverse:");
tree.inOrderTraverse(tree.root);
console.log("preorder traverse:");
tree.preOrderTraverse(tree.root);
console.log("postorder traverse:");
tree.postOrderTraverse(tree.root);
console.log("search 42:", tree.search(42));
console.log("search 98:", tree.search(98));
console.log("delete:", tree.delete(31));
console.log("inorder traverse:");
tree.inOrderTraverse(tree.root);
console.log("preorder traverse:");
tree.preOrderTraverse(tree.root);
console.log("postorder traverse:");
tree.postOrderTraverse(tree.root);
