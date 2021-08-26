class Solution:
    def reverseString(self, s: List[str]) -> None:
        temp = s[::-1]
        # for some reason, this won't let me do `s = temp`
        # or `s = temp[:]`
        s.clear()
        for i in temp:
            s.append(i)
        