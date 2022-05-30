// ./program  13.45s user 0.88s system 95% cpu 15.073 total
#include <stdlib.h>
#include <stdio.h>

typedef struct Tree {
    struct Tree* Left;
    struct Tree* Right;
} Tree;

Tree* new_tree(int depth) {
    if (depth == 0) {
        return (Tree*) NULL;
    }
    Tree* t = (Tree*) malloc(sizeof(Tree));
    t->Left = new_tree(depth -1);
    t->Right = new_tree(depth -1);
    return t;
}

int count(Tree* t) {
    if (t->Left == NULL) {
        return 1;
    }
    return 1 + count(t->Left) + count(t->Right);
}

int main() {
    for (int i = 0; i < 500; i++) {
        Tree* root = new_tree(20);
        int c = count(root);
    }
}