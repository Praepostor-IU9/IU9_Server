import java.util.Scanner;

public class Kth {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        long k = in.nextLong();
        long i = 0, u = 0, n = k, a, b;
        while (n-i >= 0)
        {
            u++;
            n -= i;
            i = (long)Math.pow(10, u-1)*u*9;
        }
        a = n / u + (long)Math.pow(10, u-1);
        b = u - (n % u) - 1;
        System.out.println((long)(a / Math.pow(10, b)) % 10);
    }
}