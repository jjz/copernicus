package lutxo

import (
	"github.com/copernet/copernicus/model/outpoint"
	"github.com/copernet/copernicus/model/utxo"
	"github.com/copernet/copernicus/util"
)

func AccessByTxid(coinsCache utxo.CacheView, hash *util.Hash) *utxo.Coin {
	out := outpoint.OutPoint{Hash: *hash, Index: 0}
	for int(out.Index) < 11000 { // todo modify to be precise
		alternate := coinsCache.GetCoin(&out)
		if !alternate.IsSpent() {
			return alternate
		}
		out.Index++
	}
	return nil
}
