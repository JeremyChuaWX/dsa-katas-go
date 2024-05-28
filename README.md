# Dsa Katas - Go

Some katas to practice DSA in Go.

## Content

### Algorithms

- \[x\] Binary search
- \[x\] Selection sort
- \[x\] Bubble sort
- \[x\] Inserstion sort
- \[x\] Merge sort
- \[x\] Quick sort
- \[x\] DFS
- \[x\] BFS
- \[ \] Dijkstra's algorithm
- \[ \] Prim's algorithm

### Data Structures

- \[ \] Binary search tree
- \[ \] AVL tree
- \[ \] AB tree
- \[ \] Heap
- \[ \] Trie
- \[ \] Hashmap
- \[ \] LRU

### Miscellaneous

- \[ \] Coin change problem
- \[ \] Comparing binary trees
- \[ \] Maximum subarray

## How to use

- To test the katas, run the command

  ```bash
  go test -v ./katas_test
  ```

- To generate stubs from template files, run the command

  ```bash
  ./scripts/generate.sh -n <name-of-kata>
  ```

  Add the `-f` flag to forcefully reset existing katas

- To clear all kata files, run the command

  ```bash
  ./scripts/clear.sh
  ```

  Add `-n <name-of-kata>` to clear a specific kata
