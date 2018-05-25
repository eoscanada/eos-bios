基于EOS.IO的区块链网络启动工具
--------------------------------------------------------

#### v1.0 ZhaoYu from EosLaoMao
#### v2.0 Ian(Yahuang Wu) from MEET.ONE

#### v2.1 Tyee Noprom from EOSpace.io

[Click me switch to English version](./README.md)

`eos-bios` 是一个命令行工具，给那些想要启动基于EOS.IO的区块链网络的人而准备.

它提供以下功能:
* 主网分阶段启动
* 启动本地开发环境Booting local development environments
* 启动测试环境Booting testnets
* 启动联合网络或者私有网络

首先你需要知道的是**discovery protocol**. 点击这个链接可以看到介绍:

[![The Disco Dance](https://i.ytimg.com/vi/8aNZ_ZnKS-A/hqdefault.jpg?sqp=-oaymwEZCNACELwBSFXyq4qpAwsIARUAAIhCGAFwAQ==&rs=AOn4CLCZeGSGv9Qix8mHX77R4-d0rzDkgA)](https://youtu.be/8aNZ_ZnKS-A)

https://youtu.be/8aNZ_ZnKS-A

[点击查看示例配置](./sample_config)

下列 **视频** 解释了`eos-bios`的其他概念:

| 备注           |                           链接 |
| ------------ | ---------------------------: |
| 发现文件的细节      | https://youtu.be/uE5J7LCqcUc |
| 网络划分算法       | https://youtu.be/I-PrnlmLNnQ |
| 启动顺序         | https://youtu.be/SbVzINnqWAE |
| 代码审查核心流程     | https://youtu.be/p1B6jvx25O0 |
| 与您的基础设施集成的钩子 | https://youtu.be/oZNV6fUoyqM |

`eos-bios`的一些问答视频:

| 备注                 |                           链接 |
| ------------------ | ---------------------------: |
| 账号是否从一个阶段延续到另一个阶段？ | https://youtu.be/amyMm5gVpLg |
| 出块节点如何就链的内容达成一致？   | https://youtu.be/WIR7nab40qk |

### 本地开发环境
-----------------------------

[点击下载`eos-bios`](https://github.com/eoscanada/eos-bios/releases),
克隆当前仓库并且复制 `sample_config` 文件夹到新的目录

修改 `my_discovery_file.yaml` 文件,把target_http_address指向本地地址:

```
target_http_address: http://localhost:8888
```

然后运行:

    ./eos-bios boot --single

这个命令会在你本地运行一个完整的开发环境，包含所有的系统合约，跟你在主网启动时的环境很类似.

这个示例配置启动了一个单节点，同时它不会指向其他的出块候选节点.


### 社区协作启动网络
------------------------------

启动社区网络的时候，*每个人*都需要执行:

    ./eos-bios orchestrate

根据特定算法和各个节点提供的 discovery 文件，各个节点将会被自动赋予相应的角色:

1. _BIOS Boot node_, 该节点将单独执行 `eos-bios boot` 命令，完成系统合约的部署，快照的导入，EOS 代币的分配等任务.
2. _Appointed Block Producer_, 这类节点将会执行 `eos-bios join --validate` 命令，加入 BIOS Boot 的网络，并负责验证 BIOS Boot 的行为.
3. _other participant_, 将会执行`eos-bios join`加入网络.

以上执行完成以后，所有节点都需要等待`seed_network_launch_block`达成共识.


### 练习加入网络


让[种子网络](#seed-networks)里面的任何一个人邀请你，他们需要执行:

    ./eos-bios invite [youraccount] [your pubkey]

这条命令会在种子网络中创建一个账号给你，修改你的`my_discovery_file.yaml`:

* `seed_network_account_name`, 需要跟你被邀请的账号一致.
* `seed_network_http_address`, 你想加入的种子网络的地址.

把你的私钥添加到`privkey.keys`, 因为需要运行发布命令:

    ./eos-bios publish

同时你也需要修改`my_discovery_file.yaml`中的这些值:

* `target_http_address` 你当前正在启动的节点HTTP地址.
* `target_p2p_address` 你当前正在启动的节点P2P地址.
* `target_appointed_block_producer_signing_key` and `target_initial_authority`: 这两个是跟你账户相关的授权信息.
* `seed_network_peers` [点击查看细节](#network-peers)

其他比较重要的字段:
* `seed_network_launch_block` 种子网络启动时的区块高度.
* `target_contents` 这里面配置的是我们已经达成共识可以写入链中的内容，比如系统合约，ERC-20快照等等.



### 练习启动网络

检查你的`my_discovery_file.yaml`,确认没问题运行:

    ./eos-bios boot

这个会测试你的所有`hook_boot_*.sh`.


查看种子网络的数据
-----------------------

如果你的配置文件已经指向了一个种子网络，你可以运行:

    ./eos-bios discover

将会打印出节点的信息, 比如节点权重等.


网络同步
-------------

discovery文件中的`seed_network_peers`段落看起来是这样的:

```
seed_network_peers:
- account: eosexample
  comment: "They are good"
  weight: 10  # Weights are between 0 and 100 (INT value)
- account: eosmore
  comment: "They are better"
  weight: 20
```
这个代表着你同意跟`eosexample`(10%的权重),`eosmore`(20%的权重)一起启动网络,`eos-bios`将会基于这些信息计算出网络结构.
这两个账户都是种子网络中的。


种子网络
-------------

这个列表是`eos-bios`启动的不同阶段的种子网络，我们会持续更新.
https://stages.eoscanada.com

## 流程化发布中的示例流程和干预

1. 每个人都运行 `eos-bios orchestrate`；
2. `eos-bios` 会下载`my_discovery_file.yaml`指向的网络拓扑，每个人都一样；
3. 网络拓扑结构按照人们在`peer`部分投票的方式进行权重排序；
4. `launch_ethereum_block`取自拓扑结构的前20％：如果他们都同意，会继续使用该数字。 否则，我们等到他们这样做（并周期性地重新绘制网络图）。

安装 / 下载
------------------

你可以在这里下载到最新的版本: https://github.com/eoscanada/eos-bios/releases 该工具是一个单独的二进制文件，支持主流的操作系统。该工具没有其他依赖，只需被赋予可执行权限，即可运行。

当然，你也可以自行编译（Go 语言）：:

    go get -v github.com/eoscanada/eos-bios/eos-bios

编译成功之后，可执行文件将默认安装到 ~/go/bin 目录下（安装Go请查阅https://golang.org/dl）


加入讨论
-------------------

加入我们的电报群: https://t.me/joinchat/GSUv1UaI5QIuifHZs8k_eA (EOSIO BIOS Boot channel)

以前的想法
--------------------

可以在README.v0.md找到之前的一些想法.

准备启动前的检查事项
-------------------

* `target_p2p_address`是否指向我准备要运行的节点?
* `target_http_address`是否指向我的节点，执行`eos-bios`的节点可以正常连接吗？


疑难杂症
---------------

* `hook_join_network.sh`中的`PRIVKEY`和`PUBKEY`是否跟`discovery_file.yaml`文件中的`target_initial_authority` and `target_initial_block_signing_key`匹配?

* 是否指向了一个旧的P2P地址?如果是的话，请将它从网络中移除。
