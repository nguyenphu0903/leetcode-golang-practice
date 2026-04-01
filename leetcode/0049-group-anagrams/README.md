# 49. Group Anagrams (Nhóm các từ đảo chữ)

## 1. Khái niệm (Concept)

Cho một mảng các chuỗi `strs`, hãy nhóm các **anagrams** (từ đảo chữ) lại với nhau. Bạn có thể trả về kết quả theo bất kỳ thứ tự nào.

**Anagram** là một từ hoặc cụm từ được hình thành bằng cách sắp xếp lại các chữ cái của một từ hoặc cụm từ khác, thường sử dụng tất cả các chữ cái gốc chính xác một lần.

Ví dụ: `["eat","tea","tan","ate","nat","bat"]`
- Nhóm 1: `["eat", "tea", "ate"]` (đều có các chữ cái a, e, t)
- Nhóm 2: `["tan", "nat"]` (đều có các chữ cái a, n, t)
- Nhóm 3: `["bat"]`
Kết quả: `[["bat"],["nat","tan"],["ate","eat","tea"]]`

## 2. Tư duy thuật toán (Intuition)

Để nhóm các từ lại với nhau, chúng ta cần một **"Đặc điểm nhận dạng" (Key)** chung cho tất cả các từ trong cùng một nhóm.

### Cách 1: Sắp xếp chuỗi (Sorting) - $O(N \cdot K \log K)$
- Với mỗi chuỗi, ta sắp xếp các chữ cái theo thứ tự bảng chữ cái.
- Ví dụ: `"eat"`, `"tea"`, `"ate"` sau khi sắp xếp đều trở thành **`"aet"`**.
- Dùng chuỗi đã sắp xếp này làm `Key` trong Hash Map.
- `Value` của Map sẽ là một danh sách các chuỗi gốc.

### Cách 2: Đếm tần suất chữ cái (Frequency Counting) - $O(N \cdot K)$
- Vì các chuỗi chỉ gồm chữ cái tiếng Anh thường (a-z), ta có thể tạo một mảng 26 phần tử để đếm số lần xuất hiện của từng chữ cái.
- Ví dụ: `"abc"` $\rightarrow$ `[1, 1, 1, 0, 0, ..., 0]`.
- Dùng mảng này (hoặc chuỗi biểu diễn mảng này) làm `Key`.
- Cách này tối ưu hơn vì không tốn thời gian sắp xếp ($O(K)$ thay vì $O(K \log K)$).

## 3. Các bước thực hiện (sử dụng Cách 1)

1. Khởi tạo một Map: `groups := make(map[string][]string)`.
2. Duyệt qua từng chuỗi `s` trong `strs`:
   - Chuyển chuỗi `s` thành mảng byte hoặc rune.
   - Sắp xếp mảng đó.
   - Chuyển mảng đã sắp xếp ngược lại thành chuỗi `sortedS`.
   - Lưu vào Map: `groups[sortedS] = append(groups[sortedS], s)`.
3. Duyệt qua Map và gom tất cả các `Value` vào một mảng 2 chiều để trả về.

---
Mời bạn đọc qua và bắt đầu giải bài toán này trong file `problem.go`! 🚀
