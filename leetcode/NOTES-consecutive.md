# 🧠 Bí Kíp Nhớ: Dạng "Chuỗi Liên Tiếp" (Consecutive Sequence)

> Tổng kết sau khi tự tay làm 128 + 485 + 674 — 3 bài cùng pattern!

## 1. Phân biệt 3 dạng

| Đặc điểm | 485. Max Consecutive Ones | 674. LCIS | 128. Longest Consecutive |
|----------|---------------------------|-----------|--------------------------|
| **Mảng** | Đã theo thứ tự, chỉ 0/1 | Đã theo thứ tự, số bất kỳ | Lộn xộn, số bất kỳ |
| **Điều kiện** | `num == 1` | `nums[i] > nums[i-1]` | `x+1` có trong Set? |
| **Kỹ thuật** | 2 biến đếm | 2 biến đếm | Hash Set + điểm bắt đầu |
| **Độ khó** | Easy | Easy | Medium |
| **Cần sort?** | Không, giữ nguyên vị trí | Không, giữ nguyên vị trí | Không, dùng Set |

## 2. Template 485 & 674 - Đếm Tuyến Tính (Nhớ: 2 biến + max cuối!)

### 485 - Đếm số 1
```go
func findMaxConsecutiveOnes(nums []int) int {
    current, longest := 0, 0
    for _, num := range nums {
        if num == 1 {
            current++
        } else {
            if current > longest { longest = current }
            current = 0
        }
    }
    if current > longest { longest = current } // QUAN TRỌNG!
    return longest
}
```

### 674 - Đếm dãy tăng liên tục (chỉ khác 1 dòng if!)
```go
func findLengthOfLCIS(nums []int) int {
    if len(nums) == 0 { return 0 }
    current, longest := 1, 1
    for i := 1; i < len(nums); i++ {
        if nums[i] > nums[i-1] {
            current++
        } else {
            if current > longest { longest = current }
            current = 1 // reset về 1, vì mỗi số tự nó dài 1
        }
    }
    if current > longest { longest = current } // QUAN TRỌNG!
    return longest
}
```

**Bẫy chung cho 485 & 674:**
- Quên `if current > longest` sau vòng for -> sai khi mảng kết thúc bằng dãy dài nhất!
- 485 khởi tạo `0,0` còn 674 khởi tạo `1,1` (vì mỗi phần tử tự nó là dãy dài 1)

## 3. Template 128 - Hash Set (Nhớ: Chỉ đếm từ đầu chuỗi)

```go
func longestConsecutive(nums []int) int {
    if len(nums) == 0 { return 0 }
    set := make(map[int]bool)
    for _, n := range nums { set[n] = true }
    
    longest := 0
    for num := range set {
        if _, ok := set[num-1]; ok { continue } // CHÌA KHÓA!
        cur, streak := num, 1
        for set[cur+1] {
            cur++
            streak++
        }
        if streak > longest { longest = streak }
    }
    return longest
}
```

**Bẫy 128:**
- `longest` phải ở NGOÀI vòng for
- `return` phải ở NGOÀI vòng for
- Chỉ đếm khi `num-1` không tồn tại -> mới O(n)!

## 4. Cách "Nhớ" Tư Duy Phỏng Vấn

**Câu hỏi:** *"Sao nghĩ ra được?"*
1. Thấy `O(n)` -> loại Sort `O(n log n)` -> nghĩ ngay **Hash Set O(1)**
2. Nghĩ brute force `O(n²)` -> hỏi "làm sao tránh đếm lại?"
3. Đáp: "Chỉ đếm từ **điểm bắt đầu** (`x-1` không có)"

**Ví dụ nhớ:** `[1,2,3,4,100,200]` -> điểm bắt đầu là `1, 100, 200` -> chỉ đếm 3 lần!

## 5. So Sánh 674 vs 300 (Để không nhầm!)

Cùng mảng `[1,3,5,4,7]`:
- **674 (Continuous):** `3` -> `[1,3,5]` (phải kề nhau)
- **300 (Không Continuous):** `4` -> `[1,3,4,7]` (được nhảy cóc, chỉ cần giữ thứ tự)

> 674 = 485 phiên bản "tăng dần", 300 = cần DP/Binary Search (khó hơn nhiều!)

## 6. Test Nhanh Để Nhớ

- `[]` -> 0
- `[0,0,1,-1]` (128) -> Set `{0,1,-1}` -> `-1,0,1` -> 3
- `[1,1,0,1,1,1]` (485) -> 3
- `[100,4,200,1,3,2]` (128) -> 4
- `[1,3,5,4,7]` (674) -> 3
- `[2,2,2,2,2]` (674) -> 1 (do `>` nghiêm ngặt)

---
*Đã nắm trọn 3 bài Consecutive! Tiếp theo: Two Pointers (167. Two Sum II) hoặc 300. LIS (Medium) 🚀*
