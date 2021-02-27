## 回溯算法

一般用来解决排列组合(Permutation & Combination & Subset)问题，找到所有可能的结果，借助深度优先遍历 dfs，找到终止条件 return 开始下一轮 dfs

```
终止条件
if ...
    return
递归调用，总共分3步:
dfs 之前做出选择
dfs()
dfs 之后撤回选择
```

以及引申出的 `C(n, k), P(n, k)` 问题

子集问题，组合问题的结果是子集问题的一部分，从 n 中取 k 个数的组合，实际上就是求 n 个元素的所有大小为 k 的子集。

## 参考资料

1. [LeetCode 例题精讲 | 08 排列组合问题：回溯法的候选集合](https://mp.weixin.qq.com/s?__biz=MzA5ODk3ODA4OQ==&mid=2648167091&idx=1&sn=82ed3bfa68f92b2826247a0bba40d8ff&chksm=88aa22f5bfddabe322bf5dafeef4f7cd56897d2d7e5b91d55e2baa2b21056694cae7da10c2b5)
2. [LeetCode 例题精讲 | 09 排列组合问题再探：回溯法的去重策略](https://mp.weixin.qq.com/s?__biz=MzA5ODk3ODA4OQ==&mid=2648167095&idx=1&sn=31f6a9115db99f142a4077ba6af308c8&chksm=88aa22f1bfddabe7e0e2d89cdf6d56ea8009388498b64b64eca64629cff29acc5925cbd41b68&scene=178#rd)
