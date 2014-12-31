#include <stdio.h>
#include <stdlib.h>
#include "tree.h"


static int delete_leaf(Tr_Node* np)
{
    Tr_Node*parent;
    int element;
    element = np->value;
    if(parent->lchild == np)
    {
        parent->lchild = NULL;
    }
    else
    {
        parent->rchild = NULL;
    }
    free(np);
    return element;

}

static void insert_node_to_nonempty_tree(Tr_Node* tr, Tr_Node* np)
{
    if(np->value <= tr->value)
    {
        if(tr->lchild == NULL)
        {
            tr->lchild = np;
            return;
        }
        else
        {
            insert_node_to_nonempty_tree(tr->lchild, np);
        }
    }
    else if(np->value > tr->value)
    {
        if(tr->rchild == NULL)
        {
            tr->rchild = np;
            return;
        }
        else
        {
            insert_node_to_nonempty_tree(tr->rchild, np);
        }

    }
}

//print values of the tree in sorted order 在分类树的顺序打印值
static void display_all(Tr_Node* np)
{
    if (np == NULL) return;
    display_all(np->lchild);
    printf("%d\n", np->value);
    display_all(np->rchild);
}

void display(Tree* tr)
{
    display_all(tr->root);
}

Tr_Node* find_min(Tree* tr)
{
    Tr_Node* np = tr->root;
    if(np == NULL)
    {
        return NULL;
    }
    while(np->lchild != NULL)
    {
        np = np->lchild;
    }
    return np;
}


Tr_Node* find_max(Tree* tr)
{
    Tr_Node* np = tr->root;
    if(np == NULL)
    {
        return NULL;
    }
    while(np->rchild != NULL)
    {
        np = np->rchild;
    }
    return np;
}

static Tr_Node* find_by_value(Tr_Node *node, int element)
{
    if (node == NULL) return NULL;
    if (node->value == element)
    {
        return node;
    }
    else if (node->value > element)
    {
        return find_by_value(node->lchild, element);
    } else {
        return find_by_value(node->rchild, element);
    }
}

Tr_Node* find(Tree* tr, int element)
{
    return find_by_value(tr->root, element);
}


Tr_Node* insert(Tree* tr, int value)
{
    Tr_Node* np;
    np = (Tr_Node*)malloc(sizeof(struct tr_node));
    np->value = value;
    np->lchild = NULL;
    np->rchild = NULL;
    if(tr == NULL)
    {
        tr->root = np;
    }
    else
    {
        insert_node_to_nonempty_tree(tr->root, np);
    }
    return tr->root;
}
