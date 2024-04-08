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
        


def main():
    solution = Solution()

    # Test Case 1
    numbers1 = [2, 7, 11, 15]
    target1 = 9
    print("Test Case 1:", solution.twoSum(numbers1, target1))  # Expected output: [0, 1]

    # Test Case 2
    numbers2 = [3, 2, 4]
    target2 = 6
    print("Test Case 2:", solution.twoSum(numbers2, target2))  # Expected output: [1, 2]

    # Test Case 3
    numbers3 = [3, 3]
    target3 = 6
    print("Test Case 3:", solution.twoSum(numbers3, target3))  # Expected output: [0, 1]

    # Test Case 4
    numbers4 = [-1, 0]
    target4 = -1
    print("Test Case 4:", solution.twoSum(numbers4, target4))  # Expected output: [0, 1]

if __name__ == "__main__":
    main()
