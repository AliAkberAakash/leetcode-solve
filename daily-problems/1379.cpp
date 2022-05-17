/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */

class Solution {
public:
    TreeNode* getTargetCopy(TreeNode* original, TreeNode* cloned, TreeNode* target) {
        queue<TreeNode*> q;
        
        q.push(cloned);
        
        while(!q.empty()){
            TreeNode* currentNode = q.front();
            q.pop();
            
            if(currentNode->val == target->val){
                cout<<&currentNode<<" "<<target<<endl;
                return currentNode;
            }
            
            if(currentNode->left != NULL){
                q.push(currentNode->left);
            }
            if(currentNode->right != NULL){
                q.push(currentNode->right);
            }
            
        }
        
        return NULL;
    }
};
