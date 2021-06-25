# ServerManagement

server-management

## Cách tổ chức packages:

- entity
- repo
- page(xử lý nghiệp vụ)
- templates(giao diện html, ...)

## Hướng dẫn code:

Ví dụ: Thêm phần quản lý lỗi trong project.

1. Tạo struct Error trong thư mục entity:
  
  // Tạo file entity > error.go
  
  ```
  type Error struct {  
      ID string  
      Description string  
  }
  ```
  
2. Xử lý database trong thư mục repo:
  
  **Lưu ý**: bảng Error trong database đã được tạo.
  
  // Tạo file repo > error_repo.go  
   // **ErrorRepo** chịu trách nhiệm xử lý việc thao tác trên database
  
  ```
  type ErrorRepo struct {  
     // Các thuộc tính
  }
  
  // Các phương thức Insert, Update, Delete, Query, ... trên bảng Error
  ```
  
3. Giả sử có 4 trang html bao gồm **xem, thêm, xóa, sửa Error**, sẽ tạo ra 4 struct để xử lý 4 trang error tương ứng.
  
  // VD: để xử lý trang Xem thông tin 1 error  
   // tạo một struct trong thư mục page để xử lý nghiệp vụ cho yêu cầu xem thông tin error từ người dùng.
  
  ```
  type ErrorInfo struct {  
      repo *ErrorRepo  
  }
  
  func (errInfo *ErrorInfo) New() {  
   // Khởi tạo các thuộc tính  
  }
  
  func (errInfo *ErrorInfo) Fetch(ErrorID string) (entity.Error, error) {  
   // Tiến hành cài đặt phương thức này và trả về thực thể Error + một lỗi nếu có.  
  }
  
  // Tương tự đối với các chức năng khác.  
  ```
  

4. Tạo giao diện html cho chức năng tương ứng.
  
5. Routing:
  

- Ví dụ người dùng gửi một request đến bằng phương thực Get kèm theo id của lỗi cần hiển thị thông tin.

- Trong file main.go, tạo một phương thức xử lí:

```
     func HandleViewError(r *gin.Engine) {  
         r.GET("/error", func (c *gin.Context) {  
             // Trong thư mục page đã tạo struct ErrorInfo tương ứng để tổng hợp thông tin phản hồi cho người dùng  
             // Giả sử không có bất kỳ lỗi nào diễn ra  
       
             errInfo := ErrorInfo{}  
             errInfo.New()  
     
             // Lấy id user gửi kèm theo trong request  
            errID := c.Query("txtID")  

            obj, _ := errInfo.Fetch(errID)  

             // Giả sử quá trình truy vấn thông tin thành công, không có lỗi nào diễn ra  
            r.LoadHTMLFiles("templates/error_view.html")  
            c.HTML(http.StatusOK, obj)  
})  
```

}

- Sau đó gọi hàm vừa tạo ra trong function setupRouter
