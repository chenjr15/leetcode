/**
 * The rand7() API is already defined in the parent class SolBase.
 * public int rand7();
 * @return a random integer in the range 1 to 7
 */
class Solution extends SolBase {
    public int rand10() {
        // (randN() - 1) * N + randN() 可以获得[1,N*N]的等概率随机数，多出来的拒接采样即可
        int a,b,c;
        do{
            a = rand7();
            b = rand7();
            c = (a-1)*7 + b ;
            // 拒绝采样大于40的值
        }while(c>40);
        // 得到的c是1-40的等概率
        return (c-1)%10 +1;
    }
}
