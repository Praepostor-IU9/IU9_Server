//
// Created by User on 31.05.2017.
//

#ifndef MODUL_3_3_SuperCalc_HPP
#define MODUL_3_3_SuperCalc_HPP

#include <iostream>
#include <vector>
#include <typeinfo>

using namespace std;

template <class T>
class Cell;

template <class T>
class Tocken {
public:
    Tocken(char c): flag(0), op(c) {};
    Tocken(const T &val): flag(1), val(val) {};
    Tocken(const Cell<T> &cell): flag(2), i(cell.i), j(cell.j) {};
private:
    friend class Cell<T>;

    int flag;
    T val;
    char op;
    Cell<T>* valCell;
    int i, j;
};

template <class T>
class Cell {
public:
    Cell(vector < Tocken<T> > v): Index(0), lexems(v), i(-1), j(-1), flag(false) {};

    Cell(int i, int j, vector< vector< Cell<T>* > >* v): Index(0), i(i), j(j), Table(v), flag(true) {}

    operator T() {
        T n = this->parseE1();
        Index = 0;
        return n;
    };

    void get() const {
        for(int i = 0; i < lexems.size(); i++) {
            if (lexems[i].flag == 0) {
                cout << lexems[i].op;
            } else if (lexems[i].flag == 1) {
                cout << lexems[i].val;
            } else if (lexems[i].flag == 2){
                cout << "CELL(" << lexems[i].i << ", " << lexems[i].j << ")";
            } else {
                cout << "nil";
            }
        }
        cout << endl;
    }

    Cell<T>& operator= (const Cell<T> &obj);
    Cell<T>& operator= (const T &val);

    Cell<T>& operator+= (const Cell<T> &obj);
    Cell<T>& operator+= (const T &val);

    Cell<T>& operator-= (const Cell<T> &obj);
    Cell<T>& operator-= (const T &val);

    Cell<T>& operator*= (const Cell<T> &obj);
    Cell<T>& operator*= (const T &val);

    Cell<T>& operator/= (const Cell<T> &obj);
    Cell<T>& operator/= (const T &val);

private:
    friend class Tocken<T>;

    vector< Tocken <T> > lexems;
    int Index;
    int i, j;
    bool flag;
    vector < vector < Cell<T> *> >* Table;

    T parseE1();
    T parseE2(T n);
    T parseT1();
    T parseT2(T n);
    T parseF();

    friend Cell<T> operator+(const Cell<T> &a, const Cell<T> &b) {
        vector < Tocken<T> > v;
        if (a.flag) {
            v.push_back(Tocken<T>(a));
        } else {
            v.push_back('(');
            v.insert(v.end(), a.lexems.begin(), a.lexems.end());
            v.push_back(')');
        }
        v.push_back(Tocken<T>('+'));
        if (b.flag) {
            v.push_back(Tocken<T>(b));
        } else {
            v.push_back('(');
            v.insert(v.end(), b.lexems.begin(), b.lexems.end());
            v.push_back(')');
        }
        return Cell<T>(v);
    }
    friend Cell<T> operator+(const T &a, const Cell<T> &b) {
        vector < Tocken<T> > v;
        v.push_back(Tocken<T>(a));
        v.push_back(Tocken<T>('+'));
        if (b.flag) {
            v.push_back(Tocken<T>(b));
        } else {
            v.push_back('(');
            v.insert(v.end(), b.lexems.begin(), b.lexems.end());
            v.push_back(')');
        }
        return Cell<T>(v);
    }
    friend Cell<T> operator+(const Cell<T> &a, const T &b) {
        vector < Tocken<T> > v;
        if (a.flag) {
            v.push_back(Tocken<T>(a));
        } else {
            v.push_back('(');
            v.insert(v.end(), a.lexems.begin(), a.lexems.end());
            v.push_back(')');
        }
        v.push_back(Tocken<T>('+'));
        v.push_back(Tocken<T>(b));
        return Cell<T>(v);
    }

    friend Cell<T> operator-(const Cell<T> &a, const Cell<T> &b) {
        vector < Tocken<T> > v;
        if (a.flag) {
            v.push_back(Tocken<T>(a));
        } else {
            v.push_back('(');
            v.insert(v.end(), a.lexems.begin(), a.lexems.end());
            v.push_back(')');
        }
        v.push_back(Tocken<T>('-'));
        if (b.flag) {
            v.push_back(Tocken<T>(b));
        } else {
            v.push_back('(');
            v.insert(v.end(), b.lexems.begin(), b.lexems.end());
            v.push_back(')');
        }
        return Cell<T>(v);
    }
    friend Cell<T> operator-(const T &a, const Cell<T> &b) {
        vector < Tocken<T> > v;
        v.push_back(Tocken<T>(a));
        v.push_back(Tocken<T>('-'));
        if (b.flag) {
            v.push_back(Tocken<T>(b));
        } else {
            v.push_back('(');
            v.insert(v.end(), b.lexems.begin(), b.lexems.end());
            v.push_back(')');
        }
        return Cell<T>(v);
    }
    friend Cell<T> operator-(const Cell<T> &a, const T &b) {
        vector < Tocken<T> > v;
        if (a.flag) {
            v.push_back(Tocken<T>(a));
        } else {
            v.insert(v.end(), a.lexems.begin(), a.lexems.end());
        }
        v.push_back(Tocken<T>('-'));
        v.push_back(Tocken<T>(b));
        return Cell<T>(v);
    }

    friend Cell<T> operator*(const Cell<T> &a, const Cell<T> &b) {
        vector < Tocken<T> > v;
        if (a.flag) {
            v.push_back(Tocken<T>(a));
        } else {
            v.push_back('(');
            v.insert(v.end(), a.lexems.begin(), a.lexems.end());
            v.push_back(')');
        }
        v.push_back(Tocken<T>('*'));
        if (b.flag) {
            v.push_back(Tocken<T>(b));
        } else {
            v.push_back('(');
            v.insert(v.end(), b.lexems.begin(), b.lexems.end());
            v.push_back(')');
        }
        return Cell<T>(v);
    }
    friend Cell<T> operator*(const T &a, const Cell<T> &b) {
        vector < Tocken<T> > v;
        v.push_back(Tocken<T>(a));
        v.push_back(Tocken<T>('*'));
        if (b.flag) {
            v.push_back(Tocken<T>(b));
        } else {
            v.push_back('(');
            v.insert(v.end(), b.lexems.begin(), b.lexems.end());
            v.push_back(')');
        }
        return Cell<T>(v);
    }
    friend Cell<T> operator*(const Cell<T> &a, const T &b) {
        vector < Tocken<T> > v;
        if (a.flag) {
            v.push_back(Tocken<T>(a));
        } else {
            v.push_back('(');
            v.insert(v.end(), a.lexems.begin(), a.lexems.end());
            v.push_back(')');
        }
        v.push_back(Tocken<T>('*'));
        v.push_back(Tocken<T>(b));
        return Cell<T>(v);
    }

    friend Cell<T> operator/(const Cell<T> &a, const Cell<T> &b) {
        vector < Tocken<T> > v;
        if (a.flag) {
            v.push_back(Tocken<T>(a));
        } else {
            v.push_back('(');
            v.insert(v.end(), a.lexems.begin(), a.lexems.end());
            v.push_back(')');
        }
        v.push_back(Tocken<T>('/'));
        if (b.flag) {
            v.push_back(Tocken<T>(b));
        } else {
            v.push_back('(');
            v.insert(v.end(), b.lexems.begin(), b.lexems.end());
            v.push_back(')');
        }
        return Cell<T>(v);
    }
    friend Cell<T> operator/(const T &a, const Cell<T> &b) {
        vector < Tocken<T> > v;
        v.push_back(Tocken<T>(a));
        v.push_back(Tocken<T>('/'));
        if (b.flag) {
            v.push_back(Tocken<T>(b));
        } else {
            v.push_back('(');
            v.insert(v.end(), b.lexems.begin(), b.lexems.end());
            v.push_back(')');
        }
        return Cell<T>(v);
    }
    friend Cell<T> operator/(const Cell<T> &a, const T &b) {
        vector < Tocken<T> > v;
        if (a.flag) {
            v.push_back(Tocken<T>(a));
        } else {
            v.push_back('(');
            v.insert(v.end(), a.lexems.begin(), a.lexems.end());
            v.push_back(')');
        }
        v.push_back(Tocken<T>('/'));
        v.push_back(Tocken<T>(b));
        return Cell<T>(v);
    }

    friend Cell<T> operator-(const Cell<T> &a) {
        vector< Tocken<T> > v;
        v.push_back(Tocken<T>('-'));
        if (a.flag) {
            v.push_back(Tocken<T>(a));
        } else {
            v.push_back('(');
            v.insert(v.end(), a.lexems.begin(), a.lexems.end());
            v.push_back(')');
        }
        return Cell<T>(v);
    }
};

template <class T>
T Cell<T>::parseE1() {
    T n;
    n = parseT1();
    n = parseE2(n);
    return n;
}
template <class T>
T Cell<T>::parseE2(T n) {
    if (Index == lexems.size()) {
        return n;
    }
    if (lexems[Index].flag == 0 && lexems[Index].op == '+') {
        Index++;
        n += parseT1();
        n = parseE2(n);
    } else if (lexems[Index].flag == 0 && lexems[Index].op == '-') {
        Index++;
        n -= parseT1();
        n = parseE2(n);
    }
    return n;
}
template <class T>
T Cell<T>::parseT1() {
    T n;
    n = parseF();
    n = parseT2(n);
    return n;
}
template <class T>
T Cell<T>::parseT2(T n) {
    if (Index == lexems.size()) {
        return n;
    }
    if (lexems[Index].flag == 0 && lexems[Index].op == '*') {
        Index++;
        n *= parseF();
        n = parseT2(n);
    } else if (lexems[Index].flag == 0 && lexems[Index].op == '/') {
        Index++;
        n /= parseF();
        n = parseT2(n);
    }
    return n;
}
template <class T>
T Cell<T>::parseF() {
    T val;
    if (Index != lexems.size()) {
        if (lexems[Index].flag == 1) {
            val = lexems[Index].val;
            Index++;
        } else if (lexems[Index].flag == 2) {
            val = (T)*(*Table)[lexems[Index].i][lexems[Index].j];
            Index++;
        } else if (lexems[Index].flag == 0 && lexems[Index].op == '-') {
            Index++;
            val = -parseF();
        }  else if (lexems[Index].flag == 0 && lexems[Index].op == '(') {
            Index++;
            val = parseE1();
            Index++;
        }
    }
    return val;
}

template <class T>
Cell<T>& Cell<T>::operator=(const Cell<T> &cell) {
    if (cell.flag) {
        lexems.clear();
        lexems.push_back(Tocken<T>(cell));
        return *this;
    }
    lexems = cell.lexems;
    return *this;
}
template <class T>
Cell<T>& Cell<T>::operator=(const T &val) {
    lexems.clear();
    lexems.push_back(Tocken<T>(val));
    return *this;
}
template <class T>
Cell<T>& Cell<T>::operator+=(const Cell<T> &obj) {
    lexems.push_back(Tocken<T>('+'));
    lexems.push_back(Tocken<T>(obj));
    return *this;
}
template <class T>
Cell<T>& Cell<T>::operator+=(const T &val) {
    lexems.push_back(Tocken<T>('+'));
    lexems.push_back(Tocken<T>(val));
    return *this;
}
template <class T>
Cell<T>& Cell<T>::operator-=(const Cell<T> &obj) {
    lexems.push_back(Tocken<T>('-'));
    lexems.push_back(Tocken<T>(obj));
    return *this;
}
template <class T>
Cell<T>& Cell<T>::operator-=(const T &val) {
    lexems.push_back(Tocken<T>('-'));
    lexems.push_back(Tocken<T>(val));
    return *this;
}
template <class T>
Cell<T>& Cell<T>::operator*=(const Cell<T> &obj) {
    lexems.push_back(Tocken<T>('*'));
    lexems.push_back(Tocken<T>(obj));
    return *this;
}
template <class T>
Cell<T>& Cell<T>::operator*=(const T &val) {
    lexems.push_back(Tocken<T>('*'));
    lexems.push_back(Tocken<T>(val));
    return *this;
}
template <class T>
Cell<T>& Cell<T>::operator/=(const Cell<T> &obj) {
    lexems.push_back(Tocken<T>('/'));
    lexems.push_back(Tocken<T>(obj));
    return *this;
}
template <class T>
Cell<T>& Cell<T>::operator/=(const T &val) {
    lexems.push_back(Tocken<T>('/'));
    lexems.push_back(Tocken<T>(val));
    return *this;
}
template <class T>
class SuperCalc {
public:
    ~SuperCalc();
    SuperCalc(int m, int n);
    Cell<T>& operator() (int i, int j);
private:
    vector< vector< Cell<T>* > > Table;
    int m, n;
};

template <class T>
SuperCalc<T>::~SuperCalc() {
    for(int i = 0; i < m; i++) {
        for(int j = 0; j < n; j++) {
            delete Table[i][j];
        }
    }
}
template <class T>
SuperCalc<T>::SuperCalc(int m, int n) {
    Table.assign(m, vector< Cell<T>* >(n));
    this->m = m;
    this->n = n;
    for(int i = 0; i < m; i++) {
        for(int j = 0; j < n; j++) {
            Table[i][j] = new Cell<T>(i, j, &Table);
        }
    }
}
template <class T>
Cell<T>& SuperCalc<T>::operator()(int i, int j) {
    return *Table[i][j];
}
#endif //MODUL_3_3_SuperCalc_HPP
