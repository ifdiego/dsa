enum Color {
    RED,
    BLACK
}

class RedBlackTreeNode<T> {
    data: T;
    color: Color;
    left: RedBlackTreeNode<T> | null;
    right: RedBlackTreeNode<T> | null;
    parent: RedBlackTreeNode<T> | null;

    constructor(data: T, color: Color) {
        this.data = data;
        this.color = color;
        this.left = null;
        this.right = null;
        this.parent = null;
    }
}

class RedBlackTree<T> {
    root: RedBlackTreeNode<T> | null;

    constructor() {
        this.root = null;
    }

    leftRotate(node: RedBlackTreeNode<T>) {
        const right = node.right!;
        node.right = right.left;
        if (right.left != null) {
            right.left.parent = node;
        }
        right.parent = node.parent;
        if (node.parent === null) {
            this.root = right;
        } else if (node === node.parent.left) {
            node.parent.left = right;
        } else {
            node.parent.right = right;
        }
        right.left = node;
        node.parent = right;
    }

    rightRotate(node: RedBlackTreeNode<T>) {
        const left = node.left!;
        node.left = left.right;
        if (left.right != null) {
            left.right.parent = node;
        }
        left.parent = node.parent;
        if (node.parent === null) {
            this.root = left;
        } else if (node === node.parent.right) {
            node.parent.right = left;
        } else {
            node.parent.left = left
        }
        left.right = node;
        node.parent = left;
    }

    fixInsertion(node: RedBlackTreeNode<T>) {
        while (node != this.root && node.parent!.color === Color.RED) {
            if (node.parent === node.parent!.parent!.left) {
                const leaf = node.parent!.parent!.right;
                if (leaf !== null && leaf.color === Color.RED) {
                    node.parent!.color = Color.BLACK;
                    leaf.color = Color.BLACK;
                    node.parent!.parent!.color = Color.RED;
                    node = node.parent!.parent!;
                } else {
                    if (node === node.parent!.right) {
                        node = node.parent!;
                        this.leftRotate(node);
                    }
                    node.parent!.color = Color.BLACK;
                    node.parent!.parent!.color = Color.RED;
                    this.rightRotate(node.parent!.parent!);
                }
            } else {
                const leaf = node.parent!.parent!.left;
                if (leaf !== null && leaf.color === Color.RED) {
                    node.parent!.color = Color.BLACK;
                    leaf.color = Color.BLACK;
                    node.parent!.parent!.color = Color.RED;
                    node = node.parent!.parent!;
                } else {
                    if (node === node.parent!.left) {
                        node = node.parent!;
                        this.rightRotate(node);
                    }
                    node.parent!.color = Color.BLACK;
                    node.parent!.parent!.color = Color.RED;
                    this.leftRotate(node.parent!.parent!);
                }
            }
        }
        this.root!.color = Color.BLACK;
    }

    insert(data: T) {
        const newNode = new RedBlackTreeNode(data, Color.RED);
        let parent = null;
        let current = this.root;
        while (current != null) {
            parent = current;
            if (data < current.data) {
                current = current.left;
            } else {
                current = current.right;
            }
        }
        newNode.parent = parent;
        if (parent === null) {
            this.root = newNode;
        } else if (data < parent.data) {
            parent.left = newNode;
        } else {
            parent.right = newNode;
        }
        this.fixInsertion(newNode);
    }

    inOrderTraversal(node: RedBlackTreeNode<T> | null) {
        if (node !== null) {
            this.inOrderTraversal(node.left);
            console.log(node.data + " " + (node.color === Color.RED ? "Red" : "Black"));
            this.inOrderTraversal(node.right);
        }
    }
}

const rbtree = new RedBlackTree();
rbtree.insert(10);
rbtree.insert(20);
rbtree.insert(5);
rbtree.insert(15);
rbtree.inOrderTraversal(rbtree.root);
