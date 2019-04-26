import java.util.Scanner;
import java.util.ArrayList;

public class Econom {
    private final String name;
    private char nameOp;
    private String nameLeft, nameRight;
    private Econom root, left, right;
    private ArrayList operands;
    private int sum;
    
    public Econom(String str) {
        this.name = str;
        if (str.length() >= 5) {
            this.nameOp = str.charAt(1);
        }
        this.nameLeft = this.nameRight = null;
        this.root = this.left = this.right = null;
        this.operands = null;
        this.sum = 0;
    }
    
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        String s = in.nextLine();
        Econom root = new Econom(s);
        root.root = root;
        root.operands = new ArrayList();
        root.growth();
        System.out.println(root.sum);
    }
    
    private void growth() {
        if (this.name.length() < 5) {
            return;
        }
        if (!this.root.operands.contains(this.name)) {
            this.root.operands.add(this.name);
            this.root.sum++;
        } else {
            return;
        }
        int i = 2, n = this.name.length()-1;
        boolean flagLeft = false, flagRight = false;
        if (this.name.charAt(i) != '(') {
            this.nameLeft = this.name.substring(i, i+1);
            i++;
        } else {
            flagLeft = true;
        }
        if (this.name.charAt(n-1) != ')') {
            this.nameRight = this.name.substring(n-1, n);
            n--;
        } else {
            flagRight = true;
        }
        if (flagLeft) {
            int u = 0;
            for (; i < n; i++) {
                if (this.name.charAt(i) == '(') {
                    u++;
                }
                if (this.name.charAt(i) == ')') {
                    u--;
                }
                if (u == 0) {
                    break;
                }
            }
            i++;
        }
        if (flagLeft) {
            String strLeft = this.name.substring(2, i);
            Econom left = new Econom(strLeft);
            this.left = left;
            left.root = this.root;
            left.growth();
        }
        if (flagRight) {
            String strRight = this.name.substring(i, n);
            Econom right = new Econom(strRight);
            this.right = right;
            right.root = this.root;
            right.growth();
        }
    }
}