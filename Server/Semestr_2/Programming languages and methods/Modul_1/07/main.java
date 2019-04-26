/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
import java.util.*;
import java.util.stream.*;
/**
 *
 * @author User
 */
enum Tag {
    NUMBER,
    VAR,
    LPAREN,
    RPAREN,
    OP,
    END;
}
class Token {
    public Tag tag;
    public String name;
    public int value;
    
    public Token(Tag tag, String name, int value) {
        this.tag = tag;
        this.name = name;
        this.value = value;
    }
}
class Lexer {
    public String text;
    public ArrayList<Token> array;
    public List<String> nameVar;
    public List<Integer> valVar;
    int INDEX;
    
    public Lexer(String text) {
        this.text = text;
        this.array = new ArrayList<>();
        this.INDEX = 0;
        this.tokenizing();
    }
    public void updateVar(List<Integer> valVar) {
        this.valVar = valVar;
    }
    
    private boolean separationToken(char c) {
        return c == ' ' || c == '+' || c == '-' || c == '/' || c == '*' || c == '(' || c == ')';
    }
    
    private void tokenizing() {
        this.nameVar = new ArrayList<>();
        int start = 0, i, countLPAREN = 0, countRPAREN = 0;
        boolean f_num = false, f_var = false;
        char c;
        for (i = 0; i < this.text.length(); i++) {
            c = this.text.charAt(i);
            if (f_var) {
                if (c >= 'a' && c <= 'z' ||
                        c >= 'A' && c <= 'Z' ||
                        c >= '0' && c <= '9') {
                    continue;
                }
                
                if (separationToken(c)) {
                    this.nameVar.add(this.text.substring(start, i));
                    this.array.add(new Token(Tag.VAR, this.text.substring(start, i), 0));
                    f_var = false;
                    //continue;
                } else {
                //ОШИБКА
                System.out.println("error");
                System.exit(0);
                }
            }
            if (f_num) {
                if (c >= '0' && c <= '9') {
                    continue;
                }
                if (separationToken(c)) {
                    this.array.add(new Token(Tag.NUMBER, null, Integer.parseInt(this.text.substring(start, i))));
                    f_num = false;
                    //continue;
                } else {
                //ОШИБКА
                System.out.println("error");
                System.exit(0);
                }
            }
            if (!f_num && !f_var) {
                if (c >= 'a' && c <= 'z' ||
                        c >= 'A' && c <= 'Z') {
                    start = i;
                    f_var = true;
                    continue;
                }
                if (c >= '0' && c <= '9') {
                    start = i;
                    f_num = true;
                    continue;
                }
                if (c == '+' || c == '-' || c == '/' || c == '*') {
                    this.array.add(new Token(Tag.OP, null, c));
                    continue;
                }
                switch (c){
                    case '(':
                        this.array.add(new Token(Tag.LPAREN, null, c));
                        countLPAREN++;
                        break;
                    case ')':
                        this.array.add(new Token(Tag.RPAREN, null, c));
                        countRPAREN++;
                        break;
                    case ' ':
                        break;
                    default:
                        //ОШИБКА
                        System.out.println("error");
                        System.exit(0);
                        break;
                }
            }
        }
        if (f_num) {
            this.array.add(new Token(Tag.NUMBER, null, Integer.parseInt(this.text.substring(start, i))));
        }
        if (f_var) {
            this.nameVar.add(this.text.substring(start, i));
            this.array.add(new Token(Tag.VAR, this.text.substring(start, i), 0));
        }
        this.array.add(new Token(Tag.END, null, 3));
        
        this.nameVar = this.nameVar.stream().distinct().collect(Collectors.toList());
        
        if (countLPAREN != countRPAREN) {
            //ОШИБКА
            System.out.println("error");
            System.exit(0);
        }
    }
    public int valOfVar(String str) {
        return valVar.get(this.nameVar.indexOf(str));
    }
    public Token next() {
        return this.array.get(this.INDEX++);
    }
}
//<E>  ::= <T> <E’>. 
//<E’> ::= + <T> <E’> | - <T> <E’> | . 
//<T>  ::= <F> <T’>. 
//<T’> ::= * <F> <T’> | / <F> <T’> | . 
//<F>  ::= <number> | <var> | ( <E> ) | - <F>.
public class Calc {
    private static Lexer lex;
    private static Token tok;
    private static void expect(Tag tag) {
        if (tok.tag != tag) {
            //ОШИБКА
            System.out.println("error");
            System.exit(0);
        } else {
            tok = lex.next();
        }
    }
    
    public static void main(String[] args) {
        int i;
        Scanner in = new Scanner(System.in);
        String text = in.nextLine();
        lex = new Lexer(text);
        List<Integer> valVar = new ArrayList<>();
        for (i = 0; i < lex.nameVar.size(); i++) {
            valVar.add(in.nextInt());
        }
        lex.updateVar(valVar);
        tok = lex.next();
        int OTVET = parseE1();
        System.out.println(OTVET);
    }
    private static int parseE1() {
        int n;
        n = parseT1();
        n = parseE2(n);
        return n;
    }
    private static int parseE2(int n) {
        if (tok.tag == Tag.OP && tok.value == '+') {
            tok = lex.next();
            n += parseT1();
            n = parseE2(n);
        } else if (tok.tag == Tag.OP && tok.value == '-') {
            tok = lex.next();
            n -= parseT1();
            n = parseE2(n);
        } else if (tok.tag == Tag.NUMBER || tok.tag == Tag.VAR) {
            //ОШИБКА
            System.out.println("error");
            System.exit(0);
        }
        return n;
    }
    private static int parseT1() {
        int n;
        n = parseF();
        n = parseT2(n);
        return n;
    }
    private static int parseT2(int n) {
        if (tok.tag == Tag.OP && tok.value == '*') {
            tok = lex.next();
            n *= parseF();
            n = parseT2(n);
        } else if (tok.tag == Tag.OP && tok.value == '/') {
            tok = lex.next();
            int m = parseF();
            if (m != 0) {
                n /= m;
            } else {
                //ОШИБКА
                System.out.println("error");
                System.exit(0);
            }
            n = parseT2(n);
        } else if (tok.tag == Tag.NUMBER || tok.tag == Tag.VAR) {
            //ОШИБКА
            System.out.println("error");
            System.exit(0);
        }
        return n;
    }
    private static int parseF() {
        int val = 0;
        if (null != tok.tag) switch (tok.tag) {
            case NUMBER:
            {
                val = tok.value;
                tok = lex.next();
                break;
            }
            case VAR:
            {
                val = lex.valOfVar(tok.name);
                tok = lex.next();
                break;
            }
            case LPAREN:
            {
                tok = lex.next();
                val = parseE1();
                expect(Tag.RPAREN);
                break;
            }
            case OP:
            {
                if (tok.value == '-') {
                    tok = lex.next();
                    val = (-1)*parseF();
                } else {
                    // ОШИБКА
                    System.out.println("error");
                    System.exit(0);
                }
                break;
            }
            default:
                // ОШИБКА
                System.out.println("error");
                System.exit(0);
                break;
        }
        return val;
    }
}
