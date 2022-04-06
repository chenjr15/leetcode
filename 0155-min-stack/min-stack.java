class MinStack {
    ArrayList<Integer> stack = new ArrayList();
    ArrayList<Integer> helper= new ArrayList();
    public MinStack() {

    }
    
    public void push(int val) {
        stack.add(val);
        if(helper.size()==0){
            helper.add(val);
        }else{
            if (helper.get(helper.size()-1)>=val){
                helper.add(val);
            }
        }
    }   
    
    public void pop() {
        int val =  stack.get(stack.size()-1);
        stack.remove(stack.size()-1);
        if (val == helper.get(helper.size()-1)){
            helper.remove(helper.size()-1);
        }
    }
    
    public int top() {
        // if (stack.size()==0){
        //     return -1;
        // }
        // 应该如果数量不对应该跳异常
        return stack.get(stack.size()-1);
    }
    
    public int getMin() {
        // 应该如果数量不对应该跳异常
        return ( helper.get(helper.size()-1));
    }
}

/**
 * Your MinStack object will be instantiated and called as such:
 * MinStack obj = new MinStack();
 * obj.push(val);
 * obj.pop();
 * int param_3 = obj.top();
 * int param_4 = obj.getMin();
 */