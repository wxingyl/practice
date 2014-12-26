#include <stdio.h>
#include <stdlib.h>
#include "tree.h"
//print values of the tree in sorted order 在分类树的顺序打印值

static int is_root(position np)
{
   return (np->parent == NULL);
}

static int is_leaf(position np)
{
    return (np->lchild == NULL && np->rchild == NULL);
}

void print_sorted_tree(TREE tr)
{
    if(tr == NULL){
        return;
    }
    print_sorted_tree(tr->lchild);
    printf("%d\n", tr->element);
    print_sorted_tree(tr->rchild);
}


position find_min(TREE tr)
{
    position np;
    np = tr;
    if(np == NULL){
        return NULL;
    }
    while(np->lchild != NULL)
    {
        np = np->lchild;
    }
    return np;
}


position find_max(TREE tr)
{
    position np;
    np = tr;
    if(np == NULL){
        return NULL;
    }
    while(np->rchild != NULL)
    {
        np = np->rchild;
    }
    return np;
}

position find_value(TREE tr, Element value)
{
    position np;
    np = tr;
    if(tr == NULL){
        return NULL;
    }
    if(tr->element == value){
        return tr;
    }
    else if(value < tr->element){
            return find_value(tr->lchild, value);
    }
    else{
        return find_value(tr->rchild, value);
    }

}


position insert_value(TREE tr, Element value)
{
    position np;
    np = (position)malloc(sizeof(struct tr_node));
    np->element = value;
    np->parent = NULL;
    np->lchild = NULL;
    np->rchild = NULL;
    if(tr == NULL){
        tr = np;
    }
    else{
        insert_node_to_nonempty_tree(tr, np);
    }
    return tr;
}


Element delete_node(position np)
{
    position replace;
    Element element;
    if(is_leaf(np)){
        return delete_leaf(np);
    }
    else{
        replace = (np->lchild != NULL) ? find_max(np->lchild) : find_min(np->rchild);
        element = np->element;
        np->element = delete_node(replace);
        return element;
    }
}



static Element delete_leaf(position np)
{
    position parent;
    Element element;
    parent = np->parent;
    element = np->element;
    if(!is_root(np)){
        if(parent->lchild == np)
        {
            parent->lchild = NULL;
        }
        else{
            parent->rchild = NULL;
        }
    free(np);
    return element;
    }
}


static void insert_node_to_nonempty_tree(TREE tr, position np)
{
    if(np->element <= tr->element)
    {
        if(tr->lchild == NULL)
		{
            tr->lchild = np;
            np->parent = tr;
            return;
		} else
		{
            insert_node_to_nonempty_tree(tr->lchild, np);
        }

    }
    else if(np->element > tr->element){
        if(tr->rchild == NULL)
		{
            tr->rchild = np;
            np->parent = tr;
            return;
        }
        else{
            insert_node_to_nonempty_tree(tr->rchild, np);
        }
    }
}
