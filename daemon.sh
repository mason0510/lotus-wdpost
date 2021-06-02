export LOTUS_SKIP_GENESIS_CHECK=_yes_
export LOTUS_PATH=~/.lotusDevnet
export LOTUS_MINER_PATH=~/.lotusminerDevnet
myFile="devgen.car"
if [ ! -f "$myFile" ];then
  echo "文件不存在"
./lotus-seed genesis new localnet.json
./lotus-seed genesis add-miner localnet.json ~/.genesis-sectors/pre-seal-t01000.json
 ./lotus daemon --lotus-make-genesis=devgen.car --genesis-template=localnet.json --bootstrap=false
  else
  echo "存在"
./lotus daemon --genesis=devgen.car --genesis-template=localnet.json --bootstrap=false
fi
