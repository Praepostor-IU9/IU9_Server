#include <iostream>
#include <fstream>
#include <vector>

using namespace std;
class dictionary;
class word {
public:
    word(string s, int n);
    word(string s);
    float measure(word other);
    void get() {
        for (int i = 0; i < bigram.size(); i++) {
            cout << bigram[i] << " ";
        }
        cout << endl;
    }
private:
    vector<int> bigram;
    string name;
    int frequency;
    friend class dictionary;
};

word::word(string s, int n) {
    name = s;
    frequency = n;
    if (s.length() == 1) {
        bigram.push_back(s[0]-'a'+1);
    } else {
        for (int i = 1; i < s.length(); i++) {
            bigram.push_back(((s[i-1]-'a'+1)<<8)+(s[i]-'a'+1));
        }
    }
}

word::word(string s) {
    name = s;
    frequency = 0;
    if (s.length() == 1) {
        bigram.push_back(s[0]-'a'+1);
    } else {
        for (int i = 1; i < s.length(); i++) {
            bigram.push_back(((s[i-1]-'a'+1)<<8)+(s[i]-'a'+1));
        }
    }
}

float word::measure(word other) {
    vector<bool> help;
    help.assign(this->bigram.size(), true);
    int Intersection = 0;
    for (int i = 0; i < other.bigram.size(); i++) {
        for (int j = 0; j < this->bigram.size(); j++) {
            if (help[j] && other.bigram[i] == this->bigram[j]) {
                Intersection++;
                help[j] = false;
                break;
            }
        }
    }
    int Union = other.bigram.size() + this->bigram.size() - Intersection;
    return (float) Intersection/Union;
}

class dictionary {
public:
    dictionary();
    void push_back(word a);
    string correction(word b);

private:
    vector<word> book;
};

dictionary::dictionary() {}

void dictionary::push_back(word a) {
    book.push_back(a);
}

string dictionary::correction(word b) {
    vector<string> total;
    int maxFrequency = 0;
    float help = 0, maxMeasure = 0;
    for(int i = 0; i < book.size(); i++) {
        help = b.measure(book[i]);
        if (help == 1) {
            return book[i].name;
        }
        if (help > maxMeasure) {
            maxMeasure = help;
            maxFrequency = book[i].frequency;
            total.clear();
            total.push_back(book[i].name);
        } else if (help == maxMeasure) {
            if (book[i].frequency > maxFrequency) {
                maxFrequency = book[i].frequency;
                total.clear();
                total.push_back(book[i].name);
            } else if (book[i].frequency == maxFrequency) {
                total.push_back(book[i].name);
            }
        }
    }
    string minStr = total[0];
    for (int i = 1; i < total.size(); i++) {
        if (minStr > total[i]) {
            minStr = total[i];
        }
    }
    return  minStr;
}

int main(){
    ifstream in("count_big.txt");
    int a;
    string s;
    dictionary dic;
    while (!in.eof()) {
        in >> s >> a;
        dic.push_back(word(s, a));
    }
    in.close();
    vector<word> input;
    while (getline(cin, s)) {
        input.push_back(word(s));
    }
    for (int i = 0; i < input.size(); i++) {
        cout << dic.correction(input[i]) << endl;
    }
    return 0;
}