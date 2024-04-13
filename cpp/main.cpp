
/* C code below */
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <iostream>
#include <vector>

bool exists(int ints[], int count, int k)
{
    // Your code goes here    
    if (count == 0)
    {
        return false;
    }

    int head = 0;
    int mid = count / 2;
    std::cout << "mid"<< mid << std::endl;
    int start = 0;

    if (k > ints[mid])
    {
        start = mid;
    }

    for (int i = start; i < count; i++)
    {
        if (ints[i] == k)
        {
            return true;
        }
    }
    return false;
}

int compute_closest_to_zero(int n, int ts[])
{
    // Write your code here
    // To debug: fprintf(stderr, "Debug messages...\n");'
    if (n == 0)
    {
        return 0;
    }

    int lowest = ts[0];
    int abs_lowest = ::abs(lowest);
    for (int i = 0; i < n; i++)
    {
        if (::abs(ts[i]) < abs_lowest)
        {
            lowest = ts[i];
            abs_lowest = ::abs(lowest);
        }
        else if (::abs(ts[i]) == abs_lowest)
        {
            if (ts[i] > lowest)
            {
                lowest = ts[i];
            }
        }

    }
    return lowest;
}

int majorityElement(int* nums, int numsSize) {
    // max size is 50k
    int map[10000] = {0};
    int maj_num = numsSize / 2;
    for (int i = 0; i < numsSize; i++)
    {
        map[nums[i]]++;
    }
    // because the majority_elementy will always be found, map necessarily 
    //  be half the size of nums
    for (int i = 0; i < 10000; i++)
    {
        if (map[i] > maj_num)
        {
            
            return i;
        }
    }
    return 0;
}
int characterReplacement(std::string s, int k) {
    int n = s.length();
    int left = 0, right = 0;
    
    std::vector<int> freq(26, 0);  // stores char frequency
    int maxFreq = 0;
    int maxLen = 0;

    while(right < n) {
        freq[s[right]]++;  // increment frequency
        maxFreq = std::max(maxFreq, freq[s[right]]);  // update maxFreq
        
        // when changes required is > k
        if(( (right-left+1) - maxFreq ) > k)
        {
            freq[s[left]]--;
            maxFreq = 0;
            for(int i=0; i<26; i++)
            {
                maxFreq = std::max(maxFreq, freq[i]);
            }
            left++;
        }

        if( (right-left+1) - maxFreq  <=  k)
        {
            maxLen = std::max(maxLen, right-left+1);
        }

        right++;
    }

    return maxLen;
}

int main(void)
{
    int arr1[1] = {-273};
    int arr2[6] = {-15, -7, -9, -14, -12, -273};
    int arr3[2] = {-10, -10};
    int arr4[6] = {15,-7,9,14,7,12 };
    std::cout << compute_closest_to_zero(1, arr1) << std::endl;
    std::cout << compute_closest_to_zero(6, arr2) << std::endl;
    std::cout << compute_closest_to_zero(2, arr3) << std::endl;
    std::cout << compute_closest_to_zero(6, arr4) << std::endl;

    //std::cout << exists(arr4, 6, -7) << std::endl;
    //std::cout << exists(arr4, 6, 273) << std::endl;
    //std::cout << exists(arr4, 6, -7) << std::endl;
    std::cout << exists(arr3, 2, -10) << std::endl;
    //std::cout << exists(arr4, 6, 10) << std::endl;
    //std::cout << exists(arr4, 6, 15) << std::endl;
    //
    int arr5[3] = {3, 2, 3};
    std::cout << majorityElement(arr5, 3);
    int arr6[7] = {2,2,1,1,1,2,2};
    std::cout << majorityElement(arr6, 7);
    std::vector<int> vec(2);
    vec.push_back(1);
    vec.push_back(1);
    vec.push_back(1);
    std::cout <<"\n" << vec.size() << std::endl;
}
