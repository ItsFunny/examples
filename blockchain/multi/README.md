# 多组织的区块链网络

- org1: 11051
- org2: 12051
- org3: 13051
- org4: 14051
- org5: 15051


- 使用方法:
    - 先生成crypto-config cryptogen_multi generate --config=crypto-config.yaml
    - 再生成 tx文件: ./prepare.sh 2 // 2代表2个channel
    - 再启动docker 网络,并且加入channel,安装链码,实例化链码
        - ./restart.sh  2  5 2 
        - 代表2个channel,5个组织,每个组织2个peer
        
- network.py: 用于生成一系列的配置文件,如configtx.yaml或者是crypto-conf.yaml,支持输入多个channel,org,peer