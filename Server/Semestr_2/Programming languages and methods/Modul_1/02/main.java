import java.util.Scanner;
import java.util.Arrays;

public class MaxNum implements Comparable<MaxNum> {
    private String ch;
    
    public MaxNum(String ch) {
        this.ch = ch;
    }
    @Override
    public String toString() { return this.ch; }
    
    public static void main (String[] args) {
        Scanner in = new Scanner(System.in);
        int k = in.nextInt();
        MaxNum[] arr = new MaxNum[k];
        for (int i = 0; i < k; i++) {
            arr[i] = new MaxNum(in.next());
        }
        Arrays.sort(arr);
        for (int i = 0; i < k; i++) {
            System.out.print(arr[i]);
        }
    }
    @Override
    public int compareTo(MaxNum obj) {
        String S1 = obj.ch + this.ch;
        String S2 = this.ch + obj.ch;
        return S1.compareTo(S2);
    }
}
