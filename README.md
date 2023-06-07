# monkey-go

## 介绍

参照《Writing An Interpreter In Go》实现的一个Monkey脚本语言的解释器，使用Go语言编写。

- let 变量
    - 支持整数、布尔、字符串、哈希、数组
    - 整数运算
    - 字符串拼接
    - 数组索引
- fn 函数
    - 一等公民
    - 支持闭包
    - 自调用
- if 分支语句
    - if else
- 内置函数
    - puts 打印
    - len 计算字符串、数组长度
    - first 取出数组索引为1的元素
    - rest 取出除数组索引为1的元素
    - last 取出数组最后一个元素
    - push 向数组中追加元素

## 代码示例

### 计算斐波那契数

```mk
let fib = fn(n){
  if(n == 0){ return 1;}
  if(n == 1){ return 1;}
  return fib(n-1) + fib(n-2);
};

let x = fib(10);
puts(x);
```

## 编译

## 参考

- https://monkeylang.org/
