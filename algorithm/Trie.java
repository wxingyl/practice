import java.util.LinkedList;
import java.util.List;

/**
 * tire[Prefix Tree] 实现
 * https://leetcode.com/problems/implement-trie-prefix-tree/
 */

public class Trie {

    private TrieNode root;

    public Trie() {
        root = new TrieNode();
    }

    // Inserts a word into the trie.
    public void insert(String word) {
        if (search(word)) return;
        TrieNode curNode = root;
        for (char c : word.toCharArray()) {
            TrieNode node = curNode.getChildNode(c);
            if (node == null) {
                node = new TrieNode();
                node.c = c;
                curNode.list.add(node);
            }
            curNode = node;
        }
        curNode.isEnd = true;
    }

    // Returns if the word is in the trie.
    public boolean search(String word) {
        TrieNode curNode = root;
        for (char c : word.toCharArray()) {
            curNode = curNode.getChildNode(c);
            if (curNode == null) return false;
        }
        return curNode.isEnd;
    }

    // Returns if there is any word in the trie
    // that starts with the given prefix.
    public boolean startsWith(String prefix) {
        TrieNode curNode = root;
        for (char c : prefix.toCharArray()) {
            curNode = curNode.getChildNode(c);
            if (curNode == null) return false;
        }
        return true;
    }


}

class TrieNode {

    char c;

    boolean isEnd;

    List<TrieNode> list;

    // Initialize your data structure here.
    public TrieNode() {
        list = new LinkedList<TrieNode>();
    }

    public TrieNode getChildNode(char c) {
        for (TrieNode e : list) {
            if (e.c == c) return e;
        }
        return null;
    }

    public String toString() {
        return c + "";
    }
}

// Your Trie object will be instantiated and called as such:
// Trie trie = new Trie();
// trie.insert("somestring");
// trie.search("key");