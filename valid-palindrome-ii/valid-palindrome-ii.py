class Solution:
    def helper(self,s,start,end):
        while start<=end:
            if s[start]!=s[end]:
                return False
            start+=1
            end-=1
        return True
    def validPalindrome(self, s: str) -> bool:
        low=0
        high=len(s)-1
        while low<=high:
            if s[low]!=s[high]:
                return (self.helper(s,low+1,high)) or (self.helper(s,low,high-1))
            low+=1
            high-=1
        return True