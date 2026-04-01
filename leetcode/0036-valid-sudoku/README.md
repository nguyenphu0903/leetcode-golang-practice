# LeetCode 36: Valid Sudoku (Sudoku hợp lệ)

## 1. Khái niệm (Concept)

Xác định xem một bảng Sudoku $9 \times 9$ có hợp lệ hay không. Chỉ các ô đã điền mới cần được xác thực theo các quy tắc sau:
1. Mỗi hàng phải chứa các chữ số từ 1-9 mà không lặp lại.
2. Mỗi cột phải chứa các chữ số từ 1-9 mà không lặp lại.
3. Mỗi trong 9 lưới con $3 \times 3$ của bảng phải chứa các chữ số từ 1-9 mà không lặp lại.

**Lưu ý:** Một bảng Sudoku hợp lệ (có thể giải được) không nhất thiết phải là bảng Sudoku đang xét. Bạn chỉ cần kiểm tra các ô đã điền có vi phạm 3 quy tắc trên hay không.

## 2. Tư duy thuật toán (Intuition)

Để kiểm tra sự trùng lặp hiệu quả, chúng ta sử dụng **Hash Set** (trong Go là `map[byte]bool` hoặc dùng mảng đánh dấu).

### Phân tích các chiều kiểm tra:
1. **Hàng (Rows):** Có 9 hàng, mỗi hàng cần 1 Set riêng.
2. **Cột (Cols):** Có 9 cột, mỗi cột cần 1 Set riêng.
3. **Ô 3x3 (Boxes):** Có 9 ô $3 \times 3$, mỗi ô cần 1 Set riêng.

### Bí quyết: Cách xác định chỉ số của ô 3x3
Làm sao để biết một ô ở tọa độ $(r, c)$ thuộc về ô $3 \times 3$ nào?
- Công thức: **`box_index = (r / 3) * 3 + (c / 3)`**
- Với công thức này, các ô sẽ được đánh số từ 0 đến 8:
  ```text
  0 0 0 | 1 1 1 | 2 2 2
  0 0 0 | 1 1 1 | 2 2 2
  0 0 0 | 1 1 1 | 2 2 2
  ------+-------+------
  3 3 3 | 4 4 4 | 5 5 5
  ... và cứ thế tiếp tục
  ```

## 3. Các bước thực hiện

1. Duyệt qua từng ô của bảng $9 \times 9$ bằng 2 vòng lặp lồng nhau.
2. Nếu ô trống (`.`), bỏ qua.
3. Nếu ô có số:
   - Kiểm tra xem số đó đã tồn tại trong Set của **Hàng** hiện tại chưa.
   - Kiểm tra xem số đó đã tồn tại trong Set của **Cột** hiện tại chưa.
   - Kiểm tra xem số đó đã tồn tại trong Set của **Ô 3x3** tương ứng chưa.
4. Nếu đã tồn tại ở bất kỳ đâu $\rightarrow$ Trả về `false`.
5. Ngược lại, thêm số đó vào cả 3 Set và tiếp tục.
6. Nếu duyệt hết toàn bảng mà không vi phạm $\rightarrow$ Trả về `true`.

---
Mời bạn xem file `problem.go` để bắt đầu thực hiện nhé! 🧩🚀
