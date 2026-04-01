# Largest Submatrix With Rearrangements

## 1. Khái niệm (Concept)

Bài toán này yêu cầu tìm diện tích lớn nhất của một ma trận con toàn số 1, với một điều kiện đặc biệt: **Bạn có thể hoán vị (rearrange) các cột của ma trận ban đầu**.

## 2. Ý tưởng giải quyết (Intuition)

Để giải bài này, ta thực hiện qua 3 bước chính:

### Bước 1: Tính "Chiều cao" tích lũy (Height Accumulation)
Với mỗi ô `(i, j)`, ta tính xem có bao nhiêu số 1 liên tiếp kết thúc tại đó (tính từ trên xuống).
- Nếu `matrix[i][j] == 1`: `height[i][j] = height[i-1][j] + 1`
- Nếu `matrix[i][j] == 0`: `height[i][j] = 0`

> Đây là một dạng biến thể của Prefix Sum, nhưng Reset về 0 khi gặp số 0.

### Bước 2: Sắp xếp các cột theo hàng (Sorting)
Tại mỗi hàng `i`, ta có một danh sách các chiều cao của các cột. Vì ta có thể hoán vị các cột, nên để tạo ra hình chữ nhật lớn nhất, ta nên xếp các cột có chiều cao lớn đứng cạnh nhau.
- **Hành động**: Sắp xếp mảng chiều cao của hàng `i` theo thứ tự **giảm dần**.

### Bước 3: Tính diện tích lớn nhất (Area Calculation)
Sau khi sắp xếp giảm dần, tại hàng `i`, cột thứ `k` (0-indexed) sẽ có chiều cao là `h = heights[k]`.
Vì các cột từ `0` đến `k` đều có chiều cao $\ge h$, nên ta có thể tạo một hình chữ nhật có:
- Chiều cao: `h`
- Chiều rộng: `k + 1`
- Diện tích: `h * (k + 1)`

Ta cập nhật `maxArea = max(maxArea, h * (k + 1))`.

## 3. Độ phức tạp (Complexity)

- **Time Complexity**: $O(M \times N \log N)$ 
  - $M \times N$ để tính chiều cao.
  - Mỗi hàng trong $M$ hàng cần sắp xếp $N$ cột ($N \log N$).
- **Space Complexity**: $O(N)$ hoặc $O(M \times N)$ tùy thuộc vào việc bạn có ghi đè lên ma trận gốc hay không.

## 4. Mẹo Interview

- Đây là bài toán mở rộng của "Largest Rectangle in Histogram". Tuy nhiên, nhờ việc được hoán vị cột, bài toán trở nên dễ hơn vì ta chỉ cần sắp xếp thay vì dùng Stack.
- Luôn hỏi xem có được phép thay đổi ma trận đầu vào không để tối ưu bộ nhớ.
