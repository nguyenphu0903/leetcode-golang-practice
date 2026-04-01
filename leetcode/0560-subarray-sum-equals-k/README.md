# 560. Subarray Sum Equals K (Tổng mảng con bằng K)

## 1. Khái niệm (Concept)

Cho một mảng số nguyên `nums` và một số nguyên `k`, hãy đếm tổng số **mảng con liên tiếp** có tổng các phần tử đúng bằng `k`.

Ví dụ 1: `nums = [1, 1, 1], k = 2`
- Các mảng con: `[1, 1]` (index 0-1) và `[1, 1]` (index 1-2).
- Kết quả: **2**.

Ví dụ 2: `nums = [1, 2, 3], k = 3`
- Các mảng con: `[1, 2]` (index 0-1) và `[3]` (index 2).
- Kết quả: **2**.

## 2. Tư duy thuật toán (Intuition)

### Cách 1: Brute Force ($O(n^2)$)
- Duyệt qua tất cả các cặp `(i, j)` và tính tổng `sum(i..j)`.
- Với $N = 2 \times 10^4$, $O(n^2)$ có thể bị Time Limit Exceeded (TLE).

### Cách 2: Prefix Sum + Hash Map ($O(n)$)
Đây là kỹ thuật mạnh mẽ nhất để giải các bài về "Subarray Sum".

Ta biết rằng: `sum(i..j) = PrefixSum(j) - PrefixSum(i-1)`.
Để `sum(i..j) == k`, ta cần:
**`PrefixSum(j) - PrefixSum(i-1) == k`**
$\Leftrightarrow$ **`PrefixSum(i-1) == PrefixSum(j) - k`**

**Ý tưởng:**
1. Khi ta đang đứng ở vị trí `j`, ta đã biết `PrefixSum(j)`.
2. Ta chỉ cần kiểm tra xem **trong quá khứ** (từ 0 đến `j-1`), đã có bao nhiêu vị trí `i-1` có `PrefixSum` bằng đúng giá trị `(PrefixSum(j) - k)`.
3. Để làm việc này nhanh chóng ($O(1)$), ta dùng một **Hash Map** để lưu: `Key: PrefixSum`, `Value: Số lần xuất hiện của nó`.

## 3. Các bước thực hiện (Steps)

1. Khởi tạo `count = 0` (biến đếm kết quả).
2. Khởi tạo `currentSum = 0` (tổng tích lũy).
3. Khởi tạo một Hash Map: `m := make(map[int]int)`.
4. **Quan trọng:** Gán `m[0] = 1` (để xử lý trường hợp mảng con bắt đầu từ index 0 có tổng đúng bằng `k`).
5. Duyệt qua mảng:
   - Cộng dồn: `currentSum += nums[i]`.
   - Kiểm tra xem trong Map có `currentSum - k` chưa. Nếu có, cộng `value` của nó vào `count`.
   - Cập nhật Map: `m[currentSum]++`.
6. Trả về `count`.

---
Mời bạn đọc qua và bắt đầu giải bài toán này! Đây là một bước nhảy vọt về tư duy thuật toán đấy! 🚀
