class Solution {
    public int[] findOrder(int numCourses, int[][] prerequisites) {
        // 找入度为0 的节点
        ArrayList<Integer>[] afterThis = new ArrayList[numCourses];
        int[] beforeThis = new int[numCourses];

        // for(int i = 0;i<numCourses;++i){
        //    beforeThis[i] =  new ArrayList<>();
        // }
        for(int[] req: prerequisites){
            int a = req[0];
            int b = req[1];
            // b before a
            if( afterThis[b] ==null){
                 afterThis[b]=  new ArrayList<>();
            }
            afterThis[b].add(a);
            beforeThis[a]++;
            // System.out.println(Arrays.toString(afterThis)+""+Arrays.toString(beforeThis));
        }
        int[] result = new int[numCourses];
        for(int i = 0;i<numCourses;++i){
            // 找入度为0的点
            boolean found=false;
            for(int j= 0;j<numCourses;++j){
                if (beforeThis[j] == 0  ){
                    found = true;
                    result[i] = j;
                    // 把该点指向的所有边删掉，再把该节点删掉
                    if (afterThis[j]!=null){
                        for(int depThis : afterThis[j]){
                            beforeThis[depThis]--;
                        }
                    }
                  
                    beforeThis[j] = -1;
                    break;
                }
            }
            if (!found){
                // 有环
                return new int[0];
            }
        }
        return result;

    }
}