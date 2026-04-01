# 238. Product of Array Except Self (Tích của mảng loại trừ chính nó)

## 1. Khái niệm (Concept)

Bài toán này là một biến thể nâng cao của **Prefix Sum** (Tổng tiền tố). Thay vì tính tổng, chúng ta tính **Tích** (Product).

Mục tiêu: Với mỗi phần tử `nums[i]`, hãy tính tích của tất cả các phần tử khác trong mảng.
Ví dụ: `nums = [1, 2, 3, 4]`
- Tại index 0: bỏ 1, tích là $2 \times 3 \times 4 = 24$
- Tại index 1: bỏ 2, tích là $1 \times 3 \times 4 = 12$
- Kết quả: `[24, 12, 8, 6]`

**Ràng buộc quan trọng:**
1. Không được dùng phép chia.
2. Thời gian phải là $O(n)$.
3. Không tính mảng kết quả thì không gian sử dụng phải là $O(1)$ (Follow-up).

## 2. Tư duy thuật toán (Intuition)

Mỗi kết quả tại vị trí `i` có thể được chia làm 2 phần:
1. **Tích bên trái (Prefix Product):** Tích của toàn bộ các số đứng trước `i`.
2. **Tích bên phải (Suffix Product):** Tích của toàn bộ các số đứng sau `i`.

`Kết quả[i] = (Tích các số bên trái i) * (Tích các số bên phải i)`

### Ví dụ: `nums = [1, 2, 3, 4]`

| Index | 0 | 1 | 2 | 3 |
| :--- | :--- | :--- | :--- | :--- |
| **Nums** | 1 | 2 | 3 | 4 |
| **Prefix Prod** | 1 | 1 | 1*2=2 | 2*3=6 |
| **Suffix Prod** | 2*3*4=24 | 3*4=12 | 4 | 1 |
| **Kết quả** | $1 \times 24 = 24$ | $1 \times 12 = 12$ | $2 \times 4 = 8$ | $6 \times 1 = 6$ |

*(Ở đây Prefix của phần tử đầu tiên và Suffix của phần tử cuối cùng mặc định là 1).*

## 3. Các bước thực hiện (Steps)

1. Duyệt từ trái sang phải để tính mảng **Prefix Product**.
2. Duyệt từ phải sang trái để tính mảng **Suffix Product**.
3. Nhân hai mảng này với nhau để ra kết quả cuối cùng.

**Tối ưu bộ nhớ ($O(1)$):**
- Thay vì dùng 2 mảng Prefix và Suffix, ta dùng chính mảng kết quả (`ans`) để lưu Prefix.
- Sau đó dùng một biến tạm `right` để lưu tích dồn từ bên phải và nhân trực tiếp vào `ans`.

---
Bạn hãy đọc kỹ phần này trước khi bắt đầu giải bài toán trong `problem.go` nhé!
