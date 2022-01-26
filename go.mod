// module github.com/forbole/bdjuno/v2

// go 1.16

// require (
// 	github.com/cosmos/cosmos-sdk v0.42.9
// 	// github.com/forbole/juno/v2 59279231c43965e770909bac243a4453cf7a8d1c
// 	github.com/forbole/juno/v2 v2.0.0-20220126130945-5f4216098ad5

// 	github.com/go-co-op/gocron v1.11.0
// 	github.com/gogo/protobuf v1.3.3
// 	github.com/jmoiron/sqlx v1.2.1-0.20200324155115-ee514944af4b
// 	github.com/lib/pq v1.10.4
// 	github.com/osmosis-labs/osmosis v1.0.4
// 	github.com/pelletier/go-toml v1.9.4
// 	github.com/proullon/ramsql v0.0.0-20181213202341-817cee58a244
// 	github.com/rs/zerolog v1.26.1
// 	github.com/spf13/cobra v1.3.0
// 	github.com/stretchr/testify v1.7.0
// 	github.com/tendermint/tendermint v0.34.14
// 	google.golang.org/grpc v1.42.0
// 	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
// )

// replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

// replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

// replace github.com/tendermint/tendermint => github.com/forbole/tendermint v0.34.13-0.20210820072129-a2a4af55563d

// // replace github.com/tendermint/tendermint => github.com/osmosis-labs/tendermint v0.34.12-0.20220109173307-59a781894ea7

// // replace github.com/osmosis-labs/osmosis => github.com/MonikaCat/osmosis/v6  5167632

// replace github.com/cosmos/cosmos-sdk => github.com/osmosis-labs/cosmos-sdk v0.44.3-osmo-1
// // v0.43.0-rc3.0.20220120015748-5df6adc097e8
// // replace github.com/cosmos/cosmos-sdk => github.com/osmosis-labs/cosmos-sdk v0.42.5

module github.com/forbole/bdjuno/v2

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.42.9
	github.com/forbole/juno/v2 v2.0.0-20220126130945-5f4216098ad5
	github.com/go-co-op/gocron v1.11.0
	github.com/gogo/protobuf v1.3.3
	github.com/jmoiron/sqlx v1.2.1-0.20200324155115-ee514944af4b
	github.com/lib/pq v1.10.4
	github.com/osmosis-labs/osmosis v1.0.3
	github.com/pelletier/go-toml v1.9.4
	github.com/proullon/ramsql v0.0.0-20181213202341-817cee58a244
	github.com/rs/zerolog v1.26.1
	github.com/spf13/cobra v1.3.0
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.14
	google.golang.org/grpc v1.42.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/tendermint/tendermint => github.com/forbole/tendermint v0.34.13-0.20210820072129-a2a4af55563d

// replace github.com/tendermint/tendermint => github.com/osmosis-labs/tendermint v0.34.12-0.20220109173307-59a781894ea7

replace github.com/cosmos/cosmos-sdk => github.com/osmosis-labs/cosmos-sdk v0.42.5-0.20210630232304-f792e47135c3
