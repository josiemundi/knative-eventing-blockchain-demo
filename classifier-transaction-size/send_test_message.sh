curl -v "http://localhost:8080" \
-X POST \
-H "X-B3-Flags: 1" \
-H "CE-SpecVersion: 0.2" \
-H "CE-Type: dev.knative.foo.bar" \
-H "CE-Time: 2018-04-05T03:56:24Z" \
-H "CE-ID: 45a8b444-3213-4758-be3f-540bf93f85ff" \
-H "CE-Source: dev.knative.example" \
-H 'Content-Type: application/json' \
-d '{
    "op": "utx",
    "x": {
        "lock_time": 625654,
        "ver": 2,
        "size": 383,
        "inputs": [
            {
                "sequence": 4294967294,
                "prev_out": {
                    "spent": false,
                    "tx_index": 0,
                    "type": 0,
                    "addr": "3ALn5Aadd2APXPQhLaiX7ctmxdZBGz7sVM",
                    "value": 60810927,
                    "n": 0,
                    "script": "a9145ee54013bd0baf77b5545ec467bec23bfd26fe6387"
                },
                "script": "160014c68942defb4274c2832991bba5f09255dcafeb69"
            }
        ],
        "time": 1586723567,
        "tx_index": 0,
        "vin_sz": 1,
        "hash": "991d8b46f49114ffd07c14b08a5f86f615156768693ba0b7d06458e5e9db8fcc",
        "vout_sz": 6,
        "relayed_by": "",
        "out": [
            {
                "spent": false,
                "tx_index": 0,
                "type": 0,
                "addr": "321qMzht8LbLnKPhPdNviywkiJwaz6wYrS",
                "value": 54100,
                "n": 0,
                "script": "a914038f16561299c4671c260e957a0111f9ee1dd44787"
            },
            {
                "spent": false,
                "tx_index": 0,
                "type": 0,
                "addr": "1CRnC2eMhVzmRiVQoMxyxQFB9vT5KvYpiJ",
                "value": 1765743,
                "n": 1,
                "script": "76a9147d5916ebe4305b2fff927414e7c7c196498fbcae88ac"
            },
            {
                "spent": false,
                "tx_index": 0,
                "type": 0,
                "addr": "1H1QDYsFqTNsEne9tfn6zujadGvaDZUCS1",
                "value": 14087444,
                "n": 2,
                "script": "76a914af956480c081aef23ab24766e66cb7a1de96763a88ac"
            },
            {
                "spent": false,
                "tx_index": 0,
                "type": 0,
                "addr": "36G6TruoEmVx73Tvbavf6We8kLoVzvggfS",
                "value": 39806752,
                "n": 3,
                "script": "a9143222050729bb1c6a6cf08db81fe456723df1fc5487"
            },
            {
                "spent": false,
                "tx_index": 0,
                "type": 0,
                "addr": "17JZHamnkSgpjQGsGt8Syh27ZVSdxjNaRp",
                "value": 3365342,
                "n": 4,
                "script": "76a9144522c6e743a17558ff68bb296ce37a505dfd2f2a88ac"
            },
            {
                "spent": false,
                "tx_index": 0,
                "type": 0,
                "addr": "1HRoBA4DPs1zhUy5enGh6Ht6h1VLT1karo",
                "value": 1730596,
                "n": 5,
                "script": "76a914b4328a68b970088a5a810a2739ec45100a58931988ac"
            }
        ]
    }
}'