class Solution {
    public String addBinary(String a, String b) {
        int pa = a.length() - 1;
        int pb = b.length() - 1;
        int carry = 0;
        int sum = 0 ;
        StringBuilder  sb = new StringBuilder();
        while(pa>-1 && pb >-1 ){
            sum = a.charAt(pa) - '0' + b.charAt(pb) -'0' + carry;
            carry = sum >>1;
            sum %=2;
            sb.append((char)('0'+sum));
            // System.out.println("" + sum%2 + " "+carry+" sb:"+sb.toString());
            
            pa--;
            pb--;
        }
        String remain = a;
        int pr = pa; 
        if( pa == -1 ){
            pr=pb;
            remain = b;
        }
        
        for(;pr>-1;pr--){
            sum = remain.charAt(pr) - '0' + carry;
            // System.out.println("" + sum%2 + " "+carry+" sb:"+sb.toString());

            sb.append((char)('0'+(sum %2)));
            carry = sum >>1;
        }
        if(carry>0){
            sb.append('1');
        }
        
        return sb.reverse().toString();
    }
}