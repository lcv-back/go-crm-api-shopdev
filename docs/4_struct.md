# Tại sao struct là thế mạnh của go ?

### **1. Tối ưu hóa về hiệu năng**

- **Struct trong Go rất nhẹ:** Chúng được thiết kế để hoạt động với hiệu năng cao, vì chúng không đi kèm với các tính năng nặng nề như class trong các ngôn ngữ lập trình hướng đối tượng khác.
- Dữ liệu trong struct được lưu trữ liên tiếp trong bộ nhớ, giúp truy cập nhanh hơn.

### **2. Sử dụng linh hoạt**

- **Không bị ràng buộc bởi hướng đối tượng:** Mặc dù Go không hỗ trợ kế thừa (inheritance) như các ngôn ngữ hướng đối tượng truyền thống, bạn có thể đạt được tính năng tái sử dụng mã thông qua **embedding** (nhúng struct trong struct).

```go
type Animal struct {
    Name string
}

type Dog struct {
    Animal
    Breed string
}

func main() {
    d := Dog{Animal: Animal{Name: "Buddy"}, Breed: "Labrador"}
    fmt.Println(d.Name) // "Buddy"
}
```

### **3. Khả năng mở rộng với phương thức (methods)**

- **Gán phương thức cho struct:** Struct có thể đi kèm với các phương thức, cho phép bạn tổ chức logic liên quan đến dữ liệu cụ thể một cách gọn gàng mà không cần sử dụng class.

```go
type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func main() {
    rect := Rectangle{Width: 5, Height: 10}
    fmt.Println(rect.Area()) // Output: 50
}
```

### **4. Hỗ trợ interface mạnh mẽ**

- Struct kết hợp chặt chẽ với interface trong Go, giúp bạn xây dựng các hệ thống linh hoạt và có thể mở rộng.
- Struct không cần phải khai báo trước là "implement interface". Chỉ cần struct có các phương thức trùng khớp với interface, Go sẽ tự động hiểu rằng struct đó đã implement interface.

```go
type Shape interface {
    Area() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func main() {
    var s Shape = Circle{Radius: 5}
    fmt.Println(s.Area()) // Output: 78.5
}
```

### **5. Đơn giản và dễ hiểu**

- Struct trong Go được thiết kế rất đơn giản, chỉ với các trường dữ liệu (fields) mà không cần thêm các khái niệm phức tạp như constructor, getter/setter bắt buộc.
- Go khuyến khích các phương pháp viết mã tối giản và dễ đọc, giúp các đội phát triển nhanh chóng hiểu rõ và làm việc với codebase.

### **6. Khả năng tương thích với JSON và các định dạng dữ liệu khác**

- Struct hỗ trợ mạnh mẽ cho việc **marshal/unmarshal** dữ liệu (như JSON, XML). Điều này rất hữu ích khi làm việc với API hoặc xử lý dữ liệu.

```go
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func main() {
    data := `{"id":1,"name":"Alice","email":"alice@example.com"}`
    var user User
    json.Unmarshal([]byte(data), &user)
    fmt.Println(user.Name) // Output: Alice
}
```

### **7. Dễ dàng kiểm soát dữ liệu bằng cách xuất/không xuất các trường**

- Các trường bắt đầu bằng chữ cái in hoa sẽ được exported (công khai) và có thể truy cập từ các package khác.
- Các trường bắt đầu bằng chữ cái thường sẽ chỉ được sử dụng trong package hiện tại **(private).**

```go
type User struct {
    Name  string // Exported
    email string // Private
}
```

### **8. Hỗ trợ các tag tùy chỉnh**

- Struct hỗ trợ gắn tags cho các trường dữ liệu, giúp tương tác với các thư viện hoặc hệ thống khác một cách linh hoạt (ví dụ: JSON, validation).

```go
type User struct {
    ID    int    `json:"id" validate:"required"`
    Name  string `json:"name" validate:"required"`
}
```

### **9. Tương thích với hệ thống concurrent (goroutines)**

- Struct hoạt động rất tốt với các tính năng concurrent của Go. Bạn có thể sử dụng **sync.Mutex** hoặc các cơ chế đồng bộ khác trong struct để kiểm soát truy cập đồng thời vào dữ liệu.

```go
import (
    "fmt"
    "sync"
)

type Counter struct {
    mu    sync.Mutex
    Count int
}

func (c *Counter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.Count++
}

func main() {
    c := &Counter{}
    c.Increment()
    fmt.Println(c.Count) // Output: 1
}
```
