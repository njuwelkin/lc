
// Definition for a Node.
class Node {
public:
    int val;
    Node* prev;
    Node* next;
    Node* child;

    Node() {}

    Node(int _val, Node* _prev, Node* _next, Node* _child) {
        val = _val;
        prev = _prev;
        next = _next;
        child = _child;
    }
};

class Solution {
        Node* next(Node* node) {
                if (node->child != nullptr) {
                        return node->child;
                }                        
                if (node->next != nullptr) {
                        return node->next;
                }                  
                for (Node* prev = node->prev; prev != nullptr; prev = prev->prev) {
                        if (prev->next != nullptr) {
                                return prev->next;
                        }                  
                }                          
                return nullptr;
        }             
public:        
    Node* flatten(Node* head) {
        Node* p; 
        for (p = head; p != nullptr; ) {
                Node* q = next(p);
                q->prev = p;
                p = q;
        }
        for (; p != head; p = p->prev) {
                p->prev->next = p;
        }
	return head
    }         
}; 

