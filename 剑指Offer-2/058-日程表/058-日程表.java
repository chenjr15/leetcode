class MyCalendar {
    // 防止区间冲突
    TreeMap<Integer,Integer> treemap ;
    public MyCalendar() {
        // key is start ,  value is end
        treemap = new TreeMap();
    }   
    
    public boolean book(int start, int end) {

        // 判断end，end 前一个，应该是start前一个，否则就有问题

        Map.Entry<Integer,Integer> beforeEnd =  treemap.floorEntry(end-1);
        // System.out.println(" "+booktimes+" Before ["+start+":"+end+"] is "+beforeEnd);
 
        if (beforeEnd!=null && beforeEnd.getKey()< end && beforeEnd.getValue()> start){
            return false;
        }
        // 插入新区间
        treemap.put(start,end);
        return true;

    }
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * MyCalendar obj = new MyCalendar();
 * boolean param_1 = obj.book(start,end);
 */