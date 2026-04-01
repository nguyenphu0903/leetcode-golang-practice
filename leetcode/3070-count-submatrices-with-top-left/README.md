# 2D Prefix Sum (Tổng tiền tố 2 chiều)

## 1. Khái niệm (Concept)

**Prefix Sum** (Tổng tiền tố) là một kỹ thuật phổ biến để tính toán nhanh tổng của một đoạn (trong mảng 1 chiều) hoặc một vùng (trong mảng 2 chiều).

Trong mảng 2 chiều (Matrix), `PrefixSum[i][j]` thường đại diện cho tổng của tất cả các phần tử từ ô `(0, 0)` đến ô `(i, j)`.

## 2. Công thức xây dựng (Construction)

Để tính `PrefixSum[i][j]`, ta sử dụng công thức Quy hoạch động (Dynamic Programming):

`PrefixSum[i][j] = grid[i][j] + PrefixSum[i-1][j] + PrefixSum[i][j-1] - PrefixSum[i-1][j-1]`

**Giải thích:**
- `PrefixSum[i-1][j]`: Tổng vùng phía trên ô hiện tại.
- `PrefixSum[i][j-1]`: Tổng vùng phía bên trái ô hiện tại.
- `PrefixSum[i-1][j-1]`: Vùng bị trùng lặp khi cộng hai vùng trên, nên cần trừ đi một lần.
- `grid[i][j]`: Giá trị tại ô hiện tại.

> [!NOTE]
> Khi `i=0` hoặc `j=0`, các chỉ số `-1` sẽ được coi là `0` để tránh lỗi biên.

## 2.1 Cách tiếp cận 2 bước (Two-Pass Intuition - Dễ hiểu nhất)

Nếu công thức trên làm bạn rối, hãy nghĩ về nó theo 2 bước đơn giản sau:

**Bước 1: Tính tổng dồn cho từng hàng (Row-wise prefix sum)**
Duyệt qua từng hàng, mỗi ô sẽ lưu tổng của chính nó và các ô bên trái nó trong cùng hàng đó.
`grid[i][j] = grid[i][j] + grid[i][j-1]`

**Bước 2: Tính tổng dồn cho từng cột (Column-wise prefix sum)**
Duyệt qua từng cột của ma trận vừa thu được ở Bước 1, mỗi ô sẽ lưu tổng của chính nó và các ô phía trên nó trong cùng cột đó.
`grid[i][j] = grid[i][j] + grid[i-1][j]`

**Kết quả**: Sau 2 bước này, mỗi ô `(i, j)` sẽ chứa chính xác tổng của toàn bộ ma trận con từ `(0, 0)` đến `(i, j)`.

## 3. Ứng dụng trong bài 3070

Bài toán yêu cầu đếm số lượng submatrices có:
1. Chứa ô đầu tiên (Top-left element: `grid[0][0]`).
2. Tổng các phần tử nhỏ hơn hoặc bằng `k`.

Vì mọi submatrix đều phải chứa `grid[0][0]`, nên các submatrix này luôn bắt đầu từ `(0, 0)` và kết thúc tại một ô `(i, j)` bất kỳ.

Do đó, bài toán quy về việc:
1. Xây dựng ma trận `PrefixSum` cho `grid`.
2. Đếm số lượng ô `(i, j)` sao cho `PrefixSum[i][j] <= k`.

## 4. Độ phức tạp (Complexity)

- **Time Complexity**: $O(m \times n)$ để duyệt qua toàn bộ ma trận một lần.
- **Space Complexity**: $O(m \times n)$ để lưu trữ ma trận Prefix Sum (có thể tối ưu xuống $O(1)$ nếu ghi đè trực tiếp lên ma trận input hoặc dùng mảng phụ).

## 5. Mẹo Interview

- Luôn kiểm tra điều kiện biên (hàng 0, cột 0).
- Chú ý đến kiểu dữ liệu của biến tổng (trong Go, hãy cẩn thận với overflow nếu `k` rất lớn, nhưng thường `int` là đủ cho LeetCode).
- Hỏi interviewer xem có thể sửa đổi ma trận input không để tiết kiệm bộ nhớ.
