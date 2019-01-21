
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
/*
包饲料定义了群体饲料。

Swarm feeds允许用户构建有关特定主题的更新feed
每次更新都不使用ENS。
更新方案建立在群块上，块键如下
可预测的、可版本控制的模式。

提要绑定到唯一标识符，该标识符由
所选主题。

提要定义为特定用户对特定主题的一系列更新。

实际数据更新也以群块的形式进行。钥匙
其中的更新是属性串联的散列，如下所示：

updateaddr=h（提要，epoch id）
其中h是sha3散列函数
feed是主题和用户地址的组合
epoch id是一个时隙。有关详细信息，请参阅查找包。

在订阅源中查找最新更新的用户只需知道主题
以及另一个用户的地址。

源更新数据为：
updatedata=feed epoch数据

进入区块负载的完整更新数据是：
updatedata符号（updatedata）

结构总结：

请求：带签名的订阅源更新
 更新：标题+数据
  标题：协议版本，保留供将来使用的占位符
  ID:有关如何定位特定更新的信息
   提要：表示用户关于特定主题的一系列出版物
    主题：更新所涉及的项目
    用户：更新源的用户
   epoch：存储更新的时隙

**/

package feed