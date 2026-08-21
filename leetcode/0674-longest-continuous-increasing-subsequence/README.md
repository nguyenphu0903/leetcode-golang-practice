# 674. Longest Continuous Increasing Subsequence

## 1. Khái niệm

Cho mảng chưa sắp xếp, tìm độ dài dãy con **tăng dần liên tục** dài nhất.
- **Liên tục** = các phần tử kề nhau trong mảng gốc
- **Tăng dần** = `nums[i] > nums[i-1]` (nghiêm ngặt)

**Ví dụ:** `[1,3,5,4,7]`
- `[1,3,5]` tăng liên tục dài 3
- `[4,7]` dài 2
- Đáp án: **3** (không được lấy `[1,3,5,7]` vì không liên tục)

## 2. Tư duy

**Giống hệt 485!** Chỉ khác điều kiện:

| Bài | Điều kiện `if` |
|-----|----------------|
| 485 | `num == 1` |
| 674 | `nums[i] > nums[i-1]` |

**Chiến lược 2 biến:**
1. `current = 1` (mỗi số tự nó là dãy dài 1)
2. `longest = 1`
3. Duyệt từ `i=1`:
   - Nếu `nums[i] > nums[i-1]` -> `current++`
   - Ngược lại -> `longest = max(longest, current)`, `current = 1`
4. Cuối: `return max(longest, current)`

## 3. Lưu ý

- Khởi tạo `current = 1` chứ không phải `0` (vì mảng có ít nhất 1 phần tử)
- Trường hợp `[2,2,2,2,2]` -> đáp án `1`

## 4. So sánh

- 485: đếm `1` liên tiếp
- 674: đếm `tăng` liên tiếp
- 128: đếm `số liên tiếp` trong Set lộn xộn

Cùng 1 pattern "đếm chuỗi"!
