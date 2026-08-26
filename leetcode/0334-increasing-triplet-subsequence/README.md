# 334. Increasing Triplet Subsequence

## 1. Khái niệm

**Dễ hơn 300 nhiều!** Không cần tìm dãy dài nhất, chỉ cần hỏi: **Có tồn tại 3 số tăng dần không?**

- `i < j < k` (giữ nguyên thứ tự)
- `nums[i] < nums[j] < nums[k]` (tăng nghiêm ngặt)

**Ví dụ:**
- `[1,2,3,4,5]` -> `true` (1<2<3)
- `[5,4,3,2,1]` -> `false` (giảm dần)
- `[2,1,5,0,4,6]` -> `true` (0<4<6)

## 2. Tư duy

**Không cần DP O(n²) như 300!** Chỉ cần **2 biến + 1 vòng for O(n)**:

- `first` = số nhỏ nhất đã gặp
- `second` = số nhỏ thứ 2 (lớn hơn first)

Duyệt `num`:
```
nếu num <= first  -> first = num  (tìm số nhỏ hơn để làm nền)
nếu num <= second -> second = num (tìm số vừa)
nếu num > second  -> return true  (tìm được số thứ 3 lớn hơn cả 2!)
```

**Ví dụ trace:** `[2,1,5,0,4,6]`
```
num=2: first=2, second=Inf
num=1: first=1 (1<2)
num=5: 5>first(1) và 5>second(Inf)? -> second=5
num=0: first=0 (0<1)
num=4: 4>first(0) và 4<=second(5)? -> second=4 (cập nhật second nhỏ hơn)
num=6: 6>second(4) -> true! (0<4<6)
```

## 3. So sánh 300 vs 334

| Bài | Hỏi gì? | Kỹ thuật |
|-----|---------|----------|
| 300 LIS | Dãy dài nhất bao nhiêu? | DP O(n²) |
| 334 Triplet | Có 3 số tăng không? | 2 biến O(n) |

334 chính là **bản rút gọn** của 300!

## 4. Bẫy

- `<=` chứ không phải `<` khi cập nhật first/second (để xử lý số bằng nhau)
- Khởi tạo `first, second = MaxInt`

---
Thử với 2 biến thôi nhé! Dễ hơn 300 nhiều!
