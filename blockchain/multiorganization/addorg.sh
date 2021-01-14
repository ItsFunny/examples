#!/bin/bash
# 第一个参数为组织名称,大小写敏感
# 第2个参数为要签名的组织的个数
# 第3个参数为channel名称
# 使用: ./addorg.sh Org6 5 demochannel1
cd ${PWD}/scripts/addneworg/
./step_1.sh ${1}
docker exec -it cli ./scripts/addneworg/step_2.sh ${2} ${3}