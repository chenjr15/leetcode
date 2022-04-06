class Solution {
    public List<List<Integer>> generate(int numRows) {
        List<List<Integer>> result = new ArrayList<List<Integer>>();
        ArrayList<Integer> list=  new ArrayList<Integer>();
        list.add(1);
        result.add(list);
        for(int row = 1;row < numRows;++row){
            int colMax = row+1;
            list=  new ArrayList<Integer>();
            list.add(1);
            for (int col = 1 ;col<colMax;col++){
                // System.out.println("row:"+row+" col:"+col);
                int val = result.get(row-1).get(col-1);
                if (col < row){
    
                    val += result.get(row-1).get(col);
                }
                list.add(val);
            }
            result.add(list);
        }
        return result;    
    }

}