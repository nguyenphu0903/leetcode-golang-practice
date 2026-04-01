# 128. Longest Consecutive Sequence (Chuỗi liên tiếp dài nhất)

## 1. Khái niệm (Concept)

Cho một mảng số nguyên chưa được sắp xếp `nums`, hãy tìm độ dài của chuỗi các phần tử liên tiếp dài nhất.

Thuật toán của bạn phải chạy trong độ phức tạp thời gian **$O(n)$**.

**Ví dụ:** `nums = [100, 4, 200, 1, 3, 2]`
- Các số liên tiếp là `[1, 2, 3, 4]`.
- Độ dài là **4**.

## 2. Tư duy thuật toán (Intuition)

### Cách 1: Sắp xếp (Sorting) - $O(n \log n)$
- Sắp xếp mảng, sau đó duyệt qua để tìm chuỗi dài nhất.
- Nhược điểm: Đề bài yêu cầu $O(n)$, mà sắp xếp thì tốn $O(n \log n)$.

### Cách 2: Hash Set - $O(n)$ - **Hợp lý nhất!**
Để đạt được $O(n)$, chúng ta cần kiểm tra sự tồn tại của một số trong $O(1)$. **Hash Set** chính là công cụ đó.

**Chiến lược "Tìm điểm bắt đầu":**
1. Đưa tất cả các số vào một Hash Set (trong Go là `map[int]bool`).
2. Duyệt qua từng số `x` trong mảng:
   - Kiểm tra xem `x` có phải là **số bắt đầu** của một chuỗi không?
   - Một số `x` là điểm bắt đầu khi và chỉ khi số **`x - 1` KHÔNG tồn tại** trong Set.
   - Nếu `x` là điểm bắt đầu:
     - Bắt đầu đếm: `x+1, x+2, x+3...` xem chuỗi này dài bao nhiêu.
     - Cập nhật độ dài lớn nhất tìm được.
   - Nếu `x` không phải điểm bắt đầu (tức là có `x - 1` trong Set): Bỏ qua, vì chuỗi chứa `x` đã hoặc sẽ được đếm khi ta xét đến số `x - 1`.

## 3. Tại sao cách này lại là $O(n)$?
Dù có vòng lặp lồng nhau, nhưng mỗi số thực chất chỉ được "thăm" tối đa 2 lần:
1. Một lần ở vòng lặp ngoài.
2. Một lần ở vòng lặp trong (khi nó là một phần của chuỗi liên tiếp).
$\rightarrow$ Tổng thời gian là tuyến tính $O(n)$.

---
Mời bạn mở file `problem.go` để cùng "về đích" nhóm Arrays & Hashing nhé! 🏁🚀
