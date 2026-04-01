# Go Data Structures & Algorithms (LeetCode Practice)

![Go](https://img.shields.io/badge/Language-Go-00ADD8?style=flat&logo=go&logoColor=white)
![Topics](https://img.shields.io/badge/Topics-Data_Structures_%7C_Algorithms-blue)
![Focus](https://img.shields.io/badge/Focus-Interview_Preparation-success)

A comprehensive, growing collection of Data Structures and Algorithms implementations, focusing on LeetCode problem solving and technical interview preparation using **Go (Golang)**.

## 🎯 Objectives

- Practice with real-world LeetCode problems.
- Master core Data Structures and Algorithms.
- Become proficient in the 14 essential competitive programming and interview patterns.
- Write clean, idiomatic Go code following industry best practices.

## 📚 Repository Structure

```text
algorithm/
├── leetcode/              # Problem solutions categorized by topic
│   ├── arrays/           # Array and string manipulation
│   ├── linkedlist/       # Linked List operations
│   ├── stack/            # Stack-based problems
│   ├── queue/            # Queue and Deque problems
│   └── hashtable/        # Hash map and set implementations
└── 01-foundations/       # Learning materials, theory, and core concepts
    ├── 01-complexity-analysis/
    └── 02-basic-data-structures/
```

## 🚀 Getting Started

### 1. Study the Theory

If you are new to a specific topic, you can begin by reviewing the foundational theories and concepts:

```bash
# Read about Arrays & Slices
cat 01-foundations/02-basic-data-structures/README.md

# Run basic examples
go run 01-foundations/02-basic-data-structures/examples.go
```

### 2. Solve LeetCode Problems

Once you're familiar with the theory, dive directly into the problems:

```bash
# Example: Running the Two Sum solution
cd leetcode/arrays/01-two-sum
go run problem.go
```

## 📖 Learning Workflow

1. **Study Theory** - Read the `README.md` and explore provided code examples.
2. **Knowledge Check** - Validate understanding with self-assessment questions.
3. **Solve Problems** - Actively code solutions for the problems within the current topic.
4. **Code Review & Optimization** - Review edge cases, optimize algorithms, and discuss time/space complexities.

## 🏃 Execution Guide

```bash
# Example: Running problems from the Arrays module

# Problem 1: Two Sum
cd leetcode/arrays/01-two-sum
go run problem.go

# Problem 2: Contains Duplicate
cd ../02-contains-duplicate
go run problem.go
```

## 📋 Problem Format

Each problem directory mirrors the LeetCode environment and includes:

- ✅ Full problem statement and constraints.
- ✅ Concrete test cases with clear input and output targets.
- ✅ Difficulty tags (Easy / Medium / Hard).
- ✅ Direct LeetCode URL references.
- ✅ Executable `main()` function with test verifications.

## 🗺️ NeetCode Roadmap Progress

Progress tracking based on the [NeetCode Roadmap](https://neetcode.io/roadmap).

### 1. Arrays & Hashing

- [x] Contains Duplicate (Easy)
- [x] Valid Anagram (Easy)
- [x] Two Sum (Easy)
- [x] Group Anagrams (Medium)
- [x] Top K Frequent Elements (Medium)
- [x] Encode and Decode Strings (Medium)
- [x] Product of Array Except Self (Medium)
- [x] Valid Sudoku (Medium)
- [ ] Longest Consecutive Sequence (Medium)

_Bonus / Extra Practice:_

- [/] Subarray Sum Equals K (Medium - Prefix Sum)

### 2. Two Pointers

- [x] Valid Palindrome (Easy)
- [ ] Two Sum II - Input Array Is Sorted (Medium)
- [ ] 3Sum (Medium)
- [ ] Container With Most Water (Medium)
- [ ] Trapping Rain Water (Hard)

### 3. Stack

- [ ] Valid Parentheses (Easy)
- [ ] Min Stack (Medium)
- [ ] Evaluate Reverse Polish Notation (Medium)
- [ ] Daily Temperatures (Medium)
- [ ] Car Fleet (Medium)

### 4. Sliding Window

- [x] Best Time to Buy and Sell Stock (Easy)
- [ ] Longest Substring Without Repeating Characters (Medium)
- [ ] Longest Repeating Character Replacement (Medium)
- [ ] Minimum Window Substring (Hard)

### 5. Binary Search

- [ ] Binary Search (Easy)
- [ ] Search a 2D Matrix (Medium)
- [ ] Koko Eating Bananas (Medium)
- [ ] Find Minimum in Rotated Sorted Array (Medium)
- [ ] Search in Rotated Sorted Array (Medium)

_(More topics will be updated as progress continues...)_

## 🎓 Recommended Resources

- [LeetCode](https://leetcode.com/) - Primary practice platform.
- [NeetCode 150](https://neetcode.io/) - A heavily curated list of essential coding interview problems.
- [Grokking the Coding Interview](https://www.educative.io/courses/grokking-the-coding-interview) - Pattern-based problem solving approach.
- [Cracking the Coding Interview](https://www.crackingthecodinginterview.com/) - The classic technical interview handbook.

---

**Happy Coding! 🚀**
