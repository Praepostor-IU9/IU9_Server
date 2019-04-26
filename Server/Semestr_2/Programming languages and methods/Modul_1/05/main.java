import java.util.*;
import java.math.*;
        
public class FastFib {
    private BigInteger leftUp, rightUp, leftDown, rightDown;
    
    private void mul(FastFib val) {
        BigInteger a, b, c, d;
        a = this.leftUp.multiply(val.leftUp).add(this.rightUp.multiply(val.leftDown));
        b = this.leftUp.multiply(val.rightUp).add(this.rightUp.multiply(val.rightDown));
        c = this.leftDown.multiply(val.leftUp).add(this.rightDown.multiply(val.leftDown));
        d = this.leftDown.multiply(val.rightUp).add(this.rightDown.multiply(val.rightDown));
        this.leftUp = a;
        this.rightUp = b;
        this.leftDown = c;
        this.rightDown = d;
    }
    
    private void exp(FastFib val) {
        this.leftUp = val.leftUp.multiply(val.leftUp).add(val.rightUp.multiply(val.leftDown));
        this.rightUp = val.leftUp.multiply(val.rightUp).add(val.rightUp.multiply(val.rightDown));
        this.leftDown = val.leftDown.multiply(val.leftUp).add(val.rightDown.multiply(val.leftDown));
        this.rightDown = val.leftDown.multiply(val.rightUp).add(val.rightDown.multiply(val.rightDown));
    }
    
    public FastFib() {
        this.leftUp = BigInteger.ONE;
        this.rightUp = BigInteger.ONE;
        this.leftDown = BigInteger.ONE;
        this.rightDown = BigInteger.ZERO;
    }
    
    private void identityMatrix() {
        this.leftUp = BigInteger.ONE;
        this.rightUp = BigInteger.ZERO;
        this.leftDown = BigInteger.ZERO;
        this.rightDown = BigInteger.ONE;
    }
    
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        long k = in.nextLong();
        if (k < 3) {
            System.out.println("1");
            return;
        }
        k -= 2;
        String str = Long.toBinaryString(k);
        FastFib[] ST = new FastFib[str.length()];
        ST[0] = new FastFib();
        for (int i = 1; i < str.length(); i++) {
            ST[i] = new FastFib();
            ST[i].exp(ST[i-1]);
        }
        FastFib fibN = new FastFib();
        fibN.identityMatrix();
        for (int i = 0; i < str.length(); i++) {
            if (str.charAt(str.length()-i-1) == '1') {
                fibN.mul(ST[i]);
            }
        }
        System.out.println(fibN.leftUp.add(fibN.rightUp));
    }
}