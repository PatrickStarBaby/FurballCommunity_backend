# FurballCommunity_backend
“Furball Community” is a community-based pet sharing and boarding platform website, which is a course design project.

“毛球社区”是一个基于社区的共享宠物寄养平台，是UCAS2023春季高级软件工程课程的课程设计项目。


### 整体架构
- 后端：
  语言采用go
  Web框架：Gin
  ORM：GORM

- 前端：
  网页端：Vue全栈+ElementUI+axios
  移动端：uni-app

- 数据库：
  mysql


### 后端项目架构

- config： 项目配置模块，将集成 mysql、swagger、token、redis 等配置部分。
- controller： 负责请求转发，接受页面过来的参数，传给 Model 处理，接到返回值，再传给页面。
- models： 对应数据表的增删查改。
- routers：处理路由。
- utils：定义项目工具组件，包括错误代码，返回类型等。



### 表设计

- user

| Id     | Password | Name   | Authority                         | Gender       | Address | Score | Intro    | ID_number | Avatar | Pet_experience | Work_time     | Pet_number           | ID_card_photo |
| ------ | -------- | ------ | --------------------------------- | ------------ | ------- | ----- | -------- | --------- | ------ | -------------- | ------------- | -------------------- | ------------- |
| 用户id | 用户密码 | 昵称   | 权限                              | 性别         | 地址    | 评分  | 个人介绍 | 身份证号  | 头像   | 养宠经历       | 可工作时间    | 最大同时照顾宠物数量 | 身份证照片    |
| int    | string   | string | int                               | int          | string  | float | string   | string    | string | string         | Int           | int                  | string        |
|        |          |        | 0：普通用户  1：可收养  2：管理员 | 0：女  1：男 |         |       |          |           |        |                | 0：<=4  1：>4 |                      |               |



- pet

| Pet_id | User_id | Name   | Gender       | Photo  | Age  | Weight | Sterilization      | Breed  | Health   | Avatar   |
| ------ | ------- | ------ | ------------ | ------ | ---- | ------ | ------------------ | ------ | -------- | -------- |
| 宠物id | 用户id  | 宠物名 | 性别         | 照片   | 年龄 | 重量   | 是否绝育           | 品种   | 健康状况 | 宠物头像 |
| Int    | Int     | string | int          | String | Int  | Int    | Int                | String | String   | string   |
|        |         |        | 0：母  1：公 |        |      |        | 0：未绝育  1：绝育 |        |          |          |



- order

| Order_id | announcer_id | Receiver_id   | Pet_id | Announce_time | Start_time | End_time | Place  | Pet_health | Status                                        | Remark | price | Evaluation    | score | check_in_status | check_in_pictures | phone    |
| -------- | ------------ | ------------- | ------ | ------------- | ---------- | -------- | ------ | ---------- | --------------------------------------------- | ------ | ----- | ------------- | ----- | --------------- | ----------------- | -------- |
| 订单id   | 发布者       | 接收者        | 宠物id | 发布时间      | 开始时间   | 结束时间 | 地点   | 健康状况   | 订单    状态                                  | 备注   | 报酬  | 评价          | 评分  | 打卡状态        | 打卡图片          | 联系电话 |
| int      | int          | int           | int    | time          | string     | string   | string | string     | int                                           | string | Int   | string        | float | string          | string            | string   |
|          |              | 当status为123 |        |               |            |          |        |            | -1：已取消  0：待付款   1：进行中   2：待评价   3：已完成 |        |       | 仅当status为3 |       |                 |                   |          |



- order_comment

| Comment_id | oeder_id | User_id | Content  | Time | Parent_id    | Is_root      |
| ---------- | -------- | ------- | -------- | ---- | ------------ | ------------ |
| 订单评论id | 订单id   | 用户id  | 评论内容 | 时间 | 母评论用户id | 是否为根评论 |
| Int        | Int      | Int     | String   | Time | int          | bool         |



- blog

| Blog_id | User_id | Content  | Time     | Title    | Banner_list | Like     |
| ------- | ------- | -------- | -------- | -------- | ----------- | -------- |
| 文章id  | 用户id  | 正文内容 | 发布时间 | 文章标题 | 轮播图列表  | 文章点赞 |
| Int     | Int     | String   | Time     | String   | Url_list    | int      |



- blog_comment

| Blog_comment_id | Blog_id | User_id | Time | Content | Avatar | User_name | Like       |
| --------------- | ------- | ------- | ---- | ------- | ------ | --------- | ---------- |
| 文章评论id      | 文章id  | 用户id  | 时间 | 内容    | 头像   | 用户名    | 评论点赞数 |
| Int             | Int     | Int     | Time | String  | String | String    | int        |



### 项目各项配置

- mysql数据库配置：

  只需要创建名为“fc”数据库，ORM模块会自动创建表

- 发送短信功能：

  使用的阿里云的短信服务，为防止AccessKey ID、Secret泄露，使用了本机系统环境变量存放身份凭证，需手动添加环境变量：

  ```
  ALIBABA_CLOUD_ACCESS_KEY_ID 和 ALIBABA_CLOUD_ACCESS_KEY_SECRET
  ```

  - **Linux和macOS系统配置方法**

  ```bash
  export ALIBABA_CLOUD_ACCESS_KEY_ID=<access_key_id>
  export ALIBABA_CLOUD_ACCESS_KEY_SECRET=<access_key_secret>
  ```

  - **Windows系统配置方法**

  ```bash
  a. 添加环境变量ALIBABA_CLOUD_ACCESS_KEY_ID和ALIBABA_CLOUD_ACCESS_KEY_SECRET，并写入已准备好的AccessKey ID和AccessKey Secret；
  b. 重启系统
  ```

- redis服务器配置：

  - 开启服务即可

- 创建静态服务器文件夹：

  - 在服务运行的父目录添加一个img文件夹，所有的静态文件，包含上传的图片都保存在这个文件夹

![image-20230706145143801](.\docs\image\image-20230706145143801.png)

​			例如：服务部署在service目录中，与service同级创建一个img文件夹