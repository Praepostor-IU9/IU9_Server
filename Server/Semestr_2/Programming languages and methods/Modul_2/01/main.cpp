#include "textstats.hpp"
#include <iostream>

void get_tokens(const string &s, const unordered_set<char> &delimiters, vector<string> &tokens) {
    int start = 0;
    bool flagStr = false;
    string myStr;
    for(int i = 0; i < s.length(); i++) {
        myStr += tolower(s[i]);
        if (flagStr) {
            if (delimiters.find(s[i]) != delimiters.end()) {

                tokens.push_back(myStr.substr(start, i-start));
                flagStr = false;
            }
        } else if (delimiters.find(s[i]) == delimiters.end()) {
            start = i;
            flagStr = true;
        }
    }
    if (flagStr) {
        tokens.push_back(myStr.substr(start, s.length()-start));
    }
};

void get_type_freq(const vector<string> &tokens,  map<string, int> &freqdi) {
    for(auto str = tokens.begin(); str != tokens.end(); ++str) {
        freqdi[*str]++;
    }
};

void get_types(const vector<string> &tokens, vector<string> &wtypes) {
    map<string, int> help;
    get_type_freq(tokens, help);
    for(auto it = help.begin(); it != help.end(); ++it) {
        wtypes.push_back(it->first);
    }
};

void get_x_length_words(const vector<string> &wtypes, int x, vector<string> &words) {
    for(auto it = wtypes.begin(); it != wtypes.end(); ++it) {
        if ((*it).length() >= x) words.push_back(*it);
    }
};

void get_x_freq_words(const map<string, int> &freqdi, int x, vector<string> &words) {
    for(auto it = freqdi.begin(); it != freqdi.end(); ++it) {
        if (it->second >= x) {
            words.push_back(it->first);
        }
    }
};

void get_words_by_length_dict(const vector<string> &wtypes, map<int, vector<string> > &lengthdi){
    for(auto it = wtypes.begin(); it != wtypes.end(); ++it) {
        lengthdi[(*it).length()].push_back(*it);
    }
};