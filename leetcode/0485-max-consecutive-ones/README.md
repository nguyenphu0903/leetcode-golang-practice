# 485. Max Consecutive Ones (Số 1 liên tiếp dài nhất)

## 1. Khái niệm (Concept)

Cho mảng nhị phân `nums` chỉ chứa `0` và `1`, tìm số lượng `1` liên tiếp nhiều nhất.

**Ví dụ:** `nums = [1,1,0,1,1,1]`
- Đoạn 1: `[1,1]` dài 2
- Đoạn 2: `[1,1,1]` dài 3
- Đáp án: **3**

## 2. Tư duy thuật toán (Intuition)

Bài này là **phiên bản siêu dễ** của 128 Longest Consecutive Sequence!

Khác ở chỗ:
- 128: số bất kỳ, phải dùng `Hash Set` để tìm `x+1`
- 485: chỉ có `0` và `1`, mảng đã sắp xếp theo thứ tự duyệt -> chỉ cần **đếm tuyến tính**!

**Chiến lược "2 biến đếm":**
1. `current` - đếm số `1` liên tiếp hiện tại
2. `longest` - lưu max tìm được

Duyệt mảng:
- Gặp `1` -> `current++`
- Gặp `0` -> `longest = max(longest, current)`, rồi `current = 0`

Cuối vòng lặp nhớ `max` thêm lần nữa vì mảng có thể kết thúc bằng `1`.

## 3. Độ phức tạp

- **Time:** `O(n)` - duyệt 1 lần
- **Space:** `O(1)` - chỉ dùng 2 biến

## 4. So sánh với 128

| Bài | Kỹ thuật | Khó |
|-----|----------|-----|
| 485 | Đếm tuyến tính | Easy |
| 128 | Hash Set + tìm điểm bắt đầu | Medium |

Làm xong 485 em sẽ thấy 128 chính là 485 phiên bản "mảng lộn xộn"!

---
Mở `problem.go` và thử sức nhé! Gợi ý đã có trong TODO. 🚀
