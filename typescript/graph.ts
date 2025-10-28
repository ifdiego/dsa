class GraphNode<T> {
    data: T;
    neighbors: GraphNode<T>[];

    constructor(data: T) {
        this.data = data;
        this.neighbors = [];
    }

    addNeighbor(node: GraphNode<T>): void {
        this.neighbors.push(node);
    }
}

class Graph<T> {
    nodes: GraphNode<T>[];

    constructor() {
        this.nodes = [];
    }

    addNode(value: T): GraphNode<T> {
        let current = this.findNode(value);
        if (!current) {
            const node = new GraphNode(value);
            this.nodes.push(node);
            return node;
        }
        return current;
    }

    addEdge(source: T, destination: T): void {
        const src = this.addNode(source);
        const dest = this.addNode(destination);
        src.addNeighbor(dest);
        dest.addNeighbor(src);
    }

    findNode(value: T): GraphNode<T> | undefined {
        return this.nodes.find(node => node.data === value);
    }

    breadthFirstSearch(node: T): void {
        const start = this.findNode(node);
        if (!start) return;
        const visited: Set<GraphNode<T>> = new Set();
        const queue: GraphNode<T>[] = [];
        visited.add(start);
        queue.push(start);
        while (queue.length > 0) {
            const current = queue.shift()!;
            console.log(current.data);
            for (const neighbor of current.neighbors) {
                if (!visited.has(neighbor)) {
                    visited.add(neighbor);
                    queue.push(neighbor);
                }
            }
        }
    }

    depthFirstSearch(node: T): void {
        const start = this.findNode(node);
        if (!start) return;
        const visited: Set<GraphNode<T>> = new Set();
        const stack: GraphNode<T>[] = [];
        stack.push(start);
        while (stack.length > 0) {
            const current = stack.pop()!;
            if (!visited.has(current)) {
                console.log(current.data);
                visited.add(current);
                for (const neighbor of current.neighbors) {
                    stack.push(neighbor);
                }
            }
        }
    }

    printGraph(): void {
        for (const node of this.nodes) {
            const neighbors = node.neighbors.map(neighbor => neighbor.data).join(", ");
            console.log(`${node.data}: [${neighbors}]`);
        }
    }
}

const graph = new Graph<number>();
graph.addEdge(1, 2);
graph.addEdge(1, 3);
graph.addEdge(2, 4);
graph.addEdge(3, 5);
graph.addEdge(4, 5);
graph.printGraph();
console.log();
graph.breadthFirstSearch(1);
console.log();
graph.depthFirstSearch(1);
