# vjudge_lab_helper
本项目是一个帮助学生在抓取Vjudge Contest题目描述和AC代码并生成实验报告的工具

# 使用教程
## 下载vl_helper_v0.9.exe和tmpl.docx文件到同一目录下
## 打开vl_helper_v0.9.exe输入vjudge用户名和密码（注意！输入密码时不显示，但程序依然在读取）
## 输入contest个数
## 输入contestID
- 多个ID需要以空格隔开一次输入
- contest在https://vjudge.net/contest可以获取，单个contest也可以在url链接中看到
## 等待程序运行
- 可能会遇到访问超时的情况，重新开启软件重复以上操作即可
- 由于时间仓促，程序未作过多的异常处理，因此可能会出现莫名其妙的闪退等问题，重开即可
- 由于vjudge返回的数据格式不统一，比较混乱，暂时没有做到完美解析，因此需要人工辅助完成实验报告
## 程序正常运行结束后会在运行目录创建./code文件夹和vlh_out.docx文件
- 实验AC代码将以：username-实验序号-题目序号.cpp的格式保存到./code文件夹
- vlh_out。docx为生成的实验报告，需要人工辅助微调格式

## TODO
- 针对vjudge返回数据开发解析模块，完美解析数据
- 在上一条的基础上开发AC代码截图保存功能
