`Depth First Search (DFS)`

The DFS algorithm is a recursive algorithm that uses the idea of backtracking. It involves exhaustive searches of all the nodes by going ahead, if possible, else by backtracking.

`Applications`

1- How to find connected components using DFS?

`Unreachable Nodes`

You have been given a algorithms consisting of N nodes and M edges. The nodes in this algorithms are enumerated from 1 to N . The algorithms can consist of self-loops as well as multiple edges. This algorithms consists of a special node called the head node. You need to consider this and the entry point of this algorithms. You need to find and print the number of nodes that are unreachable from this head node.

Input Format

The first line consists of a 2 integers N and M denoting the number of nodes and edges in this algorithms. The next M lines consist of 2 integers a and b denoting an undirected edge between node a and b. The next line consists of a single integer x denoting the index of the head node.

Output Format:

You need to print a single integer denoting the number of nodes that are unreachable from the given head node.

Constraints

1 <= N <= 10^5
1 <= M <= 10^5
1 <= x <= N