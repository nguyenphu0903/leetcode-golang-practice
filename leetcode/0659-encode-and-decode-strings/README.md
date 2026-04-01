# Encode and Decode Strings (Mã hóa và Giải mã chuỗi)

## 1. Khái niệm (Concept)

Thiết kế một thuật toán để mã hóa một danh sách các chuỗi thành một chuỗi duy nhất. Sau đó, chuỗi mã hóa này được gửi qua mạng và được giải mã lại thành danh sách các chuỗi ban đầu.

Ví dụ: 
- Input: `["lint", "code", "love", "you"]`
- Encode thành: `"4#lint4#code4#love3#you"`
- Decode ngược lại: `["lint", "code", "love", "you"]`

## 2. Thử thách (The Challenge)

Tại sao chúng ta không dùng một dấu phẩy `,` để ngăn cách các từ?
- **Vấn đề:** Nếu trong từ gốc có chứa dấu phẩy (ví dụ: `["hello", "wor,ld"]`), khi giải mã bạn sẽ bị nhầm lẫn và chia thành 3 từ `["hello", "wor", "ld"]`.

$\rightarrow$ Chúng ta cần một cách thức "đóng gói" mà không sợ nội dung bên trong từ có chứa bất kỳ ký tự đặc biệt nào.

## 3. Giải pháp: Length-Prefix Encoding (Mã hóa tiền tố độ dài)

Đây là kỹ thuật cực kỳ phổ biến trong truyền tin (như giao thức HTTP). 

### Cách mã hóa (Encode):
Với mỗi chuỗi, ta lưu theo công thức: **`[Độ dài]` + `[Ký tự đặc biệt]` + `[Nội dung chuỗi]`**.

Ví dụ: `["lint", "code"]`
- "lint" có độ dài 4. Mã hóa thành: `4#lint`
- "code" có độ dài 4. Mã hóa thành: `4#code`
- Kết quả cuối: `"4#lint4#code"`

### Cách giải mã (Decode):
1. Đọc từng ký tự cho đến khi gặp dấu `#`.
2. Con số đứng trước dấu `#` chính là số lượng ký tự bạn cần đọc tiếp theo.
3. Sau khi đọc đủ số ký tự, bạn lại lặp lại bước 1 cho đến hết chuỗi.

---
Mời bạn xem file `problem.go` để bắt đầu thực hiện nhé! 🚀
