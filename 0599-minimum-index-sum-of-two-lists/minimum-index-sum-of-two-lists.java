class Solution {
    public String[] findRestaurant(String[] list1, String[] list2) {
        if (list1 == null || list2==null ){
            return new String[0];
        }
        HashMap<String,Integer> map1 = new HashMap<>();
        ArrayList<String> common = new ArrayList();;
        int minSum = list2.length+list1.length;
        for(int i =0;i<list1.length;++i){
            map1.put(list1[i],i);
        }
        String s;
        for(int i =0;i<list2.length;++i){
           s = list2[i];
           if (s==""){
               continue;
           }
           Integer sum = map1.get(s);
           if(sum == null){
               continue;
           }
           sum += i;
           if(sum == minSum){
               common.add(s);
           }else if (sum < minSum){
               common.clear();
               common.add(s);
               minSum = sum;
           }
        //    System.out.println(common);
        }
        String[] arr = new String[common.size()];
        common.toArray(arr);
        return arr;
    }
}