# 724. Find Pivot Index (Tìm chỉ số trung tâm)

## 1. Khái niệm (Concept)

Chỉ số "Pivot" (trung tâm) của một mảng là chỉ số mà tại đó **tổng của tất cả các số bên trái** bằng **tổng của tất cả các số bên phải**.

Ví dụ: `nums = [1, 7, 3, 6, 5, 6]`
- Tại index 3 (số 6):
  - Tổng bên trái: $1 + 7 + 3 = 11$
  - Tổng bên phải: $5 + 6 = 11$
- Vậy Pivot Index là **3**.

**Lưu ý:**
- Nếu không có chỉ số nào thỏa mãn, trả về **-1**.
- Nếu chỉ số nằm ở đầu mảng (index 0), tổng bên trái mặc định là **0**. Tương tự với cuối mảng.

## 2. Tư duy thuật toán (Intuition)

Thay vì tính tổng trái và phải cho mỗi vị trí (tốn $O(n^2)$), chúng ta có thể dùng kỹ thuật **Prefix Sum** (Tổng tiền tố) để giải trong **$O(n)$**.

Giả sử ta gọi:
- `totalSum`: Tổng của toàn bộ mảng.
- `leftSum`: Tổng của các số đứng TRƯỚC vị trí `i`.
- `rightSum`: Tổng của các số đứng SAU vị trí `i`.

Tại bất kỳ vị trí `i` nào, chúng ta luôn có mối quan hệ:
`totalSum = leftSum + nums[i] + rightSum`

Từ đó suy ra:
`rightSum = totalSum - leftSum - nums[i]`

Điều kiện để `i` là Pivot Index là:
`leftSum == rightSum`
$\Leftrightarrow$ **`leftSum == totalSum - leftSum - nums[i]`**

## 3. Các bước thực hiện (Steps)

1. Tính `totalSum` của toàn bộ mảng.
2. Khởi tạo `leftSum = 0`.
3. Duyệt mảng từ trái sang phải:
   - Kiểm tra điều kiện: `leftSum == totalSum - leftSum - nums[i]`.
   - Nếu đúng, trả về `i` ngay lập tức.
   - Nếu sai, cập nhật `leftSum += nums[i]` và tiếp tục.
4. Nếu hết vòng lặp mà không tìm thấy, trả về -1.

---
Mời bạn đọc qua và bắt đầu giải bài toán trong file `problem.go`!
