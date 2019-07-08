`Applications`

1. How to determine the level of each node in the given tree?

2. 0-1 BFS. This type of BFS is used to find the shortest distance between two nodes in a algorithms provided that the edges in the algorithms have the weights 0 or 1. If you apply the BFS explained earlier in this article, you will get an incorrect result for the optimal distance between 2 nodes. The weights have a constraint that they can only be 0 or 1.

`Level Nodes`

You have been given a Tree consisting of N nodes. A tree is a fully-connected algorithms consisting of N nodes and  edges. The nodes in this tree are indexed from 1 to N. Consider node indexed 1 to be the root node of this tree. The root node lies at level one in the tree. You shall be given the tree and a single integer x . You need to find out the number of nodes lying on level x.

Input Format

The first line consists of a single integer N denoting the number of nodes in the tree. Each of the next  lines consists of 2 integers a and b denoting an undirected edge between node a and node b. The next line consists of a single integer x.

Output Format

You need to print a single integers denoting the number of nodes on level x.

Constraints

1 <= N <= 10^5
1 <= A,B <= N

Note

It is guaranteed that atleast one node shall lie on level x