class Solution {
    public int calPoints(String[] ops) {
        ArrayList<Integer> stack = new ArrayList<>();
        int total = 0;
        int point = 0;
        for (String op : ops){
            if(op==null || "".equals(op)){
                continue;
            }
            if ("+".equals(op)){
                // 前两次得分之和
                point = stack.get(stack.size()-1)+stack.get(stack.size()-2);
                stack.add(point);
                total += point;

            }else if ("D".equals(op)){
                // 前一次得分的两倍
                point = stack.get(stack.size()-1)*2;
                stack.add(point);
                total += point;

            }else if ( "C".equals(op)){
                // 前一次得分无效
                point = stack.remove(stack.size()-1);
                total-= point;

            }else {
                // 数字得分
                point = Integer.parseInt(op);
                stack.add(point);
                total+=point;
            }
        }
        // System.out.println(stack);
        return total;
    }
}