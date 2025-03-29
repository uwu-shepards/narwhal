# narwhal (unicorn testnet) 🦄

![image](https://github.com/user-attachments/assets/8c0a6698-4926-49f2-8bbe-eb8c470e608f)

### Official RPC
```
https://rpc.uwupunks.com
```

Never forget what they took from us

▄︻デ ══‐一 ❤️

▄︻デ ══‐一 ❤️

▄︻デ ══‐一 ❤️

Censorship is the enemy.

### Install

`curl https://get.ignite.com/uwu-shepards/narwhal! | sudo bash`

`wget https://github.com/CosmWasm/wasmvm/releases/download/v2.2.1/libwasmvm.x86_64.so`

```
sudo mv libwasmvm.x86_64.so /usr/local/lib/
sudo chmod 755 /usr/local/lib/libwasmvm.x86_64.so
sudo ldconfig
ldconfig -p | grep libwasmvm
```

`narwhal init uwupunks`

```
nano ~/.narwhal/config/app.toml
minimum-gas-prices = "0narwhal"
```

`nano ~/.narwhal/config/config.toml `
```
seeds = "6b37a66a808b986a4cb33c0249166010bf0b32bf@p2p.uwupunks.com:26656"
persistent_peers = "6b37a66a808b986a4cb33c0249166010bf0b32bf@p2p.uwupunks.com:26656" 
```

`nano ~/.narwhal/config/genesis.json`

```
{
  "app_name": "narwhald",
  "app_version": "0.1-ecf59183",
  "genesis_time": "2025-03-02T05:03:44.21943Z",
  "chain_id": "narwhal",
  "initial_height": 1,
  "app_hash": null,
  "app_state": {
    "06-solomachine": null,
    "07-tendermint": null,
    "auth": {
      "params": {
        "max_memo_characters": "256",
        "tx_sig_limit": "7",
        "tx_size_cost_per_byte": "10",
        "sig_verify_cost_ed25519": "590",
        "sig_verify_cost_secp256k1": "1000"
      },
      "accounts": [
        {
          "@type": "/cosmos.auth.v1beta1.BaseAccount",
          "address": "narwhal1dpq0ylat08pwrzc3gpe0wxqanjwr0e2rlhn7gx",
          "pub_key": null,
          "account_number": "0",
          "sequence": "0"
        },
        {
          "@type": "/cosmos.auth.v1beta1.BaseAccount",
          "address": "narwhal1zmlwnyf5q8m77ax57ucffn742pukxsfd7khyyf",
          "pub_key": null,
          "account_number": "1",
          "sequence": "0"
        }
      ]
    },
    "authz": {
      "authorization": []
    },
    "bank": {
      "params": {
        "send_enabled": [],
        "default_send_enabled": true
      },
      "balances": [
        {
          "address": "narwhal1zmlwnyf5q8m77ax57ucffn742pukxsfd7khyyf",
          "coins": [
            {
              "denom": "narwhal",
              "amount": "100000000"
            }
          ]
        },
        {
          "address": "narwhal1dpq0ylat08pwrzc3gpe0wxqanjwr0e2rlhn7gx",
          "coins": [
            {
              "denom": "narwhal",
              "amount": "400000000000"
            }
          ]
        }
      ],
      "supply": [
        {
          "denom": "narwhal",
          "amount": "400100000000"
        }
      ],
      "denom_metadata": [],
      "send_enabled": []
    },
    "capability": {
      "index": "1",
      "owners": []
    },
    "circuit": {
      "account_permissions": [],
      "disabled_type_urls": []
    },
    "consensus": null,
    "crisis": {
      "constant_fee": {
        "amount": "1000",
        "denom": "narwhal"
      }
    },
    "distribution": {
      "delegator_starting_infos": [],
      "delegator_withdraw_infos": [],
      "fee_pool": {
        "community_pool": []
      },
      "outstanding_rewards": [],
      "params": {
        "base_proposer_reward": "0.000000000000000000",
        "bonus_proposer_reward": "0.000000000000000000",
        "community_tax": "0.020000000000000000",
        "withdraw_addr_enabled": true
      },
      "previous_proposer": "",
      "validator_accumulated_commissions": [],
      "validator_current_rewards": [],
      "validator_historical_rewards": [],
      "validator_slash_events": []
    },
    "evidence": {
      "evidence": []
    },
    "feegrant": {
      "allowances": []
    },
    "feeibc": {
      "fee_enabled_channels": [],
      "forward_relayers": [],
      "identified_fees": [],
      "registered_counterparty_payees": [],
      "registered_payees": []
    },
    "genutil": {
      "gen_txs": [
        {
          "body": {
            "messages": [
              {
                "@type": "/cosmos.staking.v1beta1.MsgCreateValidator",
                "description": {
                  "moniker": "mynode",
                  "identity": "",
                  "website": "",
                  "security_contact": "",
                  "details": ""
                },
                "commission": {
                  "rate": "0.100000000000000000",
                  "max_rate": "0.200000000000000000",
                  "max_change_rate": "0.010000000000000000"
                },
                "min_self_delegation": "1",
                "delegator_address": "",
                "validator_address": "narwhalvaloper1dpq0ylat08pwrzc3gpe0wxqanjwr0e2r79gjpg",
                "pubkey": {
                  "@type": "/cosmos.crypto.ed25519.PubKey",
                  "key": "vH6KKJ3ZLtXcQfI5649jUEroGIgbvMWXlt7RLTveR4c="
                },
                "value": {
                  "denom": "narwhal",
                  "amount": "200000000000"
                }
              }
            ],
            "memo": "6b37a66a808b986a4cb33c0249166010bf0b32bf@192.168.50.215:26656",
            "timeout_height": "0",
            "extension_options": [],
            "non_critical_extension_options": []
          },
          "auth_info": {
            "signer_infos": [
              {
                "public_key": {
                  "@type": "/cosmos.crypto.secp256k1.PubKey",
                  "key": "A8wMQZ562fyycg3Ry7ioOa6bsR0EZ9Buyxd/SMfJi2wn"
                },
                "mode_info": {
                  "single": {
                    "mode": "SIGN_MODE_DIRECT"
                  }
                },
                "sequence": "0"
              }
            ],
            "fee": {
              "amount": [],
              "gas_limit": "200000",
              "payer": "",
              "granter": ""
            },
            "tip": null
          },
          "signatures": [
            "sEiM1N/SN783yMxUnfGe9osnHwaOfocPaIm10QiIlggiwHSQKdxksQ6nof9N6GCTNXOctHblVuT6KOBhxDV1LQ=="
          ]
        }
      ]
    },
    "gov": {
      "constitution": "",
      "deposit_params": null,
      "deposits": [],
      "params": {
        "burn_proposal_deposit_prevote": false,
        "burn_vote_quorum": false,
        "burn_vote_veto": true,
        "expedited_min_deposit": [
          {
            "amount": "50000000",
            "denom": "narwhal"
          }
        ],
        "expedited_threshold": "0.667000000000000000",
        "expedited_voting_period": "86400s",
        "max_deposit_period": "172800s",
        "min_deposit": [
          {
            "amount": "10000000",
            "denom": "narwhal"
          }
        ],
        "min_deposit_ratio": "0.010000000000000000",
        "min_initial_deposit_ratio": "0.000000000000000000",
        "proposal_cancel_dest": "",
        "proposal_cancel_ratio": "0.500000000000000000",
        "quorum": "0.334000000000000000",
        "threshold": "0.500000000000000000",
        "veto_threshold": "0.334000000000000000",
        "voting_period": "172800s"
      },
      "proposals": [],
      "starting_proposal_id": "1",
      "tally_params": null,
      "votes": [],
      "voting_params": null
    },
    "group": {
      "group_members": [],
      "group_policies": [],
      "group_policy_seq": "0",
      "group_seq": "0",
      "groups": [],
      "proposal_seq": "0",
      "proposals": [],
      "votes": []
    },
    "ibc": {
      "channel_genesis": {
        "ack_sequences": [],
        "acknowledgements": [],
        "channels": [],
        "commitments": [],
        "next_channel_sequence": "0",
        "params": {
          "upgrade_timeout": {
            "height": {
              "revision_height": "0",
              "revision_number": "0"
            },
            "timestamp": "600000000000"
          }
        },
        "receipts": [],
        "recv_sequences": [],
        "send_sequences": []
      },
      "client_genesis": {
        "clients": [],
        "clients_consensus": [],
        "clients_metadata": [],
        "create_localhost": false,
        "next_client_sequence": "0",
        "params": {
          "allowed_clients": [
            "*"
          ]
        }
      },
      "connection_genesis": {
        "client_connection_paths": [],
        "connections": [],
        "next_connection_sequence": "0",
        "params": {
          "max_expected_time_per_block": "30000000000"
        }
      }
    },
    "interchainaccounts": {
      "controller_genesis_state": {
        "active_channels": [],
        "interchain_accounts": [],
        "params": {
          "controller_enabled": true
        },
        "ports": []
      },
      "host_genesis_state": {
        "active_channels": [],
        "interchain_accounts": [],
        "params": {
          "allow_messages": [
            "*"
          ],
          "host_enabled": true
        },
        "port": "icahost"
      }
    },
    "mint": {
      "minter": {
        "annual_provisions": "0.000000000000000000",
        "inflation": "0.130000000000000000"
      },
      "params": {
        "blocks_per_year": "6311520",
        "goal_bonded": "0.670000000000000000",
        "inflation_max": "0.200000000000000000",
        "inflation_min": "0.070000000000000000",
        "inflation_rate_change": "0.130000000000000000",
        "mint_denom": "narwhal"
      }
    },
    "narwhal": {
      "params": {}
    },
    "nft": {
      "classes": [],
      "entries": []
    },
    "params": null,
    "runtime": null,
    "slashing": {
      "missed_blocks": [],
      "params": {
        "downtime_jail_duration": "600s",
        "min_signed_per_window": "0.500000000000000000",
        "signed_blocks_window": "100",
        "slash_fraction_double_sign": "0.050000000000000000",
        "slash_fraction_downtime": "0.010000000000000000"
      },
      "signing_infos": []
    },
    "staking": {
      "delegations": [],
      "exported": false,
      "last_total_power": "0",
      "last_validator_powers": [],
      "params": {
        "bond_denom": "narwhal",
        "historical_entries": 10000,
        "max_entries": 7,
        "max_validators": 100,
        "min_commission_rate": "0.000000000000000000",
        "unbonding_time": "1814400s"
      },
      "redelegations": [],
      "unbonding_delegations": [],
      "validators": []
    },
    "transfer": {
      "denom_traces": [],
      "params": {
        "receive_enabled": true,
        "send_enabled": true
      },
      "port_id": "transfer",
      "total_escrowed": []
    },
    "upgrade": {},
    "vesting": {},
    "wasm": {
      "codes": [],
      "contracts": [],
      "params": {
        "code_upload_access": {
          "addresses": [],
          "permission": "Everybody"
        },
        "instantiate_default_permission": "Everybody"
      },
      "sequences": []
    }
  },
  "consensus": {
    "params": {
      "block": {
        "max_bytes": "22020096",
        "max_gas": "-1"
      },
      "evidence": {
        "max_age_num_blocks": "100000",
        "max_age_duration": "172800000000000",
        "max_bytes": "1048576"
      },
      "validator": {
        "pub_key_types": [
          "ed25519"
        ]
      },
      "version": {
        "app": "0"
      },
      "abci": {
        "vote_extensions_enable_height": "0"
      }
    }
  }
}
```
