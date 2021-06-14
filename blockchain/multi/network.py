#!/usr/bin/python3
# -*- coding: utf-8 -*

import os
import shutil
import sys
from optparse import OptionParser

from jinja2 import Environment, FileSystemLoader

configtx_template_name = "configtx-template.yaml"
cryptogen_template_name = "crypto-config-template.yaml"
dockercompose_template_name = "docker-compose-all-template.yaml"
dockerHostIp = "172.224.2.2"
artifactName = 'artifacts'
cryptoconfigDirName = 'crypto-config'

parser = OptionParser("build", description='custom description', version='%prog 1.0')


def parse_args(args):
    # 整个清除,删除 cryptoconfig和模板等内容
    parser.add_option('--clean', dest='clean', action='store_true', default=False)
    # stop,停止网络
    parser.add_option('--stop', dest='stop', action='store_true', default=False)
    # 清除网络,并且直接重启,不重新生成
    parser.add_option('--restart', dest='restart', action='store_true', default=True)
    # 宿主机IP
    parser.add_option('--hostip', dest='hostip', action='store', default='172.224.2.2')
    # 几个channel
    parser.add_option('--channelLimit', dest='channelLimit', action='store', default=2)
    # 几个org
    parser.add_option('--orgLimit', dest='orgLimit', action='store', default=5)
    # 每个org几个peer
    parser.add_option('--peerLimit', dest='peerLimit', action='store', default=2)

    # cryptogen generate证书
    parser.add_option('--cryptogen', dest='cryptogen', action='store_true', default=False)
    # configtx生成tx文件
    parser.add_option('--configtx', dest='configtx', action='store_true', default=False)

    #
    parser.add_option('--all', dest='all', action='store_true', default=False)

    # 默认的形式: 既默认如果证书等不存在,会生成,同时如果cryptogen_multi不存在,会make,configtx也同理
    # 然后启动网络,2个channel,5个组织,每个组织2个peer,3个orderer
    parser.add_option('--defaultNetwork', dest='defaultNetwork', action='store', default=True)

    parser.parse_args(args)


def getCurrentPath():
    return os.path.split(os.path.realpath(__file__))[0]


class FabricToolHelper(object):
    def __init__(self, channel_limit, org_limit, peer_limit):
        self.org_limit = org_limit
        self.channel_limit = channel_limit
        self.peer_limit = peer_limit

    def generate_configtx(self, override):
        class ChannelNode(object):
            def __init__(self, index):
                self.index = index

        class OrgPeer(object):
            def __init__(self, host, port):
                self.port = port
                self.host = host

        class ConfigtxNode(object):
            def __init__(self, index):
                self.name = 'Org%dMSP' % index
                self.domain = 'org%d.com' % index
                self.mspId = 'Org%dMSP' % index
                self.index = index
                anchorHost = 'peer0.org%d.com' % index
                port = 10000 + 1000 * index + 51
                self.anchorPeer = OrgPeer(anchorHost, port)

        originConfigtx = os.path.join(getCurrentPath(), 'configtx.yaml')
        if os.path.exists(originConfigtx):
            if override:
                os.remove(originConfigtx)
            else:
                return

        nodes = []
        for i in range(self.org_limit):
            nodes.append(ConfigtxNode(i))
        channels = []
        for i in range(self.channel_limit):
            channels.append(ChannelNode(i))
        path = os.path.join(os.path.split(os.path.realpath(__file__))[0], "templates")
        env = Environment(loader=FileSystemLoader(path, 'utf-8'))
        template = env.get_template(configtx_template_name)
        html = template.render(nodes=nodes, channels=channels)
        self.save_templates(html, override, 'configtx.yaml')
        artifact = os.path.join(getCurrentPath(), artifactName)
        if os.path.exists(artifact):
            if override:
                shutil.rmtree(artifact)
            else:
                return

        os.makedirs(artifact)
        pwd = getCurrentPath()
        print("初始化raft 创始块")
        cmd = 'configtxgen_multi  -configPath ' + pwd + ' --profile OrdererRaftGenesis -channelID sysdemochannel ' \
                                                        ' -outputBlock ' + pwd + '/artifacts/orderer.genesis.block '
        print(cmd)
        if os.system(cmd) != 0:
            print('执行:' + cmd + ',失败,或许configtx不存在,开始make configtxgen')
            self.make_tools('configtxgen')
            os.chdir(pwd)
            if os.system(cmd) != 0:
                print('再次执行依旧失败')
                exit(-1)
        print("初始化solo 创始块")
        cmd = 'configtxgen_multi -configPath ' + pwd + ' --profile OrdererSoloGenesis -channelID sysdemochannel -outputBlock ' + pwd + '/artifacts/orderer.solo.genesis.block'
        print(cmd)
        if os.system(cmd) != 0:
            print('执行失败:' + cmd)
            exit(-1)

        for i in range(self.channel_limit):
            channelName = 'DemoChannel%d' % i
            txFileName = 'demochannel%d.tx' % i
            asChannelName = 'demochannel%d' % i

            cmd = 'configtxgen_multi  -configPath %s --profile %s  -outputCreateChannelTx %s/artifacts/%s -channelID %s' % (
                pwd, channelName, pwd, txFileName, asChannelName)
            print(cmd)
            if os.system(cmd) != 0:
                print('执行cmd:' + cmd + ",失败")
                exit(-1)

            for j in range(self.org_limit):
                orgMsp = 'Org%dMSP' % j
                anchorTxName = 'org%dmspanchors.tx' % j
                orgCmd = ' configtxgen_multi  -configPath ' + pwd + ' --profile ' + channelName + ' -outputAnchorPeersUpdate ' + pwd + '/artifacts/' + anchorTxName + ' -channelID ' + asChannelName + ' -asOrg ' + orgMsp
                print(orgCmd)
                if os.system(orgCmd) != 0:
                    print("执行orgCmd失败,cmd为:" + orgCmd)
                    exit(-1)

    def make_tools(self, name):
        gopath = os.getenv('GOPATH')
        makePath = os.path.join(gopath, 'src', 'github.com', 'hyperledger', 'fabric')
        os.chdir(makePath)
        binDir = os.path.join(makePath, 'build', 'bin')
        newName = name + '_multi'
        binPath = os.path.join(binDir, newName)
        prevBinPath = os.path.join(binDir, name)
        if os.path.exists(binDir):
            if os.path.exists(binPath):
                return
            elif os.path.exists(prevBinPath):
                os.rename(prevBinPath, binPath)
                return

        cmd = 'make %s ' % name
        if os.system(cmd) != 0:
            print(cmd + ',执行失败')
            exit(-1)
        os.rename(prevBinPath, binPath)

    def generate_cryptoconfig(self, override):
        class CryptoOrgNode(object):
            def __init__(self, name, domain):
                self.name = name
                self.domain = domain

        pwd = getCurrentPath()
        originCrypto = os.path.join(pwd, 'crypto-config.yaml')
        if os.path.exists(originCrypto):
            if override:
                os.remove(originCrypto)
            else:
                return

        nodes = []
        for i in range(self.org_limit):
            name = "Org%d" % i
            domain = "org%d.com" % i
            nodes.append(CryptoOrgNode(name, domain))

        path = os.path.join(getCurrentPath(), "templates")
        env = Environment(loader=FileSystemLoader(path, 'utf-8'))
        template = env.get_template(cryptogen_template_name)
        html = template.render(nodes=nodes)
        self.save_templates(html, override, 'crypto-config.yaml')
        cryptoconfigDir = os.path.join(pwd, cryptoconfigDirName)
        if os.path.exists(cryptoconfigDir):
            if override:
                shutil.rmtree(cryptoconfigDir)
            else:
                return

        cmd = 'cryptogen_multi generate --config=%s --output=%s ' % (
            os.path.join(pwd, 'crypto-config.yaml'), os.path.join(pwd, 'crypto-config'))
        if os.system(cmd) != 0:
            print('执行:' + cmd + ',失败,或许configtx不存在,开始make configtxgen')
            self.make_tools('cryptogen')
            os.chdir(pwd)
            if os.system(cmd) != 0:
                print('再次执行依旧失败')
                exit(-1)

    def generate_dockercompose(self, override, startIfExist):

        class DockerOrgInfo(object):
            def __init__(self, orgIndex, peerCount):
                peerNodes = []
                for i in range(peerCount):
                    last = (i == (peerCount - 1))
                    peerNodes.append(DockerOrgNodeInfo(orgIndex, i, last))
                self.name = 'org%d.com' % orgIndex
                self.mspId = 'Org%dMSP' % orgIndex
                self.nodes = peerNodes

        class DockerOrgNodeInfo(object):
            def __init__(self, orgIndex, peerIndex, is_last):
                basePort = 10000 + 1000 * orgIndex
                self.name = 'peer%d.org%d.com' % (peerIndex, orgIndex)
                self.port = basePort + 51 + (10 * peerIndex)
                self.chaincodePort = self.port + 1

                if is_last:
                    bootPort = basePort + 51 + (10 * (peerIndex - 1))
                    self.bootstrap = 'peer%d.org%d.com:%d' % (peerIndex - 1, orgIndex, bootPort)
                else:
                    bootPort = basePort + 51 + (10 * (peerIndex + 1))
                    self.bootstrap = 'peer%d.org%s.com:%d' % (peerIndex + 1, orgIndex, bootPort)

        orgs = []
        for i in range(self.org_limit):
            orgs.append(DockerOrgInfo(i, self.peer_limit))

        pwd = getCurrentPath()
        path = os.path.join(pwd, "templates")
        env = Environment(loader=FileSystemLoader(path, 'utf-8'))
        template = env.get_template(dockercompose_template_name)
        html = template.render(orgs=orgs, dockerHostIp=dockerHostIp)
        self.save_templates(html, override, 'docker-compose-all.yaml')
        if startIfExist:
            pwd = getCurrentPath()
            # 判断crypto-config在material是否存在

            cryptogenDir = os.path.join(pwd, 'crypto-config')
            if not os.path.exists(cryptogenDir) or override:
                self.generate_cryptoconfig(True)
            artificatsDir = os.path.join(pwd, 'artifacts')
            if not os.path.exists(artificatsDir) or override:
                self.generate_configtx(True)

            print('启动docker成功,生成脚本文件')
            restartShFile = os.path.join(getCurrentPath(), 'restart.sh')
            if os.path.exists(restartShFile):
                os.remove(restartShFile)
            fo = open(restartShFile, 'w')
            cmd = '''#!/usr/bin/env bash
./clean.sh
docker-compose -f docker-compose-order.yaml -f docker-compose-all.yaml   --project-name multiorganization up -d
sleep 5
docker exec -it cli ./scripts/joinchannel.sh  %d %d %d
docker cp cli:/opt/gopath/src/github.com/hyperledger/fabric/peer/demochannel1.block ../
            ''' % (
                self.channel_limit, self.org_limit, self.peer_limit)
            # cmd = 'sh -c  ' + restartShFile + ' %d %d %d' % (
            #     self.channel_limit, self.org_limit, self.peer_limit)
            fo.writelines(cmd)
            fo.close()
            os.system('chmod +x %s' % restartShFile)
            print(cmd)

    def save_templates(self, html, override, file_name):
        print(html)
        current_dir = getCurrentPath()
        if not os.path.exists(current_dir):
            os.mkdir(current_dir)
            os.chdir(current_dir)
        elif override:
            path = os.path.join(current_dir, file_name)
            if os.path.exists(path):
                os.remove(path)
            fo = open(path, "w")
            fo.writelines(html)
            fo.close()
            os.chdir(current_dir)
            print(f"create {file_name}:", "完成")


class Network(object):
    helper = FabricToolHelper(2, 5, 2)

    def __init__(self, tool_helper):
        self.helper = tool_helper

    def networkAsDefault(self, override):
        # 判断docker-compose文件是否存在
        pwd = getCurrentPath()
        self.helper.generate_dockercompose(override, True)

    def clean(self):
        pwd = getCurrentPath()
        materiDir = os.path.join(pwd, pwd)
        if os.path.exists(materiDir):
            os.removedirs(materiDir)

        print('clean successfully')


if __name__ == '__main__':
    parse_args(sys.argv[1:])
    v = parser.values
    dockerHostIp = v.hostip
    org_limit = v.orgLimit
    channel_limit = v.channelLimit
    peer_limit = v.peerLimit
    helper = FabricToolHelper(channel_limit, org_limit, peer_limit)
    networker = Network(helper)
    if v.defaultNetwork:
        networker.networkAsDefault(True)
