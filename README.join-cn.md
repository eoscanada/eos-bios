# 基于EOS.IO的区块链网络启动工具

## 1. 克隆[eos-bios repo](https://github.com/eoscanada/eos-bios)
## 2. 下载最新版本的[eos-bios release](https://github.com/eoscanada/eos-bios/releases)

或者直接从源码编译安装:
        
    go get -u -v github.com/eoscanada/eos-bios/eos-bios

## 3. 找到一个人邀请你进入种子网络
提供你12位字母的账户名以及公钥地址(publicc),他们会执行以下命令:
        
    eos-bios invite YOUR_ACCOUNT_NAME YOUR_PUBKEY

## 4. 更新你的配置文件
复制`sample_config`所有文件到新的文件夹`stageX`，X是我们正在运行的stage编号，修改`my_discovery_file.yaml`文件:
* `seed_network_account_name`, 需要跟你被邀请时提供的账户名一致
* `seed_network_http_address`, 你想要加入的种子网络的地址
* `seed_network_peers`, [点击查看详细内容](#网络节点)
* `target_http_address`, 对外提供的HTTP地址
* `target_p2p_address`, 对外提供的P2P地址
* `target_account_name`, `target_appointed_block_producer_signing_key`, `target_initial_authority`: 这三个都需要改成你的值
* `target_contents`, 我们达成一致可以写入区块链中的内容，包含系统合约，EOS映射快照
## 5. 更新`privkeys.keys`
这个文件中需要包含你的公钥地址对应的私钥

## 6. 发布你的配置文件

    eos-bios publish

## 7. 按你的实际安装环境更新`hook_init.sh`和`hook_join_network.sh`
示例配置文件使用的是Docker，你可以替换成systemd或者Kubernetes
在`hook_join_network.sh`文件中，你需要添加你的公钥私钥地址(publick key & private key)

## 8. 启动!
执行以下脚本 

    eos-bios orchestrate
    
等待boot运行

# Free Bonus

## 9. 执行以后任何一条命令看看会发生啥🤣:
* `eos-bios discover`
* `eos-bios list`
* `eos-bios discover --serve`


# peers网络

`seed_network_peers`中内容如下所示:

```
seed_network_peers:
- account: eosexample
  comment: "They are good"
  weight: 10  # Weights are between 0 and 100 (INT value)
- account: eosmore
  comment: "They are better"
  weight: 20
```

代表着你愿意与`eosexample`(10%的权重)和`eosmore`(20%的权重)一起启动网络，`eos-bios`将会基于这些信息计算出相应的图表

你可以在种子网络中找到所有的账户名

# 怎么评估我应该添加哪个peers节点

1. 他们是否完全理解启动流程，是否理解启动一条主网所需要执行的所有步骤。(boot_sequence.yaml中需要执行的步骤)

2. 他们是否知道如何编译系统合约，且保证这些被提议的合约中代码都是合法的。(target_contents中的内容)

3. 他们是否知道如何验证snapshot.csv文件是正常的，与Ethereum以太坊的快照文件内容完全一致。(snapshot.csv中的内容)

4. 他们是否能够正确的启动网络，且曾经练习过承担BIOS Boot node这个角色。

5. 他们是否能够正确的启动节点，且曾经练习过`join`加入网络。


为什么需要考虑这些因素呢？这是因为跟`eos-bios`底层设计有关，你的得票数决定了你承担什么样的角色，基于你的角色，你所做出的每一个决定都跟社区息息相关。

The reason for those is because of the design of `eos-bios` .. votes
determine who gets which role, and based on the role you have, you
have critical decisions to make and the community relies on you for
the critical things, in the order above.
