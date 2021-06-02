export LOTUS_SKIP_GENESIS_CHECK=_yes_
export LOTUS_PATH=~/.lotusDevnet
export LOTUS_MINER_PATH=~/.lotusminerDevnet
myFile="devgen.car"
if [ ! -x "$myPath"]; then
echo "devgen.car 不存在 新建中..."
	./lotus daemon --lotus-make-genesis=devgen.car --genesis-template=localnet.json --bootstrap=false
else
	echo "devgen.car 存在 daemon启动中..."
	./lotus daemon --genesis=devgen.car --genesis-template=localnet.json --bootstrap=false
fi

