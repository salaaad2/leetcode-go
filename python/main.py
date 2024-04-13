from typing import List

class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        num_indices = {}
        for i, num in enumerate(numbers):
            complement = target - num
            if complement in num_indices:
                return [num_indices[complement] + 1, i + 1]
            num_indices[num] = i
        return []

    def threeSum(self, numbers: List[int], target: int) -> List[int]:
        num_indices = {}
        for i, num in enumerate(numbers):
            complement = target - num
            if complement in num_indices:
                return [num_indices[complement], i]
            num_indices[num] = i
        return []
        
class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        buy_date: int
        sell_date: int
        buy_date = -1
        sell_date = -1
        sorted_prices = sorted(prices)
        p_len = len(prices)
        if (p_len == 1):
            return 0
        values_map = {}
        for i, value in enumerate(prices):
            values_map[value] = i
        l_i = 0
        h_i = p_len
        lowest = sorted_prices[l_i]
        highest = sorted_prices[-1]
        while l_i < h_i:
            print(f"lowest: {l_i}: {lowest}")
            print(f"highes: {h_i}: {highest}")
            if values_map[lowest] < values_map[highest]:
                return highest - lowest
            else:
                l_i += 1
                h_i -= 1
                lowest = sorted_prices[l_i + 1]
                highest = sorted_prices[h_i - 1]
        return 0
        

class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        charSet = set()
        l = 0
        res = 0

        for i in range(len(s)):
            while s[i] in charSet:
                charSet.remove(s[l])
                l += 1
            charSet.add(s[i])
            res = max(res, i - l + 1)
        return res

class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        freq_list = {}
        l = 0

        max_freq = 0
        for r in range(len(s)):
            freq_list[s[r]] = 1 + freq_list.get(s[r], 0)
            max_freq = max(max_freq, freq_list[s[r]])
            if (r - l + 1) > k:
                print("not enough iterations to make list. increase left id")
            r += 1

        return res

def main():
    solution = Solution()

    print("test case 1:", solution.characterReplacement("ABAB", 2))

if __name__ == "__main__":
    main()
