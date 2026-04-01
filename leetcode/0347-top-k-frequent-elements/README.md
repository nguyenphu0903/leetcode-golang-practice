# 347. Top K Frequent Elements (K phần tử xuất hiện nhiều nhất)

## 1. Khái niệm (Concept)

Cho một mảng số nguyên `nums` và một số nguyên `k`, hãy trả về `k` phần tử xuất hiện thường xuyên nhất. Bạn có thể trả về kết quả theo bất kỳ thứ tự nào.

Ví dụ: `nums = [1,1,1,2,2,3], k = 2`
- Số 1 xuất hiện 3 lần.
- Số 2 xuất hiện 2 lần.
- Số 3 xuất hiện 1 lần.
$\rightarrow$ `k = 2` phần tử nhiều nhất là **`[1, 2]`**.

## 2. Tư duy thuật toán (Intuition)

Mọi bài toán "Top K" đều có 2 giai đoạn:
1. **Đếm tần suất:** Dùng Hash Map `map[int]int` để đếm xem mỗi số xuất hiện bao nhiêu lần.
2. **Lọc ra Top K:** Có 3 cách tiếp cận chính:

### Cách 1: Sắp xếp (Sorting) - $O(N \log N)$
- Lấy danh sách các cặp (số, tần suất) và sắp xếp giảm dần theo tần suất.
- Lấy K phần đầu tiên.
- Nhược điểm: Chậm khi mảng lớn.

### Cách 2: Heap (Priority Queue) - $O(N \log K)$
- Duy trì một Min-Heap kích thước K để lưu các phần tử có tần suất lớn nhất.
- Ưu điểm: Hiệu quả hơn sắp xếp khi K nhỏ.

### Cách 3: Bucket Sort (Phân loại vào xô) - $O(N)$ - **Tối ưu nhất!**
- Thay vì sắp xếp theo tần suất, ta dùng **Tần suất** làm **Chỉ số (Index)** của một mảng mới.
- Mảng này gọi là `buckets`, trong đó `buckets[i]` chứa danh sách các số xuất hiện đúng `i` lần.
- Ví dụ: Với `nums = [1,1,1,2,2,3]`:
  - `buckets[3]` chứa `[1]`
  - `buckets[2]` chứa `[2]`
  - `buckets[1]` chứa `[3]`
- Ta chỉ cần duyệt ngược từ cuối mảng `buckets` về đầu để lấy đủ K phần tử.

## 3. Các bước thực hiện (Bucket Sort)

1. Đếm tần suất dùng Map: `count := map[int]int`.
2. Tạo mảng các xô: `buckets := make([][]int, len(nums) + 1)`. 
3. Duyệt Map, đưa từng số vào xô tương ứng với tần suất của nó.
4. Duyệt ngược `buckets` từ `len(nums)` về `0`, nhặt các số cho vào mảng kết quả cho đến khi đủ `k` số.

---
Mời bạn đọc qua và chuẩn bị tinh thần để "vọc" trong file `problem.go`! 🚀
