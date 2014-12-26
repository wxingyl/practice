#ifndef _TREE
#define _TREE

typedef struct tr_node
{
	int value;
	struct tr_node *left;
	struct tr_node *right;
} TreeNode;

typedef struct {
	TreeNode* root;
	int size;
} Tree;

void print_sorted_tree(TREE);
position find_min(TREE);
position find_max(TREE);
position find_value(TREE, Element);
position insert_value(TREE, Element);
Element delete_node(position);

static Element is_root(position);
static Element is_leaf(position);
static Element delete_leaf(position);
static void insert_node_to_nonempty_tree(TREE, position);

#endif // _TREE
