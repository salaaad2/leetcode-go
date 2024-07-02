struct Solution;

use std::collections::HashMap;
#[allow(unused)]
use std::collections::HashSet;

use std::time::Instant;

#[allow(unused)]
fn measure_time<F, T, R>(f: F, arg: &T) -> u128
where
    F: Fn(&T) -> R,
{
    let start = Instant::now();
    f(arg);
    let duration = start.elapsed();
    let ns = duration.as_nanos();
    ns
}

impl Solution 
{
#[allow(unused)]
    pub fn contains_duplicate(nums: &Vec<i32>) -> bool
    {
        for (i, value) in nums.iter().enumerate() 
        {
            for (j, v) in nums.iter().enumerate()
            {
                if j != i && v == value
                {
                    return true;
                }
            }
        }
        return false;
    }

#[allow(unused)]
    pub fn contains_duplicate_hashmap(nums: &Vec<i32>) -> bool
    {
        let mut map: HashMap<i32, bool> = HashMap::new();

        for i in nums
        {
            if map.contains_key(i) == false
            {
                map.insert(*i, true);
            }
            else
            {
                return true;
            }
        }
        false
    }

#[allow(unused)]
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut num_indices: HashMap<i32, i32> = HashMap::new();
        let mut complement: i32;

        for (i, val) in nums.iter().enumerate()
        {
            complement = target - val;
            if num_indices.contains_key(&complement)
            {
                return Vec::<i32>::from([num_indices[&complement], i as i32])
            }
            *num_indices.entry(*val).or_insert(0) = i as i32;
        }

        Vec::from([])
    }

    pub fn is_anagram(s: String, t: String) -> bool
    {
        if s.len() != t.len()
        {
            return false;
        }
        let mut map1: HashMap<char, i32> = HashMap::new();
        let mut map2: HashMap<char, i32> = HashMap::new();

        for c in s.chars().collect::<Vec<_>>()
        {
            *map1.entry(c).or_insert(0) += 1;
        }
        for c in t.chars().collect::<Vec<_>>()
        {
            *map2.entry(c).or_insert(0) += 1;
        }

        for (k, v) in map1
        {
            match map2.get(&k) {
                Some(item) => if *item != v { return false; },
                None => {return false;}
            }
        }
        true
    }

    pub fn group_anagrams(strs: Vec<String>) -> Vec<Vec<String>> {
        let mut output: Vec<Vec<String>> = Vec::new();

        for (i, value) in strs.iter().enumerate() 
        {
            for (j, v) in strs.iter().enumerate()
            {
                if j != i && Solution::is_anagram(v.to_string(), value.to_string())
                {
                    if let Some(inner) = output.get_mut(i)
                    {
                        inner.push(value.to_string());
                    }
                    else 
                    {
                        let mut inner = Vec::new();
                        inner.push(v.to_string());
                        inner.push(value.to_string());
                        output.insert(i, inner);
                    }
                }
            }
        }
        output
    }
}



fn main()
{
    // let nums1: Vec::<i32> = Vec::from([1,1,1,3,3,4,3,2,4,2]);
    // let nums2: Vec::<i32> = Vec::from([1,2,3,4]);
    // let nums3: Vec::<i32> = Vec::from([1,2,3,1]);
    // let nums4: Vec::<i32> = Vec::from([44, 57, 25, 32, 30, 21, 3, 86, 84, 23, 20, 85, 92, 79, 15, 71, 30, 81, 27, 73, 28, 26, 40, 14, 30, 18, 69, 86, 90, 61, 26, 82, 37, 27, 51, 43, 53, 34, 14, 23, 91, 12, 6, 92, 64, 35, 84, 70, 62, 31, 51, 24, 91, 9, 13, 98, 58, 77, 65, 64, 31, 95, 59, 14, 47, 89, 35, 53, 12, 61, 63, 90, 14, 71, 40, 99, 40, 17, 97, 6, 95, 60, 86, 57, 91, 72, 87, 91, 72, 26, 87, 59, 57, 27, 53, 99, 8, 38, 11, 83, 3, 12, 67, 93, 97, 68, 61, 16, 7, 55, 85, 97, 19, 39, 67, 87, 20, 35, 74, 4, 89, 73, 34, 91, 19, 52, 68, 71, 78, 42, 26, 35, 69, 27, 72, 20, 5, 38, 94, 60, 94, 20, 99, 54, 81, 53, 7, 6, 43, 39, 36, 26, 68, 97, 44, 31, 92, 36, 6, 21, 90, 39, 83, 4, 47, 28, 59, 45, 47, 96, 51, 28, 76, 98, 53, 64, 55, 39, 3, 1, 48, 23, 5, 53, 65, 92, 72, 74, 12, 94, 79, 43, 70, 10, 28, 86, 22, 57, 61, 13, 1, 64, 34, 72, 59, 26, 8, 2, 81, 77, 10, 33, 2, 42, 99, 63, 39, 58, 15, 87, 66, 11, 98, 73, 22, 1, 5, 72, 85, 69, 94, 27, 24, 3, 46, 32, 83, 35, 27, 10, 93, 42, 46, 81, 4, 67, 55, 27, 57, 15, 34, 39, 90, 41, 37, 34, 95, 80, 49, 53, 22, 75, 56, 6, 69, 16, 26, 76, 41, 26, 13, 35, 33, 21, 56, 60, 27, 50, 42, 11, 72, 54, 25, 69, 82, 99, 70, 39, 94, 63, 41, 85, 62, 74, 51, 40, 87, 56, 7, 67, 66, 22, 97, 84, 59, 91, 71, 86, 66, 9, 33, 68, 9, 85, 60, 96, 85, 98, 5, 92, 73, 25, 91, 3, 69, 95, 46, 79, 30, 10, 65, 21, 35, 16, 57, 60, 57, 9, 3, 39, 91, 3, 90, 78, 24, 71, 47, 11, 23, 49, 50, 28, 85, 38, 8, 51, 86, 23, 37, 33, 88, 4, 45, 58, 45, 43, 36, 51, 9, 84, 6, 70, 92, 14, 44, 26, 11, 14, 99, 22, 17, 48, 30, 26, 85, 27, 98, 92, 46, 92, 100, 29, 49, 11, 28, 58, 87, 68, 89, 74, 3, 64, 80, 77, 99, 80, 69, 74, 95, 62, 37, 86, 22, 85, 45, 15, 20, 8, 53, 28, 77, 56, 37, 63, 76, 81, 38, 82, 66, 7, 44, 66, 92, 26, 82, 80, 97, 41, 32, 17, 72, 53, 90, 23, 20, 44, 65, 91, 4, 88, 14, 48, 68, 95, 33, 28, 36, 86, 63, 68, 16, 32, 41, 65, 20, 71, 90, 88, 89, 76, 69, 73, 47, 84, 28, 5, 3, 12, 18, 24, 3, 29, 67, 7, 56, 11, 32, 40, 16, 55, 71, 40, 14, 84, 72, 63, 31, 53, 6, 64, 94, 74, 98, 37, 25, 65, 46, 3, 11, 44, 59, 29, 33, 97, 93, 74, 11, 93, 89, 52, 45, 86, 16, 63, 69, 71, 66, 64, 9, 92, 74, 16, 43, 32, 91, 54, 34, 20, 60, 93, 92, 14, 88, 17, 68, 30, 92, 9, 63, 60, 46, 93, 82, 55, 92, 34, 85, 42, 82, 76, 86, 68, 66, 54, 44, 61, 49, 82, 79, 37, 84, 26, 83, 85, 47, 21, 96, 14, 5, 2, 4, 15, 19, 9, 93, 58, 21, 4, 90, 12, 8, 83, 3, 63, 48, 32, 8, 85, 43, 22, 15, 42, 76, 46, 6, 16, 62, 76, 35, 65, 88, 41, 30, 21, 23, 67, 2, 100, 19, 94, 97, 24, 74, 20, 75, 57, 74, 95, 61, 100, 84, 40, 84, 33, 77, 29, 37, 81, 20, 88, 35, 49, 29, 55, 64, 62, 79, 74, 85, 49, 94, 4, 51, 75, 68, 15, 10, 55, 80, 52, 62, 93, 59, 1, 57, 97, 79, 75, 16, 13, 46, 40, 10, 11, 68, 47, 45, 32, 60, 44, 12, 59, 68, 89, 98, 71, 10, 57, 29, 92, 9, 7, 31, 16, 13, 92, 61, 97, 78, 17, 13, 80, 3, 58, 33, 77, 14, 86, 76, 59, 22, 99, 66, 47, 95, 21, 94, 64, 87, 91, 20, 14, 1, 81, 13, 12, 56, 74, 44, 14, 26, 97, 8, 80, 6, 46, 68, 9, 65, 71, 38, 85, 23, 42, 12, 92, 25, 26, 59, 24, 67, 99, 75, 46, 43, 26, 93, 74, 63, 7, 56, 80, 71, 60, 55, 74, 20, 31, 11, 46, 57, 98, 86, 74, 92, 23, 22, 75, 38, 71, 93, 16, 41, 22, 15, 28, 72, 78, 63, 77, 36, 62, 29, 33, 49, 77, 56, 39, 80, 22, 100, 21, 11, 59, 83, 6, 85, 42, 17, 92, 60, 69, 5, 17, 55, 26, 12, 8, 9, 37, 47, 30, 9, 12, 65, 98, 4, 1, 28, 71, 3, 80, 32, 9, 45, 100, 7, 87, 32, 78, 60, 64, 45, 37, 5, 32, 70, 15, 1, 98, 54, 26, 31, 37, 100, 28, 12, 67, 2, 58, 83, 96, 26, 42, 95, 5, 68, 60, 51, 3, 87, 15, 36, 40, 44, 38, 74, 38, 37, 32, 80, 69, 91, 60, 22, 67, 44, 6, 3, 40, 6, 68, 80, 70, 78, 92, 7, 23, 92, 49, 51, 46, 93, 15, 85, 24, 95, 12, 8, 91, 43, 41, 29, 84, 43, 74, 34, 63, 68, 95, 16, 98, 18, 92, 32, 89, 8, 2, 59, 68, 77, 42, 22, 55, 33, 86, 68, 40, 29, 76, 46, 62, 95, 34, 10, 9, 52, 69, 6, 79, 4, 47, 91, 27, 72, 31, 8, 6, 22, 31, 12, 34, 29, 50, 2, 73, 66, 67, 42, 62, 56, 52, 32, 47, 27, 2, 74, 25, 50, 77, 20, 3, 82, 66, 60, 61, 34, 3,8]);
    // println!("{}", Solution::contains_duplicate_hashmap(&&nums1));
    // println!("{}", Solution::contains_duplicate_hashmap(&&nums2));
    // println!("{}", Solution::contains_duplicate_hashmap(&&nums3));
    // println!("{}", Solution::contains_duplicate_hashmap(&&nums4));
    // let input1: String = String::from("anagram");
    // let input2: String = String::from("nagaram");
    // println!("{}", Solution::is_anagram(input1, input2));
    // let input3: String = String::from("rat");
    // let input4: String = String::from("car");
    // println!("{}", Solution::is_anagram(input3, input4));
    // let input5: Vec::<i32> = Vec::from([2,7,11,15]);
    // let input6: Vec::<i32> = Vec::from([3,2,4]);
    // let input7: Vec::<i32> = Vec::from([3,3]);

    // println!("{:?}", Solution::two_sum(input5, 9));
    // println!("{:?}", Solution::two_sum(input6, 6));
    // println!("{:?}", Solution::two_sum(input7, 6));
    let input1: Vec::<String> = vec!["eat".to_string(), "tea".to_string(), "tan".to_string(), "ate".to_string(), "nat".to_string(), "bat".to_string()];
    let input2: Vec::<String> = vec!["".to_string()];
    let input3: Vec::<String> = vec!["a".to_string()];
    println!("{:?}", Solution::group_anagrams(input1));
    println!("{:?}", Solution::group_anagrams(input2));
    println!("{:?}", Solution::group_anagrams(input3));
}
