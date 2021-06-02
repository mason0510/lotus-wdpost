export LOTUS_SKIP_GENESIS_CHECK=_yes_ 
nohup ./lotus daemon --genesis=devgen.car --genesis-template=localnet.json >>lotus.log 2>&1 &
tail -f lotus.log
