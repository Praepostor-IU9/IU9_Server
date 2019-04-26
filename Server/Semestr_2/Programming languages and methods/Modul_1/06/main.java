import java.util.*;

public class Gauss {
    private final Elem[] line;
    @Override
    public String toString() {
        return Arrays.toString(line);
    }
    public static class Elem {
        private int a;
        private int b;
        private int gcd(int a, int b) {
            return b == 0 ? Math.abs(a) : gcd(b, a%b);
        }
        private void minus(Elem val) {
            int A, B, C;
            A = this.a*val.b - val.a*this.b;
            B = this.b*val.b;
            C = gcd(A, B);
            this.a = A / C;
            this.b = B / C;
        }
        private void mul(Elem val) {
            int A, B, C;
            A = this.a*val.a;
            B = this.b*val.b;
            C = gcd(A, B);
            this.a = A / C;
            this.b = B / C;
        }
        private void inverseElem(Elem val) {
            this.a = val.b*(val.a > 0? 1 : -1);
            this.b = Math.abs(val.a)+(val.a == 0? 1 : 0);
            
        }
        private void copy(Elem val) {
            this.a = val.a;
            this.b = val.b;
        }
        private boolean zero() {
            return this.a == 0;
        }
        public Elem(int n) {
            this.a = n;
            this.b = 1;
        }
        @Override
        public String toString() {
            return this.a+"/"+this.b;
        }
    }
    public Gauss(int[] array, int n) {
        line = new Elem[n+1];
        for(int i = 0; i <= n; i++) {
            line[i] = new Elem(array[i]);
        }
    }
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        int[][] matrix = new int[n][n+1];
        Gauss[] arr = new Gauss[n];
        int i, j, u;
        for(i = 0; i < n; i++) {
            for(j = 0; j <= n; j++) {
                matrix[i][j] = in.nextInt();
            }
            arr[i] = new Gauss(matrix[i], n);
        }
        Elem k = new Elem(1);
        Elem t = new Elem(1);
        Gauss v = new Gauss(matrix[0], n);
        for(i = 0; i < n; i++) {
            if (arr[i].line[i].zero()) {
                for(j = i+1; j < n && arr[j].line[i].zero(); j++) {}
                if (j == n) {
                    System.out.print("No solution");
                    return;
                }
                v = arr[i];
                arr[i] = arr[j];
                arr[j] = v;
            }
            k.inverseElem(arr[i].line[i]);
            for(j = i; j <= n; j++) { //Приведение строки к виду 0...01...
                arr[i].line[j].mul(k);
            }
            for(j = i+1; j < n; j++) { //Вычетание строки из последующих
                k.copy(arr[j].line[i]);
                for(u = i; u <= n; u++) {
                    t.copy(arr[i].line[u]);
                    t.mul(k);
                    arr[j].line[u].minus(t);
                }
            }
        }
        t = new Elem(0);
        for(i = n-1; i >= 0; i--) { //Приведение матрицы к еденичной
            for(j = i-1; j >= 0; j--) {
                k.copy(arr[i].line[n]);
                k.mul(arr[j].line[i]);
                arr[j].line[i].mul(t);
                arr[j].line[n].minus(k);
            }
        }
        for(i = 0; i < n; i++) {
            System.out.println(arr[i].line[n]);
        }
    }
}