class Solution {
public:
    // unordered_map<char,string> table;
    string table[10] = {
        "",
        "",
        "abc",
        "def",
        "ghi",
        "jkl",
        "mno",
        "pqrs",
        "tuv",
        "wxyz",
    };
    vector<string> ret;
    Solution(){
        // table['2']="abc";
        // table['3']="def";
        // table['4']="ghi";
        // table['5']="jkl";
        // table['6']="mno";
        // table['7']="pqrs";
        // table['8']="tuv";
        // table['9']="wxyz";
    }
    void dfs(const char *digits){
        // 递归深搜方法
        if (!*digits){
            return;
        }
        auto temp = ret;
        ret.clear();
        for(auto &c :table[*digits]){
            if(!c){continue;}
            if(!temp.size()){
                ret.push_back(string(1,c));
            }
            for(auto &s:temp ){
                auto ns = s+c;
                cout<<ns<<endl;
                ret.push_back(ns);
            }
        }
        if(*(++digits))
            dfs(digits);
    }
    vector<string> letterCombinations(string digits) {
        // const char* digits_c = digits.c_str();
        // 遍历组合，看上面的深搜算法
        for(auto& digit :digits){
            auto temp = ret;
            ret.resize(0);
            for(auto &c :table[digit-'0']){
                if(!temp.size()){
                    ret.push_back(string(1,c));
                    continue;
                }
                for(auto &s:temp ){
                    auto ns = s+c;
                    ret.push_back(ns);
                }
            }
        }
        return ret;
    }
};
