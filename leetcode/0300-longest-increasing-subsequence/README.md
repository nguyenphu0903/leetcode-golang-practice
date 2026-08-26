# 300. Longest Increasing Subsequence (LIS)

## 1. Khái niệm

**Khác 674 ở 1 chữ: Continuous!**

- **674 LCIS:** Phải **kề nhau** (continuous) -> `[1,3,5,4,7]` -> `3`
- **300 LIS:** **Được nhảy cóc** (subsequence) -> `[1,3,5,4,7]` -> `4` (`[1,3,4,7]`)

> Subsequence = xóa bớt phần tử nhưng giữ nguyên thứ tự.

**Ví dụ:** `[10,9,2,5,3,7,101,18]`
- 674 (continuous): chỉ đếm kề nhau -> `[2,5]` hoặc `[3,7,101]` -> `3`
- 300 (subsequence): được nhảy -> `[2,3,7,101]` -> `4` (nhảy qua `5`)

## 2. Tư duy

### Cách 1: DP O(n²) - Dễ hiểu

- `dp[i]` = độ dài LIS **kết thúc tại i**
- `dp[i] = 1 + max(dp[j])` với `j < i` và `nums[j] < nums[i]`
- Khởi tạo `dp[i] = 1`

**Ví dụ trace:** `nums = [10,9,2,5,3,7]`
```
i=0 (10): dp[0]=1
i=1 (9):  so với 10 -> 9<10? không -> dp[1]=1
i=2 (2):  so với 10,9 -> đều lớn hơn 2 -> dp[2]=1
i=3 (5):  so với 2 -> 2<5 -> dp[3]=dp[2]+1=2
i=4 (3):  so với 2 -> 2<3 -> dp[4]=2
i=5 (7):  so với 2,5,3 -> max dp là 2 -> dp[5]=3
=> max dp = 3 hoặc 4 tùy mảng
```

### Cách 2: Binary Search O(n log n) - Follow up

Dùng mảng `tails` + binary search, khó hơn nhưng nhanh. Làm DP trước đã!

## 3. So sánh 674 vs 300

| Bài | Kề nhau? | Kỹ thuật | Độ khó |
|-----|----------|----------|--------|
| 674 | Bắt buộc | 2 biến đếm | Easy |
| 300 | Không, được nhảy | DP | Medium |

## 4. Bẫy

- Khởi tạo `dp[i]=1` chứ không phải `0`
- So sánh `nums[j] < nums[i]` (nghiêm ngặt `<`, không phải `<=`)
- Đáp án là `max(dp)`, không phải `dp[n-1]`

---
Thử DP O(n²) trước nhé! Gợi ý đã có trong `problem.go`.
