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

## 🎓 Recommended Resources

- [LeetCode](https://leetcode.com/) - Primary practice platform.
- [NeetCode 150](https://neetcode.io/) - A heavily curated list of essential coding interview problems.
- [Grokking the Coding Interview](https://www.educative.io/courses/grokking-the-coding-interview) - Pattern-based problem solving approach.
- [Cracking the Coding Interview](https://www.crackingthecodinginterview.com/) - The classic technical interview handbook.

---

**Happy Coding! 🚀**
