#ifndef _TREE_H
#define _TREE_H

typedef struct tr_node
{
    int value;
    struct tr_node *lchild;
    struct tr_node *rchild;
    struct tr_node *parent;
}Tr_Node;

typedef struct tree{
    Tr_Node *root;
    int size;
}Tree;

Tr_Node* insert(Tree*, int);
void display(Tree*);
Tr_Node* find(Tree*, int);
Tr_Node* find_min(Tree*);
Tr_Node* find_max(Tree*);
int delete(Tree*, int);
#endif // _TREE_H