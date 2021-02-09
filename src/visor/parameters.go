package visor

/*
* CODE GENERATED AUTOMATICALLY WITH FIBER COIN CREATOR
* AVOID EDITING THIS MANUALLY
*/

const (
	// MaxCoinSupply is the maximum supply of coins
	MaxCoinSupply uint64 = 500000000
	// DistributionAddressesTotal is the number of distribution addresses
	DistributionAddressesTotal uint64 = 100
	// DistributionAddressInitialBalance is the initial balance of each distribution address
	DistributionAddressInitialBalance uint64 = MaxCoinSupply / DistributionAddressesTotal
	// InitialUnlockedCount is the initial number of unlocked addresses
	InitialUnlockedCount uint64 = 25
	// UnlockAddressRate is the number of addresses to unlock per unlock time interval
	UnlockAddressRate uint64 = 5
	// UnlockTimeInterval is the distribution address unlock time interval, measured in seconds
	// Once the InitialUnlockedCount is exhausted,
	// UnlockAddressRate addresses will be unlocked per UnlockTimeInterval
	UnlockTimeInterval uint64 = 31536000 // in seconds
	// MaxDropletPrecision represents the decimal precision of droplets
	MaxDropletPrecision uint64 = 3
	//DefaultMaxBlockSize is max block size
	DefaultMaxBlockSize int = 32768 // in bytes
)

var distributionAddresses = [DistributionAddressesTotal]string{
   "2bba9U67T226V8m4xkXc7wxjZHhRBDCiShn",
   "pih9veNF5YWxMk9iTXaarwUmoCXo7touTc",
   "DNghFmd78X1rEEGJHLpwtrrpX6egSFp6i9",
   "CaXnuRJ5at8Kq67SfNHMXxXPcadg1pBUEZ",
   "qCJZJU38V32EVuXEccJ49xPEuw3m2BKWY9",
   "2hmzMohNF3wGTFqje1Hpe91LyRJv64QrzfN",
   "7QzdJeKboWGv4h3J3qkPZyXb2xbpw7ys2U",
   "2kxZNz1JinK2dsGMbhb4Byp8uNBLmUMDthJ",
   "bHYoLPsJBkZUq1AJzGngNn3yjmG6vcEm5c",
   "2fNuyEDqbR6jVW7TH3NJj35XRr1xMSsi9Ur",
   "5DEvoDEHJRW9bnEG2FhxtmKGkF7JtitUsX",
   "gWf335oNCgVgBSXywLD9GLZr6f4vLqf6MG",
   "254bkJM4SXeCkP4KtwvASFai4hY6vucLk2D",
   "2BQbGiehMChyLXQ8Bt5eJ2fEXsXMzsAmC6W",
   "QQVJFgP7qSAYohrpKeat84wZzFvVVhAX2k",
   "2dKB7zs232hazrynkZdwwdjqVyPg4C6C6eG",
   "ofodHr6ZZWzYwfnoV7wKP5WmTqGRfbgqpG",
   "2BXFzZUB3CV83AVJfMPzsypiBSm6w2rqd6U",
   "YNwaQVyDCZjoqucGCKKqKrHxiGQY6EzE5",
   "2CacZQ2a4VhHLdT6myDbVDxvbR8qPJ7Dghw",
   "2JAtxTubV6zmQmfeacvrNZaRDreqZgQUJKH",
   "2T1HyMmxBTV9CUCnunSq7dR8bFyCY7PoD17",
   "269mkepyCBxijfXzSLCDJusQAX32NXRYN1V",
   "2jD4b8CM6EK4dSRAUtNEknQ3jBntoBbqjbT",
   "4ny989wDsimkfme6jHzCf9f5gRXtvyy1WS",
   "TaNy6jHyxpfmJQH2YpCJnQwHZeiG7DmBeJ",
   "2kHVsbp6XxPL8S4M14RkkUGBSPQRXJs8bQQ",
   "xUiK2zCypJHwkaBsC3j82DUKn4GgdLZJZF",
   "249wpV1bf1D2gBbrmSsvKNvVnZssDkuivnS",
   "owUeWD5pjp1v1tAp2KGiNwDXPsjymk9bmT",
   "ENPeNvD7jRWc4rxfh2SeLvqGKEgN5P83da",
   "24t5WN95a71Dkm5BwXK9Yss47f5aPqXQRxJ",
   "GtzNDjPQrUXyNX8EHMwCDpVavDRzJrYbCg",
   "t3XFkVvNbf2hcB1jF8RzaxDZxPUUP2VTSU",
   "mFRPwEZtPYVFGSx133jVH95GbbPcfSmBUq",
   "dj48mNXfdgcvxjRHWJGA14HGjVHwaW3qYJ",
   "2V32QbDYpXeWiTj7neD8B3GeXqJhAUX7t4u",
   "T8Au47X8dQHhRCG4VNZ2jR9tMb7wDG4NFX",
   "Y3KBrGWf4ZtziYsx2AQ4XScZfTYoSbB3Fv",
   "RDn97MrGwtw5dnAU584QqXiuTNbPqKquvA",
   "2AQeoAFQEQEo7TEVrGH7w8FQCXDsP3vq76t",
   "2eUDPV3YB37EP7nwtKKM6TxhthfabaUiKbM",
   "yYbjeiaaXLMFvfzfYTRkrgC7V3avMNcPMP",
   "cxfFFvwHx4sZkHFPk87oWozDperAgg7bS",
   "FvCpUNRRRkNFAgpXupfbyXc4caiVSYHQ4F",
   "2ZgaRnEK5siC4rEohfFqvpo1p5CzW6w4Pwr",
   "aDn4PZPjqnWJc39fLt4BJ2xBy67njfVoD5",
   "gyzi71ibzi9HGmpJh6kR6AxXfp25ioz8Uh",
   "2VxqPE99aoEFgiX6vMRwny1k1rd54oX7h95",
   "Sxmgv63Fx4xANi9v7aebYrC5YtwKPaYpPv",
   "2JSSL2vNkCYmirSTCmMCnFSRxap5XwJwYzz",
   "PV5rB2Bw7LGgkDJhJ2CWJSf2SgPStobfZa",
   "2MBtuozC8Lp8MjWtgrhQy8CR3ur3BNK4DrA",
   "2X95E8c9aPZTTuas4wbFcJSc12RWGSRcFfz",
   "2jttsjgo2bCFe3AyJjn8P3zwcqbnbgMtyhP",
   "2iguKAxNevyfZd2qDa7BQFgJcT1QQA533Df",
   "2WdMJBy6EgFWHqDS6dP2RWXP1577EKpd3KQ",
   "26Jzr743xy7f4qFH7SdyW5NNnth2998X6oQ",
   "YtpT4bDHvSkta7y9LgUYywBdLCPzdckkuu",
   "2K9T1rCdrN4c9UTEF2M7Cf6tUegjtJzBTwP",
   "T8ChZqJU5EMnhdCuM3n5hT5hYupvkTAg36",
   "grmCJkbEUp76emviFHwBzf1wB2ZLU5psAu",
   "qgtwLM4dY5NcVJGkR3LtnPJeVQu3Y86adG",
   "2mh9tL7YH5aKceVubQs9sGgANLUsXP541dT",
   "22LxEfnG7WB1tjvKjqad5iZyNCWXCCo9ngv",
   "2LQjFMRsuuUAzzm9rtQS7QcvvVGVXSVRSEv",
   "ftHagonzyhLoN7ybZ3xh12VMsnJznFBGBp",
   "2cpFfd7Yo5cjBtRciyyCVFgY2o5UVYn7USs",
   "BXvArUsAcVLAqYNpdxthTfqS6wRxCxtv5n",
   "2Vv2mbTrwEy7bJq7y9kkSE3hZLqLhQFZSfB",
   "23qY2oGN38HuqVPVab8LCCAk821cbk3SXy8",
   "2BcF78wXLc58ryyqE6H4PKPjNgmBmg1GE1H",
   "2byd1vxUDphHm1Pttu1SfsGSVSWKNW7cBJB",
   "2Bcbsq1nHxZiqudKtQr1aS6B4BddjMEY66q",
   "tvmxkkGwiW8KhkD9DbEQsKY4knQw8L9R8o",
   "VpQpACgzJQvuXhSmmaQw2oWAUUuVxpwKZJ",
   "e34JEZh936tSZarJx4sVeL6FaSCrnuYZcy",
   "2ZohaNVaVqSeHVyGJzoLRxoF9wBJiGBoTZT",
   "UewUqPBF2miw32R3HnZgCy3bBFQfBe4oxV",
   "dSb1D3r94Tw5THQm7o6HmLLvNRYpaG2S78",
   "2cy5UpuA5GdqRvsjzxgeexABhdEFZXS2ETE",
   "2gC8az37og4X8if85amiAo7R2RKGjQXNWai",
   "2dnvBwp1ypmAtTRWnfBm5jwfPSLzYnykJZS",
   "9uybe4N4jBkNZa8QJQMbSZ3HDyrqdQRVKT",
   "2cNug9xGLojPhQJxDyP8qU7tkH6QZqWJ7bd",
   "RnEtym6SqAy4k9N9tztvjyx7w6uMP6Xyic",
   "2hdupistadrHPi5wNB8ZVte3drSxCYDvS4p",
   "2YShkWUEjit4ZSDP1kPWsNvQW614f7KwCuW",
   "22gT4Qk6jAzyYRSf6eR2M1QYV3xpTHomUSg",
   "kxpCoWPUvS56CYXUDu2hakPJ4QETinEivs",
   "jykGxvEaJLaH2w5ByewYBFurj25X7wdBvw",
   "GQ6nZJXX9sVNPfpmbrFtu7puF1hJXhwtLJ",
   "p2suuisQHPb8RAUWWigc8aY6QMYLcnXpSA",
   "2eiYcySyKayTdPhxYxkoAFJpdXrXvWt3WsX",
   "2cZYfa6iThZYa5kxwQ4tNTujXZnGMgAnG1j",
   "2VziFDZh5cpm9nJD5XYjNtx96VtpZUuJm9m",
   "2b3qjeYyVVnkwLy2W6mCA1RzS7xU1pm7Ku2",
   "2EXju9xQdnyKzwP7WupBq6A22GmPcg6pNY6",
   "GNvEKbb3Cci97tvdrSr6jixyeFQ9EkesZG",
   "2deM4ZtJey6Cv8wiPQkYGtNTZe8EtfSNo8S",
}
