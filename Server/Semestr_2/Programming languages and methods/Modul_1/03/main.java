import java.util.Scanner;

public class MinDist {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        String s = in.nextLine();
        char x = in.next().charAt(0), y = in.next().charAt(0);
        boolean flag_x = false, flag_y = false;
        int min = 1000001, p_x = 0, p_y = 0;
        for (int i = 0; i < s.length(); i++) {
            if (s.charAt(i) == x) {
                if (!flag_x && !flag_y) {
                    flag_x = true;
                    p_x = i;
                    continue;
                }
                if (!flag_x && flag_y) {
                    flag_x = true;
                    flag_y = false;
                    p_x = i;
                    if (p_x-p_y-1 < min) {
                        min = p_x-p_y-1;
                        if (min == 0) {
                            break;
                        }
                    }
                    continue;
                }
                if (flag_x && !flag_y) {
                    p_x = i;
                    continue;
                }
                if (flag_x && flag_y) {
                    System.out.println("Что-то пошло не так...");
                    System.out.println(1/0);
                }
            }
            if (s.charAt(i) == y) {
                if (!flag_y && !flag_x) {
                    flag_y = true;
                    p_y = i;
                    continue;
                }
                if (!flag_y && flag_x) {
                    flag_y = true;
                    flag_x = false;
                    p_y = i;
                    if (p_y-p_x-1 < min) {
                        min = p_y-p_x-1;
                        if (min == 0) {
                            break;
                        }
                    }
                    continue;
                }
                if (flag_y && !flag_x) {
                    p_y = i;
                    continue;
                }
                if (flag_y && flag_x) {
                    System.out.println("Что-то пошло не так...");
                    System.out.println(1/0);
                }
            }
        }
        if (min == 1000001) {
            System.out.println("Какой-то один символ из перечисленных не встречается в строке.");
        } else {
            System.out.println(min);
        }
    }
}