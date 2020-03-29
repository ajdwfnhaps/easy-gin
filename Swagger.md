# Swagger 笔记



##### Tags

Tags 是用来给API分组的。

##### Accept

接收的参数类型，支持表单(`mpfd`) 和 JSON(`json`)

##### Produce

返回的数据结构，一般都是`json`, 其他支持如下表：

| Mime Type                         | 声明                  |
| --------------------------------- | --------------------- |
| application/json                  | json                  |
| text/xml                          | xml                   |
| text/plain                        | plain                 |
| html                              | html                  |
| multipart/form-data               | mpfd                  |
| application/x-www-form-urlencoded | x-www-form-urlencoded |
| application/vnd.api+json          | json-api              |
| application/x-json-stream         | json-stream           |
| application/octet-stream          | octet-stream          |
| image/png                         | png                   |
| image/jpeg                        | jpeg                  |
| image/gif                         | gif                   |

##### Param

参数，从前往后分别是：

> @Param `1.参数名` ` 2.参数类型` ` 3.参数数据类型` ` 4.是否必须` `5.参数描述` `6.其他属性`

- 1.参数名

  参数名就是我们解释参数的名字。

- 2.参数类型

  参数类型主要有三种：

  - `path` 该类型参数直接拼接在URL中，如[Demo](https://github.com/razeencheng/demo-go/blob/master/swaggo-gin/handle.go)中`HandleGetFile`：

    ```
    // @Param id path integer true "文件ID"
    ```

  - `query` 该类型参数一般是组合在URL中的，如[Demo](https://github.com/razeencheng/demo-go/blob/master/swaggo-gin/handle.go)中`HandleHello`

    ```
    // @Param who query string true "人名"
    ```

  - `formData` 该类型参数一般是`POST,PUT`方法所用，如[Demo](https://github.com/razeencheng/demo-go/blob/master/swaggo-gin/handle.go)中`HandleLogin`

    ```
    // @Param user formData string true "用户名" default(admin)
    ```

- 3.参数数据类型

  数据类型主要支持一下几种：

  - string (string)
  - integer (int, uint, uint32, uint64)
  - number (float32)
  - boolean (bool)

  注意，如果你是上传文件可以使用`file`, 但参数类型一定是`formData`, 如下：

  ```
  // @Param file formData file true "文件"
  ```

- 4.是否是必须

  表明该参数是否是必须需要的，必须的在文档中会黑体标出，测试时必须填写。

- 5.参数描述

  就是参数的一些说明

- 6.其他属性

  除了上面这些属性外，我们还可以为该参数填写一些额外的属性，如枚举，默认值，值范围等。如下：

  ```
  枚举
  // @Param enumstring query string false "string enums" Enums(A, B, C)
  // @Param enumint query int false "int enums" Enums(1, 2, 3)
  // @Param enumnumber query number false "int enums" Enums(1.1, 1.2, 1.3)
  
  值添加范围
  // @Param string query string false "string valid" minlength(5) maxlength(10)
  // @Param int query int false "int valid" mininum(1) maxinum(10)
  
  设置默认值
  // @Param default query string false "string default" default(A)
  ```

  而且这些参数是可以组合使用的，如：

  ```
  // @Param enumstring query string false "string enums" Enums(A, B, C) default(A)
  ```



##### Success

指定成功响应的数据。格式为：

> // @Success `1.HTTP响应码`  `{2.响应参数类型}`  `3.响应数据类型`  `4.其他描述`

- 1.HTTP响应码

  也就是200，400，500那些。

- 2.响应参数类型 / 3.响应数据类型

  返回的数据类型，可以是自定义类型，可以是json。

  - 自定义类型

  在平常的使用中，我都会返回一些指定的模型序列化JSON的数据，这时，就可以这么写：

  ```
  // @Success 200 {object} main.File
  ```

  其中，模型直接用`包名.模型`即可。你会说，假如我返回模型数组怎么办？这时你可以这么写：

  ```
  // @Success 200 {anrry} main.File
  ```

  - json

  将如你只是返回其他的json数据可如下写：

  ```
  // @Success 200 {string} json ""
  ```

- 4.其他描述

  可以添加一些说明。



##### Failure

​	同Success。

##### Router

​	指定路由与HTTP方法。格式为：

> // @Router `/path/to/handle`  [`HTTP方法`]

​	不用加基础路径哦。



###  生成文档与测试

其实上面已经穿插的介绍了。

在`main.go`下运行`swag init`即可生成和更新文档。

点击文档中的`Try it out`即可测试。 如果部分API需要登陆，可以Try登陆接口即可。



### 优化

看到这里，基本可以使用了。但文档一般只是我们测试的时候需要，当我的产品上线后，接口文档是不应该给用户的，而且带有接口文档的包也会大很多（swaggo是直接build到二进制里的）。

想要处理这种情况，我们可以在编译的时候优化一下，如利用`build tag`来控制是否编译文档。



在`main.go`声明`swagHandler`,并在该参数不为空时才加入路由：

```go
package main

//...

var swagHandler gin.HandlerFunc

func main(){
    // ...
    
    	if swagHandler != nil {
			r.GET("/swagger/*any", swagHandler)
        }
    
    //...
}
```

同时,我们将该参数在另外加了`build tag`的包中初始化。

```go
// +build doc

package main

import (
	_ "github.com/razeencheng/demo-go/swaggo-gin/docs"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func init() {
	swagHandler = ginSwagger.WrapHandler(swaggerFiles.Handler)
}
```

之后我们就可以使用`go build -tags "doc"`来打包带文档的包，直接`go build`来打包不带文档的包。
