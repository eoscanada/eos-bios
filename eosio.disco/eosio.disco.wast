(module
 (type $FUNCSIG$viji (func (param i32 i64 i32)))
 (type $FUNCSIG$vijii (func (param i32 i64 i32 i32)))
 (type $FUNCSIG$v (func))
 (type $FUNCSIG$vjj (func (param i64 i64)))
 (type $FUNCSIG$vj (func (param i64)))
 (type $FUNCSIG$ijjjj (func (param i64 i64 i64 i64) (result i32)))
 (type $FUNCSIG$vii (func (param i32 i32)))
 (type $FUNCSIG$j (func (result i64)))
 (type $FUNCSIG$i (func (result i32)))
 (type $FUNCSIG$iiii (func (param i32 i32 i32) (result i32)))
 (type $FUNCSIG$ijjjjii (func (param i64 i64 i64 i64 i32 i32) (result i32)))
 (type $FUNCSIG$vi (func (param i32)))
 (type $FUNCSIG$iii (func (param i32 i32) (result i32)))
 (import "env" "abort" (func $abort))
 (import "env" "action_data_size" (func $action_data_size (result i32)))
 (import "env" "current_receiver" (func $current_receiver (result i64)))
 (import "env" "db_find_i64" (func $db_find_i64 (param i64 i64 i64 i64) (result i32)))
 (import "env" "db_get_i64" (func $db_get_i64 (param i32 i32 i32) (result i32)))
 (import "env" "db_store_i64" (func $db_store_i64 (param i64 i64 i64 i64 i32 i32) (result i32)))
 (import "env" "db_update_i64" (func $db_update_i64 (param i32 i64 i32 i32)))
 (import "env" "eosio_assert" (func $eosio_assert (param i32 i32)))
 (import "env" "eosio_exit" (func $eosio_exit (param i32)))
 (import "env" "memcpy" (func $memcpy (param i32 i32 i32) (result i32)))
 (import "env" "memmove" (func $memmove (param i32 i32 i32) (result i32)))
 (import "env" "memset" (func $memset (param i32 i32 i32) (result i32)))
 (import "env" "now" (func $now (result i32)))
 (import "env" "read_action_data" (func $read_action_data (param i32 i32) (result i32)))
 (import "env" "require_auth" (func $require_auth (param i64)))
 (import "env" "require_auth2" (func $require_auth2 (param i64 i64)))
 (table 3 3 anyfunc)
 (elem (i32.const 0) $__wasm_nullptr $_ZN5eosio5disco11updtgenesisEyNSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEENS1_6vectorIS7_NS5_IS7_EEEE $_ZN5eosio5disco9updtdiscoEyRKNS_14discovery_fileE)
 (memory $0 1)
 (data (i32.const 4) "\f0b\00\00")
 (data (i32.const 16) "object passed to iterator_to is not in multi_index\00")
 (data (i32.const 80) "cannot create objects in table of another contract\00")
 (data (i32.const 144) "cannot pass end iterator to modify\00")
 (data (i32.const 192) "object passed to modify is not in multi_index\00")
 (data (i32.const 240) "cannot modify objects in table of another contract\00")
 (data (i32.const 304) "updater cannot change primary key when modifying an object\00")
 (data (i32.const 368) "write\00")
 (data (i32.const 384) "error reading iterator\00")
 (data (i32.const 416) "read\00")
 (data (i32.const 432) "get\00")
 (data (i32.const 8848) "malloc_from_freed was designed to only be called after _heap was completely allocated\00")
 (export "memory" (memory $0))
 (export "_ZeqRK11checksum256S1_" (func $_ZeqRK11checksum256S1_))
 (export "_ZN5eosio12require_authERKNS_16permission_levelE" (func $_ZN5eosio12require_authERKNS_16permission_levelE))
 (export "_ZN5eosio5disco9updtdiscoEyRKNS_14discovery_fileE" (func $_ZN5eosio5disco9updtdiscoEyRKNS_14discovery_fileE))
 (export "_ZN5eosio5disco11updtgenesisEyNSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEENS1_6vectorIS7_NS5_IS7_EEEE" (func $_ZN5eosio5disco11updtgenesisEyNSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEENS1_6vectorIS7_NS5_IS7_EEEE))
 (export "apply" (func $apply))
 (export "memcmp" (func $memcmp))
 (export "malloc" (func $malloc))
 (export "free" (func $free))
 (func $_ZeqRK11checksum256S1_ (param $0 i32) (param $1 i32) (result i32)
  (i32.eqz
   (call $memcmp
    (get_local $0)
    (get_local $1)
    (i32.const 32)
   )
  )
 )
 (func $_ZN5eosio12require_authERKNS_16permission_levelE (param $0 i32)
  (call $require_auth2
   (i64.load
    (get_local $0)
   )
   (i64.load offset=8
    (get_local $0)
   )
  )
 )
 (func $_ZN5eosio5disco9updtdiscoEyRKNS_14discovery_fileE (type $FUNCSIG$viji) (param $0 i32) (param $1 i64) (param $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i32)
  (local $6 i32)
  (local $7 i32)
  (local $8 i32)
  (i32.store offset=4
   (i32.const 0)
   (tee_local $8
    (i32.sub
     (i32.load offset=4
      (i32.const 0)
     )
     (i32.const 64)
    )
   )
  )
  (i64.store offset=16
   (get_local $8)
   (get_local $1)
  )
  (call $require_auth
   (get_local $1)
  )
  (block $label$0
   (br_if $label$0
    (i32.eq
     (tee_local $7
      (i32.load
       (i32.add
        (get_local $0)
        (i32.const 36)
       )
      )
     )
     (tee_local $3
      (i32.load
       (i32.add
        (get_local $0)
        (i32.const 32)
       )
      )
     )
    )
   )
   (set_local $6
    (i32.add
     (get_local $7)
     (i32.const -24)
    )
   )
   (set_local $4
    (i32.sub
     (i32.const 0)
     (get_local $3)
    )
   )
   (loop $label$1
    (br_if $label$0
     (i64.eq
      (i64.load
       (i32.load
        (get_local $6)
       )
      )
      (get_local $1)
     )
    )
    (set_local $7
     (get_local $6)
    )
    (set_local $6
     (tee_local $5
      (i32.add
       (get_local $6)
       (i32.const -24)
      )
     )
    )
    (br_if $label$1
     (i32.ne
      (i32.add
       (get_local $5)
       (get_local $4)
      )
      (i32.const -24)
     )
    )
   )
  )
  (set_local $5
   (i32.add
    (get_local $0)
    (i32.const 8)
   )
  )
  (block $label$2
   (block $label$3
    (block $label$4
     (block $label$5
      (block $label$6
       (br_if $label$6
        (i32.eq
         (get_local $7)
         (get_local $3)
        )
       )
       (call $eosio_assert
        (i32.eq
         (i32.load offset=284
          (tee_local $6
           (i32.load
            (i32.add
             (get_local $7)
             (i32.const -24)
            )
           )
          )
         )
         (get_local $5)
        )
        (i32.const 16)
       )
       (set_local $1
        (i64.load
         (get_local $0)
        )
       )
       (br_if $label$5
        (get_local $6)
       )
       (br $label$3)
      )
      (br_if $label$4
       (i32.le_s
        (tee_local $6
         (call $db_find_i64
          (i64.load
           (i32.add
            (get_local $0)
            (i32.const 8)
           )
          )
          (i64.load
           (i32.add
            (get_local $0)
            (i32.const 16)
           )
          )
          (i64.const 5454011350977085440)
          (get_local $1)
         )
        )
        (i32.const -1)
       )
      )
      (call $eosio_assert
       (i32.eq
        (i32.load offset=284
         (tee_local $6
          (call $_ZNK5eosio11multi_indexILy5454011350977085440ENS_5disco13discovery_rowEJEE31load_object_by_primary_iteratorEl
           (get_local $5)
           (get_local $6)
          )
         )
        )
        (get_local $5)
       )
       (i32.const 16)
      )
      (set_local $1
       (i64.load
        (get_local $0)
       )
      )
     )
     (i32.store offset=32
      (get_local $8)
      (get_local $2)
     )
     (call $eosio_assert
      (i32.const 1)
      (i32.const 144)
     )
     (call $_ZN5eosio11multi_indexILy5454011350977085440ENS_5disco13discovery_rowEJEE6modifyIZNS1_9updtdiscoEyRKNS_14discovery_fileEE3$_1EEvRKS2_yOT_
      (get_local $5)
      (get_local $6)
      (get_local $1)
      (i32.add
       (get_local $8)
       (i32.const 32)
      )
     )
     (br $label$2)
    )
    (set_local $1
     (i64.load
      (get_local $0)
     )
    )
   )
   (i32.store offset=12
    (get_local $8)
    (get_local $2)
   )
   (i32.store offset=8
    (get_local $8)
    (i32.add
     (get_local $8)
     (i32.const 16)
    )
   )
   (i64.store offset=56
    (get_local $8)
    (get_local $1)
   )
   (call $eosio_assert
    (i64.eq
     (i64.load
      (i32.add
       (get_local $0)
       (i32.const 8)
      )
     )
     (call $current_receiver)
    )
    (i32.const 80)
   )
   (i32.store offset=32
    (get_local $8)
    (get_local $5)
   )
   (i32.store offset=36
    (get_local $8)
    (i32.add
     (get_local $8)
     (i32.const 8)
    )
   )
   (i32.store offset=40
    (get_local $8)
    (i32.add
     (get_local $8)
     (i32.const 56)
    )
   )
   (i32.store offset=56
    (tee_local $6
     (call $_Znwj
      (i32.const 296)
     )
    )
    (i32.const 0)
   )
   (i64.store offset=48 align=4
    (get_local $6)
    (i64.const 0)
   )
   (i32.store offset=72
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=76
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=80
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=120
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=124
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=128
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=144
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=188
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=192
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=196
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=200
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=204
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=208
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=216
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=220
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=224
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=228
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=232
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=236
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=244
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=248
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=252
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=256
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=260
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=264
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=268
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=272
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=276
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=284
    (get_local $6)
    (get_local $5)
   )
   (call $_ZZN5eosio11multi_indexILy5454011350977085440ENS_5disco13discovery_rowEJEE7emplaceIZNS1_9updtdiscoEyRKNS_14discovery_fileEE3$_0EENS3_14const_iteratorEyOT_ENKUlRSA_E_clINS3_4itemEEEDaSC_
    (i32.add
     (get_local $8)
     (i32.const 32)
    )
    (get_local $6)
   )
   (i32.store offset=48
    (get_local $8)
    (get_local $6)
   )
   (i64.store offset=32
    (get_local $8)
    (tee_local $1
     (i64.load
      (get_local $6)
     )
    )
   )
   (i32.store offset=28
    (get_local $8)
    (tee_local $7
     (i32.load offset=288
      (get_local $6)
     )
    )
   )
   (block $label$7
    (block $label$8
     (br_if $label$8
      (i32.ge_u
       (tee_local $5
        (i32.load
         (tee_local $4
          (i32.add
           (get_local $0)
           (i32.const 36)
          )
         )
        )
       )
       (i32.load
        (i32.add
         (get_local $0)
         (i32.const 40)
        )
       )
      )
     )
     (i64.store offset=8
      (get_local $5)
      (get_local $1)
     )
     (i32.store offset=16
      (get_local $5)
      (get_local $7)
     )
     (i32.store offset=48
      (get_local $8)
      (i32.const 0)
     )
     (i32.store
      (get_local $5)
      (get_local $6)
     )
     (i32.store
      (get_local $4)
      (i32.add
       (get_local $5)
       (i32.const 24)
      )
     )
     (br $label$7)
    )
    (call $_ZNSt3__16vectorIN5eosio11multi_indexILy5454011350977085440ENS1_5disco13discovery_rowEJEE8item_ptrENS_9allocatorIS6_EEE24__emplace_back_slow_pathIJNS_10unique_ptrINS5_4itemENS_14default_deleteISC_EEEERyRlEEEvDpOT_
     (i32.add
      (get_local $0)
      (i32.const 32)
     )
     (i32.add
      (get_local $8)
      (i32.const 48)
     )
     (i32.add
      (get_local $8)
      (i32.const 32)
     )
     (i32.add
      (get_local $8)
      (i32.const 28)
     )
    )
   )
   (set_local $6
    (i32.load offset=48
     (get_local $8)
    )
   )
   (i32.store offset=48
    (get_local $8)
    (i32.const 0)
   )
   (br_if $label$2
    (i32.eqz
     (get_local $6)
    )
   )
   (drop
    (call $_ZN5eosio14discovery_fileD2Ev
     (i32.add
      (get_local $6)
      (i32.const 8)
     )
    )
   )
   (call $_ZdlPv
    (get_local $6)
   )
  )
  (i32.store offset=4
   (i32.const 0)
   (i32.add
    (get_local $8)
    (i32.const 64)
   )
  )
 )
 (func $_ZNK5eosio11multi_indexILy5454011350977085440ENS_5disco13discovery_rowEJEE31load_object_by_primary_iteratorEl (param $0 i32) (param $1 i32) (result i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i64)
  (local $6 i32)
  (local $7 i32)
  (local $8 i32)
  (local $9 i32)
  (set_local $8
   (tee_local $9
    (i32.sub
     (i32.load offset=4
      (i32.const 0)
     )
     (i32.const 48)
    )
   )
  )
  (i32.store offset=4
   (i32.const 0)
   (get_local $9)
  )
  (block $label$0
   (br_if $label$0
    (i32.eq
     (tee_local $7
      (i32.load
       (i32.add
        (get_local $0)
        (i32.const 28)
       )
      )
     )
     (tee_local $2
      (i32.load offset=24
       (get_local $0)
      )
     )
    )
   )
   (set_local $3
    (i32.sub
     (i32.const 0)
     (get_local $2)
    )
   )
   (set_local $6
    (i32.add
     (get_local $7)
     (i32.const -24)
    )
   )
   (loop $label$1
    (br_if $label$0
     (i32.eq
      (i32.load
       (i32.add
        (get_local $6)
        (i32.const 16)
       )
      )
      (get_local $1)
     )
    )
    (set_local $7
     (get_local $6)
    )
    (set_local $6
     (tee_local $4
      (i32.add
       (get_local $6)
       (i32.const -24)
      )
     )
    )
    (br_if $label$1
     (i32.ne
      (i32.add
       (get_local $4)
       (get_local $3)
      )
      (i32.const -24)
     )
    )
   )
  )
  (block $label$2
   (block $label$3
    (br_if $label$3
     (i32.eq
      (get_local $7)
      (get_local $2)
     )
    )
    (set_local $6
     (i32.load
      (i32.add
       (get_local $7)
       (i32.const -24)
      )
     )
    )
    (br $label$2)
   )
   (call $eosio_assert
    (i32.xor
     (i32.shr_u
      (tee_local $6
       (call $db_get_i64
        (get_local $1)
        (i32.const 0)
        (i32.const 0)
       )
      )
      (i32.const 31)
     )
     (i32.const 1)
    )
    (i32.const 384)
   )
   (block $label$4
    (block $label$5
     (br_if $label$5
      (i32.lt_u
       (get_local $6)
       (i32.const 513)
      )
     )
     (set_local $4
      (call $malloc
       (get_local $6)
      )
     )
     (br $label$4)
    )
    (i32.store offset=4
     (i32.const 0)
     (tee_local $4
      (i32.sub
       (get_local $9)
       (i32.and
        (i32.add
         (get_local $6)
         (i32.const 15)
        )
        (i32.const -16)
       )
      )
     )
    )
   )
   (drop
    (call $db_get_i64
     (get_local $1)
     (get_local $4)
     (get_local $6)
    )
   )
   (i32.store offset=36
    (get_local $8)
    (get_local $4)
   )
   (i32.store offset=32
    (get_local $8)
    (get_local $4)
   )
   (i32.store offset=40
    (get_local $8)
    (i32.add
     (get_local $4)
     (get_local $6)
    )
   )
   (block $label$6
    (br_if $label$6
     (i32.lt_u
      (get_local $6)
      (i32.const 513)
     )
    )
    (call $free
     (get_local $4)
    )
   )
   (i32.store offset=56
    (tee_local $6
     (call $_Znwj
      (i32.const 296)
     )
    )
    (i32.const 0)
   )
   (i64.store offset=48 align=4
    (get_local $6)
    (i64.const 0)
   )
   (i32.store offset=72
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=76
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=80
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=120
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=124
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=128
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=144
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=188
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=192
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=196
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=200
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=204
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=208
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=216
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=220
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=224
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=228
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=232
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=236
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=244
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=248
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=252
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=256
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=260
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=264
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=268
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=272
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=276
    (get_local $6)
    (i32.const 0)
   )
   (i32.store offset=284
    (get_local $6)
    (get_local $0)
   )
   (call $eosio_assert
    (i32.gt_u
     (i32.sub
      (i32.load
       (i32.add
        (i32.add
         (get_local $8)
         (i32.const 32)
        )
        (i32.const 8)
       )
      )
      (tee_local $4
       (i32.load offset=36
        (get_local $8)
       )
      )
     )
     (i32.const 7)
    )
    (i32.const 416)
   )
   (drop
    (call $memcpy
     (get_local $6)
     (get_local $4)
     (i32.const 8)
    )
   )
   (i32.store offset=36
    (get_local $8)
    (i32.add
     (get_local $4)
     (i32.const 8)
    )
   )
   (call $eosio_assert
    (i32.gt_u
     (i32.sub
      (i32.load offset=8
       (tee_local $4
        (call $_ZN5eosiorsINS_10datastreamIPKcEEEERT_S6_RNS_14discovery_fileE
         (i32.add
          (get_local $8)
          (i32.const 32)
         )
         (i32.add
          (get_local $6)
          (i32.const 8)
         )
        )
       )
      )
      (i32.load offset=4
       (get_local $4)
      )
     )
     (i32.const 3)
    )
    (i32.const 416)
   )
   (drop
    (call $memcpy
     (i32.add
      (get_local $6)
      (i32.const 280)
     )
     (i32.load offset=4
      (get_local $4)
     )
     (i32.const 4)
    )
   )
   (i32.store offset=288
    (get_local $6)
    (get_local $1)
   )
   (i32.store offset=4
    (get_local $4)
    (i32.add
     (i32.load offset=4
      (get_local $4)
     )
     (i32.const 4)
    )
   )
   (i32.store offset=24
    (get_local $8)
    (get_local $6)
   )
   (i64.store offset=16
    (get_local $8)
    (tee_local $5
     (i64.load
      (get_local $6)
     )
    )
   )
   (i32.store offset=12
    (get_local $8)
    (tee_local $7
     (i32.load offset=288
      (get_local $6)
     )
    )
   )
   (block $label$7
    (block $label$8
     (br_if $label$8
      (i32.ge_u
       (tee_local $4
        (i32.load
         (tee_local $1
          (i32.add
           (get_local $0)
           (i32.const 28)
          )
         )
        )
       )
       (i32.load
        (i32.add
         (get_local $0)
         (i32.const 32)
        )
       )
      )
     )
     (i64.store offset=8
      (get_local $4)
      (get_local $5)
     )
     (i32.store offset=16
      (get_local $4)
      (get_local $7)
     )
     (i32.store offset=24
      (get_local $8)
      (i32.const 0)
     )
     (i32.store
      (get_local $4)
      (get_local $6)
     )
     (i32.store
      (get_local $1)
      (i32.add
       (get_local $4)
       (i32.const 24)
      )
     )
     (br $label$7)
    )
    (call $_ZNSt3__16vectorIN5eosio11multi_indexILy5454011350977085440ENS1_5disco13discovery_rowEJEE8item_ptrENS_9allocatorIS6_EEE24__emplace_back_slow_pathIJNS_10unique_ptrINS5_4itemENS_14default_deleteISC_EEEERyRlEEEvDpOT_
     (i32.add
      (get_local $0)
      (i32.const 24)
     )
     (i32.add
      (get_local $8)
      (i32.const 24)
     )
     (i32.add
      (get_local $8)
      (i32.const 16)
     )
     (i32.add
      (get_local $8)
      (i32.const 12)
     )
    )
   )
   (set_local $4
    (i32.load offset=24
     (get_local $8)
    )
   )
   (i32.store offset=24
    (get_local $8)
    (i32.const 0)
   )
   (br_if $label$2
    (i32.eqz
     (get_local $4)
    )
   )
   (drop
    (call $_ZN5eosio14discovery_fileD2Ev
     (i32.add
      (get_local $4)
      (i32.const 8)
     )
    )
   )
   (call $_ZdlPv
    (get_local $4)
   )
  )
  (i32.store offset=4
   (i32.const 0)
   (i32.add
    (get_local $8)
    (i32.const 48)
   )
  )
  (get_local $6)
 )
 (func $_ZZN5eosio11multi_indexILy5454011350977085440ENS_5disco13discovery_rowEJEE7emplaceIZNS1_9updtdiscoEyRKNS_14discovery_fileEE3$_0EENS3_14const_iteratorEyOT_ENKUlRSA_E_clINS3_4itemEEEDaSC_ (param $0 i32) (param $1 i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i32)
  (local $6 i64)
  (local $7 i32)
  (local $8 i32)
  (local $9 i32)
  (set_local $8
   (tee_local $9
    (i32.sub
     (i32.load offset=4
      (i32.const 0)
     )
     (i32.const 16)
    )
   )
  )
  (i32.store offset=4
   (i32.const 0)
   (get_local $9)
  )
  (i64.store
   (get_local $1)
   (i64.load
    (i32.load
     (tee_local $3
      (i32.load offset=4
       (get_local $0)
      )
     )
    )
   )
  )
  (set_local $2
   (i32.load
    (get_local $0)
   )
  )
  (block $label$0
   (block $label$1
    (br_if $label$1
     (i32.eq
      (tee_local $3
       (i32.load offset=4
        (get_local $3)
       )
      )
      (tee_local $7
       (call $memcpy
        (i32.add
         (get_local $1)
         (i32.const 8)
        )
        (get_local $3)
        (i32.const 40)
       )
      )
     )
    )
    (call $_ZNSt3__16vectorIN5eosio9peer_linkENS_9allocatorIS2_EEE6assignIPS2_EENS_9enable_ifIXaasr21__is_forward_iteratorIT_EE5valuesr16is_constructibleIS2_NS_15iterator_traitsIS9_E9referenceEEE5valueEvE4typeES9_S9_
     (i32.add
      (get_local $1)
      (i32.const 48)
     )
     (i32.load offset=40
      (get_local $3)
     )
     (i32.load
      (i32.add
       (get_local $3)
       (i32.const 44)
      )
     )
    )
    (i64.store
     (i32.add
      (get_local $1)
      (i32.const 64)
     )
     (i64.load offset=56
      (get_local $3)
     )
    )
    (call $_ZNSt3__16vectorINS_12basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEENS4_IS6_EEE6assignIPS6_EENS_9enable_ifIXaasr21__is_forward_iteratorIT_EE5valuesr16is_constructibleIS6_NS_15iterator_traitsISC_E9referenceEEE5valueEvE4typeESC_SC_
     (i32.add
      (get_local $1)
      (i32.const 72)
     )
     (i32.load offset=64
      (get_local $3)
     )
     (i32.load
      (i32.add
       (get_local $3)
       (i32.const 68)
      )
     )
    )
    (br $label$0)
   )
   (i64.store
    (i32.add
     (get_local $1)
     (i32.const 64)
    )
    (i64.load offset=56
     (get_local $3)
    )
   )
  )
  (drop
   (call $memcpy
    (i32.add
     (get_local $1)
     (i32.const 84)
    )
    (i32.add
     (get_local $3)
     (i32.const 76)
    )
    (i32.const 35)
   )
  )
  (drop
   (call $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEaSERKS5_
    (i32.add
     (get_local $1)
     (i32.const 120)
    )
    (i32.add
     (get_local $3)
     (i32.const 112)
    )
   )
  )
  (drop
   (call $memcpy
    (i32.add
     (get_local $1)
     (i32.const 136)
    )
    (i32.add
     (get_local $3)
     (i32.const 128)
    )
    (i32.const 48)
   )
  )
  (drop
   (call $_ZN5eosio13authority_mapaSERKS0_
    (i32.add
     (get_local $1)
     (i32.const 184)
    )
    (i32.add
     (get_local $3)
     (i32.const 176)
    )
   )
  )
  (block $label$2
   (br_if $label$2
    (i32.eq
     (get_local $7)
     (get_local $3)
    )
   )
   (call $_ZNSt3__16vectorIN5eosio11content_refENS_9allocatorIS2_EEE6assignIPS2_EENS_9enable_ifIXaasr21__is_forward_iteratorIT_EE5valuesr16is_constructibleIS2_NS_15iterator_traitsIS9_E9referenceEEE5valueEvE4typeES9_S9_
    (i32.add
     (get_local $1)
     (i32.const 268)
    )
    (i32.load offset=260
     (get_local $3)
    )
    (i32.load
     (i32.add
      (get_local $3)
      (i32.const 264)
     )
    )
   )
  )
  (i32.store offset=280
   (get_local $1)
   (call $now)
  )
  (i32.store
   (get_local $8)
   (i32.const 8)
  )
  (i32.store
   (tee_local $3
    (call $_ZN5eosiolsINS_10datastreamIjEEEERT_S4_RKNS_14discovery_fileE
     (get_local $8)
     (get_local $7)
    )
   )
   (i32.add
    (i32.load
     (get_local $3)
    )
    (i32.const 4)
   )
  )
  (set_local $4
   (i32.add
    (get_local $1)
    (i32.const 280)
   )
  )
  (block $label$3
   (block $label$4
    (br_if $label$4
     (i32.lt_u
      (tee_local $5
       (i32.load
        (get_local $8)
       )
      )
      (i32.const 513)
     )
    )
    (set_local $3
     (call $malloc
      (get_local $5)
     )
    )
    (br $label$3)
   )
   (i32.store offset=4
    (i32.const 0)
    (tee_local $3
     (i32.sub
      (get_local $9)
      (i32.and
       (i32.add
        (get_local $5)
        (i32.const 15)
       )
       (i32.const -16)
      )
     )
    )
   )
  )
  (i32.store
   (get_local $8)
   (get_local $3)
  )
  (i32.store offset=8
   (get_local $8)
   (i32.add
    (get_local $3)
    (get_local $5)
   )
  )
  (call $eosio_assert
   (i32.gt_s
    (get_local $5)
    (i32.const 7)
   )
   (i32.const 368)
  )
  (drop
   (call $memcpy
    (get_local $3)
    (get_local $1)
    (i32.const 8)
   )
  )
  (i32.store offset=4
   (get_local $8)
   (i32.add
    (get_local $3)
    (i32.const 8)
   )
  )
  (call $eosio_assert
   (i32.gt_s
    (i32.sub
     (i32.load offset=8
      (tee_local $7
       (call $_ZN5eosiolsINS_10datastreamIPcEEEERT_S5_RKNS_14discovery_fileE
        (get_local $8)
        (get_local $7)
       )
      )
     )
     (i32.load offset=4
      (get_local $7)
     )
    )
    (i32.const 3)
   )
   (i32.const 368)
  )
  (drop
   (call $memcpy
    (i32.load offset=4
     (get_local $7)
    )
    (get_local $4)
    (i32.const 4)
   )
  )
  (i32.store offset=4
   (get_local $7)
   (i32.add
    (i32.load offset=4
     (get_local $7)
    )
    (i32.const 4)
   )
  )
  (i32.store offset=288
   (get_local $1)
   (call $db_store_i64
    (i64.load offset=8
     (get_local $2)
    )
    (i64.const 5454011350977085440)
    (i64.load
     (i32.load offset=8
      (get_local $0)
     )
    )
    (tee_local $6
     (i64.load
      (get_local $1)
     )
    )
    (get_local $3)
    (get_local $5)
   )
  )
  (block $label$5
   (br_if $label$5
    (i32.lt_u
     (get_local $5)
     (i32.const 513)
    )
   )
   (call $free
    (get_local $3)
   )
  )
  (block $label$6
   (br_if $label$6
    (i64.lt_u
     (get_local $6)
     (i64.load offset=16
      (get_local $2)
     )
    )
   )
   (i64.store
    (i32.add
     (get_local $2)
     (i32.const 16)
    )
    (select
     (i64.const -2)
     (i64.add
      (get_local $6)
      (i64.const 1)
     )
     (i64.gt_u
      (get_local $6)
      (i64.const -3)
     )
    )
   )
  )
  (i32.store offset=4
   (i32.const 0)
   (i32.add
    (get_local $8)
    (i32.const 16)
   )
  )
 )
 (func $_ZNSt3__16vectorIN5eosio11multi_indexILy5454011350977085440ENS1_5disco13discovery_rowEJEE8item_ptrENS_9allocatorIS6_EEE24__emplace_back_slow_pathIJNS_10unique_ptrINS5_4itemENS_14default_deleteISC_EEEERyRlEEEvDpOT_ (param $0 i32) (param $1 i32) (param $2 i32) (param $3 i32)
  (local $4 i32)
  (local $5 i32)
  (local $6 i32)
  (local $7 i32)
  (block $label$0
   (block $label$1
    (br_if $label$1
     (i32.ge_u
      (tee_local $5
       (i32.add
        (tee_local $4
         (i32.div_s
          (i32.sub
           (i32.load offset=4
            (get_local $0)
           )
           (tee_local $6
            (i32.load
             (get_local $0)
            )
           )
          )
          (i32.const 24)
         )
        )
        (i32.const 1)
       )
      )
      (i32.const 178956971)
     )
    )
    (set_local $7
     (i32.const 178956970)
    )
    (block $label$2
     (block $label$3
      (br_if $label$3
       (i32.gt_u
        (tee_local $6
         (i32.div_s
          (i32.sub
           (i32.load offset=8
            (get_local $0)
           )
           (get_local $6)
          )
          (i32.const 24)
         )
        )
        (i32.const 89478484)
       )
      )
      (br_if $label$2
       (i32.eqz
        (tee_local $7
         (select
          (get_local $5)
          (tee_local $7
           (i32.shl
            (get_local $6)
            (i32.const 1)
           )
          )
          (i32.lt_u
           (get_local $7)
           (get_local $5)
          )
         )
        )
       )
      )
     )
     (set_local $6
      (call $_Znwj
       (i32.mul
        (get_local $7)
        (i32.const 24)
       )
      )
     )
     (br $label$0)
    )
    (set_local $7
     (i32.const 0)
    )
    (set_local $6
     (i32.const 0)
    )
    (br $label$0)
   )
   (call $_ZNKSt3__120__vector_base_commonILb1EE20__throw_length_errorEv
    (get_local $0)
   )
   (unreachable)
  )
  (set_local $5
   (i32.load
    (get_local $1)
   )
  )
  (i32.store
   (get_local $1)
   (i32.const 0)
  )
  (i32.store
   (tee_local $1
    (i32.add
     (get_local $6)
     (i32.mul
      (get_local $4)
      (i32.const 24)
     )
    )
   )
   (get_local $5)
  )
  (i64.store offset=8
   (get_local $1)
   (i64.load
    (get_local $2)
   )
  )
  (i32.store offset=16
   (get_local $1)
   (i32.load
    (get_local $3)
   )
  )
  (set_local $4
   (i32.add
    (get_local $6)
    (i32.mul
     (get_local $7)
     (i32.const 24)
    )
   )
  )
  (set_local $5
   (i32.add
    (get_local $1)
    (i32.const 24)
   )
  )
  (block $label$4
   (block $label$5
    (br_if $label$5
     (i32.eq
      (tee_local $6
       (i32.load
        (i32.add
         (get_local $0)
         (i32.const 4)
        )
       )
      )
      (tee_local $7
       (i32.load
        (get_local $0)
       )
      )
     )
    )
    (loop $label$6
     (set_local $3
      (i32.load
       (tee_local $2
        (i32.add
         (get_local $6)
         (i32.const -24)
        )
       )
      )
     )
     (i32.store
      (get_local $2)
      (i32.const 0)
     )
     (i32.store
      (i32.add
       (get_local $1)
       (i32.const -24)
      )
      (get_local $3)
     )
     (i32.store
      (i32.add
       (get_local $1)
       (i32.const -8)
      )
      (i32.load
       (i32.add
        (get_local $6)
        (i32.const -8)
       )
      )
     )
     (i32.store
      (i32.add
       (get_local $1)
       (i32.const -12)
      )
      (i32.load
       (i32.add
        (get_local $6)
        (i32.const -12)
       )
      )
     )
     (i32.store
      (i32.add
       (get_local $1)
       (i32.const -16)
      )
      (i32.load
       (i32.add
        (get_local $6)
        (i32.const -16)
       )
      )
     )
     (set_local $1
      (i32.add
       (get_local $1)
       (i32.const -24)
      )
     )
     (set_local $6
      (get_local $2)
     )
     (br_if $label$6
      (i32.ne
       (get_local $7)
       (get_local $2)
      )
     )
    )
    (set_local $7
     (i32.load
      (i32.add
       (get_local $0)
       (i32.const 4)
      )
     )
    )
    (set_local $6
     (i32.load
      (get_local $0)
     )
    )
    (br $label$4)
   )
   (set_local $6
    (get_local $7)
   )
  )
  (i32.store
   (get_local $0)
   (get_local $1)
  )
  (i32.store
   (i32.add
    (get_local $0)
    (i32.const 4)
   )
   (get_local $5)
  )
  (i32.store
   (i32.add
    (get_local $0)
    (i32.const 8)
   )
   (get_local $4)
  )
  (block $label$7
   (br_if $label$7
    (i32.eq
     (get_local $7)
     (get_local $6)
    )
   )
   (loop $label$8
    (set_local $1
     (i32.load
      (tee_local $7
       (i32.add
        (get_local $7)
        (i32.const -24)
       )
      )
     )
    )
    (i32.store
     (get_local $7)
     (i32.const 0)
    )
    (block $label$9
     (br_if $label$9
      (i32.eqz
       (get_local $1)
      )
     )
     (drop
      (call $_ZN5eosio14discovery_fileD2Ev
       (i32.add
        (get_local $1)
        (i32.const 8)
       )
      )
     )
     (call $_ZdlPv
      (get_local $1)
     )
    )
    (br_if $label$8
     (i32.ne
      (get_local $6)
      (get_local $7)
     )
    )
   )
  )
  (block $label$10
   (br_if $label$10
    (i32.eqz
     (get_local $6)
    )
   )
   (call $_ZdlPv
    (get_local $6)
   )
  )
 )
 (func $_ZN5eosio14discovery_fileD2Ev (param $0 i32) (result i32)
  (local $1 i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (block $label$0
   (br_if $label$0
    (i32.eqz
     (tee_local $1
      (i32.load offset=260
       (get_local $0)
      )
     )
    )
   )
   (block $label$1
    (block $label$2
     (br_if $label$2
      (i32.eq
       (tee_local $4
        (i32.load
         (tee_local $3
          (i32.add
           (get_local $0)
           (i32.const 264)
          )
         )
        )
       )
       (get_local $1)
      )
     )
     (set_local $2
      (i32.sub
       (i32.const 0)
       (get_local $1)
      )
     )
     (set_local $4
      (i32.add
       (get_local $4)
       (i32.const -36)
      )
     )
     (loop $label$3
      (block $label$4
       (br_if $label$4
        (i32.eqz
         (i32.and
          (i32.load8_u
           (i32.add
            (get_local $4)
            (i32.const 24)
           )
          )
          (i32.const 1)
         )
        )
       )
       (call $_ZdlPv
        (i32.load
         (i32.add
          (get_local $4)
          (i32.const 32)
         )
        )
       )
      )
      (block $label$5
       (br_if $label$5
        (i32.eqz
         (i32.and
          (i32.load8_u
           (i32.add
            (get_local $4)
            (i32.const 12)
           )
          )
          (i32.const 1)
         )
        )
       )
       (call $_ZdlPv
        (i32.load
         (i32.add
          (get_local $4)
          (i32.const 20)
         )
        )
       )
      )
      (block $label$6
       (br_if $label$6
        (i32.eqz
         (i32.and
          (i32.load8_u
           (get_local $4)
          )
          (i32.const 1)
         )
        )
       )
       (call $_ZdlPv
        (i32.load
         (i32.add
          (get_local $4)
          (i32.const 8)
         )
        )
       )
      )
      (br_if $label$3
       (i32.ne
        (i32.add
         (tee_local $4
          (i32.add
           (get_local $4)
           (i32.const -36)
          )
         )
         (get_local $2)
        )
        (i32.const -36)
       )
      )
     )
     (set_local $4
      (i32.load
       (i32.add
        (get_local $0)
        (i32.const 260)
       )
      )
     )
     (br $label$1)
    )
    (set_local $4
     (get_local $1)
    )
   )
   (i32.store
    (get_local $3)
    (get_local $1)
   )
   (call $_ZdlPv
    (get_local $4)
   )
  )
  (block $label$7
   (br_if $label$7
    (i32.eqz
     (tee_local $4
      (i32.load
       (i32.add
        (get_local $0)
        (i32.const 248)
       )
      )
     )
    )
   )
   (i32.store
    (i32.add
     (get_local $0)
     (i32.const 252)
    )
    (get_local $4)
   )
   (call $_ZdlPv
    (get_local $4)
   )
  )
  (block $label$8
   (br_if $label$8
    (i32.eqz
     (tee_local $4
      (i32.load
       (i32.add
        (get_local $0)
        (i32.const 236)
       )
      )
     )
    )
   )
   (i32.store
    (i32.add
     (get_local $0)
     (i32.const 240)
    )
    (get_local $4)
   )
   (call $_ZdlPv
    (get_local $4)
   )
  )
  (block $label$9
   (br_if $label$9
    (i32.eqz
     (tee_local $4
      (i32.load
       (i32.add
        (get_local $0)
        (i32.const 220)
       )
      )
     )
    )
   )
   (i32.store
    (i32.add
     (get_local $0)
     (i32.const 224)
    )
    (get_local $4)
   )
   (call $_ZdlPv
    (get_local $4)
   )
  )
  (block $label$10
   (br_if $label$10
    (i32.eqz
     (tee_local $4
      (i32.load
       (i32.add
        (get_local $0)
        (i32.const 208)
       )
      )
     )
    )
   )
   (i32.store
    (i32.add
     (get_local $0)
     (i32.const 212)
    )
    (get_local $4)
   )
   (call $_ZdlPv
    (get_local $4)
   )
  )
  (block $label$11
   (br_if $label$11
    (i32.eqz
     (tee_local $4
      (i32.load
       (i32.add
        (get_local $0)
        (i32.const 192)
       )
      )
     )
    )
   )
   (i32.store
    (i32.add
     (get_local $0)
     (i32.const 196)
    )
    (get_local $4)
   )
   (call $_ZdlPv
    (get_local $4)
   )
  )
  (block $label$12
   (br_if $label$12
    (i32.eqz
     (tee_local $4
      (i32.load
       (i32.add
        (get_local $0)
        (i32.const 180)
       )
      )
     )
    )
   )
   (i32.store
    (i32.add
     (get_local $0)
     (i32.const 184)
    )
    (get_local $4)
   )
   (call $_ZdlPv
    (get_local $4)
   )
  )
  (block $label$13
   (br_if $label$13
    (i32.eqz
     (i32.and
      (i32.load8_u offset=112
       (get_local $0)
      )
      (i32.const 1)
     )
    )
   )
   (call $_ZdlPv
    (i32.load
     (i32.add
      (get_local $0)
      (i32.const 120)
     )
    )
   )
  )
  (block $label$14
   (br_if $label$14
    (i32.eqz
     (tee_local $1
      (i32.load offset=64
       (get_local $0)
      )
     )
    )
   )
   (block $label$15
    (block $label$16
     (br_if $label$16
      (i32.eq
       (tee_local $4
        (i32.load
         (tee_local $3
          (i32.add
           (get_local $0)
           (i32.const 68)
          )
         )
        )
       )
       (get_local $1)
      )
     )
     (set_local $2
      (i32.sub
       (i32.const 0)
       (get_local $1)
      )
     )
     (set_local $4
      (i32.add
       (get_local $4)
       (i32.const -12)
      )
     )
     (loop $label$17
      (block $label$18
       (br_if $label$18
        (i32.eqz
         (i32.and
          (i32.load8_u
           (get_local $4)
          )
          (i32.const 1)
         )
        )
       )
       (call $_ZdlPv
        (i32.load
         (i32.add
          (get_local $4)
          (i32.const 8)
         )
        )
       )
      )
      (br_if $label$17
       (i32.ne
        (i32.add
         (tee_local $4
          (i32.add
           (get_local $4)
           (i32.const -12)
          )
         )
         (get_local $2)
        )
        (i32.const -12)
       )
      )
     )
     (set_local $4
      (i32.load
       (i32.add
        (get_local $0)
        (i32.const 64)
       )
      )
     )
     (br $label$15)
    )
    (set_local $4
     (get_local $1)
    )
   )
   (i32.store
    (get_local $3)
    (get_local $1)
   )
   (call $_ZdlPv
    (get_local $4)
   )
  )
  (block $label$19
   (br_if $label$19
    (i32.eqz
     (tee_local $1
      (i32.load offset=40
       (get_local $0)
      )
     )
    )
   )
   (block $label$20
    (block $label$21
     (br_if $label$21
      (i32.eq
       (tee_local $4
        (i32.load
         (tee_local $3
          (i32.add
           (get_local $0)
           (i32.const 44)
          )
         )
        )
       )
       (get_local $1)
      )
     )
     (set_local $2
      (i32.sub
       (i32.const 0)
       (get_local $1)
      )
     )
     (set_local $4
      (i32.add
       (get_local $4)
       (i32.const -24)
      )
     )
     (loop $label$22
      (block $label$23
       (br_if $label$23
        (i32.eqz
         (i32.and
          (i32.load8_u
           (get_local $4)
          )
          (i32.const 1)
         )
        )
       )
       (call $_ZdlPv
        (i32.load
         (i32.add
          (get_local $4)
          (i32.const 8)
         )
        )
       )
      )
      (br_if $label$22
       (i32.ne
        (i32.add
         (tee_local $4
          (i32.add
           (get_local $4)
           (i32.const -32)
          )
         )
         (get_local $2)
        )
        (i32.const -24)
       )
      )
     )
     (set_local $4
      (i32.load
       (i32.add
        (get_local $0)
        (i32.const 40)
       )
      )
     )
     (br $label$20)
    )
    (set_local $4
     (get_local $1)
    )
   )
   (i32.store
    (get_local $3)
    (get_local $1)
   )
   (call $_ZdlPv
    (get_local $4)
   )
  )
  (get_local $0)
 )
 (func $_ZN5eosio11multi_indexILy5454011350977085440ENS_5disco13discovery_rowEJEE6modifyIZNS1_9updtdiscoEyRKNS_14discovery_fileEE3$_1EEvRKS2_yOT_ (param $0 i32) (param $1 i32) (param $2 i64) (param $3 i32)
  (local $4 i64)
  (local $5 i32)
  (local $6 i32)
  (local $7 i32)
  (local $8 i32)
  (local $9 i32)
  (set_local $8
   (tee_local $9
    (i32.sub
     (i32.load offset=4
      (i32.const 0)
     )
     (i32.const 16)
    )
   )
  )
  (i32.store offset=4
   (i32.const 0)
   (get_local $9)
  )
  (call $eosio_assert
   (i32.eq
    (i32.load offset=284
     (get_local $1)
    )
    (get_local $0)
   )
   (i32.const 192)
  )
  (call $eosio_assert
   (i64.eq
    (i64.load
     (get_local $0)
    )
    (call $current_receiver)
   )
   (i32.const 240)
  )
  (set_local $4
   (i64.load
    (get_local $1)
   )
  )
  (block $label$0
   (block $label$1
    (br_if $label$1
     (i32.eq
      (tee_local $3
       (i32.load
        (get_local $3)
       )
      )
      (tee_local $7
       (call $memcpy
        (i32.add
         (get_local $1)
         (i32.const 8)
        )
        (get_local $3)
        (i32.const 40)
       )
      )
     )
    )
    (call $_ZNSt3__16vectorIN5eosio9peer_linkENS_9allocatorIS2_EEE6assignIPS2_EENS_9enable_ifIXaasr21__is_forward_iteratorIT_EE5valuesr16is_constructibleIS2_NS_15iterator_traitsIS9_E9referenceEEE5valueEvE4typeES9_S9_
     (i32.add
      (get_local $1)
      (i32.const 48)
     )
     (i32.load offset=40
      (get_local $3)
     )
     (i32.load
      (i32.add
       (get_local $3)
       (i32.const 44)
      )
     )
    )
    (i64.store
     (i32.add
      (get_local $1)
      (i32.const 64)
     )
     (i64.load offset=56
      (get_local $3)
     )
    )
    (call $_ZNSt3__16vectorINS_12basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEENS4_IS6_EEE6assignIPS6_EENS_9enable_ifIXaasr21__is_forward_iteratorIT_EE5valuesr16is_constructibleIS6_NS_15iterator_traitsISC_E9referenceEEE5valueEvE4typeESC_SC_
     (i32.add
      (get_local $1)
      (i32.const 72)
     )
     (i32.load offset=64
      (get_local $3)
     )
     (i32.load
      (i32.add
       (get_local $3)
       (i32.const 68)
      )
     )
    )
    (br $label$0)
   )
   (i64.store
    (i32.add
     (get_local $1)
     (i32.const 64)
    )
    (i64.load offset=56
     (get_local $3)
    )
   )
  )
  (drop
   (call $memcpy
    (i32.add
     (get_local $1)
     (i32.const 84)
    )
    (i32.add
     (get_local $3)
     (i32.const 76)
    )
    (i32.const 35)
   )
  )
  (drop
   (call $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEaSERKS5_
    (i32.add
     (get_local $1)
     (i32.const 120)
    )
    (i32.add
     (get_local $3)
     (i32.const 112)
    )
   )
  )
  (drop
   (call $memcpy
    (i32.add
     (get_local $1)
     (i32.const 136)
    )
    (i32.add
     (get_local $3)
     (i32.const 128)
    )
    (i32.const 48)
   )
  )
  (drop
   (call $_ZN5eosio13authority_mapaSERKS0_
    (i32.add
     (get_local $1)
     (i32.const 184)
    )
    (i32.add
     (get_local $3)
     (i32.const 176)
    )
   )
  )
  (block $label$2
   (br_if $label$2
    (i32.eq
     (get_local $7)
     (get_local $3)
    )
   )
   (call $_ZNSt3__16vectorIN5eosio11content_refENS_9allocatorIS2_EEE6assignIPS2_EENS_9enable_ifIXaasr21__is_forward_iteratorIT_EE5valuesr16is_constructibleIS2_NS_15iterator_traitsIS9_E9referenceEEE5valueEvE4typeES9_S9_
    (i32.add
     (get_local $1)
     (i32.const 268)
    )
    (i32.load offset=260
     (get_local $3)
    )
    (i32.load
     (i32.add
      (get_local $3)
      (i32.const 264)
     )
    )
   )
  )
  (i32.store offset=280
   (get_local $1)
   (call $now)
  )
  (call $eosio_assert
   (i64.eq
    (get_local $4)
    (i64.load
     (get_local $1)
    )
   )
   (i32.const 304)
  )
  (i32.store
   (get_local $8)
   (i32.const 8)
  )
  (i32.store
   (tee_local $3
    (call $_ZN5eosiolsINS_10datastreamIjEEEERT_S4_RKNS_14discovery_fileE
     (get_local $8)
     (get_local $7)
    )
   )
   (i32.add
    (i32.load
     (get_local $3)
    )
    (i32.const 4)
   )
  )
  (set_local $5
   (i32.add
    (get_local $1)
    (i32.const 280)
   )
  )
  (block $label$3
   (block $label$4
    (br_if $label$4
     (i32.lt_u
      (tee_local $6
       (i32.load
        (get_local $8)
       )
      )
      (i32.const 513)
     )
    )
    (set_local $3
     (call $malloc
      (get_local $6)
     )
    )
    (br $label$3)
   )
   (i32.store offset=4
    (i32.const 0)
    (tee_local $3
     (i32.sub
      (get_local $9)
      (i32.and
       (i32.add
        (get_local $6)
        (i32.const 15)
       )
       (i32.const -16)
      )
     )
    )
   )
  )
  (i32.store
   (get_local $8)
   (get_local $3)
  )
  (i32.store offset=8
   (get_local $8)
   (i32.add
    (get_local $3)
    (get_local $6)
   )
  )
  (call $eosio_assert
   (i32.gt_s
    (get_local $6)
    (i32.const 7)
   )
   (i32.const 368)
  )
  (drop
   (call $memcpy
    (get_local $3)
    (get_local $1)
    (i32.const 8)
   )
  )
  (i32.store offset=4
   (get_local $8)
   (i32.add
    (get_local $3)
    (i32.const 8)
   )
  )
  (call $eosio_assert
   (i32.gt_s
    (i32.sub
     (i32.load offset=8
      (tee_local $7
       (call $_ZN5eosiolsINS_10datastreamIPcEEEERT_S5_RKNS_14discovery_fileE
        (get_local $8)
        (get_local $7)
       )
      )
     )
     (i32.load offset=4
      (get_local $7)
     )
    )
    (i32.const 3)
   )
   (i32.const 368)
  )
  (drop
   (call $memcpy
    (i32.load offset=4
     (get_local $7)
    )
    (get_local $5)
    (i32.const 4)
   )
  )
  (i32.store offset=4
   (get_local $7)
   (i32.add
    (i32.load offset=4
     (get_local $7)
    )
    (i32.const 4)
   )
  )
  (call $db_update_i64
   (i32.load offset=288
    (get_local $1)
   )
   (get_local $2)
   (get_local $3)
   (get_local $6)
  )
  (block $label$5
   (br_if $label$5
    (i32.lt_u
     (get_local $6)
     (i32.const 513)
    )
   )
   (call $free
    (get_local $3)
   )
  )
  (block $label$6
   (br_if $label$6
    (i64.lt_u
     (get_local $4)
     (i64.load offset=16
      (get_local $0)
     )
    )
   )
   (i64.store
    (i32.add
     (get_local $0)
     (i32.const 16)
    )
    (select
     (i64.const -2)
     (i64.add
      (get_local $4)
      (i64.const 1)
     )
     (i64.gt_u
      (get_local $4)
      (i64.const -3)
     )
    )
   )
  )
  (i32.store offset=4
   (i32.const 0)
   (i32.add
    (get_local $8)
    (i32.const 16)
   )
  )
 )
 (func $_ZNSt3__16vectorIN5eosio9peer_linkENS_9allocatorIS2_EEE6assignIPS2_EENS_9enable_ifIXaasr21__is_forward_iteratorIT_EE5valuesr16is_constructibleIS2_NS_15iterator_traitsIS9_E9referenceEEE5valueEvE4typeES9_S9_ (param $0 i32) (param $1 i32) (param $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i32)
  (local $6 i32)
  (local $7 i32)
  (block $label$0
   (block $label$1
    (block $label$2
     (block $label$3
      (block $label$4
       (block $label$5
        (br_if $label$5
         (i32.le_u
          (tee_local $3
           (i32.shr_s
            (i32.sub
             (get_local $2)
             (get_local $1)
            )
            (i32.const 5)
           )
          )
          (i32.shr_s
           (i32.sub
            (tee_local $7
             (i32.load offset=8
              (get_local $0)
             )
            )
            (tee_local $6
             (i32.load
              (get_local $0)
             )
            )
           )
           (i32.const 5)
          )
         )
        )
        (br_if $label$2
         (i32.eqz
          (get_local $6)
         )
        )
        (br_if $label$4
         (i32.eq
          (tee_local $7
           (i32.load offset=4
            (get_local $0)
           )
          )
          (get_local $6)
         )
        )
        (set_local $5
         (i32.sub
          (i32.const 0)
          (get_local $6)
         )
        )
        (set_local $7
         (i32.add
          (get_local $7)
          (i32.const -24)
         )
        )
        (loop $label$6
         (block $label$7
          (br_if $label$7
           (i32.eqz
            (i32.and
             (i32.load8_u
              (get_local $7)
             )
             (i32.const 1)
            )
           )
          )
          (call $_ZdlPv
           (i32.load
            (i32.add
             (get_local $7)
             (i32.const 8)
            )
           )
          )
         )
         (br_if $label$6
          (i32.ne
           (i32.add
            (tee_local $7
             (i32.add
              (get_local $7)
              (i32.const -32)
             )
            )
            (get_local $5)
           )
           (i32.const -24)
          )
         )
        )
        (set_local $7
         (i32.load
          (get_local $0)
         )
        )
        (br $label$3)
       )
       (block $label$8
        (br_if $label$8
         (i32.eq
          (tee_local $5
           (select
            (tee_local $7
             (i32.add
              (get_local $1)
              (tee_local $5
               (i32.sub
                (i32.load offset=4
                 (get_local $0)
                )
                (get_local $6)
               )
              )
             )
            )
            (get_local $2)
            (i32.gt_u
             (get_local $3)
             (tee_local $4
              (i32.shr_s
               (get_local $5)
               (i32.const 5)
              )
             )
            )
           )
          )
          (get_local $1)
         )
        )
        (loop $label$9
         (i64.store
          (get_local $6)
          (i64.load
           (get_local $1)
          )
         )
         (drop
          (call $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEaSERKS5_
           (i32.add
            (get_local $6)
            (i32.const 8)
           )
           (i32.add
            (get_local $1)
            (i32.const 8)
           )
          )
         )
         (i64.store
          (i32.add
           (get_local $6)
           (i32.const 24)
          )
          (i64.load
           (i32.add
            (get_local $1)
            (i32.const 24)
           )
          )
         )
         (set_local $6
          (i32.add
           (get_local $6)
           (i32.const 32)
          )
         )
         (br_if $label$9
          (i32.ne
           (tee_local $1
            (i32.add
             (get_local $1)
             (i32.const 32)
            )
           )
           (get_local $5)
          )
         )
        )
       )
       (block $label$10
        (br_if $label$10
         (i32.le_u
          (get_local $3)
          (get_local $4)
         )
        )
        (br_if $label$1
         (i32.eq
          (get_local $5)
          (get_local $2)
         )
        )
        (set_local $1
         (i32.load
          (tee_local $6
           (i32.add
            (get_local $0)
            (i32.const 4)
           )
          )
         )
        )
        (loop $label$11
         (i64.store
          (get_local $1)
          (i64.load
           (get_local $7)
          )
         )
         (drop
          (call $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEC2ERKS5_
           (i32.add
            (get_local $1)
            (i32.const 8)
           )
           (i32.add
            (get_local $7)
            (i32.const 8)
           )
          )
         )
         (i64.store offset=24
          (get_local $1)
          (i64.load
           (i32.add
            (get_local $7)
            (i32.const 24)
           )
          )
         )
         (i32.store
          (get_local $6)
          (tee_local $1
           (i32.add
            (i32.load
             (get_local $6)
            )
            (i32.const 32)
           )
          )
         )
         (br_if $label$11
          (i32.ne
           (tee_local $7
            (i32.add
             (get_local $7)
             (i32.const 32)
            )
           )
           (get_local $2)
          )
         )
         (br $label$1)
        )
       )
       (block $label$12
        (br_if $label$12
         (i32.eq
          (tee_local $1
           (i32.load
            (tee_local $2
             (i32.add
              (get_local $0)
              (i32.const 4)
             )
            )
           )
          )
          (get_local $6)
         )
        )
        (set_local $7
         (i32.sub
          (i32.const 0)
          (get_local $6)
         )
        )
        (set_local $1
         (i32.add
          (get_local $1)
          (i32.const -24)
         )
        )
        (loop $label$13
         (block $label$14
          (br_if $label$14
           (i32.eqz
            (i32.and
             (i32.load8_u
              (get_local $1)
             )
             (i32.const 1)
            )
           )
          )
          (call $_ZdlPv
           (i32.load
            (i32.add
             (get_local $1)
             (i32.const 8)
            )
           )
          )
         )
         (br_if $label$13
          (i32.ne
           (i32.add
            (tee_local $1
             (i32.add
              (get_local $1)
              (i32.const -32)
             )
            )
            (get_local $7)
           )
           (i32.const -24)
          )
         )
        )
       )
       (i32.store
        (get_local $2)
        (get_local $6)
       )
       (return)
      )
      (set_local $7
       (get_local $6)
      )
     )
     (i32.store
      (i32.add
       (get_local $0)
       (i32.const 4)
      )
      (get_local $6)
     )
     (call $_ZdlPv
      (get_local $7)
     )
     (set_local $7
      (i32.const 0)
     )
     (i32.store
      (i32.add
       (get_local $0)
       (i32.const 8)
      )
      (i32.const 0)
     )
     (i64.store align=4
      (get_local $0)
      (i64.const 0)
     )
    )
    (br_if $label$0
     (i32.ge_u
      (get_local $3)
      (i32.const 134217728)
     )
    )
    (set_local $6
     (i32.const 134217727)
    )
    (block $label$15
     (br_if $label$15
      (i32.gt_u
       (i32.shr_s
        (get_local $7)
        (i32.const 5)
       )
       (i32.const 67108862)
      )
     )
     (set_local $6
      (get_local $3)
     )
     (br_if $label$15
      (i32.lt_u
       (tee_local $7
        (i32.shr_s
         (get_local $7)
         (i32.const 4)
        )
       )
       (get_local $3)
      )
     )
     (set_local $6
      (get_local $7)
     )
     (br_if $label$0
      (i32.ge_u
       (get_local $7)
       (i32.const 134217728)
      )
     )
    )
    (i32.store
     (get_local $0)
     (tee_local $6
      (call $_Znwj
       (tee_local $7
        (i32.shl
         (get_local $6)
         (i32.const 5)
        )
       )
      )
     )
    )
    (i32.store offset=4
     (get_local $0)
     (get_local $6)
    )
    (i32.store
     (i32.add
      (get_local $0)
      (i32.const 8)
     )
     (i32.add
      (get_local $6)
      (get_local $7)
     )
    )
    (br_if $label$1
     (i32.eq
      (get_local $1)
      (get_local $2)
     )
    )
    (set_local $7
     (i32.add
      (get_local $0)
      (i32.const 4)
     )
    )
    (loop $label$16
     (i64.store
      (get_local $6)
      (i64.load
       (get_local $1)
      )
     )
     (drop
      (call $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEC2ERKS5_
       (i32.add
        (get_local $6)
        (i32.const 8)
       )
       (i32.add
        (get_local $1)
        (i32.const 8)
       )
      )
     )
     (i64.store offset=24
      (get_local $6)
      (i64.load
       (i32.add
        (get_local $1)
        (i32.const 24)
       )
      )
     )
     (i32.store
      (get_local $7)
      (tee_local $6
       (i32.add
        (i32.load
         (get_local $7)
        )
        (i32.const 32)
       )
      )
     )
     (br_if $label$16
      (i32.ne
       (tee_local $1
        (i32.add
         (get_local $1)
         (i32.const 32)
        )
       )
       (get_local $2)
      )
     )
    )
   )
   (return)
  )
  (call $_ZNKSt3__120__vector_base_commonILb1EE20__throw_length_errorEv
   (get_local $0)
  )
  (unreachable)
 )
 (func $_ZNSt3__16vectorINS_12basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEENS4_IS6_EEE6assignIPS6_EENS_9enable_ifIXaasr21__is_forward_iteratorIT_EE5valuesr16is_constructibleIS6_NS_15iterator_traitsISC_E9referenceEEE5valueEvE4typeESC_SC_ (param $0 i32) (param $1 i32) (param $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i32)
  (local $6 i32)
  (local $7 i32)
  (block $label$0
   (block $label$1
    (block $label$2
     (block $label$3
      (block $label$4
       (block $label$5
        (br_if $label$5
         (i32.le_u
          (tee_local $3
           (i32.div_s
            (i32.sub
             (get_local $2)
             (get_local $1)
            )
            (i32.const 12)
           )
          )
          (i32.div_s
           (i32.sub
            (tee_local $7
             (i32.load offset=8
              (get_local $0)
             )
            )
            (tee_local $6
             (i32.load
              (get_local $0)
             )
            )
           )
           (i32.const 12)
          )
         )
        )
        (br_if $label$2
         (i32.eqz
          (get_local $6)
         )
        )
        (br_if $label$4
         (i32.eq
          (tee_local $7
           (i32.load offset=4
            (get_local $0)
           )
          )
          (get_local $6)
         )
        )
        (set_local $5
         (i32.sub
          (i32.const 0)
          (get_local $6)
         )
        )
        (set_local $7
         (i32.add
          (get_local $7)
          (i32.const -12)
         )
        )
        (loop $label$6
         (block $label$7
          (br_if $label$7
           (i32.eqz
            (i32.and
             (i32.load8_u
              (get_local $7)
             )
             (i32.const 1)
            )
           )
          )
          (call $_ZdlPv
           (i32.load
            (i32.add
             (get_local $7)
             (i32.const 8)
            )
           )
          )
         )
         (br_if $label$6
          (i32.ne
           (i32.add
            (tee_local $7
             (i32.add
              (get_local $7)
              (i32.const -12)
             )
            )
            (get_local $5)
           )
           (i32.const -12)
          )
         )
        )
        (set_local $7
         (i32.load
          (get_local $0)
         )
        )
        (br $label$3)
       )
       (block $label$8
        (br_if $label$8
         (i32.eq
          (tee_local $5
           (select
            (tee_local $7
             (i32.add
              (get_local $1)
              (i32.mul
               (tee_local $4
                (i32.div_s
                 (i32.sub
                  (i32.load offset=4
                   (get_local $0)
                  )
                  (get_local $6)
                 )
                 (i32.const 12)
                )
               )
               (i32.const 12)
              )
             )
            )
            (get_local $2)
            (i32.gt_u
             (get_local $3)
             (get_local $4)
            )
           )
          )
          (get_local $1)
         )
        )
        (loop $label$9
         (drop
          (call $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEaSERKS5_
           (get_local $6)
           (get_local $1)
          )
         )
         (set_local $6
          (i32.add
           (get_local $6)
           (i32.const 12)
          )
         )
         (br_if $label$9
          (i32.ne
           (get_local $5)
           (tee_local $1
            (i32.add
             (get_local $1)
             (i32.const 12)
            )
           )
          )
         )
        )
       )
       (block $label$10
        (br_if $label$10
         (i32.le_u
          (get_local $3)
          (get_local $4)
         )
        )
        (br_if $label$1
         (i32.eq
          (get_local $5)
          (get_local $2)
         )
        )
        (set_local $6
         (i32.load
          (tee_local $1
           (i32.add
            (get_local $0)
            (i32.const 4)
           )
          )
         )
        )
        (loop $label$11
         (drop
          (call $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEC2ERKS5_
           (get_local $6)
           (get_local $7)
          )
         )
         (i32.store
          (get_local $1)
          (tee_local $6
           (i32.add
            (i32.load
             (get_local $1)
            )
            (i32.const 12)
           )
          )
         )
         (br_if $label$11
          (i32.ne
           (get_local $2)
           (tee_local $7
            (i32.add
             (get_local $7)
             (i32.const 12)
            )
           )
          )
         )
         (br $label$1)
        )
       )
       (block $label$12
        (br_if $label$12
         (i32.eq
          (tee_local $1
           (i32.load
            (tee_local $2
             (i32.add
              (get_local $0)
              (i32.const 4)
             )
            )
           )
          )
          (get_local $6)
         )
        )
        (set_local $7
         (i32.sub
          (i32.const 0)
          (get_local $6)
         )
        )
        (set_local $1
         (i32.add
          (get_local $1)
          (i32.const -12)
         )
        )
        (loop $label$13
         (block $label$14
          (br_if $label$14
           (i32.eqz
            (i32.and
             (i32.load8_u
              (get_local $1)
             )
             (i32.const 1)
            )
           )
          )
          (call $_ZdlPv
           (i32.load
            (i32.add
             (get_local $1)
             (i32.const 8)
            )
           )
          )
         )
         (br_if $label$13
          (i32.ne
           (i32.add
            (tee_local $1
             (i32.add
              (get_local $1)
              (i32.const -12)
             )
            )
            (get_local $7)
           )
           (i32.const -12)
          )
         )
        )
       )
       (i32.store
        (get_local $2)
        (get_local $6)
       )
       (return)
      )
      (set_local $7
       (get_local $6)
      )
     )
     (i32.store
      (i32.add
       (get_local $0)
       (i32.const 4)
      )
      (get_local $6)
     )
     (call $_ZdlPv
      (get_local $7)
     )
     (set_local $7
      (i32.const 0)
     )
     (i32.store
      (i32.add
       (get_local $0)
       (i32.const 8)
      )
      (i32.const 0)
     )
     (i64.store align=4
      (get_local $0)
      (i64.const 0)
     )
    )
    (br_if $label$0
     (i32.ge_u
      (get_local $3)
      (i32.const 357913942)
     )
    )
    (set_local $6
     (i32.const 357913941)
    )
    (block $label$15
     (br_if $label$15
      (i32.gt_u
       (tee_local $7
        (i32.div_s
         (get_local $7)
         (i32.const 12)
        )
       )
       (i32.const 178956969)
      )
     )
     (set_local $6
      (select
       (get_local $3)
       (tee_local $6
        (i32.shl
         (get_local $7)
         (i32.const 1)
        )
       )
       (i32.lt_u
        (get_local $6)
        (get_local $3)
       )
      )
     )
    )
    (i32.store
     (get_local $0)
     (tee_local $6
      (call $_Znwj
       (tee_local $7
        (i32.mul
         (get_local $6)
         (i32.const 12)
        )
       )
      )
     )
    )
    (i32.store offset=4
     (get_local $0)
     (get_local $6)
    )
    (i32.store
     (i32.add
      (get_local $0)
      (i32.const 8)
     )
     (i32.add
      (get_local $6)
      (get_local $7)
     )
    )
    (br_if $label$1
     (i32.eq
      (get_local $1)
      (get_local $2)
     )
    )
    (set_local $7
     (i32.add
      (get_local $0)
      (i32.const 4)
     )
    )
    (loop $label$16
     (drop
      (call $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEC2ERKS5_
       (get_local $6)
       (get_local $1)
      )
     )
     (i32.store
      (get_local $7)
      (tee_local $6
       (i32.add
        (i32.load
         (get_local $7)
        )
        (i32.const 12)
       )
      )
     )
     (br_if $label$16
      (i32.ne
       (get_local $2)
       (tee_local $1
        (i32.add
         (get_local $1)
         (i32.const 12)
        )
       )
      )
     )
    )
   )
   (return)
  )
  (call $_ZNKSt3__120__vector_base_commonILb1EE20__throw_length_errorEv
   (get_local $0)
  )
  (unreachable)
 )
 (func $_ZN5eosio13authority_mapaSERKS0_ (param $0 i32) (param $1 i32) (result i32)
  (i32.store
   (get_local $0)
   (i32.load
    (get_local $1)
   )
  )
  (block $label$0
   (br_if $label$0
    (i32.eq
     (get_local $0)
     (get_local $1)
    )
   )
   (call $_ZNSt3__16vectorIN11eosiosystem10key_weightENS_9allocatorIS2_EEE6assignIPS2_EENS_9enable_ifIXaasr21__is_forward_iteratorIT_EE5valuesr16is_constructibleIS2_NS_15iterator_traitsIS9_E9referenceEEE5valueEvE4typeES9_S9_
    (i32.add
     (get_local $0)
     (i32.const 4)
    )
    (i32.load offset=4
     (get_local $1)
    )
    (i32.load
     (i32.add
      (get_local $1)
      (i32.const 8)
     )
    )
   )
   (call $_ZNSt3__16vectorIN11eosiosystem23permission_level_weightENS_9allocatorIS2_EEE6assignIPS2_EENS_9enable_ifIXaasr21__is_forward_iteratorIT_EE5valuesr16is_constructibleIS2_NS_15iterator_traitsIS9_E9referenceEEE5valueEvE4typeES9_S9_
    (i32.add
     (get_local $0)
     (i32.const 16)
    )
    (i32.load offset=16
     (get_local $1)
    )
    (i32.load
     (i32.add
      (get_local $1)
      (i32.const 20)
     )
    )
   )
   (i32.store offset=28
    (get_local $0)
    (i32.load offset=28
     (get_local $1)
    )
   )
   (call $_ZNSt3__16vectorIN11eosiosystem10key_weightENS_9allocatorIS2_EEE6assignIPS2_EENS_9enable_ifIXaasr21__is_forward_iteratorIT_EE5valuesr16is_constructibleIS2_NS_15iterator_traitsIS9_E9referenceEEE5valueEvE4typeES9_S9_
    (i32.add
     (get_local $0)
     (i32.const 32)
    )
    (i32.load
     (i32.add
      (get_local $1)
      (i32.const 32)
     )
    )
    (i32.load
     (i32.add
      (get_local $1)
      (i32.const 36)
     )
    )
   )
   (call $_ZNSt3__16vectorIN11eosiosystem23permission_level_weightENS_9allocatorIS2_EEE6assignIPS2_EENS_9enable_ifIXaasr21__is_forward_iteratorIT_EE5valuesr16is_constructibleIS2_NS_15iterator_traitsIS9_E9referenceEEE5valueEvE4typeES9_S9_
    (i32.add
     (get_local $0)
     (i32.const 44)
    )
    (i32.load
     (i32.add
      (get_local $1)
      (i32.const 44)
     )
    )
    (i32.load
     (i32.add
      (get_local $1)
      (i32.const 48)
     )
    )
   )
   (i32.store offset=56
    (get_local $0)
    (i32.load offset=56
     (get_local $1)
    )
   )
   (call $_ZNSt3__16vectorIN11eosiosystem10key_weightENS_9allocatorIS2_EEE6assignIPS2_EENS_9enable_ifIXaasr21__is_forward_iteratorIT_EE5valuesr16is_constructibleIS2_NS_15iterator_traitsIS9_E9referenceEEE5valueEvE4typeES9_S9_
    (i32.add
     (get_local $0)
     (i32.const 60)
    )
    (i32.load
     (i32.add
      (get_local $1)
      (i32.const 60)
     )
    )
    (i32.load
     (i32.add
      (get_local $1)
      (i32.const 64)
     )
    )
   )
   (call $_ZNSt3__16vectorIN11eosiosystem23permission_level_weightENS_9allocatorIS2_EEE6assignIPS2_EENS_9enable_ifIXaasr21__is_forward_iteratorIT_EE5valuesr16is_constructibleIS2_NS_15iterator_traitsIS9_E9referenceEEE5valueEvE4typeES9_S9_
    (i32.add
     (get_local $0)
     (i32.const 72)
    )
    (i32.load
     (i32.add
      (get_local $1)
      (i32.const 72)
     )
    )
    (i32.load
     (i32.add
      (get_local $1)
      (i32.const 76)
     )
    )
   )
   (return
    (get_local $0)
   )
  )
  (i32.store offset=28
   (get_local $0)
   (i32.load offset=28
    (get_local $1)
   )
  )
  (i32.store offset=56
   (get_local $0)
   (i32.load offset=56
    (get_local $1)
   )
  )
  (get_local $0)
 )
 (func $_ZNSt3__16vectorIN5eosio11content_refENS_9allocatorIS2_EEE6assignIPS2_EENS_9enable_ifIXaasr21__is_forward_iteratorIT_EE5valuesr16is_constructibleIS2_NS_15iterator_traitsIS9_E9referenceEEE5valueEvE4typeES9_S9_ (param $0 i32) (param $1 i32) (param $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i32)
  (local $6 i32)
  (local $7 i32)
  (block $label$0
   (block $label$1
    (block $label$2
     (block $label$3
      (block $label$4
       (block $label$5
        (br_if $label$5
         (i32.le_u
          (tee_local $3
           (i32.div_s
            (i32.sub
             (get_local $2)
             (get_local $1)
            )
            (i32.const 36)
           )
          )
          (i32.div_s
           (i32.sub
            (tee_local $7
             (i32.load offset=8
              (get_local $0)
             )
            )
            (tee_local $6
             (i32.load
              (get_local $0)
             )
            )
           )
           (i32.const 36)
          )
         )
        )
        (br_if $label$2
         (i32.eqz
          (get_local $6)
         )
        )
        (br_if $label$4
         (i32.eq
          (tee_local $7
           (i32.load offset=4
            (get_local $0)
           )
          )
          (get_local $6)
         )
        )
        (set_local $5
         (i32.sub
          (i32.const 0)
          (get_local $6)
         )
        )
        (set_local $7
         (i32.add
          (get_local $7)
          (i32.const -36)
         )
        )
        (loop $label$6
         (block $label$7
          (br_if $label$7
           (i32.eqz
            (i32.and
             (i32.load8_u
              (i32.add
               (get_local $7)
               (i32.const 24)
              )
             )
             (i32.const 1)
            )
           )
          )
          (call $_ZdlPv
           (i32.load
            (i32.add
             (get_local $7)
             (i32.const 32)
            )
           )
          )
         )
         (block $label$8
          (br_if $label$8
           (i32.eqz
            (i32.and
             (i32.load8_u
              (i32.add
               (get_local $7)
               (i32.const 12)
              )
             )
             (i32.const 1)
            )
           )
          )
          (call $_ZdlPv
           (i32.load
            (i32.add
             (get_local $7)
             (i32.const 20)
            )
           )
          )
         )
         (block $label$9
          (br_if $label$9
           (i32.eqz
            (i32.and
             (i32.load8_u
              (get_local $7)
             )
             (i32.const 1)
            )
           )
          )
          (call $_ZdlPv
           (i32.load
            (i32.add
             (get_local $7)
             (i32.const 8)
            )
           )
          )
         )
         (br_if $label$6
          (i32.ne
           (i32.add
            (tee_local $7
             (i32.add
              (get_local $7)
              (i32.const -36)
             )
            )
            (get_local $5)
           )
           (i32.const -36)
          )
         )
        )
        (set_local $7
         (i32.load
          (get_local $0)
         )
        )
        (br $label$3)
       )
       (block $label$10
        (br_if $label$10
         (i32.eq
          (tee_local $5
           (select
            (tee_local $7
             (i32.add
              (get_local $1)
              (i32.mul
               (tee_local $4
                (i32.div_s
                 (i32.sub
                  (i32.load offset=4
                   (get_local $0)
                  )
                  (get_local $6)
                 )
                 (i32.const 36)
                )
               )
               (i32.const 36)
              )
             )
            )
            (get_local $2)
            (i32.gt_u
             (get_local $3)
             (get_local $4)
            )
           )
          )
          (get_local $1)
         )
        )
        (loop $label$11
         (drop
          (call $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEaSERKS5_
           (get_local $6)
           (get_local $1)
          )
         )
         (drop
          (call $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEaSERKS5_
           (i32.add
            (get_local $6)
            (i32.const 12)
           )
           (i32.add
            (get_local $1)
            (i32.const 12)
           )
          )
         )
         (drop
          (call $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEaSERKS5_
           (i32.add
            (get_local $6)
            (i32.const 24)
           )
           (i32.add
            (get_local $1)
            (i32.const 24)
           )
          )
         )
         (set_local $6
          (i32.add
           (get_local $6)
           (i32.const 36)
          )
         )
         (br_if $label$11
          (i32.ne
           (tee_local $1
            (i32.add
             (get_local $1)
             (i32.const 36)
            )
           )
           (get_local $5)
          )
         )
        )
       )
       (block $label$12
        (br_if $label$12
         (i32.le_u
          (get_local $3)
          (get_local $4)
         )
        )
        (br_if $label$1
         (i32.eq
          (get_local $5)
          (get_local $2)
         )
        )
        (set_local $1
         (i32.load
          (tee_local $6
           (i32.add
            (get_local $0)
            (i32.const 4)
           )
          )
         )
        )
        (loop $label$13
         (drop
          (call $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEC2ERKS5_
           (get_local $1)
           (get_local $7)
          )
         )
         (drop
          (call $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEC2ERKS5_
           (i32.add
            (get_local $1)
            (i32.const 12)
           )
           (i32.add
            (get_local $7)
            (i32.const 12)
           )
          )
         )
         (drop
          (call $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEC2ERKS5_
           (i32.add
            (get_local $1)
            (i32.const 24)
           )
           (i32.add
            (get_local $7)
            (i32.const 24)
           )
          )
         )
         (i32.store
          (get_local $6)
          (tee_local $1
           (i32.add
            (i32.load
             (get_local $6)
            )
            (i32.const 36)
           )
          )
         )
         (br_if $label$13
          (i32.ne
           (tee_local $7
            (i32.add
             (get_local $7)
             (i32.const 36)
            )
           )
           (get_local $2)
          )
         )
         (br $label$1)
        )
       )
       (block $label$14
        (br_if $label$14
         (i32.eq
          (tee_local $1
           (i32.load
            (tee_local $2
             (i32.add
              (get_local $0)
              (i32.const 4)
             )
            )
           )
          )
          (get_local $6)
         )
        )
        (set_local $7
         (i32.sub
          (i32.const 0)
          (get_local $6)
         )
        )
        (set_local $1
         (i32.add
          (get_local $1)
          (i32.const -36)
         )
        )
        (loop $label$15
         (block $label$16
          (br_if $label$16
           (i32.eqz
            (i32.and
             (i32.load8_u
              (i32.add
               (get_local $1)
               (i32.const 24)
              )
             )
             (i32.const 1)
            )
           )
          )
          (call $_ZdlPv
           (i32.load
            (i32.add
             (get_local $1)
             (i32.const 32)
            )
           )
          )
         )
         (block $label$17
          (br_if $label$17
           (i32.eqz
            (i32.and
             (i32.load8_u
              (i32.add
               (get_local $1)
               (i32.const 12)
              )
             )
             (i32.const 1)
            )
           )
          )
          (call $_ZdlPv
           (i32.load
            (i32.add
             (get_local $1)
             (i32.const 20)
            )
           )
          )
         )
         (block $label$18
          (br_if $label$18
           (i32.eqz
            (i32.and
             (i32.load8_u
              (get_local $1)
             )
             (i32.const 1)
            )
           )
          )
          (call $_ZdlPv
           (i32.load
            (i32.add
             (get_local $1)
             (i32.const 8)
            )
           )
          )
         )
         (br_if $label$15
          (i32.ne
           (i32.add
            (tee_local $1
             (i32.add
              (get_local $1)
              (i32.const -36)
             )
            )
            (get_local $7)
           )
           (i32.const -36)
          )
         )
        )
       )
       (i32.store
        (get_local $2)
        (get_local $6)
       )
       (return)
      )
      (set_local $7
       (get_local $6)
      )
     )
     (i32.store
      (i32.add
       (get_local $0)
       (i32.const 4)
      )
      (get_local $6)
     )
     (call $_ZdlPv
      (get_local $7)
     )
     (set_local $7
      (i32.const 0)
     )
     (i32.store
      (i32.add
       (get_local $0)
       (i32.const 8)
      )
      (i32.const 0)
     )
     (i64.store align=4
      (get_local $0)
      (i64.const 0)
     )
    )
    (br_if $label$0
     (i32.ge_u
      (get_local $3)
      (i32.const 119304648)
     )
    )
    (set_local $6
     (i32.const 119304647)
    )
    (block $label$19
     (br_if $label$19
      (i32.gt_u
       (tee_local $7
        (i32.div_s
         (get_local $7)
         (i32.const 36)
        )
       )
       (i32.const 59652322)
      )
     )
     (set_local $6
      (select
       (get_local $3)
       (tee_local $6
        (i32.shl
         (get_local $7)
         (i32.const 1)
        )
       )
       (i32.lt_u
        (get_local $6)
        (get_local $3)
       )
      )
     )
    )
    (i32.store
     (get_local $0)
     (tee_local $6
      (call $_Znwj
       (tee_local $7
        (i32.mul
         (get_local $6)
         (i32.const 36)
        )
       )
      )
     )
    )
    (i32.store offset=4
     (get_local $0)
     (get_local $6)
    )
    (i32.store
     (i32.add
      (get_local $0)
      (i32.const 8)
     )
     (i32.add
      (get_local $6)
      (get_local $7)
     )
    )
    (br_if $label$1
     (i32.eq
      (get_local $1)
      (get_local $2)
     )
    )
    (set_local $7
     (i32.add
      (get_local $0)
      (i32.const 4)
     )
    )
    (loop $label$20
     (drop
      (call $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEC2ERKS5_
       (get_local $6)
       (get_local $1)
      )
     )
     (drop
      (call $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEC2ERKS5_
       (i32.add
        (get_local $6)
        (i32.const 12)
       )
       (i32.add
        (get_local $1)
        (i32.const 12)
       )
      )
     )
     (drop
      (call $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEC2ERKS5_
       (i32.add
        (get_local $6)
        (i32.const 24)
       )
       (i32.add
        (get_local $1)
        (i32.const 24)
       )
      )
     )
     (i32.store
      (get_local $7)
      (tee_local $6
       (i32.add
        (i32.load
         (get_local $7)
        )
        (i32.const 36)
       )
      )
     )
     (br_if $label$20
      (i32.ne
       (tee_local $1
        (i32.add
         (get_local $1)
         (i32.const 36)
        )
       )
       (get_local $2)
      )
     )
    )
   )
   (return)
  )
  (call $_ZNKSt3__120__vector_base_commonILb1EE20__throw_length_errorEv
   (get_local $0)
  )
  (unreachable)
 )
 (func $_ZN5eosiolsINS_10datastreamIjEEEERT_S4_RKNS_14discovery_fileE (param $0 i32) (param $1 i32) (result i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i64)
  (local $6 i32)
  (local $7 i32)
  (local $8 i32)
  (i32.store offset=4
   (i32.const 0)
   (tee_local $8
    (i32.sub
     (i32.load offset=4
      (i32.const 0)
     )
     (i32.const 32)
    )
   )
  )
  (i32.store
   (get_local $0)
   (tee_local $6
    (i32.add
     (i32.load
      (get_local $0)
     )
     (i32.const 40)
    )
   )
  )
  (set_local $5
   (i64.extend_u/i32
    (i32.shr_s
     (i32.sub
      (tee_local $2
       (i32.load
        (i32.add
         (get_local $1)
         (i32.const 44)
        )
       )
      )
      (tee_local $7
       (i32.load offset=40
        (get_local $1)
       )
      )
     )
     (i32.const 5)
    )
   )
  )
  (loop $label$0
   (set_local $6
    (i32.add
     (get_local $6)
     (i32.const 1)
    )
   )
   (br_if $label$0
    (i64.ne
     (tee_local $5
      (i64.shr_u
       (get_local $5)
       (i64.const 7)
      )
     )
     (i64.const 0)
    )
   )
  )
  (i32.store
   (get_local $0)
   (get_local $6)
  )
  (block $label$1
   (br_if $label$1
    (i32.eq
     (get_local $7)
     (get_local $2)
    )
   )
   (loop $label$2
    (i32.store
     (get_local $0)
     (tee_local $6
      (i32.add
       (get_local $6)
       (i32.const 8)
      )
     )
    )
    (set_local $5
     (i64.extend_u/i32
      (select
       (i32.load
        (tee_local $4
         (i32.add
          (get_local $7)
          (i32.const 12)
         )
        )
       )
       (i32.shr_u
        (tee_local $3
         (i32.load8_u offset=8
          (get_local $7)
         )
        )
        (i32.const 1)
       )
       (i32.and
        (get_local $3)
        (i32.const 1)
       )
      )
     )
    )
    (loop $label$3
     (set_local $6
      (i32.add
       (get_local $6)
       (i32.const 1)
      )
     )
     (br_if $label$3
      (i64.ne
       (tee_local $5
        (i64.shr_u
         (get_local $5)
         (i64.const 7)
        )
       )
       (i64.const 0)
      )
     )
    )
    (i32.store
     (get_local $0)
     (get_local $6)
    )
    (block $label$4
     (br_if $label$4
      (i32.eqz
       (tee_local $4
        (select
         (i32.load
          (get_local $4)
         )
         (i32.shr_u
          (tee_local $4
           (i32.load8_u
            (i32.add
             (get_local $7)
             (i32.const 8)
            )
           )
          )
          (i32.const 1)
         )
         (i32.and
          (get_local $4)
          (i32.const 1)
         )
        )
       )
      )
     )
     (i32.store
      (get_local $0)
      (tee_local $6
       (i32.add
        (get_local $4)
        (get_local $6)
       )
      )
     )
    )
    (i32.store
     (get_local $0)
     (tee_local $6
      (i32.add
       (get_local $6)
       (i32.const 8)
      )
     )
    )
    (br_if $label$2
     (i32.ne
      (tee_local $7
       (i32.add
        (get_local $7)
        (i32.const 32)
       )
      )
      (get_local $2)
     )
    )
   )
  )
  (i32.store
   (get_local $0)
   (tee_local $6
    (i32.add
     (get_local $6)
     (i32.const 8)
    )
   )
  )
  (set_local $5
   (i64.extend_u/i32
    (i32.div_s
     (i32.sub
      (tee_local $3
       (i32.load
        (i32.add
         (get_local $1)
         (i32.const 68)
        )
       )
      )
      (tee_local $7
       (i32.load offset=64
        (get_local $1)
       )
      )
     )
     (i32.const 12)
    )
   )
  )
  (loop $label$5
   (set_local $6
    (i32.add
     (get_local $6)
     (i32.const 1)
    )
   )
   (br_if $label$5
    (i64.ne
     (tee_local $5
      (i64.shr_u
       (get_local $5)
       (i64.const 7)
      )
     )
     (i64.const 0)
    )
   )
  )
  (i32.store
   (get_local $0)
   (get_local $6)
  )
  (block $label$6
   (br_if $label$6
    (i32.eq
     (get_local $7)
     (get_local $3)
    )
   )
   (loop $label$7
    (set_local $5
     (i64.extend_u/i32
      (select
       (i32.load offset=4
        (get_local $7)
       )
       (i32.shr_u
        (tee_local $4
         (i32.load8_u
          (get_local $7)
         )
        )
        (i32.const 1)
       )
       (i32.and
        (get_local $4)
        (i32.const 1)
       )
      )
     )
    )
    (loop $label$8
     (set_local $6
      (i32.add
       (get_local $6)
       (i32.const 1)
      )
     )
     (br_if $label$8
      (i64.ne
       (tee_local $5
        (i64.shr_u
         (get_local $5)
         (i64.const 7)
        )
       )
       (i64.const 0)
      )
     )
    )
    (i32.store
     (get_local $0)
     (get_local $6)
    )
    (block $label$9
     (br_if $label$9
      (i32.eqz
       (tee_local $4
        (select
         (i32.load
          (i32.add
           (get_local $7)
           (i32.const 4)
          )
         )
         (i32.shr_u
          (tee_local $4
           (i32.load8_u
            (get_local $7)
           )
          )
          (i32.const 1)
         )
         (i32.and
          (get_local $4)
          (i32.const 1)
         )
        )
       )
      )
     )
     (i32.store
      (get_local $0)
      (tee_local $6
       (i32.add
        (get_local $4)
        (get_local $6)
       )
      )
     )
    )
    (br_if $label$7
     (i32.ne
      (tee_local $7
       (i32.add
        (get_local $7)
        (i32.const 12)
       )
      )
      (get_local $3)
     )
    )
   )
  )
  (i32.store
   (get_local $0)
   (tee_local $6
    (i32.add
     (get_local $6)
     (i32.const 35)
    )
   )
  )
  (set_local $5
   (i64.extend_u/i32
    (select
     (i32.load
      (i32.add
       (get_local $1)
       (i32.const 116)
      )
     )
     (i32.shr_u
      (tee_local $7
       (i32.load8_u offset=112
        (get_local $1)
       )
      )
      (i32.const 1)
     )
     (i32.and
      (get_local $7)
      (i32.const 1)
     )
    )
   )
  )
  (loop $label$10
   (set_local $6
    (i32.add
     (get_local $6)
     (i32.const 1)
    )
   )
   (br_if $label$10
    (i64.ne
     (tee_local $5
      (i64.shr_u
       (get_local $5)
       (i64.const 7)
      )
     )
     (i64.const 0)
    )
   )
  )
  (i32.store
   (get_local $0)
   (get_local $6)
  )
  (block $label$11
   (br_if $label$11
    (i32.eqz
     (tee_local $7
      (select
       (i32.load
        (i32.add
         (get_local $1)
         (i32.const 116)
        )
       )
       (i32.shr_u
        (tee_local $7
         (i32.load8_u
          (i32.add
           (get_local $1)
           (i32.const 112)
          )
         )
        )
        (i32.const 1)
       )
       (i32.and
        (get_local $7)
        (i32.const 1)
       )
      )
     )
    )
   )
   (i32.store
    (get_local $0)
    (tee_local $6
     (i32.add
      (get_local $7)
      (get_local $6)
     )
    )
   )
  )
  (i32.store
   (get_local $0)
   (i32.add
    (get_local $6)
    (i32.const 8)
   )
  )
  (set_local $6
   (i32.add
    (get_local $6)
    (i32.const 41)
   )
  )
  (set_local $5
   (i64.load32_u offset=136
    (get_local $1)
   )
  )
  (loop $label$12
   (set_local $6
    (i32.add
     (get_local $6)
     (i32.const 1)
    )
   )
   (br_if $label$12
    (i64.ne
     (tee_local $5
      (i64.shr_u
       (get_local $5)
       (i64.const 7)
      )
     )
     (i64.const 0)
    )
   )
  )
  (i32.store
   (get_local $0)
   (get_local $6)
  )
  (set_local $7
   (i32.load
    (tee_local $0
     (call $_ZN5eosiolsINS_10datastreamIjEEEERT_S4_RKNS_13authority_mapE
      (get_local $0)
      (i32.add
       (get_local $1)
       (i32.const 176)
      )
     )
    )
   )
  )
  (set_local $5
   (i64.extend_u/i32
    (i32.div_s
     (i32.sub
      (tee_local $4
       (i32.load
        (i32.add
         (get_local $1)
         (i32.const 264)
        )
       )
      )
      (tee_local $6
       (i32.load offset=260
        (get_local $1)
       )
      )
     )
     (i32.const 36)
    )
   )
  )
  (loop $label$13
   (set_local $7
    (i32.add
     (get_local $7)
     (i32.const 1)
    )
   )
   (br_if $label$13
    (i64.ne
     (tee_local $5
      (i64.shr_u
       (get_local $5)
       (i64.const 7)
      )
     )
     (i64.const 0)
    )
   )
  )
  (i32.store
   (get_local $0)
   (get_local $7)
  )
  (block $label$14
   (br_if $label$14
    (i32.eq
     (get_local $6)
     (get_local $4)
    )
   )
   (set_local $7
    (i32.add
     (get_local $8)
     (i32.const 24)
    )
   )
   (loop $label$15
    (i32.store offset=8
     (get_local $8)
     (get_local $0)
    )
    (i32.store offset=16
     (get_local $8)
     (get_local $6)
    )
    (i32.store
     (get_local $7)
     (i32.add
      (get_local $6)
      (i32.const 24)
     )
    )
    (i32.store offset=20
     (get_local $8)
     (i32.add
      (get_local $6)
      (i32.const 12)
     )
    )
    (call $_ZN5boost3pfr6detail19for_each_field_implINS1_14sequence_tuple5tupleIJRKNSt3__112basic_stringIcNS5_11char_traitsIcEENS5_9allocatorIcEEEESD_SD_EEEZN5eosiolsINSF_10datastreamIjEENSF_11content_refELi0EEERT_SL_RKT0_EUlRKSK_E_JLj0ELj1ELj2EEEEvSL_OSM_NS5_16integer_sequenceIjJXspT1_EEEENS5_17integral_constantIbLb0EEE
     (i32.add
      (get_local $8)
      (i32.const 16)
     )
     (i32.add
      (get_local $8)
      (i32.const 8)
     )
    )
    (br_if $label$15
     (i32.ne
      (tee_local $6
       (i32.add
        (get_local $6)
        (i32.const 36)
       )
      )
      (get_local $4)
     )
    )
   )
  )
  (i32.store offset=4
   (i32.const 0)
   (i32.add
    (get_local $8)
    (i32.const 32)
   )
  )
  (get_local $0)
 )
 (func $_ZN5eosiolsINS_10datastreamIPcEEEERT_S5_RKNS_14discovery_fileE (param $0 i32) (param $1 i32) (result i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i64)
  (local $6 i32)
  (local $7 i32)
  (i32.store offset=4
   (i32.const 0)
   (tee_local $7
    (i32.sub
     (i32.load offset=4
      (i32.const 0)
     )
     (i32.const 16)
    )
   )
  )
  (call $eosio_assert
   (i32.gt_s
    (i32.sub
     (i32.load offset=8
      (get_local $0)
     )
     (i32.load offset=4
      (get_local $0)
     )
    )
    (i32.const 31)
   )
   (i32.const 368)
  )
  (drop
   (call $memcpy
    (i32.load offset=4
     (get_local $0)
    )
    (get_local $1)
    (i32.const 32)
   )
  )
  (i32.store offset=4
   (get_local $0)
   (tee_local $6
    (i32.add
     (i32.load offset=4
      (get_local $0)
     )
     (i32.const 32)
    )
   )
  )
  (call $eosio_assert
   (i32.gt_s
    (i32.sub
     (i32.load offset=8
      (get_local $0)
     )
     (get_local $6)
    )
    (i32.const 7)
   )
   (i32.const 368)
  )
  (drop
   (call $memcpy
    (i32.load offset=4
     (get_local $0)
    )
    (i32.add
     (get_local $1)
     (i32.const 32)
    )
    (i32.const 8)
   )
  )
  (i32.store offset=4
   (get_local $0)
   (i32.add
    (i32.load offset=4
     (get_local $0)
    )
    (i32.const 8)
   )
  )
  (call $eosio_assert
   (i32.gt_s
    (i32.sub
     (i32.load offset=8
      (tee_local $2
       (call $_ZN5eosiolsINS_10datastreamIPcEENS_9peer_linkEEERT_S6_RKNSt3__16vectorIT0_NS7_9allocatorIS9_EEEE
        (get_local $0)
        (i32.add
         (get_local $1)
         (i32.const 40)
        )
       )
      )
     )
     (i32.load offset=4
      (get_local $2)
     )
    )
    (i32.const 7)
   )
   (i32.const 368)
  )
  (drop
   (call $memcpy
    (i32.load offset=4
     (get_local $2)
    )
    (i32.add
     (get_local $1)
     (i32.const 56)
    )
    (i32.const 8)
   )
  )
  (i32.store offset=4
   (get_local $2)
   (tee_local $6
    (i32.add
     (i32.load offset=4
      (get_local $2)
     )
     (i32.const 8)
    )
   )
  )
  (set_local $5
   (i64.extend_u/i32
    (i32.div_s
     (i32.sub
      (i32.load
       (i32.add
        (get_local $1)
        (i32.const 68)
       )
      )
      (i32.load offset=64
       (get_local $1)
      )
     )
     (i32.const 12)
    )
   )
  )
  (set_local $0
   (i32.add
    (get_local $2)
    (i32.const 4)
   )
  )
  (loop $label$0
   (set_local $3
    (i32.wrap/i64
     (get_local $5)
    )
   )
   (i32.store8 offset=15
    (get_local $7)
    (i32.or
     (i32.shl
      (tee_local $4
       (i64.ne
        (tee_local $5
         (i64.shr_u
          (get_local $5)
          (i64.const 7)
         )
        )
        (i64.const 0)
       )
      )
      (i32.const 7)
     )
     (i32.and
      (get_local $3)
      (i32.const 127)
     )
    )
   )
   (call $eosio_assert
    (i32.gt_s
     (i32.sub
      (i32.load
       (i32.add
        (get_local $2)
        (i32.const 8)
       )
      )
      (get_local $6)
     )
     (i32.const 0)
    )
    (i32.const 368)
   )
   (drop
    (call $memcpy
     (i32.load
      (get_local $0)
     )
     (i32.add
      (get_local $7)
      (i32.const 15)
     )
     (i32.const 1)
    )
   )
   (i32.store
    (get_local $0)
    (tee_local $6
     (i32.add
      (i32.load
       (get_local $0)
      )
      (i32.const 1)
     )
    )
   )
   (br_if $label$0
    (get_local $4)
   )
  )
  (block $label$1
   (br_if $label$1
    (i32.eq
     (tee_local $0
      (i32.load
       (i32.add
        (get_local $1)
        (i32.const 64)
       )
      )
     )
     (tee_local $3
      (i32.load
       (i32.add
        (get_local $1)
        (i32.const 68)
       )
      )
     )
    )
   )
   (loop $label$2
    (drop
     (call $_ZN5eosiolsINS_10datastreamIPcEEEERT_S5_RKNSt3__112basic_stringIcNS6_11char_traitsIcEENS6_9allocatorIcEEEE
      (get_local $2)
      (get_local $0)
     )
    )
    (br_if $label$2
     (i32.ne
      (get_local $3)
      (tee_local $0
       (i32.add
        (get_local $0)
        (i32.const 12)
       )
      )
     )
    )
   )
   (set_local $6
    (i32.load
     (i32.add
      (get_local $2)
      (i32.const 4)
     )
    )
   )
  )
  (call $eosio_assert
   (i32.gt_s
    (i32.sub
     (i32.load
      (tee_local $3
       (i32.add
        (get_local $2)
        (i32.const 8)
       )
      )
     )
     (get_local $6)
    )
    (i32.const 1)
   )
   (i32.const 368)
  )
  (drop
   (call $memcpy
    (i32.load
     (tee_local $0
      (i32.add
       (get_local $2)
       (i32.const 4)
      )
     )
    )
    (i32.add
     (get_local $1)
     (i32.const 76)
    )
    (i32.const 2)
   )
  )
  (i32.store
   (get_local $0)
   (tee_local $6
    (i32.add
     (i32.load
      (get_local $0)
     )
     (i32.const 2)
    )
   )
  )
  (call $eosio_assert
   (i32.gt_s
    (i32.sub
     (i32.load
      (get_local $3)
     )
     (get_local $6)
    )
    (i32.const 0)
   )
   (i32.const 368)
  )
  (drop
   (call $memcpy
    (i32.load
     (get_local $0)
    )
    (i32.add
     (get_local $1)
     (i32.const 78)
    )
    (i32.const 1)
   )
  )
  (i32.store
   (get_local $0)
   (tee_local $6
    (i32.add
     (i32.load
      (get_local $0)
     )
     (i32.const 1)
    )
   )
  )
  (call $eosio_assert
   (i32.gt_s
    (i32.sub
     (i32.load
      (get_local $3)
     )
     (get_local $6)
    )
    (i32.const 31)
   )
   (i32.const 368)
  )
  (drop
   (call $memcpy
    (i32.load
     (get_local $0)
    )
    (i32.add
     (get_local $1)
     (i32.const 79)
    )
    (i32.const 32)
   )
  )
  (i32.store
   (get_local $0)
   (i32.add
    (i32.load
     (get_local $0)
    )
    (i32.const 32)
   )
  )
  (call $eosio_assert
   (i32.gt_s
    (i32.sub
     (i32.load offset=8
      (tee_local $0
       (call $_ZN5eosiolsINS_10datastreamIPcEEEERT_S5_RKNSt3__112basic_stringIcNS6_11char_traitsIcEENS6_9allocatorIcEEEE
        (get_local $2)
        (i32.add
         (get_local $1)
         (i32.const 112)
        )
       )
      )
     )
     (i32.load offset=4
      (get_local $0)
     )
    )
    (i32.const 7)
   )
   (i32.const 368)
  )
  (drop
   (call $memcpy
    (i32.load offset=4
     (get_local $0)
    )
    (i32.add
     (get_local $1)
     (i32.const 128)
    )
    (i32.const 8)
   )
  )
  (i32.store offset=4
   (get_local $0)
   (i32.add
    (i32.load offset=4
     (get_local $0)
    )
    (i32.const 8)
   )
  )
  (set_local $0
   (call $_ZN5eosiolsINS_10datastreamIPcEENS_11content_refEEERT_S6_RKNSt3__16vectorIT0_NS7_9allocatorIS9_EEEE
    (call $_ZN5eosiolsINS_10datastreamIPcEEEERT_S5_RKNS_13authority_mapE
     (call $_ZN5eosiolsINS_10datastreamIPcEEEERT_S5_RKNS_10public_keyE
      (get_local $0)
      (i32.add
       (get_local $1)
       (i32.const 136)
      )
     )
     (i32.add
      (get_local $1)
      (i32.const 176)
     )
    )
    (i32.add
     (get_local $1)
     (i32.const 260)
    )
   )
  )
  (i32.store offset=4
   (i32.const 0)
   (i32.add
    (get_local $7)
    (i32.const 16)
   )
  )
  (get_local $0)
 )
 (func $_ZN5eosiolsINS_10datastreamIPcEENS_9peer_linkEEERT_S6_RKNSt3__16vectorIT0_NS7_9allocatorIS9_EEEE (param $0 i32) (param $1 i32) (result i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i64)
  (local $6 i32)
  (local $7 i32)
  (local $8 i32)
  (i32.store offset=4
   (i32.const 0)
   (tee_local $8
    (i32.sub
     (i32.load offset=4
      (i32.const 0)
     )
     (i32.const 16)
    )
   )
  )
  (set_local $5
   (i64.extend_u/i32
    (i32.shr_s
     (i32.sub
      (i32.load offset=4
       (get_local $1)
      )
      (i32.load
       (get_local $1)
      )
     )
     (i32.const 5)
    )
   )
  )
  (set_local $6
   (i32.load offset=4
    (get_local $0)
   )
  )
  (set_local $3
   (i32.add
    (get_local $0)
    (i32.const 8)
   )
  )
  (set_local $4
   (i32.add
    (get_local $0)
    (i32.const 4)
   )
  )
  (loop $label$0
   (set_local $7
    (i32.wrap/i64
     (get_local $5)
    )
   )
   (i32.store8 offset=15
    (get_local $8)
    (i32.or
     (i32.shl
      (tee_local $2
       (i64.ne
        (tee_local $5
         (i64.shr_u
          (get_local $5)
          (i64.const 7)
         )
        )
        (i64.const 0)
       )
      )
      (i32.const 7)
     )
     (i32.and
      (get_local $7)
      (i32.const 127)
     )
    )
   )
   (call $eosio_assert
    (i32.gt_s
     (i32.sub
      (i32.load
       (get_local $3)
      )
      (get_local $6)
     )
     (i32.const 0)
    )
    (i32.const 368)
   )
   (drop
    (call $memcpy
     (i32.load
      (get_local $4)
     )
     (i32.add
      (get_local $8)
      (i32.const 15)
     )
     (i32.const 1)
    )
   )
   (i32.store
    (get_local $4)
    (tee_local $6
     (i32.add
      (i32.load
       (get_local $4)
      )
      (i32.const 1)
     )
    )
   )
   (br_if $label$0
    (get_local $2)
   )
  )
  (block $label$1
   (br_if $label$1
    (i32.eq
     (tee_local $7
      (i32.load
       (get_local $1)
      )
     )
     (tee_local $3
      (i32.load
       (i32.add
        (get_local $1)
        (i32.const 4)
       )
      )
     )
    )
   )
   (set_local $4
    (i32.add
     (get_local $0)
     (i32.const 4)
    )
   )
   (loop $label$2
    (call $eosio_assert
     (i32.gt_s
      (i32.sub
       (i32.load
        (tee_local $2
         (i32.add
          (get_local $0)
          (i32.const 8)
         )
        )
       )
       (get_local $6)
      )
      (i32.const 7)
     )
     (i32.const 368)
    )
    (drop
     (call $memcpy
      (i32.load
       (get_local $4)
      )
      (get_local $7)
      (i32.const 8)
     )
    )
    (i32.store
     (get_local $4)
     (i32.add
      (i32.load
       (get_local $4)
      )
      (i32.const 8)
     )
    )
    (drop
     (call $_ZN5eosiolsINS_10datastreamIPcEEEERT_S5_RKNSt3__112basic_stringIcNS6_11char_traitsIcEENS6_9allocatorIcEEEE
      (get_local $0)
      (i32.add
       (get_local $7)
       (i32.const 8)
      )
     )
    )
    (call $eosio_assert
     (i32.gt_s
      (i32.sub
       (i32.load
        (get_local $2)
       )
       (i32.load
        (get_local $4)
       )
      )
      (i32.const 7)
     )
     (i32.const 368)
    )
    (drop
     (call $memcpy
      (i32.load
       (get_local $4)
      )
      (i32.add
       (get_local $7)
       (i32.const 24)
      )
      (i32.const 8)
     )
    )
    (i32.store
     (get_local $4)
     (tee_local $6
      (i32.add
       (i32.load
        (get_local $4)
       )
       (i32.const 8)
      )
     )
    )
    (br_if $label$2
     (i32.ne
      (tee_local $7
       (i32.add
        (get_local $7)
        (i32.const 32)
       )
      )
      (get_local $3)
     )
    )
   )
  )
  (i32.store offset=4
   (i32.const 0)
   (i32.add
    (get_local $8)
    (i32.const 16)
   )
  )
  (get_local $0)
 )
 (func $_ZN5eosiolsINS_10datastreamIPcEEEERT_S5_RKNSt3__112basic_stringIcNS6_11char_traitsIcEENS6_9allocatorIcEEEE (param $0 i32) (param $1 i32) (result i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i32)
  (local $6 i64)
  (local $7 i32)
  (local $8 i32)
  (i32.store offset=4
   (i32.const 0)
   (tee_local $8
    (i32.sub
     (i32.load offset=4
      (i32.const 0)
     )
     (i32.const 16)
    )
   )
  )
  (set_local $6
   (i64.extend_u/i32
    (select
     (i32.load offset=4
      (get_local $1)
     )
     (i32.shr_u
      (tee_local $7
       (i32.load8_u
        (get_local $1)
       )
      )
      (i32.const 1)
     )
     (i32.and
      (get_local $7)
      (i32.const 1)
     )
    )
   )
  )
  (set_local $7
   (i32.load offset=4
    (get_local $0)
   )
  )
  (set_local $4
   (i32.add
    (get_local $0)
    (i32.const 8)
   )
  )
  (set_local $5
   (i32.add
    (get_local $0)
    (i32.const 4)
   )
  )
  (loop $label$0
   (set_local $2
    (i32.wrap/i64
     (get_local $6)
    )
   )
   (i32.store8 offset=15
    (get_local $8)
    (i32.or
     (i32.shl
      (tee_local $3
       (i64.ne
        (tee_local $6
         (i64.shr_u
          (get_local $6)
          (i64.const 7)
         )
        )
        (i64.const 0)
       )
      )
      (i32.const 7)
     )
     (i32.and
      (get_local $2)
      (i32.const 127)
     )
    )
   )
   (call $eosio_assert
    (i32.gt_s
     (i32.sub
      (i32.load
       (get_local $4)
      )
      (get_local $7)
     )
     (i32.const 0)
    )
    (i32.const 368)
   )
   (drop
    (call $memcpy
     (i32.load
      (get_local $5)
     )
     (i32.add
      (get_local $8)
      (i32.const 15)
     )
     (i32.const 1)
    )
   )
   (i32.store
    (get_local $5)
    (tee_local $7
     (i32.add
      (i32.load
       (get_local $5)
      )
      (i32.const 1)
     )
    )
   )
   (br_if $label$0
    (get_local $3)
   )
  )
  (block $label$1
   (br_if $label$1
    (i32.eqz
     (tee_local $3
      (select
       (i32.load
        (i32.add
         (get_local $1)
         (i32.const 4)
        )
       )
       (i32.shr_u
        (tee_local $5
         (i32.load8_u
          (get_local $1)
         )
        )
        (i32.const 1)
       )
       (tee_local $5
        (i32.and
         (get_local $5)
         (i32.const 1)
        )
       )
      )
     )
    )
   )
   (set_local $5
    (select
     (i32.load offset=8
      (get_local $1)
     )
     (i32.add
      (get_local $1)
      (i32.const 1)
     )
     (get_local $5)
    )
   )
   (set_local $4
    (i32.add
     (get_local $0)
     (i32.const 8)
    )
   )
   (set_local $2
    (i32.add
     (get_local $0)
     (i32.const 4)
    )
   )
   (loop $label$2
    (call $eosio_assert
     (i32.gt_s
      (i32.sub
       (i32.load
        (get_local $4)
       )
       (get_local $7)
      )
      (i32.const 0)
     )
     (i32.const 368)
    )
    (drop
     (call $memcpy
      (i32.load
       (get_local $2)
      )
      (get_local $5)
      (i32.const 1)
     )
    )
    (i32.store
     (get_local $2)
     (tee_local $7
      (i32.add
       (i32.load
        (get_local $2)
       )
       (i32.const 1)
      )
     )
    )
    (set_local $5
     (i32.add
      (get_local $5)
      (i32.const 1)
     )
    )
    (br_if $label$2
     (tee_local $3
      (i32.add
       (get_local $3)
       (i32.const -1)
      )
     )
    )
   )
  )
  (i32.store offset=4
   (i32.const 0)
   (i32.add
    (get_local $8)
    (i32.const 16)
   )
  )
  (get_local $0)
 )
 (func $_ZN5eosiolsINS_10datastreamIPcEEEERT_S5_RKNS_10public_keyE (param $0 i32) (param $1 i32) (result i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i32)
  (local $6 i64)
  (local $7 i32)
  (local $8 i32)
  (i32.store offset=4
   (i32.const 0)
   (tee_local $8
    (i32.sub
     (i32.load offset=4
      (i32.const 0)
     )
     (i32.const 16)
    )
   )
  )
  (set_local $7
   (i32.load offset=4
    (get_local $0)
   )
  )
  (set_local $6
   (i64.load32_u
    (get_local $1)
   )
  )
  (set_local $3
   (i32.add
    (get_local $0)
    (i32.const 8)
   )
  )
  (set_local $4
   (i32.add
    (get_local $0)
    (i32.const 4)
   )
  )
  (loop $label$0
   (set_local $5
    (i32.wrap/i64
     (get_local $6)
    )
   )
   (i32.store8 offset=15
    (get_local $8)
    (i32.or
     (i32.shl
      (tee_local $2
       (i64.ne
        (tee_local $6
         (i64.shr_u
          (get_local $6)
          (i64.const 7)
         )
        )
        (i64.const 0)
       )
      )
      (i32.const 7)
     )
     (i32.and
      (get_local $5)
      (i32.const 127)
     )
    )
   )
   (call $eosio_assert
    (i32.gt_s
     (i32.sub
      (i32.load
       (get_local $3)
      )
      (get_local $7)
     )
     (i32.const 0)
    )
    (i32.const 368)
   )
   (drop
    (call $memcpy
     (i32.load
      (get_local $4)
     )
     (i32.add
      (get_local $8)
      (i32.const 15)
     )
     (i32.const 1)
    )
   )
   (i32.store
    (get_local $4)
    (tee_local $7
     (i32.add
      (i32.load
       (get_local $4)
      )
      (i32.const 1)
     )
    )
   )
   (br_if $label$0
    (get_local $2)
   )
  )
  (set_local $2
   (i32.add
    (get_local $0)
    (i32.const 8)
   )
  )
  (set_local $4
   (i32.const 4)
  )
  (loop $label$1
   (call $eosio_assert
    (i32.gt_s
     (i32.sub
      (i32.load
       (get_local $2)
      )
      (get_local $7)
     )
     (i32.const 0)
    )
    (i32.const 368)
   )
   (drop
    (call $memcpy
     (i32.load
      (tee_local $5
       (i32.add
        (get_local $0)
        (i32.const 4)
       )
      )
     )
     (i32.add
      (get_local $1)
      (get_local $4)
     )
     (i32.const 1)
    )
   )
   (i32.store
    (get_local $5)
    (tee_local $7
     (i32.add
      (i32.load
       (get_local $5)
      )
      (i32.const 1)
     )
    )
   )
   (br_if $label$1
    (i32.ne
     (tee_local $4
      (i32.add
       (get_local $4)
       (i32.const 1)
      )
     )
     (i32.const 37)
    )
   )
  )
  (i32.store offset=4
   (i32.const 0)
   (i32.add
    (get_local $8)
    (i32.const 16)
   )
  )
  (get_local $0)
 )
 (func $_ZN5eosiolsINS_10datastreamIPcEEEERT_S5_RKNS_13authority_mapE (param $0 i32) (param $1 i32) (result i32)
  (call $eosio_assert
   (i32.gt_s
    (i32.sub
     (i32.load offset=8
      (get_local $0)
     )
     (i32.load offset=4
      (get_local $0)
     )
    )
    (i32.const 3)
   )
   (i32.const 368)
  )
  (drop
   (call $memcpy
    (i32.load offset=4
     (get_local $0)
    )
    (get_local $1)
    (i32.const 4)
   )
  )
  (i32.store offset=4
   (get_local $0)
   (i32.add
    (i32.load offset=4
     (get_local $0)
    )
    (i32.const 4)
   )
  )
  (call $eosio_assert
   (i32.gt_s
    (i32.sub
     (i32.load offset=8
      (tee_local $0
       (call $_ZN5eosiolsINS_10datastreamIPcEEN11eosiosystem23permission_level_weightEEERT_S7_RKNSt3__16vectorIT0_NS8_9allocatorISA_EEEE
        (call $_ZN5eosiolsINS_10datastreamIPcEEN11eosiosystem10key_weightEEERT_S7_RKNSt3__16vectorIT0_NS8_9allocatorISA_EEEE
         (get_local $0)
         (i32.add
          (get_local $1)
          (i32.const 4)
         )
        )
        (i32.add
         (get_local $1)
         (i32.const 16)
        )
       )
      )
     )
     (i32.load offset=4
      (get_local $0)
     )
    )
    (i32.const 3)
   )
   (i32.const 368)
  )
  (drop
   (call $memcpy
    (i32.load offset=4
     (get_local $0)
    )
    (i32.add
     (get_local $1)
     (i32.const 28)
    )
    (i32.const 4)
   )
  )
  (i32.store offset=4
   (get_local $0)
   (i32.add
    (i32.load offset=4
     (get_local $0)
    )
    (i32.const 4)
   )
  )
  (call $eosio_assert
   (i32.gt_s
    (i32.sub
     (i32.load offset=8
      (tee_local $0
       (call $_ZN5eosiolsINS_10datastreamIPcEEN11eosiosystem23permission_level_weightEEERT_S7_RKNSt3__16vectorIT0_NS8_9allocatorISA_EEEE
        (call $_ZN5eosiolsINS_10datastreamIPcEEN11eosiosystem10key_weightEEERT_S7_RKNSt3__16vectorIT0_NS8_9allocatorISA_EEEE
         (get_local $0)
         (i32.add
          (get_local $1)
          (i32.const 32)
         )
        )
        (i32.add
         (get_local $1)
         (i32.const 44)
        )
       )
      )
     )
     (i32.load offset=4
      (get_local $0)
     )
    )
    (i32.const 3)
   )
   (i32.const 368)
  )
  (drop
   (call $memcpy
    (i32.load offset=4
     (get_local $0)
    )
    (i32.add
     (get_local $1)
     (i32.const 56)
    )
    (i32.const 4)
   )
  )
  (i32.store offset=4
   (get_local $0)
   (i32.add
    (i32.load offset=4
     (get_local $0)
    )
    (i32.const 4)
   )
  )
  (call $_ZN5eosiolsINS_10datastreamIPcEEN11eosiosystem23permission_level_weightEEERT_S7_RKNSt3__16vectorIT0_NS8_9allocatorISA_EEEE
   (call $_ZN5eosiolsINS_10datastreamIPcEEN11eosiosystem10key_weightEEERT_S7_RKNSt3__16vectorIT0_NS8_9allocatorISA_EEEE
    (get_local $0)
    (i32.add
     (get_local $1)
     (i32.const 60)
    )
   )
   (i32.add
    (get_local $1)
    (i32.const 72)
   )
  )
 )
 (func $_ZN5eosiolsINS_10datastreamIPcEENS_11content_refEEERT_S6_RKNSt3__16vectorIT0_NS7_9allocatorIS9_EEEE (param $0 i32) (param $1 i32) (result i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i32)
  (local $6 i64)
  (local $7 i32)
  (local $8 i32)
  (i32.store offset=4
   (i32.const 0)
   (tee_local $8
    (i32.sub
     (i32.load offset=4
      (i32.const 0)
     )
     (i32.const 16)
    )
   )
  )
  (set_local $6
   (i64.extend_u/i32
    (i32.div_s
     (i32.sub
      (i32.load offset=4
       (get_local $1)
      )
      (i32.load
       (get_local $1)
      )
     )
     (i32.const 36)
    )
   )
  )
  (set_local $5
   (i32.load offset=4
    (get_local $0)
   )
  )
  (set_local $4
   (i32.add
    (get_local $0)
    (i32.const 8)
   )
  )
  (set_local $7
   (i32.add
    (get_local $0)
    (i32.const 4)
   )
  )
  (loop $label$0
   (set_local $2
    (i32.wrap/i64
     (get_local $6)
    )
   )
   (i32.store8 offset=15
    (get_local $8)
    (i32.or
     (i32.shl
      (tee_local $3
       (i64.ne
        (tee_local $6
         (i64.shr_u
          (get_local $6)
          (i64.const 7)
         )
        )
        (i64.const 0)
       )
      )
      (i32.const 7)
     )
     (i32.and
      (get_local $2)
      (i32.const 127)
     )
    )
   )
   (call $eosio_assert
    (i32.gt_s
     (i32.sub
      (i32.load
       (get_local $4)
      )
      (get_local $5)
     )
     (i32.const 0)
    )
    (i32.const 368)
   )
   (drop
    (call $memcpy
     (i32.load
      (get_local $7)
     )
     (i32.add
      (get_local $8)
      (i32.const 15)
     )
     (i32.const 1)
    )
   )
   (i32.store
    (get_local $7)
    (tee_local $5
     (i32.add
      (i32.load
       (get_local $7)
      )
      (i32.const 1)
     )
    )
   )
   (br_if $label$0
    (get_local $3)
   )
  )
  (block $label$1
   (br_if $label$1
    (i32.eq
     (tee_local $7
      (i32.load
       (get_local $1)
      )
     )
     (tee_local $5
      (i32.load
       (i32.add
        (get_local $1)
        (i32.const 4)
       )
      )
     )
    )
   )
   (loop $label$2
    (drop
     (call $_ZN5eosiolsINS_10datastreamIPcEEEERT_S5_RKNSt3__112basic_stringIcNS6_11char_traitsIcEENS6_9allocatorIcEEEE
      (get_local $0)
      (get_local $7)
     )
    )
    (drop
     (call $_ZN5eosiolsINS_10datastreamIPcEEEERT_S5_RKNSt3__112basic_stringIcNS6_11char_traitsIcEENS6_9allocatorIcEEEE
      (get_local $0)
      (i32.add
       (get_local $7)
       (i32.const 12)
      )
     )
    )
    (drop
     (call $_ZN5eosiolsINS_10datastreamIPcEEEERT_S5_RKNSt3__112basic_stringIcNS6_11char_traitsIcEENS6_9allocatorIcEEEE
      (get_local $0)
      (i32.add
       (get_local $7)
       (i32.const 24)
      )
     )
    )
    (br_if $label$2
     (i32.ne
      (tee_local $7
       (i32.add
        (get_local $7)
        (i32.const 36)
       )
      )
      (get_local $5)
     )
    )
   )
  )
  (i32.store offset=4
   (i32.const 0)
   (i32.add
    (get_local $8)
    (i32.const 16)
   )
  )
  (get_local $0)
 )
 (func $_ZN5eosiolsINS_10datastreamIPcEEN11eosiosystem10key_weightEEERT_S7_RKNSt3__16vectorIT0_NS8_9allocatorISA_EEEE (param $0 i32) (param $1 i32) (result i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i32)
  (local $6 i64)
  (local $7 i32)
  (local $8 i32)
  (i32.store offset=4
   (i32.const 0)
   (tee_local $8
    (i32.sub
     (i32.load offset=4
      (i32.const 0)
     )
     (i32.const 16)
    )
   )
  )
  (set_local $6
   (i64.extend_u/i32
    (i32.div_s
     (i32.sub
      (i32.load offset=4
       (get_local $1)
      )
      (i32.load
       (get_local $1)
      )
     )
     (i32.const 44)
    )
   )
  )
  (set_local $7
   (i32.load offset=4
    (get_local $0)
   )
  )
  (set_local $4
   (i32.add
    (get_local $0)
    (i32.const 8)
   )
  )
  (set_local $5
   (i32.add
    (get_local $0)
    (i32.const 4)
   )
  )
  (loop $label$0
   (set_local $2
    (i32.wrap/i64
     (get_local $6)
    )
   )
   (i32.store8 offset=15
    (get_local $8)
    (i32.or
     (i32.shl
      (tee_local $3
       (i64.ne
        (tee_local $6
         (i64.shr_u
          (get_local $6)
          (i64.const 7)
         )
        )
        (i64.const 0)
       )
      )
      (i32.const 7)
     )
     (i32.and
      (get_local $2)
      (i32.const 127)
     )
    )
   )
   (call $eosio_assert
    (i32.gt_s
     (i32.sub
      (i32.load
       (get_local $4)
      )
      (get_local $7)
     )
     (i32.const 0)
    )
    (i32.const 368)
   )
   (drop
    (call $memcpy
     (i32.load
      (get_local $5)
     )
     (i32.add
      (get_local $8)
      (i32.const 15)
     )
     (i32.const 1)
    )
   )
   (i32.store
    (get_local $5)
    (tee_local $7
     (i32.add
      (i32.load
       (get_local $5)
      )
      (i32.const 1)
     )
    )
   )
   (br_if $label$0
    (get_local $3)
   )
  )
  (block $label$1
   (br_if $label$1
    (i32.eq
     (tee_local $7
      (i32.load
       (get_local $1)
      )
     )
     (tee_local $2
      (i32.load
       (i32.add
        (get_local $1)
        (i32.const 4)
       )
      )
     )
    )
   )
   (loop $label$2
    (call $eosio_assert
     (i32.gt_s
      (i32.sub
       (i32.load offset=8
        (tee_local $5
         (call $_ZN5eosiolsINS_10datastreamIPcEEEERT_S5_RKNS_10public_keyE
          (get_local $0)
          (get_local $7)
         )
        )
       )
       (i32.load offset=4
        (get_local $5)
       )
      )
      (i32.const 1)
     )
     (i32.const 368)
    )
    (drop
     (call $memcpy
      (i32.load offset=4
       (get_local $5)
      )
      (i32.add
       (get_local $7)
       (i32.const 40)
      )
      (i32.const 2)
     )
    )
    (i32.store offset=4
     (get_local $5)
     (i32.add
      (i32.load offset=4
       (get_local $5)
      )
      (i32.const 2)
     )
    )
    (br_if $label$2
     (i32.ne
      (tee_local $7
       (i32.add
        (get_local $7)
        (i32.const 44)
       )
      )
      (get_local $2)
     )
    )
   )
  )
  (i32.store offset=4
   (i32.const 0)
   (i32.add
    (get_local $8)
    (i32.const 16)
   )
  )
  (get_local $0)
 )
 (func $_ZN5eosiolsINS_10datastreamIPcEEN11eosiosystem23permission_level_weightEEERT_S7_RKNSt3__16vectorIT0_NS8_9allocatorISA_EEEE (param $0 i32) (param $1 i32) (result i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i64)
  (local $6 i32)
  (local $7 i32)
  (local $8 i32)
  (i32.store offset=4
   (i32.const 0)
   (tee_local $8
    (i32.sub
     (i32.load offset=4
      (i32.const 0)
     )
     (i32.const 16)
    )
   )
  )
  (set_local $5
   (i64.extend_u/i32
    (i32.div_s
     (i32.sub
      (i32.load offset=4
       (get_local $1)
      )
      (i32.load
       (get_local $1)
      )
     )
     (i32.const 24)
    )
   )
  )
  (set_local $6
   (i32.load offset=4
    (get_local $0)
   )
  )
  (set_local $3
   (i32.add
    (get_local $0)
    (i32.const 8)
   )
  )
  (set_local $4
   (i32.add
    (get_local $0)
    (i32.const 4)
   )
  )
  (loop $label$0
   (set_local $7
    (i32.wrap/i64
     (get_local $5)
    )
   )
   (i32.store8 offset=15
    (get_local $8)
    (i32.or
     (i32.shl
      (tee_local $2
       (i64.ne
        (tee_local $5
         (i64.shr_u
          (get_local $5)
          (i64.const 7)
         )
        )
        (i64.const 0)
       )
      )
      (i32.const 7)
     )
     (i32.and
      (get_local $7)
      (i32.const 127)
     )
    )
   )
   (call $eosio_assert
    (i32.gt_s
     (i32.sub
      (i32.load
       (get_local $3)
      )
      (get_local $6)
     )
     (i32.const 0)
    )
    (i32.const 368)
   )
   (drop
    (call $memcpy
     (i32.load
      (get_local $4)
     )
     (i32.add
      (get_local $8)
      (i32.const 15)
     )
     (i32.const 1)
    )
   )
   (i32.store
    (get_local $4)
    (tee_local $6
     (i32.add
      (i32.load
       (get_local $4)
      )
      (i32.const 1)
     )
    )
   )
   (br_if $label$0
    (get_local $2)
   )
  )
  (block $label$1
   (br_if $label$1
    (i32.eq
     (tee_local $7
      (i32.load
       (get_local $1)
      )
     )
     (tee_local $3
      (i32.load
       (i32.add
        (get_local $1)
        (i32.const 4)
       )
      )
     )
    )
   )
   (set_local $4
    (i32.add
     (get_local $0)
     (i32.const 4)
    )
   )
   (loop $label$2
    (call $eosio_assert
     (i32.gt_s
      (i32.sub
       (i32.load
        (tee_local $2
         (i32.add
          (get_local $0)
          (i32.const 8)
         )
        )
       )
       (get_local $6)
      )
      (i32.const 7)
     )
     (i32.const 368)
    )
    (drop
     (call $memcpy
      (i32.load
       (get_local $4)
      )
      (get_local $7)
      (i32.const 8)
     )
    )
    (i32.store
     (get_local $4)
     (tee_local $6
      (i32.add
       (i32.load
        (get_local $4)
       )
       (i32.const 8)
      )
     )
    )
    (call $eosio_assert
     (i32.gt_s
      (i32.sub
       (i32.load
        (get_local $2)
       )
       (get_local $6)
      )
      (i32.const 7)
     )
     (i32.const 368)
    )
    (drop
     (call $memcpy
      (i32.load
       (get_local $4)
      )
      (i32.add
       (get_local $7)
       (i32.const 8)
      )
      (i32.const 8)
     )
    )
    (i32.store
     (get_local $4)
     (tee_local $6
      (i32.add
       (i32.load
        (get_local $4)
       )
       (i32.const 8)
      )
     )
    )
    (call $eosio_assert
     (i32.gt_s
      (i32.sub
       (i32.load
        (get_local $2)
       )
       (get_local $6)
      )
      (i32.const 1)
     )
     (i32.const 368)
    )
    (drop
     (call $memcpy
      (i32.load
       (get_local $4)
      )
      (i32.add
       (get_local $7)
       (i32.const 16)
      )
      (i32.const 2)
     )
    )
    (i32.store
     (get_local $4)
     (tee_local $6
      (i32.add
       (i32.load
        (get_local $4)
       )
       (i32.const 2)
      )
     )
    )
    (br_if $label$2
     (i32.ne
      (tee_local $7
       (i32.add
        (get_local $7)
        (i32.const 24)
       )
      )
      (get_local $3)
     )
    )
   )
  )
  (i32.store offset=4
   (i32.const 0)
   (i32.add
    (get_local $8)
    (i32.const 16)
   )
  )
  (get_local $0)
 )
 (func $_ZN5eosiolsINS_10datastreamIjEEEERT_S4_RKNS_13authority_mapE (param $0 i32) (param $1 i32) (result i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i32)
  (local $6 i64)
  (i32.store
   (get_local $0)
   (tee_local $5
    (i32.add
     (i32.load
      (get_local $0)
     )
     (i32.const 4)
    )
   )
  )
  (set_local $6
   (i64.extend_u/i32
    (i32.div_s
     (i32.sub
      (tee_local $2
       (i32.load
        (i32.add
         (get_local $1)
         (i32.const 8)
        )
       )
      )
      (tee_local $4
       (i32.load offset=4
        (get_local $1)
       )
      )
     )
     (i32.const 44)
    )
   )
  )
  (loop $label$0
   (set_local $5
    (i32.add
     (get_local $5)
     (i32.const 1)
    )
   )
   (br_if $label$0
    (i64.ne
     (tee_local $6
      (i64.shr_u
       (get_local $6)
       (i64.const 7)
      )
     )
     (i64.const 0)
    )
   )
  )
  (i32.store
   (get_local $0)
   (get_local $5)
  )
  (block $label$1
   (br_if $label$1
    (i32.eq
     (get_local $4)
     (get_local $2)
    )
   )
   (loop $label$2
    (set_local $5
     (i32.add
      (get_local $5)
      (i32.const 35)
     )
    )
    (set_local $6
     (i64.load32_u
      (get_local $4)
     )
    )
    (loop $label$3
     (set_local $5
      (i32.add
       (get_local $5)
       (i32.const 1)
      )
     )
     (br_if $label$3
      (i64.ne
       (tee_local $6
        (i64.shr_u
         (get_local $6)
         (i64.const 7)
        )
       )
       (i64.const 0)
      )
     )
    )
    (br_if $label$2
     (i32.ne
      (tee_local $4
       (i32.add
        (get_local $4)
        (i32.const 44)
       )
      )
      (get_local $2)
     )
    )
   )
   (i32.store
    (get_local $0)
    (get_local $5)
   )
  )
  (set_local $6
   (i64.extend_u/i32
    (i32.div_s
     (tee_local $3
      (i32.sub
       (tee_local $4
        (i32.load
         (i32.add
          (get_local $1)
          (i32.const 20)
         )
        )
       )
       (tee_local $2
        (i32.load offset=16
         (get_local $1)
        )
       )
      )
     )
     (i32.const 24)
    )
   )
  )
  (loop $label$4
   (set_local $5
    (i32.add
     (get_local $5)
     (i32.const 1)
    )
   )
   (br_if $label$4
    (i64.ne
     (tee_local $6
      (i64.shr_u
       (get_local $6)
       (i64.const 7)
      )
     )
     (i64.const 0)
    )
   )
  )
  (block $label$5
   (br_if $label$5
    (i32.eq
     (get_local $2)
     (get_local $4)
    )
   )
   (set_local $5
    (i32.add
     (i32.add
      (i32.mul
       (i32.div_u
        (i32.add
         (get_local $3)
         (i32.const -24)
        )
        (i32.const 24)
       )
       (i32.const 18)
      )
      (get_local $5)
     )
     (i32.const 18)
    )
   )
  )
  (i32.store
   (get_local $0)
   (tee_local $5
    (i32.add
     (get_local $5)
     (i32.const 4)
    )
   )
  )
  (set_local $6
   (i64.extend_u/i32
    (i32.div_s
     (i32.sub
      (tee_local $2
       (i32.load
        (i32.add
         (get_local $1)
         (i32.const 36)
        )
       )
      )
      (tee_local $4
       (i32.load
        (i32.add
         (get_local $1)
         (i32.const 32)
        )
       )
      )
     )
     (i32.const 44)
    )
   )
  )
  (loop $label$6
   (set_local $5
    (i32.add
     (get_local $5)
     (i32.const 1)
    )
   )
   (br_if $label$6
    (i64.ne
     (tee_local $6
      (i64.shr_u
       (get_local $6)
       (i64.const 7)
      )
     )
     (i64.const 0)
    )
   )
  )
  (i32.store
   (get_local $0)
   (get_local $5)
  )
  (block $label$7
   (br_if $label$7
    (i32.eq
     (get_local $4)
     (get_local $2)
    )
   )
   (loop $label$8
    (set_local $5
     (i32.add
      (get_local $5)
      (i32.const 35)
     )
    )
    (set_local $6
     (i64.load32_u
      (get_local $4)
     )
    )
    (loop $label$9
     (set_local $5
      (i32.add
       (get_local $5)
       (i32.const 1)
      )
     )
     (br_if $label$9
      (i64.ne
       (tee_local $6
        (i64.shr_u
         (get_local $6)
         (i64.const 7)
        )
       )
       (i64.const 0)
      )
     )
    )
    (br_if $label$8
     (i32.ne
      (tee_local $4
       (i32.add
        (get_local $4)
        (i32.const 44)
       )
      )
      (get_local $2)
     )
    )
   )
   (i32.store
    (get_local $0)
    (get_local $5)
   )
  )
  (set_local $6
   (i64.extend_u/i32
    (i32.div_s
     (tee_local $3
      (i32.sub
       (tee_local $4
        (i32.load
         (i32.add
          (get_local $1)
          (i32.const 48)
         )
        )
       )
       (tee_local $2
        (i32.load
         (i32.add
          (get_local $1)
          (i32.const 44)
         )
        )
       )
      )
     )
     (i32.const 24)
    )
   )
  )
  (loop $label$10
   (set_local $5
    (i32.add
     (get_local $5)
     (i32.const 1)
    )
   )
   (br_if $label$10
    (i64.ne
     (tee_local $6
      (i64.shr_u
       (get_local $6)
       (i64.const 7)
      )
     )
     (i64.const 0)
    )
   )
  )
  (block $label$11
   (br_if $label$11
    (i32.eq
     (get_local $2)
     (get_local $4)
    )
   )
   (set_local $5
    (i32.add
     (i32.add
      (i32.mul
       (i32.div_u
        (i32.add
         (get_local $3)
         (i32.const -24)
        )
        (i32.const 24)
       )
       (i32.const 18)
      )
      (get_local $5)
     )
     (i32.const 18)
    )
   )
  )
  (i32.store
   (get_local $0)
   (tee_local $5
    (i32.add
     (get_local $5)
     (i32.const 4)
    )
   )
  )
  (set_local $6
   (i64.extend_u/i32
    (i32.div_s
     (i32.sub
      (tee_local $2
       (i32.load
        (i32.add
         (get_local $1)
         (i32.const 64)
        )
       )
      )
      (tee_local $4
       (i32.load
        (i32.add
         (get_local $1)
         (i32.const 60)
        )
       )
      )
     )
     (i32.const 44)
    )
   )
  )
  (loop $label$12
   (set_local $5
    (i32.add
     (get_local $5)
     (i32.const 1)
    )
   )
   (br_if $label$12
    (i64.ne
     (tee_local $6
      (i64.shr_u
       (get_local $6)
       (i64.const 7)
      )
     )
     (i64.const 0)
    )
   )
  )
  (i32.store
   (get_local $0)
   (get_local $5)
  )
  (block $label$13
   (br_if $label$13
    (i32.eq
     (get_local $4)
     (get_local $2)
    )
   )
   (loop $label$14
    (set_local $5
     (i32.add
      (get_local $5)
      (i32.const 35)
     )
    )
    (set_local $6
     (i64.load32_u
      (get_local $4)
     )
    )
    (loop $label$15
     (set_local $5
      (i32.add
       (get_local $5)
       (i32.const 1)
      )
     )
     (br_if $label$15
      (i64.ne
       (tee_local $6
        (i64.shr_u
         (get_local $6)
         (i64.const 7)
        )
       )
       (i64.const 0)
      )
     )
    )
    (br_if $label$14
     (i32.ne
      (tee_local $4
       (i32.add
        (get_local $4)
        (i32.const 44)
       )
      )
      (get_local $2)
     )
    )
   )
   (i32.store
    (get_local $0)
    (get_local $5)
   )
  )
  (set_local $6
   (i64.extend_u/i32
    (i32.div_s
     (tee_local $1
      (i32.sub
       (tee_local $4
        (i32.load
         (i32.add
          (get_local $1)
          (i32.const 76)
         )
        )
       )
       (tee_local $2
        (i32.load
         (i32.add
          (get_local $1)
          (i32.const 72)
         )
        )
       )
      )
     )
     (i32.const 24)
    )
   )
  )
  (loop $label$16
   (set_local $5
    (i32.add
     (get_local $5)
     (i32.const 1)
    )
   )
   (br_if $label$16
    (i64.ne
     (tee_local $6
      (i64.shr_u
       (get_local $6)
       (i64.const 7)
      )
     )
     (i64.const 0)
    )
   )
  )
  (block $label$17
   (br_if $label$17
    (i32.eq
     (get_local $2)
     (get_local $4)
    )
   )
   (set_local $5
    (i32.add
     (i32.add
      (i32.mul
       (i32.div_u
        (i32.add
         (get_local $1)
         (i32.const -24)
        )
        (i32.const 24)
       )
       (i32.const 18)
      )
      (get_local $5)
     )
     (i32.const 18)
    )
   )
  )
  (i32.store
   (get_local $0)
   (get_local $5)
  )
  (get_local $0)
 )
 (func $_ZN5boost3pfr6detail19for_each_field_implINS1_14sequence_tuple5tupleIJRKNSt3__112basic_stringIcNS5_11char_traitsIcEENS5_9allocatorIcEEEESD_SD_EEEZN5eosiolsINSF_10datastreamIjEENSF_11content_refELi0EEERT_SL_RKT0_EUlRKSK_E_JLj0ELj1ELj2EEEEvSL_OSM_NS5_16integer_sequenceIjJXspT1_EEEENS5_17integral_constantIbLb0EEE (param $0 i32) (param $1 i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i64)
  (set_local $5
   (i64.extend_u/i32
    (select
     (i32.load offset=4
      (tee_local $2
       (i32.load
        (get_local $0)
       )
      )
     )
     (i32.shr_u
      (tee_local $4
       (i32.load8_u
        (get_local $2)
       )
      )
      (i32.const 1)
     )
     (i32.and
      (get_local $4)
      (i32.const 1)
     )
    )
   )
  )
  (set_local $4
   (i32.load
    (tee_local $3
     (i32.load
      (get_local $1)
     )
    )
   )
  )
  (loop $label$0
   (set_local $4
    (i32.add
     (get_local $4)
     (i32.const 1)
    )
   )
   (br_if $label$0
    (i64.ne
     (tee_local $5
      (i64.shr_u
       (get_local $5)
       (i64.const 7)
      )
     )
     (i64.const 0)
    )
   )
  )
  (i32.store
   (get_local $3)
   (get_local $4)
  )
  (block $label$1
   (br_if $label$1
    (i32.eqz
     (tee_local $2
      (select
       (i32.load
        (i32.add
         (get_local $2)
         (i32.const 4)
        )
       )
       (i32.shr_u
        (tee_local $2
         (i32.load8_u
          (get_local $2)
         )
        )
        (i32.const 1)
       )
       (i32.and
        (get_local $2)
        (i32.const 1)
       )
      )
     )
    )
   )
   (i32.store
    (get_local $3)
    (i32.add
     (get_local $2)
     (get_local $4)
    )
   )
  )
  (set_local $5
   (i64.extend_u/i32
    (select
     (i32.load offset=4
      (tee_local $2
       (i32.load offset=4
        (get_local $0)
       )
      )
     )
     (i32.shr_u
      (tee_local $4
       (i32.load8_u
        (get_local $2)
       )
      )
      (i32.const 1)
     )
     (i32.and
      (get_local $4)
      (i32.const 1)
     )
    )
   )
  )
  (set_local $4
   (i32.load
    (tee_local $3
     (i32.load
      (get_local $1)
     )
    )
   )
  )
  (loop $label$2
   (set_local $4
    (i32.add
     (get_local $4)
     (i32.const 1)
    )
   )
   (br_if $label$2
    (i64.ne
     (tee_local $5
      (i64.shr_u
       (get_local $5)
       (i64.const 7)
      )
     )
     (i64.const 0)
    )
   )
  )
  (i32.store
   (get_local $3)
   (get_local $4)
  )
  (block $label$3
   (br_if $label$3
    (i32.eqz
     (tee_local $2
      (select
       (i32.load
        (i32.add
         (get_local $2)
         (i32.const 4)
        )
       )
       (i32.shr_u
        (tee_local $2
         (i32.load8_u
          (get_local $2)
         )
        )
        (i32.const 1)
       )
       (i32.and
        (get_local $2)
        (i32.const 1)
       )
      )
     )
    )
   )
   (i32.store
    (get_local $3)
    (i32.add
     (get_local $2)
     (get_local $4)
    )
   )
  )
  (set_local $5
   (i64.extend_u/i32
    (select
     (i32.load offset=4
      (tee_local $0
       (i32.load offset=8
        (get_local $0)
       )
      )
     )
     (i32.shr_u
      (tee_local $4
       (i32.load8_u
        (get_local $0)
       )
      )
      (i32.const 1)
     )
     (i32.and
      (get_local $4)
      (i32.const 1)
     )
    )
   )
  )
  (set_local $4
   (i32.load
    (tee_local $1
     (i32.load
      (get_local $1)
     )
    )
   )
  )
  (loop $label$4
   (set_local $4
    (i32.add
     (get_local $4)
     (i32.const 1)
    )
   )
   (br_if $label$4
    (i64.ne
     (tee_local $5
      (i64.shr_u
       (get_local $5)
       (i64.const 7)
      )
     )
     (i64.const 0)
    )
   )
  )
  (i32.store
   (get_local $1)
   (get_local $4)
  )
  (block $label$5
   (br_if $label$5
    (i32.eqz
     (tee_local $0
      (select
       (i32.load
        (i32.add
         (get_local $0)
         (i32.const 4)
        )
       )
       (i32.shr_u
        (tee_local $0
         (i32.load8_u
          (get_local $0)
         )
        )
        (i32.const 1)
       )
       (i32.and
        (get_local $0)
        (i32.const 1)
       )
      )
     )
    )
   )
   (i32.store
    (get_local $1)
    (i32.add
     (get_local $0)
     (get_local $4)
    )
   )
  )
 )
 (func $_ZNSt3__16vectorIN11eosiosystem10key_weightENS_9allocatorIS2_EEE6assignIPS2_EENS_9enable_ifIXaasr21__is_forward_iteratorIT_EE5valuesr16is_constructibleIS2_NS_15iterator_traitsIS9_E9referenceEEE5valueEvE4typeES9_S9_ (param $0 i32) (param $1 i32) (param $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i32)
  (local $6 i32)
  (local $7 i32)
  (local $8 i32)
  (block $label$0
   (block $label$1
    (block $label$2
     (block $label$3
      (br_if $label$3
       (i32.le_u
        (tee_local $4
         (i32.div_s
          (tee_local $3
           (i32.sub
            (get_local $2)
            (get_local $1)
           )
          )
          (i32.const 44)
         )
        )
        (i32.div_s
         (i32.sub
          (tee_local $8
           (i32.load offset=8
            (get_local $0)
           )
          )
          (tee_local $5
           (i32.load
            (get_local $0)
           )
          )
         )
         (i32.const 44)
        )
       )
      )
      (block $label$4
       (br_if $label$4
        (i32.eqz
         (get_local $5)
        )
       )
       (i32.store offset=4
        (get_local $0)
        (get_local $5)
       )
       (call $_ZdlPv
        (get_local $5)
       )
       (set_local $8
        (i32.const 0)
       )
       (i32.store
        (i32.add
         (get_local $0)
         (i32.const 8)
        )
        (i32.const 0)
       )
       (i64.store align=4
        (get_local $0)
        (i64.const 0)
       )
      )
      (br_if $label$0
       (i32.ge_u
        (get_local $4)
        (i32.const 97612894)
       )
      )
      (set_local $5
       (i32.const 97612893)
      )
      (block $label$5
       (br_if $label$5
        (i32.gt_u
         (tee_local $2
          (i32.div_s
           (get_local $8)
           (i32.const 44)
          )
         )
         (i32.const 48806445)
        )
       )
       (set_local $5
        (select
         (get_local $4)
         (tee_local $5
          (i32.shl
           (get_local $2)
           (i32.const 1)
          )
         )
         (i32.lt_u
          (get_local $5)
          (get_local $4)
         )
        )
       )
      )
      (i32.store
       (get_local $0)
       (tee_local $5
        (call $_Znwj
         (tee_local $4
          (i32.mul
           (get_local $5)
           (i32.const 44)
          )
         )
        )
       )
      )
      (i32.store offset=4
       (get_local $0)
       (get_local $5)
      )
      (i32.store
       (i32.add
        (get_local $0)
        (i32.const 8)
       )
       (i32.add
        (get_local $5)
        (get_local $4)
       )
      )
      (br_if $label$2
       (i32.lt_s
        (get_local $3)
        (i32.const 1)
       )
      )
      (drop
       (call $memcpy
        (get_local $5)
        (get_local $1)
        (get_local $3)
       )
      )
      (i32.store
       (tee_local $0
        (i32.add
         (get_local $0)
         (i32.const 4)
        )
       )
       (i32.add
        (i32.load
         (get_local $0)
        )
        (i32.mul
         (i32.div_u
          (get_local $3)
          (i32.const 44)
         )
         (i32.const 44)
        )
       )
      )
      (return)
     )
     (set_local $7
      (i32.div_s
       (tee_local $8
        (i32.sub
         (tee_local $6
          (select
           (i32.add
            (get_local $1)
            (i32.mul
             (tee_local $3
              (i32.div_s
               (i32.sub
                (i32.load offset=4
                 (get_local $0)
                )
                (get_local $5)
               )
               (i32.const 44)
              )
             )
             (i32.const 44)
            )
           )
           (get_local $2)
           (i32.gt_u
            (get_local $4)
            (get_local $3)
           )
          )
         )
         (get_local $1)
        )
       )
       (i32.const 44)
      )
     )
     (block $label$6
      (br_if $label$6
       (i32.eqz
        (get_local $8)
       )
      )
      (drop
       (call $memmove
        (get_local $5)
        (get_local $1)
        (get_local $8)
       )
      )
     )
     (br_if $label$1
      (i32.le_u
       (get_local $4)
       (get_local $3)
      )
     )
     (br_if $label$2
      (i32.lt_s
       (tee_local $1
        (i32.sub
         (get_local $2)
         (get_local $6)
        )
       )
       (i32.const 1)
      )
     )
     (drop
      (call $memcpy
       (i32.load
        (tee_local $0
         (i32.add
          (get_local $0)
          (i32.const 4)
         )
        )
       )
       (get_local $6)
       (get_local $1)
      )
     )
     (i32.store
      (get_local $0)
      (i32.add
       (i32.load
        (get_local $0)
       )
       (i32.mul
        (i32.div_u
         (get_local $1)
         (i32.const 44)
        )
        (i32.const 44)
       )
      )
     )
     (return)
    )
    (return)
   )
   (i32.store
    (i32.add
     (get_local $0)
     (i32.const 4)
    )
    (i32.add
     (get_local $5)
     (i32.mul
      (get_local $7)
      (i32.const 44)
     )
    )
   )
   (return)
  )
  (call $_ZNKSt3__120__vector_base_commonILb1EE20__throw_length_errorEv
   (get_local $0)
  )
  (unreachable)
 )
 (func $_ZNSt3__16vectorIN11eosiosystem23permission_level_weightENS_9allocatorIS2_EEE6assignIPS2_EENS_9enable_ifIXaasr21__is_forward_iteratorIT_EE5valuesr16is_constructibleIS2_NS_15iterator_traitsIS9_E9referenceEEE5valueEvE4typeES9_S9_ (param $0 i32) (param $1 i32) (param $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i32)
  (local $6 i32)
  (local $7 i32)
  (local $8 i32)
  (block $label$0
   (block $label$1
    (block $label$2
     (block $label$3
      (br_if $label$3
       (i32.le_u
        (tee_local $4
         (i32.div_s
          (tee_local $3
           (i32.sub
            (get_local $2)
            (get_local $1)
           )
          )
          (i32.const 24)
         )
        )
        (i32.div_s
         (i32.sub
          (tee_local $8
           (i32.load offset=8
            (get_local $0)
           )
          )
          (tee_local $5
           (i32.load
            (get_local $0)
           )
          )
         )
         (i32.const 24)
        )
       )
      )
      (block $label$4
       (br_if $label$4
        (i32.eqz
         (get_local $5)
        )
       )
       (i32.store offset=4
        (get_local $0)
        (get_local $5)
       )
       (call $_ZdlPv
        (get_local $5)
       )
       (set_local $8
        (i32.const 0)
       )
       (i32.store
        (i32.add
         (get_local $0)
         (i32.const 8)
        )
        (i32.const 0)
       )
       (i64.store align=4
        (get_local $0)
        (i64.const 0)
       )
      )
      (br_if $label$0
       (i32.ge_u
        (get_local $4)
        (i32.const 178956971)
       )
      )
      (set_local $5
       (i32.const 178956970)
      )
      (block $label$5
       (br_if $label$5
        (i32.gt_u
         (tee_local $2
          (i32.div_s
           (get_local $8)
           (i32.const 24)
          )
         )
         (i32.const 89478484)
        )
       )
       (set_local $5
        (select
         (get_local $4)
         (tee_local $5
          (i32.shl
           (get_local $2)
           (i32.const 1)
          )
         )
         (i32.lt_u
          (get_local $5)
          (get_local $4)
         )
        )
       )
      )
      (i32.store
       (get_local $0)
       (tee_local $5
        (call $_Znwj
         (tee_local $4
          (i32.mul
           (get_local $5)
           (i32.const 24)
          )
         )
        )
       )
      )
      (i32.store offset=4
       (get_local $0)
       (get_local $5)
      )
      (i32.store
       (i32.add
        (get_local $0)
        (i32.const 8)
       )
       (i32.add
        (get_local $5)
        (get_local $4)
       )
      )
      (br_if $label$2
       (i32.lt_s
        (get_local $3)
        (i32.const 1)
       )
      )
      (drop
       (call $memcpy
        (get_local $5)
        (get_local $1)
        (get_local $3)
       )
      )
      (i32.store
       (tee_local $0
        (i32.add
         (get_local $0)
         (i32.const 4)
        )
       )
       (i32.add
        (i32.load
         (get_local $0)
        )
        (i32.mul
         (i32.div_u
          (get_local $3)
          (i32.const 24)
         )
         (i32.const 24)
        )
       )
      )
      (return)
     )
     (set_local $7
      (i32.div_s
       (tee_local $8
        (i32.sub
         (tee_local $6
          (select
           (i32.add
            (get_local $1)
            (i32.mul
             (tee_local $3
              (i32.div_s
               (i32.sub
                (i32.load offset=4
                 (get_local $0)
                )
                (get_local $5)
               )
               (i32.const 24)
              )
             )
             (i32.const 24)
            )
           )
           (get_local $2)
           (i32.gt_u
            (get_local $4)
            (get_local $3)
           )
          )
         )
         (get_local $1)
        )
       )
       (i32.const 24)
      )
     )
     (block $label$6
      (br_if $label$6
       (i32.eqz
        (get_local $8)
       )
      )
      (drop
       (call $memmove
        (get_local $5)
        (get_local $1)
        (get_local $8)
       )
      )
     )
     (br_if $label$1
      (i32.le_u
       (get_local $4)
       (get_local $3)
      )
     )
     (br_if $label$2
      (i32.lt_s
       (tee_local $1
        (i32.sub
         (get_local $2)
         (get_local $6)
        )
       )
       (i32.const 1)
      )
     )
     (drop
      (call $memcpy
       (i32.load
        (tee_local $0
         (i32.add
          (get_local $0)
          (i32.const 4)
         )
        )
       )
       (get_local $6)
       (get_local $1)
      )
     )
     (i32.store
      (get_local $0)
      (i32.add
       (i32.load
        (get_local $0)
       )
       (i32.mul
        (i32.div_u
         (get_local $1)
         (i32.const 24)
        )
        (i32.const 24)
       )
      )
     )
     (return)
    )
    (return)
   )
   (i32.store
    (i32.add
     (get_local $0)
     (i32.const 4)
    )
    (i32.add
     (get_local $5)
     (i32.mul
      (get_local $7)
      (i32.const 24)
     )
    )
   )
   (return)
  )
  (call $_ZNKSt3__120__vector_base_commonILb1EE20__throw_length_errorEv
   (get_local $0)
  )
  (unreachable)
 )
 (func $_ZN5eosiorsINS_10datastreamIPKcEEEERT_S6_RNS_14discovery_fileE (param $0 i32) (param $1 i32) (result i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i32)
  (local $6 i64)
  (local $7 i32)
  (call $eosio_assert
   (i32.gt_u
    (i32.sub
     (i32.load offset=8
      (get_local $0)
     )
     (i32.load offset=4
      (get_local $0)
     )
    )
    (i32.const 31)
   )
   (i32.const 416)
  )
  (drop
   (call $memcpy
    (get_local $1)
    (i32.load offset=4
     (get_local $0)
    )
    (i32.const 32)
   )
  )
  (i32.store offset=4
   (get_local $0)
   (tee_local $7
    (i32.add
     (i32.load offset=4
      (get_local $0)
     )
     (i32.const 32)
    )
   )
  )
  (call $eosio_assert
   (i32.gt_u
    (i32.sub
     (i32.load offset=8
      (get_local $0)
     )
     (get_local $7)
    )
    (i32.const 7)
   )
   (i32.const 416)
  )
  (drop
   (call $memcpy
    (i32.add
     (get_local $1)
     (i32.const 32)
    )
    (i32.load offset=4
     (get_local $0)
    )
    (i32.const 8)
   )
  )
  (i32.store offset=4
   (get_local $0)
   (i32.add
    (i32.load offset=4
     (get_local $0)
    )
    (i32.const 8)
   )
  )
  (call $eosio_assert
   (i32.gt_u
    (i32.sub
     (i32.load offset=8
      (tee_local $0
       (call $_ZN5eosiorsINS_10datastreamIPKcEENS_9peer_linkEEERT_S7_RNSt3__16vectorIT0_NS8_9allocatorISA_EEEE
        (get_local $0)
        (i32.add
         (get_local $1)
         (i32.const 40)
        )
       )
      )
     )
     (i32.load offset=4
      (get_local $0)
     )
    )
    (i32.const 7)
   )
   (i32.const 416)
  )
  (drop
   (call $memcpy
    (i32.add
     (get_local $1)
     (i32.const 56)
    )
    (i32.load offset=4
     (get_local $0)
    )
    (i32.const 8)
   )
  )
  (i32.store offset=4
   (get_local $0)
   (i32.add
    (i32.load offset=4
     (get_local $0)
    )
    (i32.const 8)
   )
  )
  (call $eosio_assert
   (i32.gt_u
    (i32.sub
     (i32.load offset=8
      (tee_local $0
       (call $_ZN5eosiorsINS_10datastreamIPKcEENSt3__112basic_stringIcNS5_11char_traitsIcEENS5_9allocatorIcEEEEEERT_SD_RNS5_6vectorIT0_NS9_ISF_EEEE
        (get_local $0)
        (i32.add
         (get_local $1)
         (i32.const 64)
        )
       )
      )
     )
     (i32.load offset=4
      (get_local $0)
     )
    )
    (i32.const 1)
   )
   (i32.const 416)
  )
  (drop
   (call $memcpy
    (i32.add
     (get_local $1)
     (i32.const 76)
    )
    (i32.load offset=4
     (get_local $0)
    )
    (i32.const 2)
   )
  )
  (i32.store offset=4
   (get_local $0)
   (tee_local $7
    (i32.add
     (i32.load offset=4
      (get_local $0)
     )
     (i32.const 2)
    )
   )
  )
  (call $eosio_assert
   (i32.ne
    (i32.load offset=8
     (get_local $0)
    )
    (get_local $7)
   )
   (i32.const 416)
  )
  (drop
   (call $memcpy
    (i32.add
     (get_local $1)
     (i32.const 78)
    )
    (i32.load offset=4
     (get_local $0)
    )
    (i32.const 1)
   )
  )
  (i32.store offset=4
   (get_local $0)
   (tee_local $7
    (i32.add
     (i32.load offset=4
      (get_local $0)
     )
     (i32.const 1)
    )
   )
  )
  (call $eosio_assert
   (i32.gt_u
    (i32.sub
     (i32.load offset=8
      (get_local $0)
     )
     (get_local $7)
    )
    (i32.const 31)
   )
   (i32.const 416)
  )
  (drop
   (call $memcpy
    (i32.add
     (get_local $1)
     (i32.const 79)
    )
    (i32.load offset=4
     (get_local $0)
    )
    (i32.const 32)
   )
  )
  (i32.store offset=4
   (get_local $0)
   (i32.add
    (i32.load offset=4
     (get_local $0)
    )
    (i32.const 32)
   )
  )
  (call $eosio_assert
   (i32.gt_u
    (i32.sub
     (i32.load offset=8
      (tee_local $2
       (call $_ZN5eosiorsINS_10datastreamIPKcEEEERT_S6_RNSt3__112basic_stringIcNS7_11char_traitsIcEENS7_9allocatorIcEEEE
        (get_local $0)
        (i32.add
         (get_local $1)
         (i32.const 112)
        )
       )
      )
     )
     (i32.load offset=4
      (get_local $2)
     )
    )
    (i32.const 7)
   )
   (i32.const 416)
  )
  (drop
   (call $memcpy
    (i32.add
     (get_local $1)
     (i32.const 128)
    )
    (i32.load offset=4
     (get_local $2)
    )
    (i32.const 8)
   )
  )
  (i32.store offset=4
   (get_local $2)
   (tee_local $0
    (i32.add
     (i32.load offset=4
      (get_local $2)
     )
     (i32.const 8)
    )
   )
  )
  (set_local $5
   (i32.const 0)
  )
  (set_local $6
   (i64.const 0)
  )
  (set_local $3
   (i32.add
    (get_local $2)
    (i32.const 4)
   )
  )
  (loop $label$0
   (call $eosio_assert
    (i32.lt_u
     (get_local $0)
     (i32.load
      (i32.add
       (get_local $2)
       (i32.const 8)
      )
     )
    )
    (i32.const 432)
   )
   (set_local $7
    (i32.load8_u
     (tee_local $0
      (i32.load
       (get_local $3)
      )
     )
    )
   )
   (i32.store
    (get_local $3)
    (tee_local $0
     (i32.add
      (get_local $0)
      (i32.const 1)
     )
    )
   )
   (set_local $6
    (i64.or
     (i64.extend_u/i32
      (i32.shl
       (i32.and
        (get_local $7)
        (i32.const 127)
       )
       (tee_local $5
        (i32.and
         (get_local $5)
         (i32.const 255)
        )
       )
      )
     )
     (get_local $6)
    )
   )
   (set_local $5
    (i32.add
     (get_local $5)
     (i32.const 7)
    )
   )
   (br_if $label$0
    (i32.shr_u
     (get_local $7)
     (i32.const 7)
    )
   )
  )
  (i64.store32
   (i32.add
    (get_local $1)
    (i32.const 136)
   )
   (get_local $6)
  )
  (set_local $3
   (i32.add
    (get_local $1)
    (i32.const 140)
   )
  )
  (set_local $7
   (i32.const 0)
  )
  (set_local $4
   (i32.add
    (get_local $2)
    (i32.const 8)
   )
  )
  (set_local $5
   (i32.add
    (get_local $2)
    (i32.const 4)
   )
  )
  (loop $label$1
   (call $eosio_assert
    (i32.ne
     (i32.load
      (get_local $4)
     )
     (get_local $0)
    )
    (i32.const 416)
   )
   (drop
    (call $memcpy
     (i32.add
      (get_local $3)
      (get_local $7)
     )
     (i32.load
      (get_local $5)
     )
     (i32.const 1)
    )
   )
   (i32.store
    (get_local $5)
    (tee_local $0
     (i32.add
      (i32.load
       (get_local $5)
      )
      (i32.const 1)
     )
    )
   )
   (br_if $label$1
    (i32.ne
     (tee_local $7
      (i32.add
       (get_local $7)
       (i32.const 1)
      )
     )
     (i32.const 33)
    )
   )
  )
  (call $_ZN5eosiorsINS_10datastreamIPKcEENS_11content_refEEERT_S7_RNSt3__16vectorIT0_NS8_9allocatorISA_EEEE
   (call $_ZN5eosiorsINS_10datastreamIPKcEEEERT_S6_RNS_13authority_mapE
    (get_local $2)
    (i32.add
     (get_local $1)
     (i32.const 176)
    )
   )
   (i32.add
    (get_local $1)
    (i32.const 260)
   )
  )
 )
 (func $_ZN5eosiorsINS_10datastreamIPKcEENS_9peer_linkEEERT_S7_RNSt3__16vectorIT0_NS8_9allocatorISA_EEEE (param $0 i32) (param $1 i32) (result i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i64)
  (local $6 i32)
  (local $7 i32)
  (set_local $7
   (i32.load offset=4
    (get_local $0)
   )
  )
  (set_local $6
   (i32.const 0)
  )
  (set_local $5
   (i64.const 0)
  )
  (set_local $2
   (i32.add
    (get_local $0)
    (i32.const 8)
   )
  )
  (set_local $3
   (i32.add
    (get_local $0)
    (i32.const 4)
   )
  )
  (loop $label$0
   (call $eosio_assert
    (i32.lt_u
     (get_local $7)
     (i32.load
      (get_local $2)
     )
    )
    (i32.const 432)
   )
   (set_local $4
    (i32.load8_u
     (tee_local $7
      (i32.load
       (get_local $3)
      )
     )
    )
   )
   (i32.store
    (get_local $3)
    (tee_local $7
     (i32.add
      (get_local $7)
      (i32.const 1)
     )
    )
   )
   (set_local $5
    (i64.or
     (i64.extend_u/i32
      (i32.shl
       (i32.and
        (get_local $4)
        (i32.const 127)
       )
       (tee_local $6
        (i32.and
         (get_local $6)
         (i32.const 255)
        )
       )
      )
     )
     (get_local $5)
    )
   )
   (set_local $6
    (i32.add
     (get_local $6)
     (i32.const 7)
    )
   )
   (br_if $label$0
    (i32.shr_u
     (get_local $4)
     (i32.const 7)
    )
   )
  )
  (block $label$1
   (block $label$2
    (br_if $label$2
     (i32.le_u
      (tee_local $4
       (i32.wrap/i64
        (get_local $5)
       )
      )
      (tee_local $7
       (i32.shr_s
        (i32.sub
         (tee_local $2
          (i32.load offset=4
           (get_local $1)
          )
         )
         (tee_local $6
          (i32.load
           (get_local $1)
          )
         )
        )
        (i32.const 5)
       )
      )
     )
    )
    (call $_ZNSt3__16vectorIN5eosio9peer_linkENS_9allocatorIS2_EEE8__appendEj
     (get_local $1)
     (i32.sub
      (get_local $4)
      (get_local $7)
     )
    )
    (set_local $2
     (i32.load
      (i32.add
       (get_local $1)
       (i32.const 4)
      )
     )
    )
    (br $label$1)
   )
   (br_if $label$1
    (i32.ge_u
     (get_local $4)
     (get_local $7)
    )
   )
   (block $label$3
    (br_if $label$3
     (i32.eq
      (get_local $2)
      (tee_local $3
       (i32.add
        (get_local $6)
        (tee_local $4
         (i32.shl
          (get_local $4)
          (i32.const 5)
         )
        )
       )
      )
     )
    )
    (set_local $7
     (i32.sub
      (i32.sub
       (i32.const 0)
       (get_local $6)
      )
      (get_local $4)
     )
    )
    (set_local $4
     (i32.add
      (get_local $2)
      (i32.const -24)
     )
    )
    (loop $label$4
     (block $label$5
      (br_if $label$5
       (i32.eqz
        (i32.and
         (i32.load8_u
          (get_local $4)
         )
         (i32.const 1)
        )
       )
      )
      (call $_ZdlPv
       (i32.load
        (i32.add
         (get_local $4)
         (i32.const 8)
        )
       )
      )
     )
     (br_if $label$4
      (i32.ne
       (i32.add
        (tee_local $4
         (i32.add
          (get_local $4)
          (i32.const -32)
         )
        )
        (get_local $7)
       )
       (i32.const -24)
      )
     )
    )
   )
   (i32.store
    (i32.add
     (get_local $1)
     (i32.const 4)
    )
    (get_local $3)
   )
   (set_local $2
    (get_local $3)
   )
  )
  (block $label$6
   (br_if $label$6
    (i32.eq
     (tee_local $7
      (i32.load
       (get_local $1)
      )
     )
     (get_local $2)
    )
   )
   (set_local $6
    (i32.load
     (tee_local $4
      (i32.add
       (get_local $0)
       (i32.const 4)
      )
     )
    )
   )
   (loop $label$7
    (call $eosio_assert
     (i32.gt_u
      (i32.sub
       (i32.load
        (tee_local $3
         (i32.add
          (get_local $0)
          (i32.const 8)
         )
        )
       )
       (get_local $6)
      )
      (i32.const 7)
     )
     (i32.const 416)
    )
    (drop
     (call $memcpy
      (get_local $7)
      (i32.load
       (get_local $4)
      )
      (i32.const 8)
     )
    )
    (i32.store
     (get_local $4)
     (i32.add
      (i32.load
       (get_local $4)
      )
      (i32.const 8)
     )
    )
    (drop
     (call $_ZN5eosiorsINS_10datastreamIPKcEEEERT_S6_RNSt3__112basic_stringIcNS7_11char_traitsIcEENS7_9allocatorIcEEEE
      (get_local $0)
      (i32.add
       (get_local $7)
       (i32.const 8)
      )
     )
    )
    (call $eosio_assert
     (i32.gt_u
      (i32.sub
       (i32.load
        (get_local $3)
       )
       (i32.load
        (get_local $4)
       )
      )
      (i32.const 7)
     )
     (i32.const 416)
    )
    (drop
     (call $memcpy
      (i32.add
       (get_local $7)
       (i32.const 24)
      )
      (i32.load
       (get_local $4)
      )
      (i32.const 8)
     )
    )
    (i32.store
     (get_local $4)
     (tee_local $6
      (i32.add
       (i32.load
        (get_local $4)
       )
       (i32.const 8)
      )
     )
    )
    (br_if $label$7
     (i32.ne
      (tee_local $7
       (i32.add
        (get_local $7)
        (i32.const 32)
       )
      )
      (get_local $2)
     )
    )
   )
  )
  (get_local $0)
 )
 (func $_ZN5eosiorsINS_10datastreamIPKcEENSt3__112basic_stringIcNS5_11char_traitsIcEENS5_9allocatorIcEEEEEERT_SD_RNS5_6vectorIT0_NS9_ISF_EEEE (param $0 i32) (param $1 i32) (result i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i32)
  (local $6 i64)
  (local $7 i32)
  (set_local $5
   (i32.load offset=4
    (get_local $0)
   )
  )
  (set_local $7
   (i32.const 0)
  )
  (set_local $6
   (i64.const 0)
  )
  (set_local $2
   (i32.add
    (get_local $0)
    (i32.const 8)
   )
  )
  (set_local $3
   (i32.add
    (get_local $0)
    (i32.const 4)
   )
  )
  (loop $label$0
   (call $eosio_assert
    (i32.lt_u
     (get_local $5)
     (i32.load
      (get_local $2)
     )
    )
    (i32.const 432)
   )
   (set_local $4
    (i32.load8_u
     (tee_local $5
      (i32.load
       (get_local $3)
      )
     )
    )
   )
   (i32.store
    (get_local $3)
    (tee_local $5
     (i32.add
      (get_local $5)
      (i32.const 1)
     )
    )
   )
   (set_local $6
    (i64.or
     (i64.extend_u/i32
      (i32.shl
       (i32.and
        (get_local $4)
        (i32.const 127)
       )
       (tee_local $7
        (i32.and
         (get_local $7)
         (i32.const 255)
        )
       )
      )
     )
     (get_local $6)
    )
   )
   (set_local $7
    (i32.add
     (get_local $7)
     (i32.const 7)
    )
   )
   (br_if $label$0
    (i32.shr_u
     (get_local $4)
     (i32.const 7)
    )
   )
  )
  (block $label$1
   (block $label$2
    (br_if $label$2
     (i32.le_u
      (tee_local $4
       (i32.wrap/i64
        (get_local $6)
       )
      )
      (tee_local $7
       (i32.div_s
        (i32.sub
         (tee_local $5
          (i32.load offset=4
           (get_local $1)
          )
         )
         (tee_local $3
          (i32.load
           (get_local $1)
          )
         )
        )
        (i32.const 12)
       )
      )
     )
    )
    (call $_ZNSt3__16vectorINS_12basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEENS4_IS6_EEE8__appendEj
     (get_local $1)
     (i32.sub
      (get_local $4)
      (get_local $7)
     )
    )
    (set_local $5
     (i32.load
      (i32.add
       (get_local $1)
       (i32.const 4)
      )
     )
    )
    (br $label$1)
   )
   (br_if $label$1
    (i32.ge_u
     (get_local $4)
     (get_local $7)
    )
   )
   (block $label$3
    (br_if $label$3
     (i32.eq
      (get_local $5)
      (tee_local $2
       (i32.add
        (get_local $3)
        (tee_local $4
         (i32.mul
          (get_local $4)
          (i32.const 12)
         )
        )
       )
      )
     )
    )
    (set_local $7
     (i32.sub
      (i32.sub
       (i32.const 0)
       (get_local $3)
      )
      (get_local $4)
     )
    )
    (set_local $4
     (i32.add
      (get_local $5)
      (i32.const -12)
     )
    )
    (loop $label$4
     (block $label$5
      (br_if $label$5
       (i32.eqz
        (i32.and
         (i32.load8_u
          (get_local $4)
         )
         (i32.const 1)
        )
       )
      )
      (call $_ZdlPv
       (i32.load
        (i32.add
         (get_local $4)
         (i32.const 8)
        )
       )
      )
     )
     (br_if $label$4
      (i32.ne
       (i32.add
        (tee_local $4
         (i32.add
          (get_local $4)
          (i32.const -12)
         )
        )
        (get_local $7)
       )
       (i32.const -12)
      )
     )
    )
   )
   (i32.store
    (i32.add
     (get_local $1)
     (i32.const 4)
    )
    (get_local $2)
   )
   (set_local $5
    (get_local $2)
   )
  )
  (block $label$6
   (br_if $label$6
    (i32.eq
     (tee_local $4
      (i32.load
       (get_local $1)
      )
     )
     (get_local $5)
    )
   )
   (loop $label$7
    (drop
     (call $_ZN5eosiorsINS_10datastreamIPKcEEEERT_S6_RNSt3__112basic_stringIcNS7_11char_traitsIcEENS7_9allocatorIcEEEE
      (get_local $0)
      (get_local $4)
     )
    )
    (br_if $label$7
     (i32.ne
      (get_local $5)
      (tee_local $4
       (i32.add
        (get_local $4)
        (i32.const 12)
       )
      )
     )
    )
   )
  )
  (get_local $0)
 )
 (func $_ZN5eosiorsINS_10datastreamIPKcEEEERT_S6_RNSt3__112basic_stringIcNS7_11char_traitsIcEENS7_9allocatorIcEEEE (param $0 i32) (param $1 i32) (result i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i32)
  (local $6 i64)
  (local $7 i32)
  (set_local $5
   (i32.load offset=4
    (get_local $0)
   )
  )
  (set_local $7
   (i32.const 0)
  )
  (set_local $6
   (i64.const 0)
  )
  (set_local $2
   (i32.add
    (get_local $0)
    (i32.const 8)
   )
  )
  (set_local $3
   (i32.add
    (get_local $0)
    (i32.const 4)
   )
  )
  (loop $label$0
   (call $eosio_assert
    (i32.lt_u
     (get_local $5)
     (i32.load
      (get_local $2)
     )
    )
    (i32.const 432)
   )
   (set_local $4
    (i32.load8_u
     (tee_local $5
      (i32.load
       (get_local $3)
      )
     )
    )
   )
   (i32.store
    (get_local $3)
    (tee_local $5
     (i32.add
      (get_local $5)
      (i32.const 1)
     )
    )
   )
   (set_local $6
    (i64.or
     (i64.extend_u/i32
      (i32.shl
       (i32.and
        (get_local $4)
        (i32.const 127)
       )
       (tee_local $7
        (i32.and
         (get_local $7)
         (i32.const 255)
        )
       )
      )
     )
     (get_local $6)
    )
   )
   (set_local $7
    (i32.add
     (get_local $7)
     (i32.const 7)
    )
   )
   (br_if $label$0
    (i32.shr_u
     (get_local $4)
     (i32.const 7)
    )
   )
  )
  (call $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEE6resizeEjc
   (get_local $1)
   (i32.wrap/i64
    (get_local $6)
   )
   (i32.const 0)
  )
  (block $label$1
   (br_if $label$1
    (i32.eqz
     (tee_local $7
      (select
       (i32.load offset=4
        (get_local $1)
       )
       (i32.shr_u
        (tee_local $4
         (i32.load8_u
          (get_local $1)
         )
        )
        (i32.const 1)
       )
       (tee_local $4
        (i32.and
         (get_local $4)
         (i32.const 1)
        )
       )
      )
     )
    )
   )
   (set_local $4
    (select
     (i32.load offset=8
      (get_local $1)
     )
     (i32.add
      (get_local $1)
      (i32.const 1)
     )
     (get_local $4)
    )
   )
   (set_local $3
    (i32.load
     (tee_local $5
      (i32.add
       (get_local $0)
       (i32.const 4)
      )
     )
    )
   )
   (set_local $2
    (i32.add
     (get_local $0)
     (i32.const 8)
    )
   )
   (loop $label$2
    (call $eosio_assert
     (i32.ne
      (i32.load
       (get_local $2)
      )
      (get_local $3)
     )
     (i32.const 416)
    )
    (drop
     (call $memcpy
      (get_local $4)
      (i32.load
       (get_local $5)
      )
      (i32.const 1)
     )
    )
    (i32.store
     (get_local $5)
     (tee_local $3
      (i32.add
       (i32.load
        (get_local $5)
       )
       (i32.const 1)
      )
     )
    )
    (set_local $4
     (i32.add
      (get_local $4)
      (i32.const 1)
     )
    )
    (br_if $label$2
     (tee_local $7
      (i32.add
       (get_local $7)
       (i32.const -1)
      )
     )
    )
   )
  )
  (get_local $0)
 )
 (func $_ZN5eosiorsINS_10datastreamIPKcEEEERT_S6_RNS_13authority_mapE (param $0 i32) (param $1 i32) (result i32)
  (call $eosio_assert
   (i32.gt_u
    (i32.sub
     (i32.load offset=8
      (get_local $0)
     )
     (i32.load offset=4
      (get_local $0)
     )
    )
    (i32.const 3)
   )
   (i32.const 416)
  )
  (drop
   (call $memcpy
    (get_local $1)
    (i32.load offset=4
     (get_local $0)
    )
    (i32.const 4)
   )
  )
  (i32.store offset=4
   (get_local $0)
   (i32.add
    (i32.load offset=4
     (get_local $0)
    )
    (i32.const 4)
   )
  )
  (call $eosio_assert
   (i32.gt_u
    (i32.sub
     (i32.load offset=8
      (tee_local $0
       (call $_ZN5eosiorsINS_10datastreamIPKcEEN11eosiosystem23permission_level_weightEEERT_S8_RNSt3__16vectorIT0_NS9_9allocatorISB_EEEE
        (call $_ZN5eosiorsINS_10datastreamIPKcEEN11eosiosystem10key_weightEEERT_S8_RNSt3__16vectorIT0_NS9_9allocatorISB_EEEE
         (get_local $0)
         (i32.add
          (get_local $1)
          (i32.const 4)
         )
        )
        (i32.add
         (get_local $1)
         (i32.const 16)
        )
       )
      )
     )
     (i32.load offset=4
      (get_local $0)
     )
    )
    (i32.const 3)
   )
   (i32.const 416)
  )
  (drop
   (call $memcpy
    (i32.add
     (get_local $1)
     (i32.const 28)
    )
    (i32.load offset=4
     (get_local $0)
    )
    (i32.const 4)
   )
  )
  (i32.store offset=4
   (get_local $0)
   (i32.add
    (i32.load offset=4
     (get_local $0)
    )
    (i32.const 4)
   )
  )
  (call $eosio_assert
   (i32.gt_u
    (i32.sub
     (i32.load offset=8
      (tee_local $0
       (call $_ZN5eosiorsINS_10datastreamIPKcEEN11eosiosystem23permission_level_weightEEERT_S8_RNSt3__16vectorIT0_NS9_9allocatorISB_EEEE
        (call $_ZN5eosiorsINS_10datastreamIPKcEEN11eosiosystem10key_weightEEERT_S8_RNSt3__16vectorIT0_NS9_9allocatorISB_EEEE
         (get_local $0)
         (i32.add
          (get_local $1)
          (i32.const 32)
         )
        )
        (i32.add
         (get_local $1)
         (i32.const 44)
        )
       )
      )
     )
     (i32.load offset=4
      (get_local $0)
     )
    )
    (i32.const 3)
   )
   (i32.const 416)
  )
  (drop
   (call $memcpy
    (i32.add
     (get_local $1)
     (i32.const 56)
    )
    (i32.load offset=4
     (get_local $0)
    )
    (i32.const 4)
   )
  )
  (i32.store offset=4
   (get_local $0)
   (i32.add
    (i32.load offset=4
     (get_local $0)
    )
    (i32.const 4)
   )
  )
  (call $_ZN5eosiorsINS_10datastreamIPKcEEN11eosiosystem23permission_level_weightEEERT_S8_RNSt3__16vectorIT0_NS9_9allocatorISB_EEEE
   (call $_ZN5eosiorsINS_10datastreamIPKcEEN11eosiosystem10key_weightEEERT_S8_RNSt3__16vectorIT0_NS9_9allocatorISB_EEEE
    (get_local $0)
    (i32.add
     (get_local $1)
     (i32.const 60)
    )
   )
   (i32.add
    (get_local $1)
    (i32.const 72)
   )
  )
 )
 (func $_ZN5eosiorsINS_10datastreamIPKcEENS_11content_refEEERT_S7_RNSt3__16vectorIT0_NS8_9allocatorISA_EEEE (param $0 i32) (param $1 i32) (result i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i64)
  (local $6 i32)
  (local $7 i32)
  (set_local $4
   (i32.load offset=4
    (get_local $0)
   )
  )
  (set_local $6
   (i32.const 0)
  )
  (set_local $5
   (i64.const 0)
  )
  (set_local $2
   (i32.add
    (get_local $0)
    (i32.const 8)
   )
  )
  (set_local $3
   (i32.add
    (get_local $0)
    (i32.const 4)
   )
  )
  (loop $label$0
   (call $eosio_assert
    (i32.lt_u
     (get_local $4)
     (i32.load
      (get_local $2)
     )
    )
    (i32.const 432)
   )
   (set_local $7
    (i32.load8_u
     (tee_local $4
      (i32.load
       (get_local $3)
      )
     )
    )
   )
   (i32.store
    (get_local $3)
    (tee_local $4
     (i32.add
      (get_local $4)
      (i32.const 1)
     )
    )
   )
   (set_local $5
    (i64.or
     (i64.extend_u/i32
      (i32.shl
       (i32.and
        (get_local $7)
        (i32.const 127)
       )
       (tee_local $6
        (i32.and
         (get_local $6)
         (i32.const 255)
        )
       )
      )
     )
     (get_local $5)
    )
   )
   (set_local $6
    (i32.add
     (get_local $6)
     (i32.const 7)
    )
   )
   (br_if $label$0
    (i32.shr_u
     (get_local $7)
     (i32.const 7)
    )
   )
  )
  (call $_ZNSt3__16vectorIN5eosio11content_refENS_9allocatorIS2_EEE6resizeEj
   (get_local $1)
   (i32.wrap/i64
    (get_local $5)
   )
  )
  (block $label$1
   (br_if $label$1
    (i32.eq
     (tee_local $7
      (i32.load
       (get_local $1)
      )
     )
     (tee_local $4
      (i32.load offset=4
       (get_local $1)
      )
     )
    )
   )
   (loop $label$2
    (drop
     (call $_ZN5eosiorsINS_10datastreamIPKcEEEERT_S6_RNSt3__112basic_stringIcNS7_11char_traitsIcEENS7_9allocatorIcEEEE
      (get_local $0)
      (get_local $7)
     )
    )
    (drop
     (call $_ZN5eosiorsINS_10datastreamIPKcEEEERT_S6_RNSt3__112basic_stringIcNS7_11char_traitsIcEENS7_9allocatorIcEEEE
      (get_local $0)
      (i32.add
       (get_local $7)
       (i32.const 12)
      )
     )
    )
    (drop
     (call $_ZN5eosiorsINS_10datastreamIPKcEEEERT_S6_RNSt3__112basic_stringIcNS7_11char_traitsIcEENS7_9allocatorIcEEEE
      (get_local $0)
      (i32.add
       (get_local $7)
       (i32.const 24)
      )
     )
    )
    (br_if $label$2
     (i32.ne
      (tee_local $7
       (i32.add
        (get_local $7)
        (i32.const 36)
       )
      )
      (get_local $4)
     )
    )
   )
  )
  (get_local $0)
 )
 (func $_ZNSt3__16vectorIN5eosio11content_refENS_9allocatorIS2_EEE6resizeEj (param $0 i32) (param $1 i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i32)
  (block $label$0
   (br_if $label$0
    (i32.ge_u
     (tee_local $5
      (i32.div_s
       (i32.sub
        (tee_local $2
         (i32.load offset=4
          (get_local $0)
         )
        )
        (tee_local $3
         (i32.load
          (get_local $0)
         )
        )
       )
       (i32.const 36)
      )
     )
     (get_local $1)
    )
   )
   (call $_ZNSt3__16vectorIN5eosio11content_refENS_9allocatorIS2_EEE8__appendEj
    (get_local $0)
    (i32.sub
     (get_local $1)
     (get_local $5)
    )
   )
   (return)
  )
  (block $label$1
   (br_if $label$1
    (i32.le_u
     (get_local $5)
     (get_local $1)
    )
   )
   (block $label$2
    (br_if $label$2
     (i32.eq
      (get_local $2)
      (tee_local $4
       (i32.add
        (get_local $3)
        (tee_local $1
         (i32.mul
          (get_local $1)
          (i32.const 36)
         )
        )
       )
      )
     )
    )
    (set_local $5
     (i32.sub
      (i32.sub
       (i32.const 0)
       (get_local $3)
      )
      (get_local $1)
     )
    )
    (set_local $1
     (i32.add
      (get_local $2)
      (i32.const -36)
     )
    )
    (loop $label$3
     (block $label$4
      (br_if $label$4
       (i32.eqz
        (i32.and
         (i32.load8_u
          (i32.add
           (get_local $1)
           (i32.const 24)
          )
         )
         (i32.const 1)
        )
       )
      )
      (call $_ZdlPv
       (i32.load
        (i32.add
         (get_local $1)
         (i32.const 32)
        )
       )
      )
     )
     (block $label$5
      (br_if $label$5
       (i32.eqz
        (i32.and
         (i32.load8_u
          (i32.add
           (get_local $1)
           (i32.const 12)
          )
         )
         (i32.const 1)
        )
       )
      )
      (call $_ZdlPv
       (i32.load
        (i32.add
         (get_local $1)
         (i32.const 20)
        )
       )
      )
     )
     (block $label$6
      (br_if $label$6
       (i32.eqz
        (i32.and
         (i32.load8_u
          (get_local $1)
         )
         (i32.const 1)
        )
       )
      )
      (call $_ZdlPv
       (i32.load
        (i32.add
         (get_local $1)
         (i32.const 8)
        )
       )
      )
     )
     (br_if $label$3
      (i32.ne
       (i32.add
        (tee_local $1
         (i32.add
          (get_local $1)
          (i32.const -36)
         )
        )
        (get_local $5)
       )
       (i32.const -36)
      )
     )
    )
   )
   (i32.store
    (i32.add
     (get_local $0)
     (i32.const 4)
    )
    (get_local $4)
   )
  )
 )
 (func $_ZNSt3__16vectorIN5eosio11content_refENS_9allocatorIS2_EEE8__appendEj (param $0 i32) (param $1 i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i32)
  (local $6 i32)
  (local $7 i32)
  (block $label$0
   (block $label$1
    (block $label$2
     (block $label$3
      (br_if $label$3
       (i32.ge_u
        (i32.div_s
         (i32.sub
          (tee_local $7
           (i32.load offset=8
            (get_local $0)
           )
          )
          (tee_local $5
           (i32.load offset=4
            (get_local $0)
           )
          )
         )
         (i32.const 36)
        )
        (get_local $1)
       )
      )
      (br_if $label$1
       (i32.ge_u
        (tee_local $2
         (i32.add
          (tee_local $5
           (i32.div_s
            (i32.sub
             (get_local $5)
             (tee_local $4
              (i32.load
               (get_local $0)
              )
             )
            )
            (i32.const 36)
           )
          )
          (get_local $1)
         )
        )
        (i32.const 119304648)
       )
      )
      (set_local $6
       (i32.const 119304647)
      )
      (block $label$4
       (br_if $label$4
        (i32.gt_u
         (tee_local $7
          (i32.div_s
           (i32.sub
            (get_local $7)
            (get_local $4)
           )
           (i32.const 36)
          )
         )
         (i32.const 59652322)
        )
       )
       (br_if $label$2
        (i32.eqz
         (tee_local $6
          (select
           (get_local $2)
           (tee_local $6
            (i32.shl
             (get_local $7)
             (i32.const 1)
            )
           )
           (i32.lt_u
            (get_local $6)
            (get_local $2)
           )
          )
         )
        )
       )
      )
      (set_local $7
       (call $_Znwj
        (i32.mul
         (get_local $6)
         (i32.const 36)
        )
       )
      )
      (br $label$0)
     )
     (set_local $6
      (get_local $5)
     )
     (set_local $7
      (get_local $1)
     )
     (loop $label$5
      (i64.store align=4
       (get_local $6)
       (i64.const 0)
      )
      (i64.store align=4
       (i32.add
        (get_local $6)
        (i32.const 8)
       )
       (i64.const 0)
      )
      (i64.store align=4
       (i32.add
        (get_local $6)
        (i32.const 16)
       )
       (i64.const 0)
      )
      (i64.store align=4
       (i32.add
        (get_local $6)
        (i32.const 24)
       )
       (i64.const 0)
      )
      (i32.store
       (i32.add
        (get_local $6)
        (i32.const 32)
       )
       (i32.const 0)
      )
      (set_local $6
       (i32.add
        (get_local $6)
        (i32.const 36)
       )
      )
      (br_if $label$5
       (tee_local $7
        (i32.add
         (get_local $7)
         (i32.const -1)
        )
       )
      )
     )
     (i32.store
      (i32.add
       (get_local $0)
       (i32.const 4)
      )
      (i32.add
       (get_local $5)
       (i32.mul
        (get_local $1)
        (i32.const 36)
       )
      )
     )
     (return)
    )
    (set_local $6
     (i32.const 0)
    )
    (set_local $7
     (i32.const 0)
    )
    (br $label$0)
   )
   (call $_ZNKSt3__120__vector_base_commonILb1EE20__throw_length_errorEv
    (get_local $0)
   )
   (unreachable)
  )
  (set_local $2
   (i32.add
    (get_local $7)
    (i32.mul
     (get_local $6)
     (i32.const 36)
    )
   )
  )
  (set_local $6
   (tee_local $7
    (i32.add
     (get_local $7)
     (i32.mul
      (get_local $5)
      (i32.const 36)
     )
    )
   )
  )
  (set_local $5
   (get_local $1)
  )
  (loop $label$6
   (i64.store align=4
    (get_local $6)
    (i64.const 0)
   )
   (i64.store align=4
    (i32.add
     (get_local $6)
     (i32.const 8)
    )
    (i64.const 0)
   )
   (i64.store align=4
    (i32.add
     (get_local $6)
     (i32.const 16)
    )
    (i64.const 0)
   )
   (i64.store align=4
    (i32.add
     (get_local $6)
     (i32.const 24)
    )
    (i64.const 0)
   )
   (i32.store
    (i32.add
     (get_local $6)
     (i32.const 32)
    )
    (i32.const 0)
   )
   (set_local $6
    (i32.add
     (get_local $6)
     (i32.const 36)
    )
   )
   (br_if $label$6
    (tee_local $5
     (i32.add
      (get_local $5)
      (i32.const -1)
     )
    )
   )
  )
  (set_local $3
   (i32.add
    (get_local $7)
    (i32.mul
     (get_local $1)
     (i32.const 36)
    )
   )
  )
  (block $label$7
   (block $label$8
    (br_if $label$8
     (i32.eq
      (tee_local $5
       (i32.load
        (i32.add
         (get_local $0)
         (i32.const 4)
        )
       )
      )
      (tee_local $6
       (i32.load
        (get_local $0)
       )
      )
     )
    )
    (set_local $4
     (i32.sub
      (i32.const 0)
      (get_local $6)
     )
    )
    (set_local $6
     (i32.add
      (get_local $5)
      (i32.const -36)
     )
    )
    (loop $label$9
     (i64.store align=4
      (i32.add
       (get_local $7)
       (i32.const -36)
      )
      (i64.load align=4
       (get_local $6)
      )
     )
     (i32.store
      (i32.add
       (get_local $7)
       (i32.const -28)
      )
      (i32.load
       (tee_local $5
        (i32.add
         (get_local $6)
         (i32.const 8)
        )
       )
      )
     )
     (i32.store
      (get_local $6)
      (i32.const 0)
     )
     (i32.store
      (i32.add
       (get_local $6)
       (i32.const 4)
      )
      (i32.const 0)
     )
     (i32.store
      (get_local $5)
      (i32.const 0)
     )
     (i32.store
      (i32.add
       (get_local $7)
       (i32.const -16)
      )
      (i32.load
       (tee_local $5
        (i32.add
         (get_local $6)
         (i32.const 20)
        )
       )
      )
     )
     (i64.store align=4
      (i32.add
       (get_local $7)
       (i32.const -24)
      )
      (i64.load align=4
       (tee_local $1
        (i32.add
         (get_local $6)
         (i32.const 12)
        )
       )
      )
     )
     (i32.store
      (get_local $1)
      (i32.const 0)
     )
     (i32.store
      (i32.add
       (get_local $6)
       (i32.const 16)
      )
      (i32.const 0)
     )
     (i32.store
      (get_local $5)
      (i32.const 0)
     )
     (i32.store
      (i32.add
       (get_local $7)
       (i32.const -4)
      )
      (i32.load
       (tee_local $5
        (i32.add
         (get_local $6)
         (i32.const 32)
        )
       )
      )
     )
     (i64.store align=4
      (i32.add
       (get_local $7)
       (i32.const -12)
      )
      (i64.load align=4
       (tee_local $1
        (i32.add
         (get_local $6)
         (i32.const 24)
        )
       )
      )
     )
     (i32.store
      (get_local $1)
      (i32.const 0)
     )
     (i32.store
      (i32.add
       (get_local $6)
       (i32.const 28)
      )
      (i32.const 0)
     )
     (i32.store
      (get_local $5)
      (i32.const 0)
     )
     (set_local $7
      (i32.add
       (get_local $7)
       (i32.const -36)
      )
     )
     (br_if $label$9
      (i32.ne
       (i32.add
        (tee_local $6
         (i32.add
          (get_local $6)
          (i32.const -36)
         )
        )
        (get_local $4)
       )
       (i32.const -36)
      )
     )
    )
    (set_local $6
     (i32.load
      (i32.add
       (get_local $0)
       (i32.const 4)
      )
     )
    )
    (set_local $5
     (i32.load
      (get_local $0)
     )
    )
    (br $label$7)
   )
   (set_local $5
    (get_local $6)
   )
  )
  (i32.store
   (get_local $0)
   (get_local $7)
  )
  (i32.store
   (i32.add
    (get_local $0)
    (i32.const 4)
   )
   (get_local $3)
  )
  (i32.store
   (i32.add
    (get_local $0)
    (i32.const 8)
   )
   (get_local $2)
  )
  (block $label$10
   (br_if $label$10
    (i32.eq
     (get_local $6)
     (get_local $5)
    )
   )
   (set_local $7
    (i32.sub
     (i32.const 0)
     (get_local $5)
    )
   )
   (set_local $6
    (i32.add
     (get_local $6)
     (i32.const -36)
    )
   )
   (loop $label$11
    (block $label$12
     (br_if $label$12
      (i32.eqz
       (i32.and
        (i32.load8_u
         (i32.add
          (get_local $6)
          (i32.const 24)
         )
        )
        (i32.const 1)
       )
      )
     )
     (call $_ZdlPv
      (i32.load
       (i32.add
        (get_local $6)
        (i32.const 32)
       )
      )
     )
    )
    (block $label$13
     (br_if $label$13
      (i32.eqz
       (i32.and
        (i32.load8_u
         (i32.add
          (get_local $6)
          (i32.const 12)
         )
        )
        (i32.const 1)
       )
      )
     )
     (call $_ZdlPv
      (i32.load
       (i32.add
        (get_local $6)
        (i32.const 20)
       )
      )
     )
    )
    (block $label$14
     (br_if $label$14
      (i32.eqz
       (i32.and
        (i32.load8_u
         (get_local $6)
        )
        (i32.const 1)
       )
      )
     )
     (call $_ZdlPv
      (i32.load
       (i32.add
        (get_local $6)
        (i32.const 8)
       )
      )
     )
    )
    (br_if $label$11
     (i32.ne
      (i32.add
       (tee_local $6
        (i32.add
         (get_local $6)
         (i32.const -36)
        )
       )
       (get_local $7)
      )
      (i32.const -36)
     )
    )
   )
  )
  (block $label$15
   (br_if $label$15
    (i32.eqz
     (get_local $5)
    )
   )
   (call $_ZdlPv
    (get_local $5)
   )
  )
 )
 (func $_ZN5eosiorsINS_10datastreamIPKcEEN11eosiosystem10key_weightEEERT_S8_RNSt3__16vectorIT0_NS9_9allocatorISB_EEEE (param $0 i32) (param $1 i32) (result i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i32)
  (local $6 i64)
  (local $7 i32)
  (set_local $5
   (i32.load offset=4
    (get_local $0)
   )
  )
  (set_local $7
   (i32.const 0)
  )
  (set_local $6
   (i64.const 0)
  )
  (set_local $2
   (i32.add
    (get_local $0)
    (i32.const 8)
   )
  )
  (set_local $3
   (i32.add
    (get_local $0)
    (i32.const 4)
   )
  )
  (loop $label$0
   (call $eosio_assert
    (i32.lt_u
     (get_local $5)
     (i32.load
      (get_local $2)
     )
    )
    (i32.const 432)
   )
   (set_local $4
    (i32.load8_u
     (tee_local $5
      (i32.load
       (get_local $3)
      )
     )
    )
   )
   (i32.store
    (get_local $3)
    (tee_local $5
     (i32.add
      (get_local $5)
      (i32.const 1)
     )
    )
   )
   (set_local $6
    (i64.or
     (i64.extend_u/i32
      (i32.shl
       (i32.and
        (get_local $4)
        (i32.const 127)
       )
       (tee_local $7
        (i32.and
         (get_local $7)
         (i32.const 255)
        )
       )
      )
     )
     (get_local $6)
    )
   )
   (set_local $7
    (i32.add
     (get_local $7)
     (i32.const 7)
    )
   )
   (br_if $label$0
    (i32.shr_u
     (get_local $4)
     (i32.const 7)
    )
   )
  )
  (block $label$1
   (block $label$2
    (block $label$3
     (br_if $label$3
      (i32.le_u
       (tee_local $7
        (i32.wrap/i64
         (get_local $6)
        )
       )
       (tee_local $3
        (i32.div_s
         (i32.sub
          (tee_local $5
           (i32.load offset=4
            (get_local $1)
           )
          )
          (tee_local $4
           (i32.load
            (get_local $1)
           )
          )
         )
         (i32.const 44)
        )
       )
      )
     )
     (call $_ZNSt3__16vectorIN11eosiosystem10key_weightENS_9allocatorIS2_EEE8__appendEj
      (get_local $1)
      (i32.sub
       (get_local $7)
       (get_local $3)
      )
     )
     (br_if $label$2
      (i32.ne
       (tee_local $4
        (i32.load
         (get_local $1)
        )
       )
       (tee_local $5
        (i32.load
         (i32.add
          (get_local $1)
          (i32.const 4)
         )
        )
       )
      )
     )
     (br $label$1)
    )
    (block $label$4
     (br_if $label$4
      (i32.ge_u
       (get_local $7)
       (get_local $3)
      )
     )
     (i32.store
      (i32.add
       (get_local $1)
       (i32.const 4)
      )
      (tee_local $5
       (i32.add
        (get_local $4)
        (i32.mul
         (get_local $7)
         (i32.const 44)
        )
       )
      )
     )
    )
    (br_if $label$1
     (i32.eq
      (get_local $4)
      (get_local $5)
     )
    )
   )
   (loop $label$5
    (drop
     (call $_ZN11eosiosystemrsIN5eosio10datastreamIPKcEEEERT_S7_RNS_10key_weightE
      (get_local $0)
      (get_local $4)
     )
    )
    (br_if $label$5
     (i32.ne
      (get_local $5)
      (tee_local $4
       (i32.add
        (get_local $4)
        (i32.const 44)
       )
      )
     )
    )
   )
  )
  (get_local $0)
 )
 (func $_ZN5eosiorsINS_10datastreamIPKcEEN11eosiosystem23permission_level_weightEEERT_S8_RNSt3__16vectorIT0_NS9_9allocatorISB_EEEE (param $0 i32) (param $1 i32) (result i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i64)
  (local $6 i32)
  (local $7 i32)
  (set_local $7
   (i32.load offset=4
    (get_local $0)
   )
  )
  (set_local $6
   (i32.const 0)
  )
  (set_local $5
   (i64.const 0)
  )
  (set_local $2
   (i32.add
    (get_local $0)
    (i32.const 8)
   )
  )
  (set_local $3
   (i32.add
    (get_local $0)
    (i32.const 4)
   )
  )
  (loop $label$0
   (call $eosio_assert
    (i32.lt_u
     (get_local $7)
     (i32.load
      (get_local $2)
     )
    )
    (i32.const 432)
   )
   (set_local $4
    (i32.load8_u
     (tee_local $7
      (i32.load
       (get_local $3)
      )
     )
    )
   )
   (i32.store
    (get_local $3)
    (tee_local $7
     (i32.add
      (get_local $7)
      (i32.const 1)
     )
    )
   )
   (set_local $5
    (i64.or
     (i64.extend_u/i32
      (i32.shl
       (i32.and
        (get_local $4)
        (i32.const 127)
       )
       (tee_local $6
        (i32.and
         (get_local $6)
         (i32.const 255)
        )
       )
      )
     )
     (get_local $5)
    )
   )
   (set_local $6
    (i32.add
     (get_local $6)
     (i32.const 7)
    )
   )
   (br_if $label$0
    (i32.shr_u
     (get_local $4)
     (i32.const 7)
    )
   )
  )
  (block $label$1
   (block $label$2
    (block $label$3
     (br_if $label$3
      (i32.le_u
       (tee_local $4
        (i32.wrap/i64
         (get_local $5)
        )
       )
       (tee_local $6
        (i32.div_s
         (i32.sub
          (tee_local $2
           (i32.load offset=4
            (get_local $1)
           )
          )
          (tee_local $7
           (i32.load
            (get_local $1)
           )
          )
         )
         (i32.const 24)
        )
       )
      )
     )
     (call $_ZNSt3__16vectorIN11eosiosystem23permission_level_weightENS_9allocatorIS2_EEE8__appendEj
      (get_local $1)
      (i32.sub
       (get_local $4)
       (get_local $6)
      )
     )
     (br_if $label$2
      (i32.ne
       (tee_local $7
        (i32.load
         (get_local $1)
        )
       )
       (tee_local $2
        (i32.load
         (i32.add
          (get_local $1)
          (i32.const 4)
         )
        )
       )
      )
     )
     (br $label$1)
    )
    (block $label$4
     (br_if $label$4
      (i32.ge_u
       (get_local $4)
       (get_local $6)
      )
     )
     (i32.store
      (i32.add
       (get_local $1)
       (i32.const 4)
      )
      (tee_local $2
       (i32.add
        (get_local $7)
        (i32.mul
         (get_local $4)
         (i32.const 24)
        )
       )
      )
     )
    )
    (br_if $label$1
     (i32.eq
      (get_local $7)
      (get_local $2)
     )
    )
   )
   (set_local $3
    (i32.load
     (tee_local $4
      (i32.add
       (get_local $0)
       (i32.const 4)
      )
     )
    )
   )
   (loop $label$5
    (call $eosio_assert
     (i32.gt_u
      (i32.sub
       (i32.load
        (tee_local $6
         (i32.add
          (get_local $0)
          (i32.const 8)
         )
        )
       )
       (get_local $3)
      )
      (i32.const 7)
     )
     (i32.const 416)
    )
    (drop
     (call $memcpy
      (get_local $7)
      (i32.load
       (get_local $4)
      )
      (i32.const 8)
     )
    )
    (i32.store
     (get_local $4)
     (tee_local $3
      (i32.add
       (i32.load
        (get_local $4)
       )
       (i32.const 8)
      )
     )
    )
    (call $eosio_assert
     (i32.gt_u
      (i32.sub
       (i32.load
        (get_local $6)
       )
       (get_local $3)
      )
      (i32.const 7)
     )
     (i32.const 416)
    )
    (drop
     (call $memcpy
      (i32.add
       (get_local $7)
       (i32.const 8)
      )
      (i32.load
       (get_local $4)
      )
      (i32.const 8)
     )
    )
    (i32.store
     (get_local $4)
     (tee_local $3
      (i32.add
       (i32.load
        (get_local $4)
       )
       (i32.const 8)
      )
     )
    )
    (call $eosio_assert
     (i32.gt_u
      (i32.sub
       (i32.load
        (get_local $6)
       )
       (get_local $3)
      )
      (i32.const 1)
     )
     (i32.const 416)
    )
    (drop
     (call $memcpy
      (i32.add
       (get_local $7)
       (i32.const 16)
      )
      (i32.load
       (get_local $4)
      )
      (i32.const 2)
     )
    )
    (i32.store
     (get_local $4)
     (tee_local $3
      (i32.add
       (i32.load
        (get_local $4)
       )
       (i32.const 2)
      )
     )
    )
    (br_if $label$5
     (i32.ne
      (tee_local $7
       (i32.add
        (get_local $7)
        (i32.const 24)
       )
      )
      (get_local $2)
     )
    )
   )
  )
  (get_local $0)
 )
 (func $_ZNSt3__16vectorIN11eosiosystem23permission_level_weightENS_9allocatorIS2_EEE8__appendEj (param $0 i32) (param $1 i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i32)
  (local $6 i32)
  (block $label$0
   (block $label$1
    (block $label$2
     (block $label$3
      (block $label$4
       (br_if $label$4
        (i32.ge_u
         (i32.div_s
          (i32.sub
           (tee_local $2
            (i32.load offset=8
             (get_local $0)
            )
           )
           (tee_local $6
            (i32.load offset=4
             (get_local $0)
            )
           )
          )
          (i32.const 24)
         )
         (get_local $1)
        )
       )
       (br_if $label$2
        (i32.ge_u
         (tee_local $4
          (i32.add
           (tee_local $3
            (i32.div_s
             (i32.sub
              (get_local $6)
              (tee_local $5
               (i32.load
                (get_local $0)
               )
              )
             )
             (i32.const 24)
            )
           )
           (get_local $1)
          )
         )
         (i32.const 178956971)
        )
       )
       (set_local $6
        (i32.const 178956970)
       )
       (block $label$5
        (br_if $label$5
         (i32.gt_u
          (tee_local $2
           (i32.div_s
            (i32.sub
             (get_local $2)
             (get_local $5)
            )
            (i32.const 24)
           )
          )
          (i32.const 89478484)
         )
        )
        (br_if $label$3
         (i32.eqz
          (tee_local $6
           (select
            (get_local $4)
            (tee_local $6
             (i32.shl
              (get_local $2)
              (i32.const 1)
             )
            )
            (i32.lt_u
             (get_local $6)
             (get_local $4)
            )
           )
          )
         )
        )
       )
       (set_local $2
        (call $_Znwj
         (i32.mul
          (get_local $6)
          (i32.const 24)
         )
        )
       )
       (br $label$1)
      )
      (set_local $0
       (i32.add
        (get_local $0)
        (i32.const 4)
       )
      )
      (loop $label$6
       (i64.store
        (get_local $6)
        (i64.const 0)
       )
       (i64.store
        (i32.add
         (get_local $6)
         (i32.const 16)
        )
        (i64.const 0)
       )
       (i64.store
        (i32.add
         (get_local $6)
         (i32.const 8)
        )
        (i64.const 0)
       )
       (i32.store
        (get_local $0)
        (tee_local $6
         (i32.add
          (i32.load
           (get_local $0)
          )
          (i32.const 24)
         )
        )
       )
       (br_if $label$6
        (tee_local $1
         (i32.add
          (get_local $1)
          (i32.const -1)
         )
        )
       )
       (br $label$0)
      )
     )
     (set_local $6
      (i32.const 0)
     )
     (set_local $2
      (i32.const 0)
     )
     (br $label$1)
    )
    (call $_ZNKSt3__120__vector_base_commonILb1EE20__throw_length_errorEv
     (get_local $0)
    )
    (unreachable)
   )
   (set_local $4
    (i32.add
     (get_local $2)
     (i32.mul
      (get_local $6)
      (i32.const 24)
     )
    )
   )
   (set_local $6
    (tee_local $5
     (i32.add
      (get_local $2)
      (i32.mul
       (get_local $3)
       (i32.const 24)
      )
     )
    )
   )
   (loop $label$7
    (i64.store
     (get_local $6)
     (i64.const 0)
    )
    (i64.store
     (i32.add
      (get_local $6)
      (i32.const 16)
     )
     (i64.const 0)
    )
    (i64.store
     (i32.add
      (get_local $6)
      (i32.const 8)
     )
     (i64.const 0)
    )
    (set_local $6
     (i32.add
      (get_local $6)
      (i32.const 24)
     )
    )
    (br_if $label$7
     (tee_local $1
      (i32.add
       (get_local $1)
       (i32.const -1)
      )
     )
    )
   )
   (set_local $5
    (i32.add
     (get_local $5)
     (i32.mul
      (i32.div_s
       (tee_local $2
        (i32.sub
         (i32.load
          (tee_local $3
           (i32.add
            (get_local $0)
            (i32.const 4)
           )
          )
         )
         (tee_local $1
          (i32.load
           (get_local $0)
          )
         )
        )
       )
       (i32.const -24)
      )
      (i32.const 24)
     )
    )
   )
   (block $label$8
    (br_if $label$8
     (i32.lt_s
      (get_local $2)
      (i32.const 1)
     )
    )
    (drop
     (call $memcpy
      (get_local $5)
      (get_local $1)
      (get_local $2)
     )
    )
    (set_local $1
     (i32.load
      (get_local $0)
     )
    )
   )
   (i32.store
    (get_local $0)
    (get_local $5)
   )
   (i32.store
    (get_local $3)
    (get_local $6)
   )
   (i32.store
    (i32.add
     (get_local $0)
     (i32.const 8)
    )
    (get_local $4)
   )
   (br_if $label$0
    (i32.eqz
     (get_local $1)
    )
   )
   (call $_ZdlPv
    (get_local $1)
   )
   (return)
  )
 )
 (func $_ZNSt3__16vectorIN11eosiosystem10key_weightENS_9allocatorIS2_EEE8__appendEj (param $0 i32) (param $1 i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i32)
  (local $6 i32)
  (block $label$0
   (block $label$1
    (block $label$2
     (block $label$3
      (block $label$4
       (br_if $label$4
        (i32.ge_u
         (i32.div_s
          (i32.sub
           (tee_local $2
            (i32.load offset=8
             (get_local $0)
            )
           )
           (tee_local $6
            (i32.load offset=4
             (get_local $0)
            )
           )
          )
          (i32.const 44)
         )
         (get_local $1)
        )
       )
       (br_if $label$2
        (i32.ge_u
         (tee_local $4
          (i32.add
           (tee_local $3
            (i32.div_s
             (i32.sub
              (get_local $6)
              (tee_local $5
               (i32.load
                (get_local $0)
               )
              )
             )
             (i32.const 44)
            )
           )
           (get_local $1)
          )
         )
         (i32.const 97612894)
        )
       )
       (set_local $6
        (i32.const 97612893)
       )
       (block $label$5
        (br_if $label$5
         (i32.gt_u
          (tee_local $2
           (i32.div_s
            (i32.sub
             (get_local $2)
             (get_local $5)
            )
            (i32.const 44)
           )
          )
          (i32.const 48806445)
         )
        )
        (br_if $label$3
         (i32.eqz
          (tee_local $6
           (select
            (get_local $4)
            (tee_local $6
             (i32.shl
              (get_local $2)
              (i32.const 1)
             )
            )
            (i32.lt_u
             (get_local $6)
             (get_local $4)
            )
           )
          )
         )
        )
       )
       (set_local $2
        (call $_Znwj
         (i32.mul
          (get_local $6)
          (i32.const 44)
         )
        )
       )
       (br $label$1)
      )
      (loop $label$6
       (drop
        (call $memset
         (i32.add
          (get_local $6)
          (i32.const 4)
         )
         (i32.const 0)
         (i32.const 40)
        )
       )
       (i32.store
        (get_local $6)
        (i32.const 0)
       )
       (i32.store
        (tee_local $6
         (i32.add
          (get_local $0)
          (i32.const 4)
         )
        )
        (tee_local $6
         (i32.add
          (i32.load
           (get_local $6)
          )
          (i32.const 44)
         )
        )
       )
       (br_if $label$6
        (tee_local $1
         (i32.add
          (get_local $1)
          (i32.const -1)
         )
        )
       )
       (br $label$0)
      )
     )
     (set_local $6
      (i32.const 0)
     )
     (set_local $2
      (i32.const 0)
     )
     (br $label$1)
    )
    (call $_ZNKSt3__120__vector_base_commonILb1EE20__throw_length_errorEv
     (get_local $0)
    )
    (unreachable)
   )
   (set_local $4
    (i32.add
     (get_local $2)
     (i32.mul
      (get_local $6)
      (i32.const 44)
     )
    )
   )
   (set_local $6
    (tee_local $5
     (i32.add
      (get_local $2)
      (i32.mul
       (get_local $3)
       (i32.const 44)
      )
     )
    )
   )
   (loop $label$7
    (drop
     (call $memset
      (i32.add
       (get_local $6)
       (i32.const 4)
      )
      (i32.const 0)
      (i32.const 40)
     )
    )
    (i32.store
     (get_local $6)
     (i32.const 0)
    )
    (set_local $6
     (i32.add
      (get_local $6)
      (i32.const 44)
     )
    )
    (br_if $label$7
     (tee_local $1
      (i32.add
       (get_local $1)
       (i32.const -1)
      )
     )
    )
   )
   (set_local $5
    (i32.add
     (get_local $5)
     (i32.mul
      (i32.div_s
       (tee_local $2
        (i32.sub
         (i32.load
          (tee_local $3
           (i32.add
            (get_local $0)
            (i32.const 4)
           )
          )
         )
         (tee_local $1
          (i32.load
           (get_local $0)
          )
         )
        )
       )
       (i32.const -44)
      )
      (i32.const 44)
     )
    )
   )
   (block $label$8
    (br_if $label$8
     (i32.lt_s
      (get_local $2)
      (i32.const 1)
     )
    )
    (drop
     (call $memcpy
      (get_local $5)
      (get_local $1)
      (get_local $2)
     )
    )
    (set_local $1
     (i32.load
      (get_local $0)
     )
    )
   )
   (i32.store
    (get_local $0)
    (get_local $5)
   )
   (i32.store
    (get_local $3)
    (get_local $6)
   )
   (i32.store
    (i32.add
     (get_local $0)
     (i32.const 8)
    )
    (get_local $4)
   )
   (br_if $label$0
    (i32.eqz
     (get_local $1)
    )
   )
   (call $_ZdlPv
    (get_local $1)
   )
   (return)
  )
 )
 (func $_ZN11eosiosystemrsIN5eosio10datastreamIPKcEEEERT_S7_RNS_10key_weightE (param $0 i32) (param $1 i32) (result i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i64)
  (local $6 i32)
  (local $7 i32)
  (set_local $7
   (i32.load offset=4
    (get_local $0)
   )
  )
  (set_local $4
   (i32.const 0)
  )
  (set_local $5
   (i64.const 0)
  )
  (set_local $2
   (i32.add
    (get_local $0)
    (i32.const 8)
   )
  )
  (set_local $3
   (i32.add
    (get_local $0)
    (i32.const 4)
   )
  )
  (loop $label$0
   (call $eosio_assert
    (i32.lt_u
     (get_local $7)
     (i32.load
      (get_local $2)
     )
    )
    (i32.const 432)
   )
   (set_local $6
    (i32.load8_u
     (tee_local $7
      (i32.load
       (get_local $3)
      )
     )
    )
   )
   (i32.store
    (get_local $3)
    (tee_local $7
     (i32.add
      (get_local $7)
      (i32.const 1)
     )
    )
   )
   (set_local $5
    (i64.or
     (i64.extend_u/i32
      (i32.shl
       (i32.and
        (get_local $6)
        (i32.const 127)
       )
       (tee_local $4
        (i32.and
         (get_local $4)
         (i32.const 255)
        )
       )
      )
     )
     (get_local $5)
    )
   )
   (set_local $4
    (i32.add
     (get_local $4)
     (i32.const 7)
    )
   )
   (br_if $label$0
    (i32.shr_u
     (get_local $6)
     (i32.const 7)
    )
   )
  )
  (i64.store32
   (get_local $1)
   (get_local $5)
  )
  (set_local $3
   (i32.add
    (get_local $0)
    (i32.const 8)
   )
  )
  (set_local $6
   (i32.const 4)
  )
  (loop $label$1
   (call $eosio_assert
    (i32.ne
     (i32.load
      (get_local $3)
     )
     (get_local $7)
    )
    (i32.const 416)
   )
   (drop
    (call $memcpy
     (i32.add
      (get_local $1)
      (get_local $6)
     )
     (i32.load
      (tee_local $4
       (i32.add
        (get_local $0)
        (i32.const 4)
       )
      )
     )
     (i32.const 1)
    )
   )
   (i32.store
    (get_local $4)
    (tee_local $7
     (i32.add
      (i32.load
       (get_local $4)
      )
      (i32.const 1)
     )
    )
   )
   (br_if $label$1
    (i32.ne
     (tee_local $6
      (i32.add
       (get_local $6)
       (i32.const 1)
      )
     )
     (i32.const 37)
    )
   )
  )
  (call $eosio_assert
   (i32.gt_u
    (i32.sub
     (i32.load
      (i32.add
       (get_local $0)
       (i32.const 8)
      )
     )
     (get_local $7)
    )
    (i32.const 1)
   )
   (i32.const 416)
  )
  (drop
   (call $memcpy
    (i32.add
     (get_local $1)
     (i32.const 40)
    )
    (i32.load
     (tee_local $7
      (i32.add
       (get_local $0)
       (i32.const 4)
      )
     )
    )
    (i32.const 2)
   )
  )
  (i32.store
   (get_local $7)
   (i32.add
    (i32.load
     (get_local $7)
    )
    (i32.const 2)
   )
  )
  (get_local $0)
 )
 (func $_ZNSt3__16vectorINS_12basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEENS4_IS6_EEE8__appendEj (param $0 i32) (param $1 i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i32)
  (local $6 i32)
  (block $label$0
   (block $label$1
    (block $label$2
     (block $label$3
      (br_if $label$3
       (i32.ge_u
        (i32.div_s
         (i32.sub
          (tee_local $6
           (i32.load offset=8
            (get_local $0)
           )
          )
          (tee_local $5
           (i32.load offset=4
            (get_local $0)
           )
          )
         )
         (i32.const 12)
        )
        (get_local $1)
       )
      )
      (br_if $label$1
       (i32.ge_u
        (tee_local $3
         (i32.add
          (tee_local $5
           (i32.div_s
            (i32.sub
             (get_local $5)
             (tee_local $2
              (i32.load
               (get_local $0)
              )
             )
            )
            (i32.const 12)
           )
          )
          (get_local $1)
         )
        )
        (i32.const 357913942)
       )
      )
      (set_local $4
       (i32.const 357913941)
      )
      (block $label$4
       (br_if $label$4
        (i32.gt_u
         (tee_local $6
          (i32.div_s
           (i32.sub
            (get_local $6)
            (get_local $2)
           )
           (i32.const 12)
          )
         )
         (i32.const 178956969)
        )
       )
       (br_if $label$2
        (i32.eqz
         (tee_local $4
          (select
           (get_local $3)
           (tee_local $4
            (i32.shl
             (get_local $6)
             (i32.const 1)
            )
           )
           (i32.lt_u
            (get_local $4)
            (get_local $3)
           )
          )
         )
        )
       )
      )
      (set_local $6
       (call $_Znwj
        (i32.mul
         (get_local $4)
         (i32.const 12)
        )
       )
      )
      (br $label$0)
     )
     (set_local $4
      (get_local $5)
     )
     (set_local $6
      (get_local $1)
     )
     (loop $label$5
      (i64.store align=4
       (get_local $4)
       (i64.const 0)
      )
      (i32.store
       (i32.add
        (get_local $4)
        (i32.const 8)
       )
       (i32.const 0)
      )
      (set_local $4
       (i32.add
        (get_local $4)
        (i32.const 12)
       )
      )
      (br_if $label$5
       (tee_local $6
        (i32.add
         (get_local $6)
         (i32.const -1)
        )
       )
      )
     )
     (i32.store
      (i32.add
       (get_local $0)
       (i32.const 4)
      )
      (i32.add
       (get_local $5)
       (i32.mul
        (get_local $1)
        (i32.const 12)
       )
      )
     )
     (return)
    )
    (set_local $4
     (i32.const 0)
    )
    (set_local $6
     (i32.const 0)
    )
    (br $label$0)
   )
   (call $_ZNKSt3__120__vector_base_commonILb1EE20__throw_length_errorEv
    (get_local $0)
   )
   (unreachable)
  )
  (set_local $2
   (i32.add
    (get_local $6)
    (i32.mul
     (get_local $4)
     (i32.const 12)
    )
   )
  )
  (set_local $4
   (tee_local $6
    (i32.add
     (get_local $6)
     (i32.mul
      (get_local $5)
      (i32.const 12)
     )
    )
   )
  )
  (set_local $5
   (get_local $1)
  )
  (loop $label$6
   (i64.store align=4
    (get_local $4)
    (i64.const 0)
   )
   (i32.store
    (i32.add
     (get_local $4)
     (i32.const 8)
    )
    (i32.const 0)
   )
   (set_local $4
    (i32.add
     (get_local $4)
     (i32.const 12)
    )
   )
   (br_if $label$6
    (tee_local $5
     (i32.add
      (get_local $5)
      (i32.const -1)
     )
    )
   )
  )
  (set_local $3
   (i32.add
    (get_local $6)
    (i32.mul
     (get_local $1)
     (i32.const 12)
    )
   )
  )
  (block $label$7
   (block $label$8
    (br_if $label$8
     (i32.eq
      (tee_local $5
       (i32.load
        (i32.add
         (get_local $0)
         (i32.const 4)
        )
       )
      )
      (tee_local $4
       (i32.load
        (get_local $0)
       )
      )
     )
    )
    (set_local $1
     (i32.sub
      (i32.const 0)
      (get_local $4)
     )
    )
    (set_local $4
     (i32.add
      (get_local $5)
      (i32.const -12)
     )
    )
    (loop $label$9
     (i64.store align=4
      (i32.add
       (get_local $6)
       (i32.const -12)
      )
      (i64.load align=4
       (get_local $4)
      )
     )
     (i32.store
      (i32.add
       (get_local $6)
       (i32.const -4)
      )
      (i32.load
       (tee_local $5
        (i32.add
         (get_local $4)
         (i32.const 8)
        )
       )
      )
     )
     (i32.store
      (get_local $4)
      (i32.const 0)
     )
     (i32.store
      (i32.add
       (get_local $4)
       (i32.const 4)
      )
      (i32.const 0)
     )
     (i32.store
      (get_local $5)
      (i32.const 0)
     )
     (set_local $6
      (i32.add
       (get_local $6)
       (i32.const -12)
      )
     )
     (br_if $label$9
      (i32.ne
       (i32.add
        (tee_local $4
         (i32.add
          (get_local $4)
          (i32.const -12)
         )
        )
        (get_local $1)
       )
       (i32.const -12)
      )
     )
    )
    (set_local $4
     (i32.load
      (i32.add
       (get_local $0)
       (i32.const 4)
      )
     )
    )
    (set_local $5
     (i32.load
      (get_local $0)
     )
    )
    (br $label$7)
   )
   (set_local $5
    (get_local $4)
   )
  )
  (i32.store
   (get_local $0)
   (get_local $6)
  )
  (i32.store
   (i32.add
    (get_local $0)
    (i32.const 4)
   )
   (get_local $3)
  )
  (i32.store
   (i32.add
    (get_local $0)
    (i32.const 8)
   )
   (get_local $2)
  )
  (block $label$10
   (br_if $label$10
    (i32.eq
     (get_local $4)
     (get_local $5)
    )
   )
   (set_local $6
    (i32.sub
     (i32.const 0)
     (get_local $5)
    )
   )
   (set_local $4
    (i32.add
     (get_local $4)
     (i32.const -12)
    )
   )
   (loop $label$11
    (block $label$12
     (br_if $label$12
      (i32.eqz
       (i32.and
        (i32.load8_u
         (get_local $4)
        )
        (i32.const 1)
       )
      )
     )
     (call $_ZdlPv
      (i32.load
       (i32.add
        (get_local $4)
        (i32.const 8)
       )
      )
     )
    )
    (br_if $label$11
     (i32.ne
      (i32.add
       (tee_local $4
        (i32.add
         (get_local $4)
         (i32.const -12)
        )
       )
       (get_local $6)
      )
      (i32.const -12)
     )
    )
   )
  )
  (block $label$13
   (br_if $label$13
    (i32.eqz
     (get_local $5)
    )
   )
   (call $_ZdlPv
    (get_local $5)
   )
  )
 )
 (func $_ZNSt3__16vectorIN5eosio9peer_linkENS_9allocatorIS2_EEE8__appendEj (param $0 i32) (param $1 i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i32)
  (local $6 i32)
  (local $7 i32)
  (block $label$0
   (block $label$1
    (block $label$2
     (block $label$3
      (block $label$4
       (block $label$5
        (br_if $label$5
         (i32.ge_u
          (i32.shr_s
           (i32.sub
            (tee_local $7
             (i32.load offset=8
              (get_local $0)
             )
            )
            (tee_local $6
             (i32.load offset=4
              (get_local $0)
             )
            )
           )
           (i32.const 5)
          )
          (get_local $1)
         )
        )
        (br_if $label$3
         (i32.ge_u
          (tee_local $5
           (i32.add
            (tee_local $3
             (i32.shr_s
              (i32.sub
               (get_local $6)
               (tee_local $4
                (i32.load
                 (get_local $0)
                )
               )
              )
              (i32.const 5)
             )
            )
            (get_local $1)
           )
          )
          (i32.const 134217728)
         )
        )
        (set_local $6
         (i32.const 134217727)
        )
        (block $label$6
         (br_if $label$6
          (i32.gt_u
           (i32.shr_s
            (tee_local $7
             (i32.sub
              (get_local $7)
              (get_local $4)
             )
            )
            (i32.const 5)
           )
           (i32.const 67108862)
          )
         )
         (br_if $label$4
          (i32.eqz
           (tee_local $6
            (select
             (get_local $5)
             (tee_local $6
              (i32.shr_s
               (get_local $7)
               (i32.const 4)
              )
             )
             (i32.lt_u
              (get_local $6)
              (get_local $5)
             )
            )
           )
          )
         )
         (br_if $label$2
          (i32.ge_u
           (get_local $6)
           (i32.const 134217728)
          )
         )
        )
        (set_local $7
         (call $_Znwj
          (i32.shl
           (get_local $6)
           (i32.const 5)
          )
         )
        )
        (br $label$1)
       )
       (set_local $7
        (i32.add
         (get_local $0)
         (i32.const 4)
        )
       )
       (loop $label$7
        (i64.store
         (tee_local $4
          (i32.add
           (get_local $6)
           (i32.const 8)
          )
         )
         (i64.const 0)
        )
        (i64.store
         (tee_local $5
          (i32.add
           (get_local $6)
           (i32.const 16)
          )
         )
         (i64.const 0)
        )
        (i64.store
         (get_local $6)
         (i64.const 0)
        )
        (i64.store
         (i32.add
          (get_local $6)
          (i32.const 24)
         )
         (i64.const 0)
        )
        (i32.store
         (get_local $4)
         (i32.const 0)
        )
        (i32.store
         (get_local $5)
         (i32.const 0)
        )
        (i32.store
         (get_local $7)
         (tee_local $6
          (i32.add
           (i32.load
            (get_local $7)
           )
           (i32.const 32)
          )
         )
        )
        (br_if $label$7
         (tee_local $1
          (i32.add
           (get_local $1)
           (i32.const -1)
          )
         )
        )
        (br $label$0)
       )
      )
      (set_local $6
       (i32.const 0)
      )
      (set_local $7
       (i32.const 0)
      )
      (br $label$1)
     )
     (call $_ZNKSt3__120__vector_base_commonILb1EE20__throw_length_errorEv
      (get_local $0)
     )
     (unreachable)
    )
    (call $abort)
    (unreachable)
   )
   (set_local $2
    (i32.add
     (get_local $7)
     (i32.shl
      (get_local $6)
      (i32.const 5)
     )
    )
   )
   (set_local $6
    (tee_local $7
     (i32.add
      (get_local $7)
      (i32.shl
       (get_local $3)
       (i32.const 5)
      )
     )
    )
   )
   (loop $label$8
    (i64.store
     (tee_local $4
      (i32.add
       (get_local $6)
       (i32.const 8)
      )
     )
     (i64.const 0)
    )
    (i64.store
     (tee_local $5
      (i32.add
       (get_local $6)
       (i32.const 16)
      )
     )
     (i64.const 0)
    )
    (i64.store
     (get_local $6)
     (i64.const 0)
    )
    (i64.store
     (i32.add
      (get_local $6)
      (i32.const 24)
     )
     (i64.const 0)
    )
    (i32.store
     (get_local $4)
     (i32.const 0)
    )
    (i32.store
     (get_local $5)
     (i32.const 0)
    )
    (set_local $6
     (i32.add
      (get_local $6)
      (i32.const 32)
     )
    )
    (br_if $label$8
     (tee_local $1
      (i32.add
       (get_local $1)
       (i32.const -1)
      )
     )
    )
   )
   (block $label$9
    (block $label$10
     (br_if $label$10
      (i32.eq
       (tee_local $1
        (i32.load
         (i32.add
          (get_local $0)
          (i32.const 4)
         )
        )
       )
       (tee_local $4
        (i32.load
         (get_local $0)
        )
       )
      )
     )
     (set_local $3
      (i32.sub
       (i32.const 0)
       (get_local $4)
      )
     )
     (set_local $1
      (i32.add
       (get_local $1)
       (i32.const -16)
      )
     )
     (loop $label$11
      (i64.store
       (i32.add
        (get_local $7)
        (i32.const -32)
       )
       (i64.load
        (i32.add
         (get_local $1)
         (i32.const -16)
        )
       )
      )
      (i64.store align=4
       (i32.add
        (get_local $7)
        (i32.const -20)
       )
       (i64.load align=4
        (tee_local $4
         (i32.add
          (get_local $1)
          (i32.const -4)
         )
        )
       )
      )
      (i32.store
       (i32.add
        (get_local $7)
        (i32.const -24)
       )
       (i32.load
        (tee_local $5
         (i32.add
          (get_local $1)
          (i32.const -8)
         )
        )
       )
      )
      (i32.store
       (get_local $5)
       (i32.const 0)
      )
      (i32.store
       (get_local $4)
       (i32.const 0)
      )
      (i32.store
       (get_local $1)
       (i32.const 0)
      )
      (i64.store
       (i32.add
        (get_local $7)
        (i32.const -8)
       )
       (i64.load
        (i32.add
         (get_local $1)
         (i32.const 8)
        )
       )
      )
      (set_local $7
       (i32.add
        (get_local $7)
        (i32.const -32)
       )
      )
      (br_if $label$11
       (i32.ne
        (i32.add
         (tee_local $1
          (i32.add
           (get_local $1)
           (i32.const -32)
          )
         )
         (get_local $3)
        )
        (i32.const -16)
       )
      )
     )
     (set_local $4
      (i32.load
       (i32.add
        (get_local $0)
        (i32.const 4)
       )
      )
     )
     (set_local $5
      (i32.load
       (get_local $0)
      )
     )
     (br $label$9)
    )
    (set_local $5
     (get_local $4)
    )
   )
   (i32.store
    (get_local $0)
    (get_local $7)
   )
   (i32.store
    (i32.add
     (get_local $0)
     (i32.const 4)
    )
    (get_local $6)
   )
   (i32.store
    (i32.add
     (get_local $0)
     (i32.const 8)
    )
    (get_local $2)
   )
   (block $label$12
    (br_if $label$12
     (i32.eq
      (get_local $4)
      (get_local $5)
     )
    )
    (set_local $1
     (i32.sub
      (i32.const 0)
      (get_local $5)
     )
    )
    (set_local $6
     (i32.add
      (get_local $4)
      (i32.const -24)
     )
    )
    (loop $label$13
     (block $label$14
      (br_if $label$14
       (i32.eqz
        (i32.and
         (i32.load8_u
          (get_local $6)
         )
         (i32.const 1)
        )
       )
      )
      (call $_ZdlPv
       (i32.load
        (i32.add
         (get_local $6)
         (i32.const 8)
        )
       )
      )
     )
     (br_if $label$13
      (i32.ne
       (i32.add
        (tee_local $6
         (i32.add
          (get_local $6)
          (i32.const -32)
         )
        )
        (get_local $1)
       )
       (i32.const -24)
      )
     )
    )
   )
   (br_if $label$0
    (i32.eqz
     (get_local $5)
    )
   )
   (call $_ZdlPv
    (get_local $5)
   )
  )
 )
 (func $_ZN5eosio5disco11updtgenesisEyNSt3__112basic_stringIcNS1_11char_traitsIcEENS1_9allocatorIcEEEENS1_6vectorIS7_NS5_IS7_EEEE (type $FUNCSIG$vijii) (param $0 i32) (param $1 i64) (param $2 i32) (param $3 i32)
  (local $4 i32)
  (local $5 i32)
  (local $6 i32)
  (local $7 i32)
  (local $8 i32)
  (local $9 i32)
  (i32.store offset=4
   (i32.const 0)
   (tee_local $9
    (i32.sub
     (i32.load offset=4
      (i32.const 0)
     )
     (i32.const 64)
    )
   )
  )
  (i64.store offset=16
   (get_local $9)
   (get_local $1)
  )
  (call $require_auth
   (get_local $1)
  )
  (block $label$0
   (br_if $label$0
    (i32.eq
     (tee_local $8
      (i32.load
       (i32.add
        (get_local $0)
        (i32.const 76)
       )
      )
     )
     (tee_local $4
      (i32.load
       (i32.add
        (get_local $0)
        (i32.const 72)
       )
      )
     )
    )
   )
   (set_local $7
    (i32.add
     (get_local $8)
     (i32.const -24)
    )
   )
   (set_local $5
    (i32.sub
     (i32.const 0)
     (get_local $4)
    )
   )
   (loop $label$1
    (br_if $label$0
     (i64.eq
      (i64.load
       (i32.load
        (get_local $7)
       )
      )
      (get_local $1)
     )
    )
    (set_local $8
     (get_local $7)
    )
    (set_local $7
     (tee_local $6
      (i32.add
       (get_local $7)
       (i32.const -24)
      )
     )
    )
    (br_if $label$1
     (i32.ne
      (i32.add
       (get_local $6)
       (get_local $5)
      )
      (i32.const -24)
     )
    )
   )
  )
  (set_local $7
   (i32.add
    (get_local $0)
    (i32.const 48)
   )
  )
  (block $label$2
   (block $label$3
    (block $label$4
     (block $label$5
      (block $label$6
       (br_if $label$6
        (i32.eq
         (get_local $8)
         (get_local $4)
        )
       )
       (call $eosio_assert
        (i32.eq
         (i32.load offset=32
          (tee_local $6
           (i32.load
            (i32.add
             (get_local $8)
             (i32.const -24)
            )
           )
          )
         )
         (get_local $7)
        )
        (i32.const 16)
       )
       (set_local $1
        (i64.load
         (get_local $0)
        )
       )
       (br_if $label$5
        (get_local $6)
       )
       (br $label$3)
      )
      (br_if $label$4
       (i32.le_s
        (tee_local $6
         (call $db_find_i64
          (i64.load
           (i32.add
            (get_local $0)
            (i32.const 48)
           )
          )
          (i64.load
           (i32.add
            (get_local $0)
            (i32.const 56)
           )
          )
          (i64.const 7108558431253954560)
          (get_local $1)
         )
        )
        (i32.const -1)
       )
      )
      (call $eosio_assert
       (i32.eq
        (i32.load offset=32
         (tee_local $6
          (call $_ZNK5eosio11multi_indexILy7108558431253954560ENS_5disco11genesis_rowEJEE31load_object_by_primary_iteratorEl
           (get_local $7)
           (get_local $6)
          )
         )
        )
        (get_local $7)
       )
       (i32.const 16)
      )
      (set_local $1
       (i64.load
        (get_local $0)
       )
      )
     )
     (i32.store offset=36
      (get_local $9)
      (get_local $3)
     )
     (i32.store offset=32
      (get_local $9)
      (get_local $2)
     )
     (call $eosio_assert
      (i32.const 1)
      (i32.const 144)
     )
     (call $_ZN5eosio11multi_indexILy7108558431253954560ENS_5disco11genesis_rowEJEE6modifyIZNS1_11updtgenesisEyNSt3__112basic_stringIcNS5_11char_traitsIcEENS5_9allocatorIcEEEENS5_6vectorISB_NS9_ISB_EEEEE3$_3EEvRKS2_yOT_
      (get_local $7)
      (get_local $6)
      (get_local $1)
      (i32.add
       (get_local $9)
       (i32.const 32)
      )
     )
     (br $label$2)
    )
    (set_local $1
     (i64.load
      (get_local $0)
     )
    )
   )
   (i32.store offset=4
    (get_local $9)
    (get_local $2)
   )
   (i32.store offset=8
    (get_local $9)
    (get_local $3)
   )
   (i32.store
    (get_local $9)
    (i32.add
     (get_local $9)
     (i32.const 16)
    )
   )
   (i64.store offset=56
    (get_local $9)
    (get_local $1)
   )
   (call $eosio_assert
    (i64.eq
     (i64.load
      (i32.add
       (get_local $0)
       (i32.const 48)
      )
     )
     (call $current_receiver)
    )
    (i32.const 80)
   )
   (i32.store offset=32
    (get_local $9)
    (get_local $7)
   )
   (i32.store offset=36
    (get_local $9)
    (get_local $9)
   )
   (i32.store offset=40
    (get_local $9)
    (i32.add
     (get_local $9)
     (i32.const 56)
    )
   )
   (i64.store offset=8 align=4
    (tee_local $6
     (call $_Znwj
      (i32.const 48)
     )
    )
    (i64.const 0)
   )
   (i64.store offset=16 align=4
    (get_local $6)
    (i64.const 0)
   )
   (i64.store offset=24 align=4
    (get_local $6)
    (i64.const 0)
   )
   (i32.store offset=32
    (get_local $6)
    (get_local $7)
   )
   (call $_ZZN5eosio11multi_indexILy7108558431253954560ENS_5disco11genesis_rowEJEE7emplaceIZNS1_11updtgenesisEyNSt3__112basic_stringIcNS5_11char_traitsIcEENS5_9allocatorIcEEEENS5_6vectorISB_NS9_ISB_EEEEE3$_2EENS3_14const_iteratorEyOT_ENKUlRSH_E_clINS3_4itemEEEDaSJ_
    (i32.add
     (get_local $9)
     (i32.const 32)
    )
    (get_local $6)
   )
   (i32.store offset=48
    (get_local $9)
    (get_local $6)
   )
   (i64.store offset=32
    (get_local $9)
    (tee_local $1
     (i64.load
      (get_local $6)
     )
    )
   )
   (i32.store offset=28
    (get_local $9)
    (tee_local $8
     (i32.load offset=36
      (get_local $6)
     )
    )
   )
   (block $label$7
    (block $label$8
     (br_if $label$8
      (i32.ge_u
       (tee_local $7
        (i32.load
         (tee_local $5
          (i32.add
           (get_local $0)
           (i32.const 76)
          )
         )
        )
       )
       (i32.load
        (i32.add
         (get_local $0)
         (i32.const 80)
        )
       )
      )
     )
     (i64.store offset=8
      (get_local $7)
      (get_local $1)
     )
     (i32.store offset=16
      (get_local $7)
      (get_local $8)
     )
     (i32.store offset=48
      (get_local $9)
      (i32.const 0)
     )
     (i32.store
      (get_local $7)
      (get_local $6)
     )
     (i32.store
      (get_local $5)
      (i32.add
       (get_local $7)
       (i32.const 24)
      )
     )
     (br $label$7)
    )
    (call $_ZNSt3__16vectorIN5eosio11multi_indexILy7108558431253954560ENS1_5disco11genesis_rowEJEE8item_ptrENS_9allocatorIS6_EEE24__emplace_back_slow_pathIJNS_10unique_ptrINS5_4itemENS_14default_deleteISC_EEEERyRlEEEvDpOT_
     (i32.add
      (get_local $0)
      (i32.const 72)
     )
     (i32.add
      (get_local $9)
      (i32.const 48)
     )
     (i32.add
      (get_local $9)
      (i32.const 32)
     )
     (i32.add
      (get_local $9)
      (i32.const 28)
     )
    )
   )
   (set_local $8
    (i32.load offset=48
     (get_local $9)
    )
   )
   (i32.store offset=48
    (get_local $9)
    (i32.const 0)
   )
   (br_if $label$2
    (i32.eqz
     (get_local $8)
    )
   )
   (block $label$9
    (br_if $label$9
     (i32.eqz
      (tee_local $5
       (i32.load offset=20
        (get_local $8)
       )
      )
     )
    )
    (block $label$10
     (block $label$11
      (br_if $label$11
       (i32.eq
        (tee_local $7
         (i32.load
          (tee_local $0
           (i32.add
            (get_local $8)
            (i32.const 24)
           )
          )
         )
        )
        (get_local $5)
       )
      )
      (set_local $6
       (i32.sub
        (i32.const 0)
        (get_local $5)
       )
      )
      (set_local $7
       (i32.add
        (get_local $7)
        (i32.const -12)
       )
      )
      (loop $label$12
       (block $label$13
        (br_if $label$13
         (i32.eqz
          (i32.and
           (i32.load8_u
            (get_local $7)
           )
           (i32.const 1)
          )
         )
        )
        (call $_ZdlPv
         (i32.load
          (i32.add
           (get_local $7)
           (i32.const 8)
          )
         )
        )
       )
       (br_if $label$12
        (i32.ne
         (i32.add
          (tee_local $7
           (i32.add
            (get_local $7)
            (i32.const -12)
           )
          )
          (get_local $6)
         )
         (i32.const -12)
        )
       )
      )
      (set_local $7
       (i32.load
        (i32.add
         (get_local $8)
         (i32.const 20)
        )
       )
      )
      (br $label$10)
     )
     (set_local $7
      (get_local $5)
     )
    )
    (i32.store
     (get_local $0)
     (get_local $5)
    )
    (call $_ZdlPv
     (get_local $7)
    )
   )
   (block $label$14
    (br_if $label$14
     (i32.eqz
      (i32.and
       (i32.load8_u offset=8
        (get_local $8)
       )
       (i32.const 1)
      )
     )
    )
    (call $_ZdlPv
     (i32.load
      (i32.add
       (get_local $8)
       (i32.const 16)
      )
     )
    )
   )
   (call $_ZdlPv
    (get_local $8)
   )
  )
  (i32.store offset=4
   (i32.const 0)
   (i32.add
    (get_local $9)
    (i32.const 64)
   )
  )
 )
 (func $_ZNK5eosio11multi_indexILy7108558431253954560ENS_5disco11genesis_rowEJEE31load_object_by_primary_iteratorEl (param $0 i32) (param $1 i32) (result i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i64)
  (local $6 i32)
  (local $7 i32)
  (local $8 i32)
  (local $9 i32)
  (set_local $8
   (tee_local $9
    (i32.sub
     (i32.load offset=4
      (i32.const 0)
     )
     (i32.const 48)
    )
   )
  )
  (i32.store offset=4
   (i32.const 0)
   (get_local $9)
  )
  (block $label$0
   (br_if $label$0
    (i32.eq
     (tee_local $7
      (i32.load
       (i32.add
        (get_local $0)
        (i32.const 28)
       )
      )
     )
     (tee_local $2
      (i32.load offset=24
       (get_local $0)
      )
     )
    )
   )
   (set_local $3
    (i32.sub
     (i32.const 0)
     (get_local $2)
    )
   )
   (set_local $6
    (i32.add
     (get_local $7)
     (i32.const -24)
    )
   )
   (loop $label$1
    (br_if $label$0
     (i32.eq
      (i32.load
       (i32.add
        (get_local $6)
        (i32.const 16)
       )
      )
      (get_local $1)
     )
    )
    (set_local $7
     (get_local $6)
    )
    (set_local $6
     (tee_local $4
      (i32.add
       (get_local $6)
       (i32.const -24)
      )
     )
    )
    (br_if $label$1
     (i32.ne
      (i32.add
       (get_local $4)
       (get_local $3)
      )
      (i32.const -24)
     )
    )
   )
  )
  (block $label$2
   (block $label$3
    (br_if $label$3
     (i32.eq
      (get_local $7)
      (get_local $2)
     )
    )
    (set_local $4
     (i32.load
      (i32.add
       (get_local $7)
       (i32.const -24)
      )
     )
    )
    (br $label$2)
   )
   (call $eosio_assert
    (i32.xor
     (i32.shr_u
      (tee_local $4
       (call $db_get_i64
        (get_local $1)
        (i32.const 0)
        (i32.const 0)
       )
      )
      (i32.const 31)
     )
     (i32.const 1)
    )
    (i32.const 384)
   )
   (block $label$4
    (block $label$5
     (br_if $label$5
      (i32.lt_u
       (get_local $4)
       (i32.const 513)
      )
     )
     (set_local $6
      (call $malloc
       (get_local $4)
      )
     )
     (br $label$4)
    )
    (i32.store offset=4
     (i32.const 0)
     (tee_local $6
      (i32.sub
       (get_local $9)
       (i32.and
        (i32.add
         (get_local $4)
         (i32.const 15)
        )
        (i32.const -16)
       )
      )
     )
    )
   )
   (drop
    (call $db_get_i64
     (get_local $1)
     (get_local $6)
     (get_local $4)
    )
   )
   (i32.store offset=36
    (get_local $8)
    (get_local $6)
   )
   (i32.store offset=32
    (get_local $8)
    (get_local $6)
   )
   (i32.store offset=40
    (get_local $8)
    (tee_local $7
     (i32.add
      (get_local $6)
      (get_local $4)
     )
    )
   )
   (block $label$6
    (br_if $label$6
     (i32.le_u
      (get_local $4)
      (i32.const 512)
     )
    )
    (call $free
     (get_local $6)
    )
    (set_local $7
     (i32.load
      (i32.add
       (get_local $8)
       (i32.const 40)
      )
     )
    )
    (set_local $6
     (i32.load offset=36
      (get_local $8)
     )
    )
   )
   (i64.store offset=8 align=4
    (tee_local $4
     (call $_Znwj
      (i32.const 48)
     )
    )
    (i64.const 0)
   )
   (i64.store offset=16 align=4
    (get_local $4)
    (i64.const 0)
   )
   (i64.store offset=24 align=4
    (get_local $4)
    (i64.const 0)
   )
   (i32.store offset=32
    (get_local $4)
    (get_local $0)
   )
   (call $eosio_assert
    (i32.gt_u
     (i32.sub
      (get_local $7)
      (get_local $6)
     )
     (i32.const 7)
    )
    (i32.const 416)
   )
   (drop
    (call $memcpy
     (get_local $4)
     (get_local $6)
     (i32.const 8)
    )
   )
   (i32.store offset=36
    (get_local $8)
    (i32.add
     (get_local $6)
     (i32.const 8)
    )
   )
   (drop
    (call $_ZN5eosiorsINS_10datastreamIPKcEENSt3__112basic_stringIcNS5_11char_traitsIcEENS5_9allocatorIcEEEEEERT_SD_RNS5_6vectorIT0_NS9_ISF_EEEE
     (call $_ZN5eosiorsINS_10datastreamIPKcEEEERT_S6_RNSt3__112basic_stringIcNS7_11char_traitsIcEENS7_9allocatorIcEEEE
      (i32.add
       (get_local $8)
       (i32.const 32)
      )
      (i32.add
       (get_local $4)
       (i32.const 8)
      )
     )
     (i32.add
      (get_local $4)
      (i32.const 20)
     )
    )
   )
   (i32.store offset=36
    (get_local $4)
    (get_local $1)
   )
   (i32.store offset=24
    (get_local $8)
    (get_local $4)
   )
   (i64.store offset=16
    (get_local $8)
    (tee_local $5
     (i64.load
      (get_local $4)
     )
    )
   )
   (i32.store offset=12
    (get_local $8)
    (tee_local $7
     (i32.load offset=36
      (get_local $4)
     )
    )
   )
   (block $label$7
    (block $label$8
     (br_if $label$8
      (i32.ge_u
       (tee_local $6
        (i32.load
         (tee_local $1
          (i32.add
           (get_local $0)
           (i32.const 28)
          )
         )
        )
       )
       (i32.load
        (i32.add
         (get_local $0)
         (i32.const 32)
        )
       )
      )
     )
     (i64.store offset=8
      (get_local $6)
      (get_local $5)
     )
     (i32.store offset=16
      (get_local $6)
      (get_local $7)
     )
     (i32.store offset=24
      (get_local $8)
      (i32.const 0)
     )
     (i32.store
      (get_local $6)
      (get_local $4)
     )
     (i32.store
      (get_local $1)
      (i32.add
       (get_local $6)
       (i32.const 24)
      )
     )
     (br $label$7)
    )
    (call $_ZNSt3__16vectorIN5eosio11multi_indexILy7108558431253954560ENS1_5disco11genesis_rowEJEE8item_ptrENS_9allocatorIS6_EEE24__emplace_back_slow_pathIJNS_10unique_ptrINS5_4itemENS_14default_deleteISC_EEEERyRlEEEvDpOT_
     (i32.add
      (get_local $0)
      (i32.const 24)
     )
     (i32.add
      (get_local $8)
      (i32.const 24)
     )
     (i32.add
      (get_local $8)
      (i32.const 16)
     )
     (i32.add
      (get_local $8)
      (i32.const 12)
     )
    )
   )
   (set_local $1
    (i32.load offset=24
     (get_local $8)
    )
   )
   (i32.store offset=24
    (get_local $8)
    (i32.const 0)
   )
   (br_if $label$2
    (i32.eqz
     (get_local $1)
    )
   )
   (block $label$9
    (br_if $label$9
     (i32.eqz
      (tee_local $3
       (i32.load offset=20
        (get_local $1)
       )
      )
     )
    )
    (block $label$10
     (block $label$11
      (br_if $label$11
       (i32.eq
        (tee_local $6
         (i32.load
          (tee_local $0
           (i32.add
            (get_local $1)
            (i32.const 24)
           )
          )
         )
        )
        (get_local $3)
       )
      )
      (set_local $7
       (i32.sub
        (i32.const 0)
        (get_local $3)
       )
      )
      (set_local $6
       (i32.add
        (get_local $6)
        (i32.const -12)
       )
      )
      (loop $label$12
       (block $label$13
        (br_if $label$13
         (i32.eqz
          (i32.and
           (i32.load8_u
            (get_local $6)
           )
           (i32.const 1)
          )
         )
        )
        (call $_ZdlPv
         (i32.load
          (i32.add
           (get_local $6)
           (i32.const 8)
          )
         )
        )
       )
       (br_if $label$12
        (i32.ne
         (i32.add
          (tee_local $6
           (i32.add
            (get_local $6)
            (i32.const -12)
           )
          )
          (get_local $7)
         )
         (i32.const -12)
        )
       )
      )
      (set_local $6
       (i32.load
        (i32.add
         (get_local $1)
         (i32.const 20)
        )
       )
      )
      (br $label$10)
     )
     (set_local $6
      (get_local $3)
     )
    )
    (i32.store
     (get_local $0)
     (get_local $3)
    )
    (call $_ZdlPv
     (get_local $6)
    )
   )
   (block $label$14
    (br_if $label$14
     (i32.eqz
      (i32.and
       (i32.load8_u offset=8
        (get_local $1)
       )
       (i32.const 1)
      )
     )
    )
    (call $_ZdlPv
     (i32.load
      (i32.add
       (get_local $1)
       (i32.const 16)
      )
     )
    )
   )
   (call $_ZdlPv
    (get_local $1)
   )
  )
  (i32.store offset=4
   (i32.const 0)
   (i32.add
    (get_local $8)
    (i32.const 48)
   )
  )
  (get_local $4)
 )
 (func $_ZZN5eosio11multi_indexILy7108558431253954560ENS_5disco11genesis_rowEJEE7emplaceIZNS1_11updtgenesisEyNSt3__112basic_stringIcNS5_11char_traitsIcEENS5_9allocatorIcEEEENS5_6vectorISB_NS9_ISB_EEEEE3$_2EENS3_14const_iteratorEyOT_ENKUlRSH_E_clINS3_4itemEEEDaSJ_ (param $0 i32) (param $1 i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i64)
  (local $5 i32)
  (local $6 i32)
  (local $7 i32)
  (set_local $6
   (tee_local $7
    (i32.sub
     (i32.load offset=4
      (i32.const 0)
     )
     (i32.const 16)
    )
   )
  )
  (i32.store offset=4
   (i32.const 0)
   (get_local $7)
  )
  (i64.store
   (get_local $1)
   (i64.load
    (i32.load
     (tee_local $5
      (i32.load offset=4
       (get_local $0)
      )
     )
    )
   )
  )
  (set_local $2
   (i32.load
    (get_local $0)
   )
  )
  (drop
   (call $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEaSERKS5_
    (i32.add
     (get_local $1)
     (i32.const 8)
    )
    (i32.load offset=4
     (get_local $5)
    )
   )
  )
  (block $label$0
   (br_if $label$0
    (i32.eq
     (tee_local $3
      (i32.add
       (get_local $1)
       (i32.const 20)
      )
     )
     (tee_local $5
      (i32.load offset=8
       (get_local $5)
      )
     )
    )
   )
   (call $_ZNSt3__16vectorINS_12basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEENS4_IS6_EEE6assignIPS6_EENS_9enable_ifIXaasr21__is_forward_iteratorIT_EE5valuesr16is_constructibleIS6_NS_15iterator_traitsISC_E9referenceEEE5valueEvE4typeESC_SC_
    (get_local $3)
    (i32.load
     (get_local $5)
    )
    (i32.load offset=4
     (get_local $5)
    )
   )
  )
  (i32.store
   (get_local $6)
   (i32.const 0)
  )
  (drop
   (call $_ZN5eosiolsINS_10datastreamIjEEEERT_S4_RKNS_5disco11genesis_rowE
    (get_local $6)
    (get_local $1)
   )
  )
  (block $label$1
   (block $label$2
    (br_if $label$2
     (i32.lt_u
      (tee_local $3
       (i32.load
        (get_local $6)
       )
      )
      (i32.const 513)
     )
    )
    (set_local $5
     (call $malloc
      (get_local $3)
     )
    )
    (br $label$1)
   )
   (i32.store offset=4
    (i32.const 0)
    (tee_local $5
     (i32.sub
      (get_local $7)
      (i32.and
       (i32.add
        (get_local $3)
        (i32.const 15)
       )
       (i32.const -16)
      )
     )
    )
   )
  )
  (i32.store offset=4
   (get_local $6)
   (get_local $5)
  )
  (i32.store
   (get_local $6)
   (get_local $5)
  )
  (i32.store offset=8
   (get_local $6)
   (i32.add
    (get_local $5)
    (get_local $3)
   )
  )
  (drop
   (call $_ZN5eosiolsINS_10datastreamIPcEEEERT_S5_RKNS_5disco11genesis_rowE
    (get_local $6)
    (get_local $1)
   )
  )
  (i32.store offset=36
   (get_local $1)
   (call $db_store_i64
    (i64.load offset=8
     (get_local $2)
    )
    (i64.const 7108558431253954560)
    (i64.load
     (i32.load offset=8
      (get_local $0)
     )
    )
    (tee_local $4
     (i64.load
      (get_local $1)
     )
    )
    (get_local $5)
    (get_local $3)
   )
  )
  (block $label$3
   (br_if $label$3
    (i32.lt_u
     (get_local $3)
     (i32.const 513)
    )
   )
   (call $free
    (get_local $5)
   )
  )
  (block $label$4
   (br_if $label$4
    (i64.lt_u
     (get_local $4)
     (i64.load offset=16
      (get_local $2)
     )
    )
   )
   (i64.store
    (i32.add
     (get_local $2)
     (i32.const 16)
    )
    (select
     (i64.const -2)
     (i64.add
      (get_local $4)
      (i64.const 1)
     )
     (i64.gt_u
      (get_local $4)
      (i64.const -3)
     )
    )
   )
  )
  (i32.store offset=4
   (i32.const 0)
   (i32.add
    (get_local $6)
    (i32.const 16)
   )
  )
 )
 (func $_ZNSt3__16vectorIN5eosio11multi_indexILy7108558431253954560ENS1_5disco11genesis_rowEJEE8item_ptrENS_9allocatorIS6_EEE24__emplace_back_slow_pathIJNS_10unique_ptrINS5_4itemENS_14default_deleteISC_EEEERyRlEEEvDpOT_ (param $0 i32) (param $1 i32) (param $2 i32) (param $3 i32)
  (local $4 i32)
  (local $5 i32)
  (local $6 i64)
  (local $7 i32)
  (local $8 i32)
  (local $9 i32)
  (local $10 i32)
  (i32.store offset=4
   (i32.const 0)
   (tee_local $10
    (i32.sub
     (i32.load offset=4
      (i32.const 0)
     )
     (i32.const 32)
    )
   )
  )
  (block $label$0
   (block $label$1
    (br_if $label$1
     (i32.ge_u
      (tee_local $8
       (i32.add
        (tee_local $4
         (i32.div_s
          (i32.sub
           (i32.load offset=4
            (get_local $0)
           )
           (tee_local $9
            (i32.load
             (get_local $0)
            )
           )
          )
          (i32.const 24)
         )
        )
        (i32.const 1)
       )
      )
      (i32.const 178956971)
     )
    )
    (set_local $5
     (i32.add
      (get_local $0)
      (i32.const 8)
     )
    )
    (block $label$2
     (block $label$3
      (block $label$4
       (br_if $label$4
        (i32.ge_u
         (tee_local $9
          (i32.div_s
           (i32.sub
            (i32.load offset=8
             (get_local $0)
            )
            (get_local $9)
           )
           (i32.const 24)
          )
         )
         (i32.const 89478485)
        )
       )
       (i32.store
        (i32.add
         (get_local $10)
         (i32.const 24)
        )
        (get_local $5)
       )
       (set_local $5
        (i32.const 0)
       )
       (i32.store offset=20
        (get_local $10)
        (i32.const 0)
       )
       (set_local $7
        (i32.add
         (get_local $10)
         (i32.const 20)
        )
       )
       (br_if $label$2
        (i32.eqz
         (tee_local $9
          (select
           (get_local $8)
           (tee_local $9
            (i32.shl
             (get_local $9)
             (i32.const 1)
            )
           )
           (i32.lt_u
            (get_local $9)
            (get_local $8)
           )
          )
         )
        )
       )
       (set_local $5
        (get_local $9)
       )
       (br $label$3)
      )
      (i32.store
       (i32.add
        (get_local $10)
        (i32.const 24)
       )
       (get_local $5)
      )
      (i32.store offset=20
       (get_local $10)
       (i32.const 0)
      )
      (set_local $7
       (i32.add
        (get_local $10)
        (i32.const 20)
       )
      )
      (set_local $5
       (i32.const 178956970)
      )
     )
     (set_local $8
      (call $_Znwj
       (i32.mul
        (get_local $5)
        (i32.const 24)
       )
      )
     )
     (br $label$0)
    )
    (set_local $8
     (i32.const 0)
    )
    (br $label$0)
   )
   (call $_ZNKSt3__120__vector_base_commonILb1EE20__throw_length_errorEv
    (get_local $0)
   )
   (unreachable)
  )
  (i32.store offset=8
   (get_local $10)
   (get_local $8)
  )
  (i32.store offset=12
   (get_local $10)
   (tee_local $9
    (i32.add
     (get_local $8)
     (i32.mul
      (get_local $4)
      (i32.const 24)
     )
    )
   )
  )
  (i32.store
   (get_local $7)
   (tee_local $5
    (i32.add
     (get_local $8)
     (i32.mul
      (get_local $5)
      (i32.const 24)
     )
    )
   )
  )
  (set_local $8
   (i32.load
    (get_local $1)
   )
  )
  (i32.store
   (get_local $1)
   (i32.const 0)
  )
  (set_local $1
   (i32.load
    (get_local $3)
   )
  )
  (set_local $6
   (i64.load
    (get_local $2)
   )
  )
  (i32.store
   (get_local $9)
   (get_local $8)
  )
  (i64.store offset=8
   (get_local $9)
   (get_local $6)
  )
  (i32.store offset=16
   (get_local $9)
   (get_local $1)
  )
  (i32.store offset=16
   (get_local $10)
   (tee_local $1
    (i32.add
     (get_local $9)
     (i32.const 24)
    )
   )
  )
  (block $label$5
   (br_if $label$5
    (i32.eq
     (tee_local $8
      (i32.load
       (i32.add
        (get_local $0)
        (i32.const 4)
       )
      )
     )
     (tee_local $2
      (i32.load
       (get_local $0)
      )
     )
    )
   )
   (loop $label$6
    (set_local $1
     (i32.load
      (tee_local $5
       (i32.add
        (get_local $8)
        (i32.const -24)
       )
      )
     )
    )
    (i32.store
     (get_local $5)
     (i32.const 0)
    )
    (i32.store
     (i32.add
      (get_local $9)
      (i32.const -24)
     )
     (get_local $1)
    )
    (i32.store
     (i32.add
      (get_local $9)
      (i32.const -8)
     )
     (i32.load
      (i32.add
       (get_local $8)
       (i32.const -8)
      )
     )
    )
    (i32.store
     (i32.add
      (get_local $9)
      (i32.const -12)
     )
     (i32.load
      (i32.add
       (get_local $8)
       (i32.const -12)
      )
     )
    )
    (i32.store
     (i32.add
      (get_local $9)
      (i32.const -16)
     )
     (i32.load
      (i32.add
       (get_local $8)
       (i32.const -16)
      )
     )
    )
    (i32.store offset=12
     (get_local $10)
     (tee_local $9
      (i32.add
       (i32.load offset=12
        (get_local $10)
       )
       (i32.const -24)
      )
     )
    )
    (set_local $8
     (get_local $5)
    )
    (br_if $label$6
     (i32.ne
      (get_local $2)
      (get_local $5)
     )
    )
   )
   (set_local $8
    (i32.load
     (i32.add
      (get_local $0)
      (i32.const 4)
     )
    )
   )
   (set_local $5
    (i32.load
     (get_local $7)
    )
   )
   (set_local $2
    (i32.load
     (get_local $0)
    )
   )
   (set_local $1
    (i32.load
     (i32.add
      (get_local $10)
      (i32.const 16)
     )
    )
   )
  )
  (i32.store
   (get_local $0)
   (get_local $9)
  )
  (i32.store
   (i32.add
    (get_local $0)
    (i32.const 4)
   )
   (get_local $1)
  )
  (i32.store
   (i32.add
    (i32.add
     (get_local $10)
     (i32.const 8)
    )
    (i32.const 8)
   )
   (get_local $8)
  )
  (set_local $8
   (i32.load
    (tee_local $9
     (i32.add
      (get_local $0)
      (i32.const 8)
     )
    )
   )
  )
  (i32.store
   (get_local $9)
   (get_local $5)
  )
  (i32.store offset=12
   (get_local $10)
   (get_local $2)
  )
  (i32.store
   (get_local $7)
   (get_local $8)
  )
  (i32.store offset=8
   (get_local $10)
   (get_local $2)
  )
  (drop
   (call $_ZNSt3__114__split_bufferIN5eosio11multi_indexILy7108558431253954560ENS1_5disco11genesis_rowEJEE8item_ptrERNS_9allocatorIS6_EEED2Ev
    (i32.add
     (get_local $10)
     (i32.const 8)
    )
   )
  )
  (i32.store offset=4
   (i32.const 0)
   (i32.add
    (get_local $10)
    (i32.const 32)
   )
  )
 )
 (func $_ZN5eosio11multi_indexILy7108558431253954560ENS_5disco11genesis_rowEJEE6modifyIZNS1_11updtgenesisEyNSt3__112basic_stringIcNS5_11char_traitsIcEENS5_9allocatorIcEEEENS5_6vectorISB_NS9_ISB_EEEEE3$_3EEvRKS2_yOT_ (param $0 i32) (param $1 i32) (param $2 i64) (param $3 i32)
  (local $4 i64)
  (local $5 i32)
  (local $6 i32)
  (local $7 i32)
  (set_local $6
   (tee_local $7
    (i32.sub
     (i32.load offset=4
      (i32.const 0)
     )
     (i32.const 16)
    )
   )
  )
  (i32.store offset=4
   (i32.const 0)
   (get_local $7)
  )
  (call $eosio_assert
   (i32.eq
    (i32.load offset=32
     (get_local $1)
    )
    (get_local $0)
   )
   (i32.const 192)
  )
  (call $eosio_assert
   (i64.eq
    (i64.load
     (get_local $0)
    )
    (call $current_receiver)
   )
   (i32.const 240)
  )
  (set_local $4
   (i64.load
    (get_local $1)
   )
  )
  (drop
   (call $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEaSERKS5_
    (i32.add
     (get_local $1)
     (i32.const 8)
    )
    (i32.load
     (get_local $3)
    )
   )
  )
  (block $label$0
   (br_if $label$0
    (i32.eq
     (tee_local $5
      (i32.add
       (get_local $1)
       (i32.const 20)
      )
     )
     (tee_local $3
      (i32.load offset=4
       (get_local $3)
      )
     )
    )
   )
   (call $_ZNSt3__16vectorINS_12basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEENS4_IS6_EEE6assignIPS6_EENS_9enable_ifIXaasr21__is_forward_iteratorIT_EE5valuesr16is_constructibleIS6_NS_15iterator_traitsISC_E9referenceEEE5valueEvE4typeESC_SC_
    (get_local $5)
    (i32.load
     (get_local $3)
    )
    (i32.load offset=4
     (get_local $3)
    )
   )
  )
  (call $eosio_assert
   (i64.eq
    (get_local $4)
    (i64.load
     (get_local $1)
    )
   )
   (i32.const 304)
  )
  (i32.store
   (get_local $6)
   (i32.const 0)
  )
  (drop
   (call $_ZN5eosiolsINS_10datastreamIjEEEERT_S4_RKNS_5disco11genesis_rowE
    (get_local $6)
    (get_local $1)
   )
  )
  (block $label$1
   (block $label$2
    (br_if $label$2
     (i32.lt_u
      (tee_local $5
       (i32.load
        (get_local $6)
       )
      )
      (i32.const 513)
     )
    )
    (set_local $3
     (call $malloc
      (get_local $5)
     )
    )
    (br $label$1)
   )
   (i32.store offset=4
    (i32.const 0)
    (tee_local $3
     (i32.sub
      (get_local $7)
      (i32.and
       (i32.add
        (get_local $5)
        (i32.const 15)
       )
       (i32.const -16)
      )
     )
    )
   )
  )
  (i32.store offset=4
   (get_local $6)
   (get_local $3)
  )
  (i32.store
   (get_local $6)
   (get_local $3)
  )
  (i32.store offset=8
   (get_local $6)
   (i32.add
    (get_local $3)
    (get_local $5)
   )
  )
  (drop
   (call $_ZN5eosiolsINS_10datastreamIPcEEEERT_S5_RKNS_5disco11genesis_rowE
    (get_local $6)
    (get_local $1)
   )
  )
  (call $db_update_i64
   (i32.load offset=36
    (get_local $1)
   )
   (get_local $2)
   (get_local $3)
   (get_local $5)
  )
  (block $label$3
   (br_if $label$3
    (i32.lt_u
     (get_local $5)
     (i32.const 513)
    )
   )
   (call $free
    (get_local $3)
   )
  )
  (block $label$4
   (br_if $label$4
    (i64.lt_u
     (get_local $4)
     (i64.load offset=16
      (get_local $0)
     )
    )
   )
   (i64.store
    (i32.add
     (get_local $0)
     (i32.const 16)
    )
    (select
     (i64.const -2)
     (i64.add
      (get_local $4)
      (i64.const 1)
     )
     (i64.gt_u
      (get_local $4)
      (i64.const -3)
     )
    )
   )
  )
  (i32.store offset=4
   (i32.const 0)
   (i32.add
    (get_local $6)
    (i32.const 16)
   )
  )
 )
 (func $_ZN5eosiolsINS_10datastreamIjEEEERT_S4_RKNS_5disco11genesis_rowE (param $0 i32) (param $1 i32) (result i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i64)
  (i32.store
   (get_local $0)
   (tee_local $4
    (i32.add
     (i32.load
      (get_local $0)
     )
     (i32.const 8)
    )
   )
  )
  (set_local $5
   (i64.extend_u/i32
    (select
     (i32.load
      (i32.add
       (get_local $1)
       (i32.const 12)
      )
     )
     (i32.shr_u
      (tee_local $3
       (i32.load8_u offset=8
        (get_local $1)
       )
      )
      (i32.const 1)
     )
     (i32.and
      (get_local $3)
      (i32.const 1)
     )
    )
   )
  )
  (loop $label$0
   (set_local $4
    (i32.add
     (get_local $4)
     (i32.const 1)
    )
   )
   (br_if $label$0
    (i64.ne
     (tee_local $5
      (i64.shr_u
       (get_local $5)
       (i64.const 7)
      )
     )
     (i64.const 0)
    )
   )
  )
  (i32.store
   (get_local $0)
   (get_local $4)
  )
  (block $label$1
   (br_if $label$1
    (i32.eqz
     (tee_local $3
      (select
       (i32.load
        (i32.add
         (get_local $1)
         (i32.const 12)
        )
       )
       (i32.shr_u
        (tee_local $3
         (i32.load8_u
          (i32.add
           (get_local $1)
           (i32.const 8)
          )
         )
        )
        (i32.const 1)
       )
       (i32.and
        (get_local $3)
        (i32.const 1)
       )
      )
     )
    )
   )
   (i32.store
    (get_local $0)
    (tee_local $4
     (i32.add
      (get_local $3)
      (get_local $4)
     )
    )
   )
  )
  (set_local $5
   (i64.extend_u/i32
    (i32.div_s
     (i32.sub
      (tee_local $2
       (i32.load
        (i32.add
         (get_local $1)
         (i32.const 24)
        )
       )
      )
      (tee_local $1
       (i32.load offset=20
        (get_local $1)
       )
      )
     )
     (i32.const 12)
    )
   )
  )
  (loop $label$2
   (set_local $4
    (i32.add
     (get_local $4)
     (i32.const 1)
    )
   )
   (br_if $label$2
    (i64.ne
     (tee_local $5
      (i64.shr_u
       (get_local $5)
       (i64.const 7)
      )
     )
     (i64.const 0)
    )
   )
  )
  (i32.store
   (get_local $0)
   (get_local $4)
  )
  (block $label$3
   (br_if $label$3
    (i32.eq
     (get_local $1)
     (get_local $2)
    )
   )
   (loop $label$4
    (set_local $5
     (i64.extend_u/i32
      (select
       (i32.load offset=4
        (get_local $1)
       )
       (i32.shr_u
        (tee_local $3
         (i32.load8_u
          (get_local $1)
         )
        )
        (i32.const 1)
       )
       (i32.and
        (get_local $3)
        (i32.const 1)
       )
      )
     )
    )
    (loop $label$5
     (set_local $4
      (i32.add
       (get_local $4)
       (i32.const 1)
      )
     )
     (br_if $label$5
      (i64.ne
       (tee_local $5
        (i64.shr_u
         (get_local $5)
         (i64.const 7)
        )
       )
       (i64.const 0)
      )
     )
    )
    (i32.store
     (get_local $0)
     (get_local $4)
    )
    (block $label$6
     (br_if $label$6
      (i32.eqz
       (tee_local $3
        (select
         (i32.load
          (i32.add
           (get_local $1)
           (i32.const 4)
          )
         )
         (i32.shr_u
          (tee_local $3
           (i32.load8_u
            (get_local $1)
           )
          )
          (i32.const 1)
         )
         (i32.and
          (get_local $3)
          (i32.const 1)
         )
        )
       )
      )
     )
     (i32.store
      (get_local $0)
      (tee_local $4
       (i32.add
        (get_local $3)
        (get_local $4)
       )
      )
     )
    )
    (br_if $label$4
     (i32.ne
      (tee_local $1
       (i32.add
        (get_local $1)
        (i32.const 12)
       )
      )
      (get_local $2)
     )
    )
   )
  )
  (get_local $0)
 )
 (func $_ZN5eosiolsINS_10datastreamIPcEEEERT_S5_RKNS_5disco11genesis_rowE (param $0 i32) (param $1 i32) (result i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i32)
  (local $6 i64)
  (local $7 i32)
  (i32.store offset=4
   (i32.const 0)
   (tee_local $7
    (i32.sub
     (i32.load offset=4
      (i32.const 0)
     )
     (i32.const 16)
    )
   )
  )
  (call $eosio_assert
   (i32.gt_s
    (i32.sub
     (i32.load offset=8
      (get_local $0)
     )
     (i32.load offset=4
      (get_local $0)
     )
    )
    (i32.const 7)
   )
   (i32.const 368)
  )
  (drop
   (call $memcpy
    (i32.load offset=4
     (get_local $0)
    )
    (get_local $1)
    (i32.const 8)
   )
  )
  (i32.store offset=4
   (get_local $0)
   (i32.add
    (i32.load offset=4
     (get_local $0)
    )
    (i32.const 8)
   )
  )
  (set_local $5
   (i32.load offset=4
    (tee_local $2
     (call $_ZN5eosiolsINS_10datastreamIPcEEEERT_S5_RKNSt3__112basic_stringIcNS6_11char_traitsIcEENS6_9allocatorIcEEEE
      (get_local $0)
      (i32.add
       (get_local $1)
       (i32.const 8)
      )
     )
    )
   )
  )
  (set_local $6
   (i64.extend_u/i32
    (i32.div_s
     (i32.sub
      (i32.load
       (i32.add
        (get_local $1)
        (i32.const 24)
       )
      )
      (i32.load offset=20
       (get_local $1)
      )
     )
     (i32.const 12)
    )
   )
  )
  (set_local $0
   (i32.add
    (get_local $2)
    (i32.const 4)
   )
  )
  (loop $label$0
   (set_local $3
    (i32.wrap/i64
     (get_local $6)
    )
   )
   (i32.store8 offset=15
    (get_local $7)
    (i32.or
     (i32.shl
      (tee_local $4
       (i64.ne
        (tee_local $6
         (i64.shr_u
          (get_local $6)
          (i64.const 7)
         )
        )
        (i64.const 0)
       )
      )
      (i32.const 7)
     )
     (i32.and
      (get_local $3)
      (i32.const 127)
     )
    )
   )
   (call $eosio_assert
    (i32.gt_s
     (i32.sub
      (i32.load
       (i32.add
        (get_local $2)
        (i32.const 8)
       )
      )
      (get_local $5)
     )
     (i32.const 0)
    )
    (i32.const 368)
   )
   (drop
    (call $memcpy
     (i32.load
      (get_local $0)
     )
     (i32.add
      (get_local $7)
      (i32.const 15)
     )
     (i32.const 1)
    )
   )
   (i32.store
    (get_local $0)
    (tee_local $5
     (i32.add
      (i32.load
       (get_local $0)
      )
      (i32.const 1)
     )
    )
   )
   (br_if $label$0
    (get_local $4)
   )
  )
  (block $label$1
   (br_if $label$1
    (i32.eq
     (tee_local $0
      (i32.load
       (i32.add
        (get_local $1)
        (i32.const 20)
       )
      )
     )
     (tee_local $5
      (i32.load
       (i32.add
        (get_local $1)
        (i32.const 24)
       )
      )
     )
    )
   )
   (loop $label$2
    (drop
     (call $_ZN5eosiolsINS_10datastreamIPcEEEERT_S5_RKNSt3__112basic_stringIcNS6_11char_traitsIcEENS6_9allocatorIcEEEE
      (get_local $2)
      (get_local $0)
     )
    )
    (br_if $label$2
     (i32.ne
      (get_local $5)
      (tee_local $0
       (i32.add
        (get_local $0)
        (i32.const 12)
       )
      )
     )
    )
   )
  )
  (i32.store offset=4
   (i32.const 0)
   (i32.add
    (get_local $7)
    (i32.const 16)
   )
  )
  (get_local $2)
 )
 (func $_ZNSt3__114__split_bufferIN5eosio11multi_indexILy7108558431253954560ENS1_5disco11genesis_rowEJEE8item_ptrERNS_9allocatorIS6_EEED2Ev (param $0 i32) (result i32)
  (local $1 i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i32)
  (local $6 i32)
  (local $7 i32)
  (block $label$0
   (br_if $label$0
    (i32.eq
     (tee_local $7
      (i32.load offset=8
       (get_local $0)
      )
     )
     (tee_local $1
      (i32.load offset=4
       (get_local $0)
      )
     )
    )
   )
   (set_local $5
    (i32.add
     (get_local $0)
     (i32.const 8)
    )
   )
   (loop $label$1
    (i32.store
     (get_local $5)
     (tee_local $7
      (i32.add
       (get_local $7)
       (i32.const -24)
      )
     )
    )
    (set_local $2
     (i32.load
      (get_local $7)
     )
    )
    (i32.store
     (get_local $7)
     (i32.const 0)
    )
    (block $label$2
     (br_if $label$2
      (i32.eqz
       (get_local $2)
      )
     )
     (block $label$3
      (br_if $label$3
       (i32.eqz
        (tee_local $3
         (i32.load offset=20
          (get_local $2)
         )
        )
       )
      )
      (block $label$4
       (block $label$5
        (br_if $label$5
         (i32.eq
          (tee_local $7
           (i32.load
            (tee_local $6
             (i32.add
              (get_local $2)
              (i32.const 24)
             )
            )
           )
          )
          (get_local $3)
         )
        )
        (set_local $4
         (i32.sub
          (i32.const 0)
          (get_local $3)
         )
        )
        (set_local $7
         (i32.add
          (get_local $7)
          (i32.const -12)
         )
        )
        (loop $label$6
         (block $label$7
          (br_if $label$7
           (i32.eqz
            (i32.and
             (i32.load8_u
              (get_local $7)
             )
             (i32.const 1)
            )
           )
          )
          (call $_ZdlPv
           (i32.load
            (i32.add
             (get_local $7)
             (i32.const 8)
            )
           )
          )
         )
         (br_if $label$6
          (i32.ne
           (i32.add
            (tee_local $7
             (i32.add
              (get_local $7)
              (i32.const -12)
             )
            )
            (get_local $4)
           )
           (i32.const -12)
          )
         )
        )
        (set_local $7
         (i32.load
          (i32.add
           (get_local $2)
           (i32.const 20)
          )
         )
        )
        (br $label$4)
       )
       (set_local $7
        (get_local $3)
       )
      )
      (i32.store
       (get_local $6)
       (get_local $3)
      )
      (call $_ZdlPv
       (get_local $7)
      )
     )
     (block $label$8
      (br_if $label$8
       (i32.eqz
        (i32.and
         (i32.load8_u offset=8
          (get_local $2)
         )
         (i32.const 1)
        )
       )
      )
      (call $_ZdlPv
       (i32.load
        (i32.add
         (get_local $2)
         (i32.const 16)
        )
       )
      )
     )
     (call $_ZdlPv
      (get_local $2)
     )
    )
    (br_if $label$1
     (i32.ne
      (tee_local $7
       (i32.load
        (get_local $5)
       )
      )
      (get_local $1)
     )
    )
   )
  )
  (block $label$9
   (br_if $label$9
    (i32.eqz
     (tee_local $7
      (i32.load
       (get_local $0)
      )
     )
    )
   )
   (call $_ZdlPv
    (get_local $7)
   )
  )
  (get_local $0)
 )
 (func $apply (param $0 i64) (param $1 i64) (param $2 i64)
  (local $3 i32)
  (local $4 i32)
  (local $5 i32)
  (local $6 i32)
  (local $7 i32)
  (i32.store offset=4
   (i32.const 0)
   (tee_local $7
    (i32.sub
     (i32.load offset=4
      (i32.const 0)
     )
     (i32.const 128)
    )
   )
  )
  (block $label$0
   (block $label$1
    (br_if $label$1
     (i64.ne
      (get_local $1)
      (get_local $0)
     )
    )
    (i64.store
     (i32.add
      (get_local $7)
      (i32.const 56)
     )
     (get_local $1)
    )
    (i64.store
     (i32.add
      (get_local $7)
      (i32.const 64)
     )
     (i64.const -1)
    )
    (i64.store
     (i32.add
      (get_local $7)
      (i32.const 72)
     )
     (i64.const 0)
    )
    (i32.store
     (i32.add
      (get_local $7)
      (i32.const 80)
     )
     (i32.const 0)
    )
    (i64.store offset=48
     (get_local $7)
     (get_local $1)
    )
    (i64.store offset=40
     (get_local $7)
     (get_local $1)
    )
    (i64.store offset=88
     (get_local $7)
     (get_local $1)
    )
    (i64.store
     (i32.add
      (get_local $7)
      (i32.const 96)
     )
     (get_local $1)
    )
    (i64.store
     (i32.add
      (get_local $7)
      (i32.const 104)
     )
     (i64.const -1)
    )
    (i32.store
     (i32.add
      (get_local $7)
      (i32.const 112)
     )
     (i32.const 0)
    )
    (i32.store
     (i32.add
      (get_local $7)
      (i32.const 116)
     )
     (i32.const 0)
    )
    (i32.store
     (i32.add
      (get_local $7)
      (i32.const 120)
     )
     (i32.const 0)
    )
    (block $label$2
     (block $label$3
      (br_if $label$3
       (i64.eq
        (get_local $2)
        (i64.const -3074950589539418112)
       )
      )
      (br_if $label$0
       (i64.ne
        (get_local $2)
        (i64.const -3074949011639914496)
       )
      )
      (i32.store offset=36
       (get_local $7)
       (i32.const 0)
      )
      (i32.store offset=32
       (get_local $7)
       (i32.const 1)
      )
      (i64.store offset=8 align=4
       (get_local $7)
       (i64.load offset=32
        (get_local $7)
       )
      )
      (drop
       (call $_ZN5eosio14execute_actionINS_5discoES1_JyNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEENS2_6vectorIS8_NS6_IS8_EEEEEEEbPT_MT0_FvDpT1_E
        (i32.add
         (get_local $7)
         (i32.const 40)
        )
        (i32.add
         (get_local $7)
         (i32.const 8)
        )
       )
      )
      (br $label$2)
     )
     (i32.store offset=28
      (get_local $7)
      (i32.const 0)
     )
     (i32.store offset=24
      (get_local $7)
      (i32.const 2)
     )
     (i64.store offset=16 align=4
      (get_local $7)
      (i64.load offset=24
       (get_local $7)
      )
     )
     (drop
      (call $_ZN5eosio14execute_actionINS_5discoES1_JyRKNS_14discovery_fileEEEEbPT_MT0_FvDpT1_E
       (i32.add
        (get_local $7)
        (i32.const 40)
       )
       (i32.add
        (get_local $7)
        (i32.const 16)
       )
      )
     )
    )
    (drop
     (call $_ZNSt3__113__vector_baseIN5eosio11multi_indexILy7108558431253954560ENS1_5disco11genesis_rowEJEE8item_ptrENS_9allocatorIS6_EEED2Ev
      (i32.add
       (get_local $7)
       (i32.const 112)
      )
     )
    )
    (br_if $label$1
     (i32.eqz
      (tee_local $3
       (i32.load
        (i32.add
         (get_local $7)
         (i32.const 72)
        )
       )
      )
     )
    )
    (block $label$4
     (block $label$5
      (br_if $label$5
       (i32.eq
        (tee_local $6
         (i32.load
          (tee_local $5
           (i32.add
            (get_local $7)
            (i32.const 76)
           )
          )
         )
        )
        (get_local $3)
       )
      )
      (loop $label$6
       (set_local $4
        (i32.load
         (tee_local $6
          (i32.add
           (get_local $6)
           (i32.const -24)
          )
         )
        )
       )
       (i32.store
        (get_local $6)
        (i32.const 0)
       )
       (block $label$7
        (br_if $label$7
         (i32.eqz
          (get_local $4)
         )
        )
        (drop
         (call $_ZN5eosio14discovery_fileD2Ev
          (i32.add
           (get_local $4)
           (i32.const 8)
          )
         )
        )
        (call $_ZdlPv
         (get_local $4)
        )
       )
       (br_if $label$6
        (i32.ne
         (get_local $3)
         (get_local $6)
        )
       )
      )
      (set_local $6
       (i32.load
        (i32.add
         (get_local $7)
         (i32.const 72)
        )
       )
      )
      (br $label$4)
     )
     (set_local $6
      (get_local $3)
     )
    )
    (i32.store
     (get_local $5)
     (get_local $3)
    )
    (call $_ZdlPv
     (get_local $6)
    )
   )
   (i32.store offset=4
    (i32.const 0)
    (i32.add
     (get_local $7)
     (i32.const 128)
    )
   )
   (return)
  )
  (call $eosio_exit
   (i32.const 0)
  )
  (unreachable)
 )
 (func $_ZN5eosio14execute_actionINS_5discoES1_JyNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEENS2_6vectorIS8_NS6_IS8_EEEEEEEbPT_MT0_FvDpT1_E (param $0 i32) (param $1 i32) (result i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (i32.store offset=4
   (i32.const 0)
   (tee_local $2
    (i32.sub
     (i32.load offset=4
      (i32.const 0)
     )
     (i32.const 64)
    )
   )
  )
  (i32.store offset=44
   (tee_local $4
    (get_local $2)
   )
   (get_local $0)
  )
  (i32.store offset=32
   (get_local $4)
   (i32.load
    (get_local $1)
   )
  )
  (i32.store offset=36
   (get_local $4)
   (i32.load offset=4
    (get_local $1)
   )
  )
  (block $label$0
   (block $label$1
    (br_if $label$1
     (i32.lt_u
      (tee_local $0
       (call $action_data_size)
      )
      (i32.const 513)
     )
    )
    (set_local $1
     (call $malloc
      (get_local $0)
     )
    )
    (br $label$0)
   )
   (i32.store offset=4
    (i32.const 0)
    (tee_local $1
     (i32.sub
      (get_local $2)
      (i32.and
       (i32.add
        (get_local $0)
        (i32.const 15)
       )
       (i32.const -16)
      )
     )
    )
   )
  )
  (drop
   (call $read_action_data
    (get_local $1)
    (get_local $0)
   )
  )
  (i64.store
   (i32.add
    (get_local $4)
    (i32.const 16)
   )
   (i64.const 0)
  )
  (i64.store
   (i32.add
    (get_local $4)
    (i32.const 24)
   )
   (i64.const 0)
  )
  (i64.store
   (get_local $4)
   (i64.const 0)
  )
  (i64.store offset=8
   (get_local $4)
   (i64.const 0)
  )
  (i32.store offset=56
   (get_local $4)
   (i32.add
    (get_local $1)
    (get_local $0)
   )
  )
  (i32.store offset=48
   (get_local $4)
   (get_local $1)
  )
  (call $eosio_assert
   (i32.gt_u
    (get_local $0)
    (i32.const 7)
   )
   (i32.const 416)
  )
  (drop
   (call $memcpy
    (get_local $4)
    (get_local $1)
    (i32.const 8)
   )
  )
  (i32.store offset=52
   (get_local $4)
   (i32.add
    (get_local $1)
    (i32.const 8)
   )
  )
  (drop
   (call $_ZN5eosiorsINS_10datastreamIPKcEEEERT_S6_RNSt3__112basic_stringIcNS7_11char_traitsIcEENS7_9allocatorIcEEEE
    (i32.add
     (get_local $4)
     (i32.const 48)
    )
    (i32.add
     (get_local $4)
     (i32.const 8)
    )
   )
  )
  (drop
   (call $_ZN5eosiorsINS_10datastreamIPKcEENSt3__112basic_stringIcNS5_11char_traitsIcEENS5_9allocatorIcEEEEEERT_SD_RNS5_6vectorIT0_NS9_ISF_EEEE
    (i32.add
     (get_local $4)
     (i32.const 48)
    )
    (i32.add
     (get_local $4)
     (i32.const 20)
    )
   )
  )
  (block $label$2
   (br_if $label$2
    (i32.lt_u
     (get_local $0)
     (i32.const 513)
    )
   )
   (call $free
    (get_local $1)
   )
  )
  (i32.store offset=52
   (get_local $4)
   (i32.add
    (get_local $4)
    (i32.const 32)
   )
  )
  (i32.store offset=48
   (get_local $4)
   (i32.add
    (get_local $4)
    (i32.const 44)
   )
  )
  (call $_ZN5boost4mp116detail16tuple_apply_implIRZN5eosio14execute_actionINS3_5discoES5_JyNSt3__112basic_stringIcNS6_11char_traitsIcEENS6_9allocatorIcEEEENS6_6vectorISC_NSA_ISC_EEEEEEEbPT_MT0_FvDpT1_EEUlDpT_E_RNS6_5tupleIJySC_SF_EEEJLj0ELj1ELj2EEEEDTclclsr3stdE7forwardISG_Efp_Espclsr3stdE3getIXT1_EEclsr3stdE7forwardISI_Efp0_EEEEOSG_OSI_NS0_16integer_sequenceIjJXspT1_EEEE
   (i32.add
    (get_local $4)
    (i32.const 48)
   )
   (get_local $4)
  )
  (block $label$3
   (br_if $label$3
    (i32.eqz
     (tee_local $2
      (i32.load offset=20
       (get_local $4)
      )
     )
    )
   )
   (block $label$4
    (block $label$5
     (br_if $label$5
      (i32.eq
       (tee_local $1
        (i32.load
         (tee_local $3
          (i32.add
           (get_local $4)
           (i32.const 24)
          )
         )
        )
       )
       (get_local $2)
      )
     )
     (set_local $0
      (i32.sub
       (i32.const 0)
       (get_local $2)
      )
     )
     (set_local $1
      (i32.add
       (get_local $1)
       (i32.const -12)
      )
     )
     (loop $label$6
      (block $label$7
       (br_if $label$7
        (i32.eqz
         (i32.and
          (i32.load8_u
           (get_local $1)
          )
          (i32.const 1)
         )
        )
       )
       (call $_ZdlPv
        (i32.load
         (i32.add
          (get_local $1)
          (i32.const 8)
         )
        )
       )
      )
      (br_if $label$6
       (i32.ne
        (i32.add
         (tee_local $1
          (i32.add
           (get_local $1)
           (i32.const -12)
          )
         )
         (get_local $0)
        )
        (i32.const -12)
       )
      )
     )
     (set_local $1
      (i32.load
       (i32.add
        (get_local $4)
        (i32.const 20)
       )
      )
     )
     (br $label$4)
    )
    (set_local $1
     (get_local $2)
    )
   )
   (i32.store
    (get_local $3)
    (get_local $2)
   )
   (call $_ZdlPv
    (get_local $1)
   )
  )
  (block $label$8
   (br_if $label$8
    (i32.eqz
     (i32.and
      (i32.load8_u offset=8
       (get_local $4)
      )
      (i32.const 1)
     )
    )
   )
   (call $_ZdlPv
    (i32.load
     (i32.add
      (get_local $4)
      (i32.const 16)
     )
    )
   )
  )
  (i32.store offset=4
   (i32.const 0)
   (i32.add
    (get_local $4)
    (i32.const 64)
   )
  )
  (i32.const 1)
 )
 (func $_ZN5eosio14execute_actionINS_5discoES1_JyRKNS_14discovery_fileEEEEbPT_MT0_FvDpT1_E (param $0 i32) (param $1 i32) (result i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i64)
  (local $6 i32)
  (local $7 i32)
  (set_local $7
   (tee_local $2
    (i32.sub
     (i32.load offset=4
      (i32.const 0)
     )
     (i32.const 560)
    )
   )
  )
  (i32.store offset=4
   (i32.const 0)
   (get_local $2)
  )
  (set_local $3
   (i32.load offset=4
    (get_local $1)
   )
  )
  (set_local $6
   (i32.load
    (get_local $1)
   )
  )
  (block $label$0
   (block $label$1
    (br_if $label$1
     (i32.lt_u
      (tee_local $4
       (call $action_data_size)
      )
      (i32.const 513)
     )
    )
    (set_local $1
     (call $malloc
      (get_local $4)
     )
    )
    (br $label$0)
   )
   (i32.store offset=4
    (i32.const 0)
    (tee_local $1
     (i32.sub
      (get_local $2)
      (i32.and
       (i32.add
        (get_local $4)
        (i32.const 15)
       )
       (i32.const -16)
      )
     )
    )
   )
  )
  (drop
   (call $read_action_data
    (get_local $1)
    (get_local $4)
   )
  )
  (i64.store offset=8
   (get_local $7)
   (i64.const 0)
  )
  (set_local $2
   (call $memset
    (i32.add
     (i32.add
      (get_local $7)
      (i32.const 8)
     )
     (i32.const 8)
    )
    (i32.const 0)
    (i32.const 272)
   )
  )
  (i32.store
   (i32.add
    (get_local $7)
    (i32.const 64)
   )
   (i32.const 0)
  )
  (i64.store
   (i32.add
    (get_local $7)
    (i32.const 56)
   )
   (i64.const 0)
  )
  (i32.store
   (i32.add
    (get_local $7)
    (i32.const 80)
   )
   (i32.const 0)
  )
  (i32.store
   (i32.add
    (get_local $7)
    (i32.const 84)
   )
   (i32.const 0)
  )
  (i32.store
   (i32.add
    (get_local $7)
    (i32.const 88)
   )
   (i32.const 0)
  )
  (i32.store
   (i32.add
    (get_local $7)
    (i32.const 128)
   )
   (i32.const 0)
  )
  (i32.store
   (i32.add
    (get_local $7)
    (i32.const 132)
   )
   (i32.const 0)
  )
  (i32.store
   (i32.add
    (get_local $7)
    (i32.const 136)
   )
   (i32.const 0)
  )
  (i32.store
   (i32.add
    (get_local $7)
    (i32.const 152)
   )
   (i32.const 0)
  )
  (i32.store
   (i32.add
    (get_local $7)
    (i32.const 196)
   )
   (i32.const 0)
  )
  (i32.store
   (i32.add
    (get_local $7)
    (i32.const 200)
   )
   (i32.const 0)
  )
  (i32.store
   (i32.add
    (get_local $7)
    (i32.const 204)
   )
   (i32.const 0)
  )
  (i32.store
   (i32.add
    (get_local $7)
    (i32.const 208)
   )
   (i32.const 0)
  )
  (i32.store
   (i32.add
    (get_local $7)
    (i32.const 212)
   )
   (i32.const 0)
  )
  (i32.store
   (i32.add
    (get_local $7)
    (i32.const 216)
   )
   (i32.const 0)
  )
  (i32.store
   (i32.add
    (get_local $7)
    (i32.const 224)
   )
   (i32.const 0)
  )
  (i32.store
   (i32.add
    (get_local $7)
    (i32.const 228)
   )
   (i32.const 0)
  )
  (i32.store
   (i32.add
    (get_local $7)
    (i32.const 232)
   )
   (i32.const 0)
  )
  (i32.store
   (i32.add
    (get_local $7)
    (i32.const 236)
   )
   (i32.const 0)
  )
  (i32.store
   (i32.add
    (get_local $7)
    (i32.const 240)
   )
   (i32.const 0)
  )
  (i32.store
   (i32.add
    (get_local $7)
    (i32.const 244)
   )
   (i32.const 0)
  )
  (i32.store
   (i32.add
    (get_local $7)
    (i32.const 252)
   )
   (i32.const 0)
  )
  (i32.store
   (i32.add
    (get_local $7)
    (i32.const 256)
   )
   (i32.const 0)
  )
  (i32.store
   (i32.add
    (get_local $7)
    (i32.const 260)
   )
   (i32.const 0)
  )
  (i32.store
   (i32.add
    (get_local $7)
    (i32.const 264)
   )
   (i32.const 0)
  )
  (i32.store
   (i32.add
    (get_local $7)
    (i32.const 268)
   )
   (i32.const 0)
  )
  (i32.store
   (i32.add
    (get_local $7)
    (i32.const 272)
   )
   (i32.const 0)
  )
  (i32.store
   (i32.add
    (get_local $7)
    (i32.const 276)
   )
   (i32.const 0)
  )
  (i32.store
   (i32.add
    (i32.add
     (get_local $7)
     (i32.const 8)
    )
    (i32.const 272)
   )
   (i32.const 0)
  )
  (i32.store
   (i32.add
    (get_local $7)
    (i32.const 284)
   )
   (i32.const 0)
  )
  (i32.store offset=296
   (get_local $7)
   (i32.add
    (get_local $1)
    (get_local $4)
   )
  )
  (i32.store offset=288
   (get_local $7)
   (get_local $1)
  )
  (call $eosio_assert
   (i32.gt_u
    (get_local $4)
    (i32.const 7)
   )
   (i32.const 416)
  )
  (drop
   (call $memcpy
    (i32.add
     (get_local $7)
     (i32.const 8)
    )
    (get_local $1)
    (i32.const 8)
   )
  )
  (i32.store offset=292
   (get_local $7)
   (i32.add
    (get_local $1)
    (i32.const 8)
   )
  )
  (drop
   (call $_ZN5eosiorsINS_10datastreamIPKcEEEERT_S6_RNS_14discovery_fileE
    (i32.add
     (get_local $7)
     (i32.const 288)
    )
    (get_local $2)
   )
  )
  (block $label$2
   (br_if $label$2
    (i32.lt_u
     (get_local $4)
     (i32.const 513)
    )
   )
   (call $free
    (get_local $1)
   )
  )
  (set_local $4
   (i32.add
    (get_local $0)
    (i32.shr_s
     (get_local $3)
     (i32.const 1)
    )
   )
  )
  (set_local $5
   (i64.load offset=8
    (get_local $7)
   )
  )
  (set_local $1
   (call $_ZN5eosio14discovery_fileC2ERKS0_
    (i32.add
     (get_local $7)
     (i32.const 288)
    )
    (get_local $2)
   )
  )
  (block $label$3
   (br_if $label$3
    (i32.eqz
     (i32.and
      (get_local $3)
      (i32.const 1)
     )
    )
   )
   (set_local $6
    (i32.load
     (i32.add
      (i32.load
       (get_local $4)
      )
      (get_local $6)
     )
    )
   )
  )
  (call_indirect $FUNCSIG$viji
   (get_local $4)
   (get_local $5)
   (get_local $1)
   (get_local $6)
  )
  (drop
   (call $_ZN5eosio14discovery_fileD2Ev
    (get_local $1)
   )
  )
  (drop
   (call $_ZN5eosio14discovery_fileD2Ev
    (get_local $2)
   )
  )
  (i32.store offset=4
   (i32.const 0)
   (i32.add
    (get_local $7)
    (i32.const 560)
   )
  )
  (i32.const 1)
 )
 (func $_ZNSt3__113__vector_baseIN5eosio11multi_indexILy7108558431253954560ENS1_5disco11genesis_rowEJEE8item_ptrENS_9allocatorIS6_EEED2Ev (param $0 i32) (result i32)
  (local $1 i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i32)
  (local $6 i32)
  (local $7 i32)
  (block $label$0
   (br_if $label$0
    (i32.eqz
     (tee_local $1
      (i32.load
       (get_local $0)
      )
     )
    )
   )
   (block $label$1
    (block $label$2
     (br_if $label$2
      (i32.eq
       (tee_local $6
        (i32.load offset=4
         (get_local $0)
        )
       )
       (get_local $1)
      )
     )
     (loop $label$3
      (set_local $2
       (i32.load
        (tee_local $6
         (i32.add
          (get_local $6)
          (i32.const -24)
         )
        )
       )
      )
      (i32.store
       (get_local $6)
       (i32.const 0)
      )
      (block $label$4
       (br_if $label$4
        (i32.eqz
         (get_local $2)
        )
       )
       (block $label$5
        (br_if $label$5
         (i32.eqz
          (tee_local $3
           (i32.load offset=20
            (get_local $2)
           )
          )
         )
        )
        (block $label$6
         (block $label$7
          (br_if $label$7
           (i32.eq
            (tee_local $7
             (i32.load
              (tee_local $5
               (i32.add
                (get_local $2)
                (i32.const 24)
               )
              )
             )
            )
            (get_local $3)
           )
          )
          (set_local $4
           (i32.sub
            (i32.const 0)
            (get_local $3)
           )
          )
          (set_local $7
           (i32.add
            (get_local $7)
            (i32.const -12)
           )
          )
          (loop $label$8
           (block $label$9
            (br_if $label$9
             (i32.eqz
              (i32.and
               (i32.load8_u
                (get_local $7)
               )
               (i32.const 1)
              )
             )
            )
            (call $_ZdlPv
             (i32.load
              (i32.add
               (get_local $7)
               (i32.const 8)
              )
             )
            )
           )
           (br_if $label$8
            (i32.ne
             (i32.add
              (tee_local $7
               (i32.add
                (get_local $7)
                (i32.const -12)
               )
              )
              (get_local $4)
             )
             (i32.const -12)
            )
           )
          )
          (set_local $7
           (i32.load
            (i32.add
             (get_local $2)
             (i32.const 20)
            )
           )
          )
          (br $label$6)
         )
         (set_local $7
          (get_local $3)
         )
        )
        (i32.store
         (get_local $5)
         (get_local $3)
        )
        (call $_ZdlPv
         (get_local $7)
        )
       )
       (block $label$10
        (br_if $label$10
         (i32.eqz
          (i32.and
           (i32.load8_u offset=8
            (get_local $2)
           )
           (i32.const 1)
          )
         )
        )
        (call $_ZdlPv
         (i32.load
          (i32.add
           (get_local $2)
           (i32.const 16)
          )
         )
        )
       )
       (call $_ZdlPv
        (get_local $2)
       )
      )
      (br_if $label$3
       (i32.ne
        (get_local $6)
        (get_local $1)
       )
      )
     )
     (set_local $7
      (i32.load
       (get_local $0)
      )
     )
     (br $label$1)
    )
    (set_local $7
     (get_local $1)
    )
   )
   (i32.store
    (i32.add
     (get_local $0)
     (i32.const 4)
    )
    (get_local $1)
   )
   (call $_ZdlPv
    (get_local $7)
   )
  )
  (get_local $0)
 )
 (func $_ZN5eosio14discovery_fileC2ERKS0_ (param $0 i32) (param $1 i32) (result i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i32)
  (i64.store offset=40 align=4
   (tee_local $2
    (call $memcpy
     (get_local $0)
     (get_local $1)
     (i32.const 40)
    )
   )
   (i64.const 0)
  )
  (i32.store
   (i32.add
    (get_local $2)
    (i32.const 48)
   )
   (i32.const 0)
  )
  (block $label$0
   (block $label$1
    (block $label$2
     (br_if $label$2
      (i32.eqz
       (tee_local $0
        (i32.shr_s
         (tee_local $5
          (i32.sub
           (i32.load
            (i32.add
             (get_local $1)
             (i32.const 44)
            )
           )
           (i32.load offset=40
            (get_local $1)
           )
          )
         )
         (i32.const 5)
        )
       )
      )
     )
     (br_if $label$1
      (i32.ge_u
       (get_local $0)
       (i32.const 134217728)
      )
     )
     (i32.store
      (tee_local $4
       (i32.add
        (get_local $2)
        (i32.const 44)
       )
      )
      (tee_local $5
       (call $_Znwj
        (get_local $5)
       )
      )
     )
     (i32.store
      (i32.add
       (get_local $2)
       (i32.const 48)
      )
      (i32.add
       (get_local $5)
       (i32.shl
        (get_local $0)
        (i32.const 5)
       )
      )
     )
     (i32.store
      (i32.add
       (get_local $2)
       (i32.const 40)
      )
      (get_local $5)
     )
     (br_if $label$2
      (i32.eq
       (tee_local $0
        (i32.load
         (i32.add
          (get_local $1)
          (i32.const 40)
         )
        )
       )
       (tee_local $3
        (i32.load
         (i32.add
          (get_local $1)
          (i32.const 44)
         )
        )
       )
      )
     )
     (loop $label$3
      (i64.store
       (get_local $5)
       (i64.load
        (get_local $0)
       )
      )
      (drop
       (call $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEC2ERKS5_
        (i32.add
         (get_local $5)
         (i32.const 8)
        )
        (i32.add
         (get_local $0)
         (i32.const 8)
        )
       )
      )
      (i64.store offset=24
       (get_local $5)
       (i64.load
        (i32.add
         (get_local $0)
         (i32.const 24)
        )
       )
      )
      (i32.store
       (get_local $4)
       (tee_local $5
        (i32.add
         (i32.load
          (get_local $4)
         )
         (i32.const 32)
        )
       )
      )
      (br_if $label$3
       (i32.ne
        (tee_local $0
         (i32.add
          (get_local $0)
          (i32.const 32)
         )
        )
        (get_local $3)
       )
      )
     )
    )
    (i64.store align=4
     (i32.add
      (get_local $2)
      (i32.const 68)
     )
     (i64.const 0)
    )
    (i32.store offset=64
     (get_local $2)
     (i32.const 0)
    )
    (i64.store offset=56
     (get_local $2)
     (i64.load offset=56
      (get_local $1)
     )
    )
    (set_local $3
     (i32.div_s
      (tee_local $0
       (i32.sub
        (i32.load
         (i32.add
          (get_local $1)
          (i32.const 68)
         )
        )
        (i32.load offset=64
         (get_local $1)
        )
       )
      )
      (i32.const 12)
     )
    )
    (block $label$4
     (br_if $label$4
      (i32.eqz
       (get_local $0)
      )
     )
     (br_if $label$0
      (i32.ge_u
       (get_local $3)
       (i32.const 357913942)
      )
     )
     (i32.store
      (tee_local $4
       (i32.add
        (get_local $2)
        (i32.const 68)
       )
      )
      (tee_local $5
       (call $_Znwj
        (get_local $0)
       )
      )
     )
     (i32.store
      (i32.add
       (get_local $2)
       (i32.const 72)
      )
      (i32.add
       (get_local $5)
       (i32.mul
        (get_local $3)
        (i32.const 12)
       )
      )
     )
     (i32.store
      (i32.add
       (get_local $2)
       (i32.const 64)
      )
      (get_local $5)
     )
     (br_if $label$4
      (i32.eq
       (tee_local $0
        (i32.load
         (i32.add
          (get_local $1)
          (i32.const 64)
         )
        )
       )
       (tee_local $3
        (i32.load
         (i32.add
          (get_local $1)
          (i32.const 68)
         )
        )
       )
      )
     )
     (loop $label$5
      (drop
       (call $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEC2ERKS5_
        (get_local $5)
        (get_local $0)
       )
      )
      (i32.store
       (get_local $4)
       (tee_local $5
        (i32.add
         (i32.load
          (get_local $4)
         )
         (i32.const 12)
        )
       )
      )
      (br_if $label$5
       (i32.ne
        (get_local $3)
        (tee_local $0
         (i32.add
          (get_local $0)
          (i32.const 12)
         )
        )
       )
      )
     )
    )
    (drop
     (call $memcpy
      (i32.add
       (get_local $2)
       (i32.const 76)
      )
      (i32.add
       (get_local $1)
       (i32.const 76)
      )
      (i32.const 35)
     )
    )
    (drop
     (call $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEC2ERKS5_
      (i32.add
       (get_local $2)
       (i32.const 112)
      )
      (i32.add
       (get_local $1)
       (i32.const 112)
      )
     )
    )
    (drop
     (call $memcpy
      (i32.add
       (get_local $2)
       (i32.const 128)
      )
      (i32.add
       (get_local $1)
       (i32.const 128)
      )
      (i32.const 48)
     )
    )
    (drop
     (call $_ZN11eosiosystem9authorityC2ERKS0_
      (i32.add
       (get_local $2)
       (i32.const 176)
      )
      (i32.add
       (get_local $1)
       (i32.const 176)
      )
     )
    )
    (drop
     (call $_ZN11eosiosystem9authorityC2ERKS0_
      (i32.add
       (get_local $2)
       (i32.const 204)
      )
      (i32.add
       (get_local $1)
       (i32.const 204)
      )
     )
    )
    (drop
     (call $_ZN11eosiosystem9authorityC2ERKS0_
      (i32.add
       (get_local $2)
       (i32.const 232)
      )
      (i32.add
       (get_local $1)
       (i32.const 232)
      )
     )
    )
    (drop
     (call $_ZNSt3__16vectorIN5eosio11content_refENS_9allocatorIS2_EEEC2ERKS5_
      (i32.add
       (get_local $2)
       (i32.const 260)
      )
      (i32.add
       (get_local $1)
       (i32.const 260)
      )
     )
    )
    (return
     (get_local $2)
    )
   )
   (call $_ZNKSt3__120__vector_base_commonILb1EE20__throw_length_errorEv
    (i32.add
     (get_local $2)
     (i32.const 40)
    )
   )
   (unreachable)
  )
  (call $_ZNKSt3__120__vector_base_commonILb1EE20__throw_length_errorEv
   (i32.add
    (get_local $2)
    (i32.const 64)
   )
  )
  (unreachable)
 )
 (func $_ZN11eosiosystem9authorityC2ERKS0_ (param $0 i32) (param $1 i32) (result i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i32)
  (i64.store align=4
   (i32.add
    (get_local $0)
    (i32.const 8)
   )
   (i64.const 0)
  )
  (i32.store offset=4
   (get_local $0)
   (i32.const 0)
  )
  (i32.store
   (get_local $0)
   (i32.load
    (get_local $1)
   )
  )
  (set_local $3
   (i32.div_s
    (tee_local $2
     (i32.sub
      (i32.load
       (i32.add
        (get_local $1)
        (i32.const 8)
       )
      )
      (i32.load offset=4
       (get_local $1)
      )
     )
    )
    (i32.const 44)
   )
  )
  (block $label$0
   (block $label$1
    (block $label$2
     (br_if $label$2
      (i32.eqz
       (get_local $2)
      )
     )
     (br_if $label$1
      (i32.ge_u
       (get_local $3)
       (i32.const 97612894)
      )
     )
     (i32.store
      (i32.add
       (get_local $0)
       (i32.const 4)
      )
      (tee_local $2
       (call $_Znwj
        (get_local $2)
       )
      )
     )
     (i32.store
      (i32.add
       (get_local $0)
       (i32.const 12)
      )
      (i32.add
       (get_local $2)
       (i32.mul
        (get_local $3)
        (i32.const 44)
       )
      )
     )
     (i32.store
      (tee_local $3
       (i32.add
        (get_local $0)
        (i32.const 8)
       )
      )
      (get_local $2)
     )
     (br_if $label$2
      (i32.lt_s
       (tee_local $5
        (i32.sub
         (i32.load
          (i32.add
           (get_local $1)
           (i32.const 8)
          )
         )
         (tee_local $4
          (i32.load
           (i32.add
            (get_local $1)
            (i32.const 4)
           )
          )
         )
        )
       )
       (i32.const 1)
      )
     )
     (drop
      (call $memcpy
       (get_local $2)
       (get_local $4)
       (get_local $5)
      )
     )
     (i32.store
      (get_local $3)
      (i32.add
       (i32.load
        (get_local $3)
       )
       (i32.mul
        (i32.div_u
         (get_local $5)
         (i32.const 44)
        )
        (i32.const 44)
       )
      )
     )
    )
    (i64.store offset=16 align=4
     (get_local $0)
     (i64.const 0)
    )
    (i32.store
     (i32.add
      (get_local $0)
      (i32.const 24)
     )
     (i32.const 0)
    )
    (set_local $3
     (i32.div_s
      (tee_local $2
       (i32.sub
        (i32.load
         (i32.add
          (get_local $1)
          (i32.const 20)
         )
        )
        (i32.load offset=16
         (get_local $1)
        )
       )
      )
      (i32.const 24)
     )
    )
    (block $label$3
     (br_if $label$3
      (i32.eqz
       (get_local $2)
      )
     )
     (br_if $label$0
      (i32.ge_u
       (get_local $3)
       (i32.const 178956971)
      )
     )
     (i32.store
      (i32.add
       (get_local $0)
       (i32.const 16)
      )
      (tee_local $2
       (call $_Znwj
        (get_local $2)
       )
      )
     )
     (i32.store
      (i32.add
       (get_local $0)
       (i32.const 24)
      )
      (i32.add
       (get_local $2)
       (i32.mul
        (get_local $3)
        (i32.const 24)
       )
      )
     )
     (i32.store
      (tee_local $3
       (i32.add
        (get_local $0)
        (i32.const 20)
       )
      )
      (get_local $2)
     )
     (br_if $label$3
      (i32.lt_s
       (tee_local $1
        (i32.sub
         (i32.load
          (i32.add
           (get_local $1)
           (i32.const 20)
          )
         )
         (tee_local $5
          (i32.load
           (i32.add
            (get_local $1)
            (i32.const 16)
           )
          )
         )
        )
       )
       (i32.const 1)
      )
     )
     (drop
      (call $memcpy
       (get_local $2)
       (get_local $5)
       (get_local $1)
      )
     )
     (i32.store
      (get_local $3)
      (i32.add
       (i32.load
        (get_local $3)
       )
       (i32.mul
        (i32.div_u
         (get_local $1)
         (i32.const 24)
        )
        (i32.const 24)
       )
      )
     )
    )
    (return
     (get_local $0)
    )
   )
   (call $_ZNKSt3__120__vector_base_commonILb1EE20__throw_length_errorEv
    (i32.add
     (get_local $0)
     (i32.const 4)
    )
   )
   (unreachable)
  )
  (call $_ZNKSt3__120__vector_base_commonILb1EE20__throw_length_errorEv
   (i32.add
    (get_local $0)
    (i32.const 16)
   )
  )
  (unreachable)
 )
 (func $_ZNSt3__16vectorIN5eosio11content_refENS_9allocatorIS2_EEEC2ERKS5_ (param $0 i32) (param $1 i32) (result i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i32)
  (i64.store align=4
   (get_local $0)
   (i64.const 0)
  )
  (i32.store offset=8
   (get_local $0)
   (i32.const 0)
  )
  (set_local $2
   (i32.div_s
    (tee_local $5
     (i32.sub
      (i32.load offset=4
       (get_local $1)
      )
      (i32.load
       (get_local $1)
      )
     )
    )
    (i32.const 36)
   )
  )
  (block $label$0
   (block $label$1
    (br_if $label$1
     (i32.eqz
      (get_local $5)
     )
    )
    (br_if $label$0
     (i32.ge_u
      (get_local $2)
      (i32.const 119304648)
     )
    )
    (i32.store
     (tee_local $3
      (i32.add
       (get_local $0)
       (i32.const 4)
      )
     )
     (tee_local $4
      (call $_Znwj
       (get_local $5)
      )
     )
    )
    (i32.store
     (i32.add
      (get_local $0)
      (i32.const 8)
     )
     (i32.add
      (get_local $4)
      (i32.mul
       (get_local $2)
       (i32.const 36)
      )
     )
    )
    (i32.store
     (get_local $0)
     (get_local $4)
    )
    (br_if $label$1
     (i32.eq
      (tee_local $5
       (i32.load
        (get_local $1)
       )
      )
      (tee_local $1
       (i32.load
        (i32.add
         (get_local $1)
         (i32.const 4)
        )
       )
      )
     )
    )
    (loop $label$2
     (drop
      (call $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEC2ERKS5_
       (get_local $4)
       (get_local $5)
      )
     )
     (drop
      (call $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEC2ERKS5_
       (i32.add
        (get_local $4)
        (i32.const 12)
       )
       (i32.add
        (get_local $5)
        (i32.const 12)
       )
      )
     )
     (drop
      (call $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEC2ERKS5_
       (i32.add
        (get_local $4)
        (i32.const 24)
       )
       (i32.add
        (get_local $5)
        (i32.const 24)
       )
      )
     )
     (i32.store
      (get_local $3)
      (tee_local $4
       (i32.add
        (i32.load
         (get_local $3)
        )
        (i32.const 36)
       )
      )
     )
     (br_if $label$2
      (i32.ne
       (tee_local $5
        (i32.add
         (get_local $5)
         (i32.const 36)
        )
       )
       (get_local $1)
      )
     )
    )
   )
   (return
    (get_local $0)
   )
  )
  (call $_ZNKSt3__120__vector_base_commonILb1EE20__throw_length_errorEv
   (get_local $0)
  )
  (unreachable)
 )
 (func $_ZN5boost4mp116detail16tuple_apply_implIRZN5eosio14execute_actionINS3_5discoES5_JyNSt3__112basic_stringIcNS6_11char_traitsIcEENS6_9allocatorIcEEEENS6_6vectorISC_NSA_ISC_EEEEEEEbPT_MT0_FvDpT1_EEUlDpT_E_RNS6_5tupleIJySC_SF_EEEJLj0ELj1ELj2EEEEDTclclsr3stdE7forwardISG_Efp_Espclsr3stdE3getIXT1_EEclsr3stdE7forwardISI_Efp0_EEEEOSG_OSI_NS0_16integer_sequenceIjJXspT1_EEEE (param $0 i32) (param $1 i32)
  (local $2 i64)
  (local $3 i32)
  (local $4 i32)
  (local $5 i32)
  (local $6 i32)
  (i32.store offset=4
   (i32.const 0)
   (tee_local $6
    (i32.sub
     (i32.load offset=4
      (i32.const 0)
     )
     (i32.const 32)
    )
   )
  )
  (set_local $2
   (i64.load
    (get_local $1)
   )
  )
  (drop
   (call $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEC2ERKS5_
    (i32.add
     (get_local $6)
     (i32.const 16)
    )
    (i32.add
     (get_local $1)
     (i32.const 8)
    )
   )
  )
  (i32.store offset=8
   (get_local $6)
   (i32.const 0)
  )
  (i64.store
   (get_local $6)
   (i64.const 0)
  )
  (set_local $3
   (i32.div_s
    (tee_local $5
     (i32.sub
      (i32.load
       (i32.add
        (get_local $1)
        (i32.const 24)
       )
      )
      (i32.load offset=20
       (get_local $1)
      )
     )
    )
    (i32.const 12)
   )
  )
  (block $label$0
   (block $label$1
    (br_if $label$1
     (i32.eqz
      (get_local $5)
     )
    )
    (br_if $label$0
     (i32.ge_u
      (get_local $3)
      (i32.const 357913942)
     )
    )
    (i32.store
     (i32.add
      (get_local $6)
      (i32.const 8)
     )
     (i32.add
      (tee_local $4
       (call $_Znwj
        (get_local $5)
       )
      )
      (i32.mul
       (get_local $3)
       (i32.const 12)
      )
     )
    )
    (i32.store
     (get_local $6)
     (get_local $4)
    )
    (i32.store offset=4
     (get_local $6)
     (get_local $4)
    )
    (br_if $label$1
     (i32.eq
      (tee_local $5
       (i32.load
        (i32.add
         (get_local $1)
         (i32.const 20)
        )
       )
      )
      (tee_local $1
       (i32.load
        (i32.add
         (get_local $1)
         (i32.const 24)
        )
       )
      )
     )
    )
    (loop $label$2
     (drop
      (call $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEC2ERKS5_
       (get_local $4)
       (get_local $5)
      )
     )
     (i32.store offset=4
      (get_local $6)
      (tee_local $4
       (i32.add
        (i32.load offset=4
         (get_local $6)
        )
        (i32.const 12)
       )
      )
     )
     (br_if $label$2
      (i32.ne
       (get_local $1)
       (tee_local $5
        (i32.add
         (get_local $5)
         (i32.const 12)
        )
       )
      )
     )
    )
   )
   (call $_ZZN5eosio14execute_actionINS_5discoES1_JyNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEENS2_6vectorIS8_NS6_IS8_EEEEEEEbPT_MT0_FvDpT1_EENKUlDpT_E_clIJyS8_SB_EEEDaSK_
    (get_local $0)
    (get_local $2)
    (i32.add
     (get_local $6)
     (i32.const 16)
    )
    (get_local $6)
   )
   (block $label$3
    (br_if $label$3
     (i32.eqz
      (tee_local $1
       (i32.load
        (get_local $6)
       )
      )
     )
    )
    (block $label$4
     (block $label$5
      (br_if $label$5
       (i32.eq
        (tee_local $5
         (i32.load offset=4
          (get_local $6)
         )
        )
        (get_local $1)
       )
      )
      (set_local $4
       (i32.sub
        (i32.const 0)
        (get_local $1)
       )
      )
      (set_local $5
       (i32.add
        (get_local $5)
        (i32.const -12)
       )
      )
      (loop $label$6
       (block $label$7
        (br_if $label$7
         (i32.eqz
          (i32.and
           (i32.load8_u
            (get_local $5)
           )
           (i32.const 1)
          )
         )
        )
        (call $_ZdlPv
         (i32.load
          (i32.add
           (get_local $5)
           (i32.const 8)
          )
         )
        )
       )
       (br_if $label$6
        (i32.ne
         (i32.add
          (tee_local $5
           (i32.add
            (get_local $5)
            (i32.const -12)
           )
          )
          (get_local $4)
         )
         (i32.const -12)
        )
       )
      )
      (set_local $5
       (i32.load
        (get_local $6)
       )
      )
      (br $label$4)
     )
     (set_local $5
      (get_local $1)
     )
    )
    (i32.store offset=4
     (get_local $6)
     (get_local $1)
    )
    (call $_ZdlPv
     (get_local $5)
    )
   )
   (block $label$8
    (br_if $label$8
     (i32.eqz
      (i32.and
       (i32.load8_u offset=16
        (get_local $6)
       )
       (i32.const 1)
      )
     )
    )
    (call $_ZdlPv
     (i32.load offset=24
      (get_local $6)
     )
    )
   )
   (i32.store offset=4
    (i32.const 0)
    (i32.add
     (get_local $6)
     (i32.const 32)
    )
   )
   (return)
  )
  (call $_ZNKSt3__120__vector_base_commonILb1EE20__throw_length_errorEv
   (get_local $6)
  )
  (unreachable)
 )
 (func $_ZZN5eosio14execute_actionINS_5discoES1_JyNSt3__112basic_stringIcNS2_11char_traitsIcEENS2_9allocatorIcEEEENS2_6vectorIS8_NS6_IS8_EEEEEEEbPT_MT0_FvDpT1_EENKUlDpT_E_clIJyS8_SB_EEEDaSK_ (param $0 i32) (param $1 i64) (param $2 i32) (param $3 i32)
  (local $4 i32)
  (local $5 i32)
  (local $6 i32)
  (local $7 i32)
  (i32.store offset=4
   (i32.const 0)
   (tee_local $7
    (i32.sub
     (i32.load offset=4
      (i32.const 0)
     )
     (i32.const 32)
    )
   )
  )
  (set_local $4
   (i32.add
    (i32.load
     (i32.load
      (get_local $0)
     )
    )
    (i32.shr_s
     (tee_local $5
      (i32.load offset=4
       (tee_local $0
        (i32.load offset=4
         (get_local $0)
        )
       )
      )
     )
     (i32.const 1)
    )
   )
  )
  (set_local $6
   (i32.load
    (get_local $0)
   )
  )
  (block $label$0
   (br_if $label$0
    (i32.eqz
     (i32.and
      (get_local $5)
      (i32.const 1)
     )
    )
   )
   (set_local $6
    (i32.load
     (i32.add
      (i32.load
       (get_local $4)
      )
      (get_local $6)
     )
    )
   )
  )
  (drop
   (call $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEC2ERKS5_
    (i32.add
     (get_local $7)
     (i32.const 16)
    )
    (get_local $2)
   )
  )
  (i32.store offset=8
   (get_local $7)
   (i32.const 0)
  )
  (i64.store
   (get_local $7)
   (i64.const 0)
  )
  (set_local $5
   (i32.div_s
    (tee_local $0
     (i32.sub
      (i32.load offset=4
       (get_local $3)
      )
      (i32.load
       (get_local $3)
      )
     )
    )
    (i32.const 12)
   )
  )
  (block $label$1
   (block $label$2
    (br_if $label$2
     (i32.eqz
      (get_local $0)
     )
    )
    (br_if $label$1
     (i32.ge_u
      (get_local $5)
      (i32.const 357913942)
     )
    )
    (i32.store
     (i32.add
      (get_local $7)
      (i32.const 8)
     )
     (i32.add
      (tee_local $2
       (call $_Znwj
        (get_local $0)
       )
      )
      (i32.mul
       (get_local $5)
       (i32.const 12)
      )
     )
    )
    (i32.store
     (get_local $7)
     (get_local $2)
    )
    (i32.store offset=4
     (get_local $7)
     (get_local $2)
    )
    (br_if $label$2
     (i32.eq
      (tee_local $0
       (i32.load
        (get_local $3)
       )
      )
      (tee_local $3
       (i32.load
        (i32.add
         (get_local $3)
         (i32.const 4)
        )
       )
      )
     )
    )
    (loop $label$3
     (drop
      (call $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEC2ERKS5_
       (get_local $2)
       (get_local $0)
      )
     )
     (i32.store offset=4
      (get_local $7)
      (tee_local $2
       (i32.add
        (i32.load offset=4
         (get_local $7)
        )
        (i32.const 12)
       )
      )
     )
     (br_if $label$3
      (i32.ne
       (get_local $3)
       (tee_local $0
        (i32.add
         (get_local $0)
         (i32.const 12)
        )
       )
      )
     )
    )
   )
   (call_indirect $FUNCSIG$vijii
    (get_local $4)
    (get_local $1)
    (i32.add
     (get_local $7)
     (i32.const 16)
    )
    (get_local $7)
    (get_local $6)
   )
   (block $label$4
    (br_if $label$4
     (i32.eqz
      (tee_local $3
       (i32.load
        (get_local $7)
       )
      )
     )
    )
    (block $label$5
     (block $label$6
      (br_if $label$6
       (i32.eq
        (tee_local $0
         (i32.load offset=4
          (get_local $7)
         )
        )
        (get_local $3)
       )
      )
      (set_local $2
       (i32.sub
        (i32.const 0)
        (get_local $3)
       )
      )
      (set_local $0
       (i32.add
        (get_local $0)
        (i32.const -12)
       )
      )
      (loop $label$7
       (block $label$8
        (br_if $label$8
         (i32.eqz
          (i32.and
           (i32.load8_u
            (get_local $0)
           )
           (i32.const 1)
          )
         )
        )
        (call $_ZdlPv
         (i32.load
          (i32.add
           (get_local $0)
           (i32.const 8)
          )
         )
        )
       )
       (br_if $label$7
        (i32.ne
         (i32.add
          (tee_local $0
           (i32.add
            (get_local $0)
            (i32.const -12)
           )
          )
          (get_local $2)
         )
         (i32.const -12)
        )
       )
      )
      (set_local $0
       (i32.load
        (get_local $7)
       )
      )
      (br $label$5)
     )
     (set_local $0
      (get_local $3)
     )
    )
    (i32.store offset=4
     (get_local $7)
     (get_local $3)
    )
    (call $_ZdlPv
     (get_local $0)
    )
   )
   (block $label$9
    (br_if $label$9
     (i32.eqz
      (i32.and
       (i32.load8_u offset=16
        (get_local $7)
       )
       (i32.const 1)
      )
     )
    )
    (call $_ZdlPv
     (i32.load offset=24
      (get_local $7)
     )
    )
   )
   (i32.store offset=4
    (i32.const 0)
    (i32.add
     (get_local $7)
     (i32.const 32)
    )
   )
   (return)
  )
  (call $_ZNKSt3__120__vector_base_commonILb1EE20__throw_length_errorEv
   (get_local $7)
  )
  (unreachable)
 )
 (func $memcmp (param $0 i32) (param $1 i32) (param $2 i32) (result i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i32)
  (set_local $5
   (i32.const 0)
  )
  (block $label$0
   (br_if $label$0
    (i32.eqz
     (get_local $2)
    )
   )
   (block $label$1
    (loop $label$2
     (br_if $label$1
      (i32.ne
       (tee_local $3
        (i32.load8_u
         (get_local $0)
        )
       )
       (tee_local $4
        (i32.load8_u
         (get_local $1)
        )
       )
      )
     )
     (set_local $1
      (i32.add
       (get_local $1)
       (i32.const 1)
      )
     )
     (set_local $0
      (i32.add
       (get_local $0)
       (i32.const 1)
      )
     )
     (br_if $label$2
      (tee_local $2
       (i32.add
        (get_local $2)
        (i32.const -1)
       )
      )
     )
     (br $label$0)
    )
   )
   (set_local $5
    (i32.sub
     (get_local $3)
     (get_local $4)
    )
   )
  )
  (get_local $5)
 )
 (func $_Znwj (param $0 i32) (result i32)
  (local $1 i32)
  (local $2 i32)
  (block $label$0
   (br_if $label$0
    (tee_local $0
     (call $malloc
      (tee_local $1
       (select
        (get_local $0)
        (i32.const 1)
        (get_local $0)
       )
      )
     )
    )
   )
   (loop $label$1
    (set_local $0
     (i32.const 0)
    )
    (br_if $label$0
     (i32.eqz
      (tee_local $2
       (i32.load offset=436
        (i32.const 0)
       )
      )
     )
    )
    (call_indirect $FUNCSIG$v
     (get_local $2)
    )
    (br_if $label$1
     (i32.eqz
      (tee_local $0
       (call $malloc
        (get_local $1)
       )
      )
     )
    )
   )
  )
  (get_local $0)
 )
 (func $_ZdlPv (param $0 i32)
  (block $label$0
   (br_if $label$0
    (i32.eqz
     (get_local $0)
    )
   )
   (call $free
    (get_local $0)
   )
  )
 )
 (func $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEaSERKS5_ (param $0 i32) (param $1 i32) (result i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i32)
  (local $6 i32)
  (block $label$0
   (block $label$1
    (block $label$2
     (block $label$3
      (br_if $label$3
       (i32.eq
        (get_local $0)
        (get_local $1)
       )
      )
      (set_local $2
       (select
        (i32.load offset=4
         (get_local $1)
        )
        (i32.shr_u
         (tee_local $2
          (i32.load8_u
           (get_local $1)
          )
         )
         (i32.const 1)
        )
        (tee_local $4
         (i32.and
          (get_local $2)
          (i32.const 1)
         )
        )
       )
      )
      (set_local $5
       (i32.add
        (get_local $1)
        (i32.const 1)
       )
      )
      (set_local $6
       (i32.load offset=8
        (get_local $1)
       )
      )
      (set_local $1
       (i32.const 10)
      )
      (block $label$4
       (br_if $label$4
        (i32.eqz
         (i32.and
          (tee_local $3
           (i32.load8_u
            (get_local $0)
           )
          )
          (i32.const 1)
         )
        )
       )
       (set_local $1
        (i32.add
         (i32.and
          (tee_local $3
           (i32.load
            (get_local $0)
           )
          )
          (i32.const -2)
         )
         (i32.const -1)
        )
       )
      )
      (set_local $5
       (select
        (get_local $6)
        (get_local $5)
        (get_local $4)
       )
      )
      (set_local $4
       (i32.and
        (get_local $3)
        (i32.const 1)
       )
      )
      (block $label$5
       (block $label$6
        (block $label$7
         (br_if $label$7
          (i32.le_u
           (get_local $2)
           (get_local $1)
          )
         )
         (br_if $label$6
          (get_local $4)
         )
         (set_local $3
          (i32.shr_u
           (i32.and
            (get_local $3)
            (i32.const 254)
           )
           (i32.const 1)
          )
         )
         (br $label$5)
        )
        (br_if $label$2
         (get_local $4)
        )
        (set_local $1
         (i32.add
          (get_local $0)
          (i32.const 1)
         )
        )
        (br_if $label$1
         (get_local $2)
        )
        (br $label$0)
       )
       (set_local $3
        (i32.load offset=4
         (get_local $0)
        )
       )
      )
      (call $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEE21__grow_by_and_replaceEjjjjjjPKc
       (get_local $0)
       (get_local $1)
       (i32.sub
        (get_local $2)
        (get_local $1)
       )
       (get_local $3)
       (i32.const 0)
       (get_local $3)
       (get_local $2)
       (get_local $5)
      )
     )
     (return
      (get_local $0)
     )
    )
    (set_local $1
     (i32.load offset=8
      (get_local $0)
     )
    )
    (br_if $label$0
     (i32.eqz
      (get_local $2)
     )
    )
   )
   (drop
    (call $memmove
     (get_local $1)
     (get_local $5)
     (get_local $2)
    )
   )
  )
  (i32.store8
   (i32.add
    (get_local $1)
    (get_local $2)
   )
   (i32.const 0)
  )
  (block $label$8
   (br_if $label$8
    (i32.and
     (i32.load8_u
      (get_local $0)
     )
     (i32.const 1)
    )
   )
   (i32.store8
    (get_local $0)
    (i32.shl
     (get_local $2)
     (i32.const 1)
    )
   )
   (return
    (get_local $0)
   )
  )
  (i32.store offset=4
   (get_local $0)
   (get_local $2)
  )
  (get_local $0)
 )
 (func $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEE21__grow_by_and_replaceEjjjjjjPKc (param $0 i32) (param $1 i32) (param $2 i32) (param $3 i32) (param $4 i32) (param $5 i32) (param $6 i32) (param $7 i32)
  (local $8 i32)
  (local $9 i32)
  (local $10 i32)
  (block $label$0
   (br_if $label$0
    (i32.lt_u
     (i32.sub
      (i32.const -18)
      (get_local $1)
     )
     (get_local $2)
    )
   )
   (block $label$1
    (block $label$2
     (br_if $label$2
      (i32.and
       (i32.load8_u
        (get_local $0)
       )
       (i32.const 1)
      )
     )
     (set_local $9
      (i32.add
       (get_local $0)
       (i32.const 1)
      )
     )
     (br $label$1)
    )
    (set_local $9
     (i32.load offset=8
      (get_local $0)
     )
    )
   )
   (set_local $10
    (i32.const -17)
   )
   (block $label$3
    (br_if $label$3
     (i32.gt_u
      (get_local $1)
      (i32.const 2147483622)
     )
    )
    (set_local $10
     (i32.const 11)
    )
    (br_if $label$3
     (i32.lt_u
      (tee_local $2
       (select
        (tee_local $8
         (i32.shl
          (get_local $1)
          (i32.const 1)
         )
        )
        (tee_local $2
         (i32.add
          (get_local $2)
          (get_local $1)
         )
        )
        (i32.lt_u
         (get_local $2)
         (get_local $8)
        )
       )
      )
      (i32.const 11)
     )
    )
    (set_local $10
     (i32.and
      (i32.add
       (get_local $2)
       (i32.const 16)
      )
      (i32.const -16)
     )
    )
   )
   (set_local $2
    (call $_Znwj
     (get_local $10)
    )
   )
   (block $label$4
    (br_if $label$4
     (i32.eqz
      (get_local $4)
     )
    )
    (drop
     (call $memcpy
      (get_local $2)
      (get_local $9)
      (get_local $4)
     )
    )
   )
   (block $label$5
    (br_if $label$5
     (i32.eqz
      (get_local $6)
     )
    )
    (drop
     (call $memcpy
      (i32.add
       (get_local $2)
       (get_local $4)
      )
      (get_local $7)
      (get_local $6)
     )
    )
   )
   (block $label$6
    (br_if $label$6
     (i32.eqz
      (tee_local $7
       (i32.sub
        (tee_local $3
         (i32.sub
          (get_local $3)
          (get_local $5)
         )
        )
        (get_local $4)
       )
      )
     )
    )
    (drop
     (call $memcpy
      (i32.add
       (i32.add
        (get_local $2)
        (get_local $4)
       )
       (get_local $6)
      )
      (i32.add
       (i32.add
        (get_local $9)
        (get_local $4)
       )
       (get_local $5)
      )
      (get_local $7)
     )
    )
   )
   (block $label$7
    (br_if $label$7
     (i32.eq
      (get_local $1)
      (i32.const 10)
     )
    )
    (call $_ZdlPv
     (get_local $9)
    )
   )
   (i32.store offset=8
    (get_local $0)
    (get_local $2)
   )
   (i32.store
    (get_local $0)
    (i32.or
     (get_local $10)
     (i32.const 1)
    )
   )
   (i32.store offset=4
    (get_local $0)
    (tee_local $4
     (i32.add
      (get_local $3)
      (get_local $6)
     )
    )
   )
   (i32.store8
    (i32.add
     (get_local $2)
     (get_local $4)
    )
    (i32.const 0)
   )
   (return)
  )
  (call $abort)
  (unreachable)
 )
 (func $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEE6resizeEjc (param $0 i32) (param $1 i32) (param $2 i32)
  (local $3 i32)
  (local $4 i32)
  (block $label$0
   (block $label$1
    (block $label$2
     (br_if $label$2
      (tee_local $4
       (i32.and
        (tee_local $3
         (i32.load8_u
          (get_local $0)
         )
        )
        (i32.const 1)
       )
      )
     )
     (br_if $label$1
      (i32.lt_u
       (tee_local $3
        (i32.shr_u
         (get_local $3)
         (i32.const 1)
        )
       )
       (get_local $1)
      )
     )
     (br $label$0)
    )
    (br_if $label$0
     (i32.ge_u
      (tee_local $3
       (i32.load offset=4
        (get_local $0)
       )
      )
      (get_local $1)
     )
    )
   )
   (drop
    (call $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEE6appendEjc
     (get_local $0)
     (i32.sub
      (get_local $1)
      (get_local $3)
     )
     (get_local $2)
    )
   )
   (return)
  )
  (block $label$3
   (br_if $label$3
    (get_local $4)
   )
   (i32.store8
    (i32.add
     (i32.add
      (get_local $0)
      (get_local $1)
     )
     (i32.const 1)
    )
    (i32.const 0)
   )
   (i32.store8
    (get_local $0)
    (i32.shl
     (get_local $1)
     (i32.const 1)
    )
   )
   (return)
  )
  (i32.store8
   (i32.add
    (i32.load offset=8
     (get_local $0)
    )
    (get_local $1)
   )
   (i32.const 0)
  )
  (i32.store offset=4
   (get_local $0)
   (get_local $1)
  )
 )
 (func $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEE6appendEjc (param $0 i32) (param $1 i32) (param $2 i32) (result i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i32)
  (block $label$0
   (br_if $label$0
    (i32.eqz
     (get_local $1)
    )
   )
   (set_local $3
    (i32.const 10)
   )
   (block $label$1
    (br_if $label$1
     (i32.eqz
      (i32.and
       (tee_local $5
        (i32.load8_u
         (get_local $0)
        )
       )
       (i32.const 1)
      )
     )
    )
    (set_local $3
     (i32.add
      (i32.and
       (tee_local $5
        (i32.load
         (get_local $0)
        )
       )
       (i32.const -2)
      )
      (i32.const -1)
     )
    )
   )
   (block $label$2
    (block $label$3
     (br_if $label$3
      (i32.and
       (get_local $5)
       (i32.const 1)
      )
     )
     (set_local $4
      (i32.shr_u
       (i32.and
        (get_local $5)
        (i32.const 254)
       )
       (i32.const 1)
      )
     )
     (br $label$2)
    )
    (set_local $4
     (i32.load offset=4
      (get_local $0)
     )
    )
   )
   (block $label$4
    (br_if $label$4
     (i32.ge_u
      (i32.sub
       (get_local $3)
       (get_local $4)
      )
      (get_local $1)
     )
    )
    (call $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEE9__grow_byEjjjjjj
     (get_local $0)
     (get_local $3)
     (i32.add
      (i32.sub
       (get_local $1)
       (get_local $3)
      )
      (get_local $4)
     )
     (get_local $4)
     (get_local $4)
     (i32.const 0)
     (i32.const 0)
    )
    (set_local $5
     (i32.load8_u
      (get_local $0)
     )
    )
   )
   (block $label$5
    (block $label$6
     (br_if $label$6
      (i32.and
       (get_local $5)
       (i32.const 1)
      )
     )
     (set_local $5
      (i32.add
       (get_local $0)
       (i32.const 1)
      )
     )
     (br $label$5)
    )
    (set_local $5
     (i32.load offset=8
      (get_local $0)
     )
    )
   )
   (drop
    (call $memset
     (i32.add
      (get_local $5)
      (get_local $4)
     )
     (i32.and
      (get_local $2)
      (i32.const 255)
     )
     (get_local $1)
    )
   )
   (set_local $1
    (i32.add
     (get_local $4)
     (get_local $1)
    )
   )
   (block $label$7
    (block $label$8
     (br_if $label$8
      (i32.and
       (i32.load8_u
        (get_local $0)
       )
       (i32.const 1)
      )
     )
     (i32.store8
      (get_local $0)
      (i32.shl
       (get_local $1)
       (i32.const 1)
      )
     )
     (br $label$7)
    )
    (i32.store offset=4
     (get_local $0)
     (get_local $1)
    )
   )
   (i32.store8
    (i32.add
     (get_local $5)
     (get_local $1)
    )
    (i32.const 0)
   )
  )
  (get_local $0)
 )
 (func $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEE9__grow_byEjjjjjj (param $0 i32) (param $1 i32) (param $2 i32) (param $3 i32) (param $4 i32) (param $5 i32) (param $6 i32)
  (local $7 i32)
  (local $8 i32)
  (local $9 i32)
  (block $label$0
   (br_if $label$0
    (i32.lt_u
     (i32.sub
      (i32.const -17)
      (get_local $1)
     )
     (get_local $2)
    )
   )
   (block $label$1
    (block $label$2
     (br_if $label$2
      (i32.and
       (i32.load8_u
        (get_local $0)
       )
       (i32.const 1)
      )
     )
     (set_local $8
      (i32.add
       (get_local $0)
       (i32.const 1)
      )
     )
     (br $label$1)
    )
    (set_local $8
     (i32.load offset=8
      (get_local $0)
     )
    )
   )
   (set_local $9
    (i32.const -17)
   )
   (block $label$3
    (br_if $label$3
     (i32.gt_u
      (get_local $1)
      (i32.const 2147483622)
     )
    )
    (set_local $9
     (i32.const 11)
    )
    (br_if $label$3
     (i32.lt_u
      (tee_local $2
       (select
        (tee_local $7
         (i32.shl
          (get_local $1)
          (i32.const 1)
         )
        )
        (tee_local $2
         (i32.add
          (get_local $2)
          (get_local $1)
         )
        )
        (i32.lt_u
         (get_local $2)
         (get_local $7)
        )
       )
      )
      (i32.const 11)
     )
    )
    (set_local $9
     (i32.and
      (i32.add
       (get_local $2)
       (i32.const 16)
      )
      (i32.const -16)
     )
    )
   )
   (set_local $2
    (call $_Znwj
     (get_local $9)
    )
   )
   (block $label$4
    (br_if $label$4
     (i32.eqz
      (get_local $4)
     )
    )
    (drop
     (call $memcpy
      (get_local $2)
      (get_local $8)
      (get_local $4)
     )
    )
   )
   (block $label$5
    (br_if $label$5
     (i32.eqz
      (tee_local $3
       (i32.sub
        (i32.sub
         (get_local $3)
         (get_local $5)
        )
        (get_local $4)
       )
      )
     )
    )
    (drop
     (call $memcpy
      (i32.add
       (i32.add
        (get_local $2)
        (get_local $4)
       )
       (get_local $6)
      )
      (i32.add
       (i32.add
        (get_local $8)
        (get_local $4)
       )
       (get_local $5)
      )
      (get_local $3)
     )
    )
   )
   (block $label$6
    (br_if $label$6
     (i32.eq
      (get_local $1)
      (i32.const 10)
     )
    )
    (call $_ZdlPv
     (get_local $8)
    )
   )
   (i32.store offset=8
    (get_local $0)
    (get_local $2)
   )
   (i32.store
    (get_local $0)
    (i32.or
     (get_local $9)
     (i32.const 1)
    )
   )
   (return)
  )
  (call $abort)
  (unreachable)
 )
 (func $_ZNKSt3__120__vector_base_commonILb1EE20__throw_length_errorEv (param $0 i32)
  (call $abort)
  (unreachable)
 )
 (func $_ZNSt3__112basic_stringIcNS_11char_traitsIcEENS_9allocatorIcEEEC2ERKS5_ (param $0 i32) (param $1 i32) (result i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (i64.store align=4
   (get_local $0)
   (i64.const 0)
  )
  (i32.store
   (tee_local $3
    (i32.add
     (get_local $0)
     (i32.const 8)
    )
   )
   (i32.const 0)
  )
  (block $label$0
   (br_if $label$0
    (i32.and
     (i32.load8_u
      (get_local $1)
     )
     (i32.const 1)
    )
   )
   (i64.store align=4
    (get_local $0)
    (i64.load align=4
     (get_local $1)
    )
   )
   (i32.store
    (get_local $3)
    (i32.load
     (i32.add
      (get_local $1)
      (i32.const 8)
     )
    )
   )
   (return
    (get_local $0)
   )
  )
  (block $label$1
   (br_if $label$1
    (i32.ge_u
     (tee_local $3
      (i32.load offset=4
       (get_local $1)
      )
     )
     (i32.const -16)
    )
   )
   (set_local $2
    (i32.load offset=8
     (get_local $1)
    )
   )
   (block $label$2
    (block $label$3
     (block $label$4
      (br_if $label$4
       (i32.ge_u
        (get_local $3)
        (i32.const 11)
       )
      )
      (i32.store8
       (get_local $0)
       (i32.shl
        (get_local $3)
        (i32.const 1)
       )
      )
      (set_local $1
       (i32.add
        (get_local $0)
        (i32.const 1)
       )
      )
      (br_if $label$3
       (get_local $3)
      )
      (br $label$2)
     )
     (set_local $1
      (call $_Znwj
       (tee_local $4
        (i32.and
         (i32.add
          (get_local $3)
          (i32.const 16)
         )
         (i32.const -16)
        )
       )
      )
     )
     (i32.store
      (get_local $0)
      (i32.or
       (get_local $4)
       (i32.const 1)
      )
     )
     (i32.store offset=8
      (get_local $0)
      (get_local $1)
     )
     (i32.store offset=4
      (get_local $0)
      (get_local $3)
     )
    )
    (drop
     (call $memcpy
      (get_local $1)
      (get_local $2)
      (get_local $3)
     )
    )
   )
   (i32.store8
    (i32.add
     (get_local $1)
     (get_local $3)
    )
    (i32.const 0)
   )
   (return
    (get_local $0)
   )
  )
  (call $abort)
  (unreachable)
 )
 (func $malloc (param $0 i32) (result i32)
  (call $_ZN5eosio14memory_manager6mallocEm
   (i32.const 440)
   (get_local $0)
  )
 )
 (func $_ZN5eosio14memory_manager6mallocEm (param $0 i32) (param $1 i32) (result i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i32)
  (local $6 i32)
  (local $7 i32)
  (local $8 i32)
  (local $9 i32)
  (local $10 i32)
  (local $11 i32)
  (local $12 i32)
  (local $13 i32)
  (block $label$0
   (br_if $label$0
    (i32.eqz
     (get_local $1)
    )
   )
   (block $label$1
    (br_if $label$1
     (tee_local $13
      (i32.load offset=8384
       (get_local $0)
      )
     )
    )
    (set_local $13
     (i32.const 16)
    )
    (i32.store
     (i32.add
      (get_local $0)
      (i32.const 8384)
     )
     (i32.const 16)
    )
   )
   (set_local $2
    (select
     (i32.sub
      (i32.add
       (get_local $1)
       (i32.const 8)
      )
      (tee_local $2
       (i32.and
        (i32.add
         (get_local $1)
         (i32.const 4)
        )
        (i32.const 7)
       )
      )
     )
     (get_local $1)
     (get_local $2)
    )
   )
   (block $label$2
    (block $label$3
     (block $label$4
      (br_if $label$4
       (i32.ge_u
        (tee_local $10
         (i32.load offset=8388
          (get_local $0)
         )
        )
        (get_local $13)
       )
      )
      (set_local $1
       (i32.add
        (i32.add
         (get_local $0)
         (i32.mul
          (get_local $10)
          (i32.const 12)
         )
        )
        (i32.const 8192)
       )
      )
      (block $label$5
       (br_if $label$5
        (get_local $10)
       )
       (br_if $label$5
        (i32.load
         (tee_local $13
          (i32.add
           (get_local $0)
           (i32.const 8196)
          )
         )
        )
       )
       (i32.store
        (get_local $1)
        (i32.const 8192)
       )
       (i32.store
        (get_local $13)
        (get_local $0)
       )
      )
      (set_local $10
       (i32.add
        (get_local $2)
        (i32.const 4)
       )
      )
      (loop $label$6
       (block $label$7
        (br_if $label$7
         (i32.gt_u
          (i32.add
           (tee_local $13
            (i32.load offset=8
             (get_local $1)
            )
           )
           (get_local $10)
          )
          (i32.load
           (get_local $1)
          )
         )
        )
        (i32.store
         (tee_local $13
          (i32.add
           (i32.load offset=4
            (get_local $1)
           )
           (get_local $13)
          )
         )
         (i32.or
          (i32.and
           (i32.load
            (get_local $13)
           )
           (i32.const -2147483648)
          )
          (get_local $2)
         )
        )
        (i32.store
         (tee_local $1
          (i32.add
           (get_local $1)
           (i32.const 8)
          )
         )
         (i32.add
          (i32.load
           (get_local $1)
          )
          (get_local $10)
         )
        )
        (i32.store
         (get_local $13)
         (i32.or
          (i32.load
           (get_local $13)
          )
          (i32.const -2147483648)
         )
        )
        (br_if $label$3
         (tee_local $1
          (i32.add
           (get_local $13)
           (i32.const 4)
          )
         )
        )
       )
       (br_if $label$6
        (tee_local $1
         (call $_ZN5eosio14memory_manager16next_active_heapEv
          (get_local $0)
         )
        )
       )
      )
     )
     (set_local $4
      (i32.sub
       (i32.const 2147483644)
       (get_local $2)
      )
     )
     (set_local $11
      (i32.add
       (get_local $0)
       (i32.const 8392)
      )
     )
     (set_local $12
      (i32.add
       (get_local $0)
       (i32.const 8384)
      )
     )
     (set_local $13
      (tee_local $3
       (i32.load offset=8392
        (get_local $0)
       )
      )
     )
     (loop $label$8
      (call $eosio_assert
       (i32.eq
        (i32.load
         (i32.add
          (tee_local $1
           (i32.add
            (get_local $0)
            (i32.mul
             (get_local $13)
             (i32.const 12)
            )
           )
          )
          (i32.const 8200)
         )
        )
        (i32.load
         (tee_local $5
          (i32.add
           (get_local $1)
           (i32.const 8192)
          )
         )
        )
       )
       (i32.const 8848)
      )
      (set_local $13
       (i32.add
        (tee_local $6
         (i32.load
          (i32.add
           (get_local $1)
           (i32.const 8196)
          )
         )
        )
        (i32.const 4)
       )
      )
      (loop $label$9
       (set_local $7
        (i32.add
         (get_local $6)
         (i32.load
          (get_local $5)
         )
        )
       )
       (set_local $1
        (i32.and
         (tee_local $9
          (i32.load
           (tee_local $8
            (i32.add
             (get_local $13)
             (i32.const -4)
            )
           )
          )
         )
         (i32.const 2147483647)
        )
       )
       (block $label$10
        (br_if $label$10
         (i32.lt_s
          (get_local $9)
          (i32.const 0)
         )
        )
        (block $label$11
         (br_if $label$11
          (i32.ge_u
           (get_local $1)
           (get_local $2)
          )
         )
         (loop $label$12
          (br_if $label$11
           (i32.ge_u
            (tee_local $10
             (i32.add
              (get_local $13)
              (get_local $1)
             )
            )
            (get_local $7)
           )
          )
          (br_if $label$11
           (i32.lt_s
            (tee_local $10
             (i32.load
              (get_local $10)
             )
            )
            (i32.const 0)
           )
          )
          (br_if $label$12
           (i32.lt_u
            (tee_local $1
             (i32.add
              (i32.add
               (get_local $1)
               (i32.and
                (get_local $10)
                (i32.const 2147483647)
               )
              )
              (i32.const 4)
             )
            )
            (get_local $2)
           )
          )
         )
        )
        (i32.store
         (get_local $8)
         (i32.or
          (select
           (get_local $1)
           (get_local $2)
           (i32.lt_u
            (get_local $1)
            (get_local $2)
           )
          )
          (i32.and
           (get_local $9)
           (i32.const -2147483648)
          )
         )
        )
        (block $label$13
         (br_if $label$13
          (i32.le_u
           (get_local $1)
           (get_local $2)
          )
         )
         (i32.store
          (i32.add
           (get_local $13)
           (get_local $2)
          )
          (i32.and
           (i32.add
            (get_local $4)
            (get_local $1)
           )
           (i32.const 2147483647)
          )
         )
        )
        (br_if $label$2
         (i32.ge_u
          (get_local $1)
          (get_local $2)
         )
        )
       )
       (br_if $label$9
        (i32.lt_u
         (tee_local $13
          (i32.add
           (i32.add
            (get_local $13)
            (get_local $1)
           )
           (i32.const 4)
          )
         )
         (get_local $7)
        )
       )
      )
      (set_local $1
       (i32.const 0)
      )
      (i32.store
       (get_local $11)
       (tee_local $13
        (select
         (i32.const 0)
         (tee_local $13
          (i32.add
           (i32.load
            (get_local $11)
           )
           (i32.const 1)
          )
         )
         (i32.eq
          (get_local $13)
          (i32.load
           (get_local $12)
          )
         )
        )
       )
      )
      (br_if $label$8
       (i32.ne
        (get_local $13)
        (get_local $3)
       )
      )
     )
    )
    (return
     (get_local $1)
    )
   )
   (i32.store
    (get_local $8)
    (i32.or
     (i32.load
      (get_local $8)
     )
     (i32.const -2147483648)
    )
   )
   (return
    (get_local $13)
   )
  )
  (i32.const 0)
 )
 (func $_ZN5eosio14memory_manager16next_active_heapEv (param $0 i32) (result i32)
  (local $1 i32)
  (local $2 i32)
  (local $3 i32)
  (local $4 i32)
  (local $5 i32)
  (local $6 i32)
  (local $7 i32)
  (local $8 i32)
  (set_local $1
   (i32.load offset=8388
    (get_local $0)
   )
  )
  (block $label$0
   (block $label$1
    (br_if $label$1
     (i32.eqz
      (i32.load8_u offset=8934
       (i32.const 0)
      )
     )
    )
    (set_local $7
     (i32.load offset=8936
      (i32.const 0)
     )
    )
    (br $label$0)
   )
   (set_local $7
    (current_memory)
   )
   (i32.store8 offset=8934
    (i32.const 0)
    (i32.const 1)
   )
   (i32.store offset=8936
    (i32.const 0)
    (tee_local $7
     (i32.shl
      (get_local $7)
      (i32.const 16)
     )
    )
   )
  )
  (set_local $3
   (get_local $7)
  )
  (block $label$2
   (block $label$3
    (block $label$4
     (block $label$5
      (br_if $label$5
       (i32.le_u
        (tee_local $2
         (i32.shr_u
          (i32.add
           (get_local $7)
           (i32.const 65535)
          )
          (i32.const 16)
         )
        )
        (tee_local $8
         (current_memory)
        )
       )
      )
      (drop
       (grow_memory
        (i32.sub
         (get_local $2)
         (get_local $8)
        )
       )
      )
      (set_local $8
       (i32.const 0)
      )
      (br_if $label$4
       (i32.ne
        (get_local $2)
        (current_memory)
       )
      )
      (set_local $3
       (i32.load offset=8936
        (i32.const 0)
       )
      )
     )
     (set_local $8
      (i32.const 0)
     )
     (i32.store offset=8936
      (i32.const 0)
      (get_local $3)
     )
     (br_if $label$4
      (i32.lt_s
       (get_local $7)
       (i32.const 0)
      )
     )
     (set_local $2
      (i32.add
       (get_local $0)
       (i32.mul
        (get_local $1)
        (i32.const 12)
       )
      )
     )
     (set_local $7
      (i32.sub
       (i32.sub
        (i32.add
         (get_local $7)
         (select
          (i32.const 65536)
          (i32.const 131072)
          (tee_local $6
           (i32.lt_u
            (tee_local $8
             (i32.and
              (get_local $7)
              (i32.const 65535)
             )
            )
            (i32.const 64513)
           )
          )
         )
        )
        (select
         (get_local $8)
         (i32.and
          (get_local $7)
          (i32.const 131071)
         )
         (get_local $6)
        )
       )
       (get_local $7)
      )
     )
     (block $label$6
      (br_if $label$6
       (i32.load8_u offset=8934
        (i32.const 0)
       )
      )
      (set_local $3
       (current_memory)
      )
      (i32.store8 offset=8934
       (i32.const 0)
       (i32.const 1)
      )
      (i32.store offset=8936
       (i32.const 0)
       (tee_local $3
        (i32.shl
         (get_local $3)
         (i32.const 16)
        )
       )
      )
     )
     (set_local $2
      (i32.add
       (get_local $2)
       (i32.const 8192)
      )
     )
     (br_if $label$3
      (i32.lt_s
       (get_local $7)
       (i32.const 0)
      )
     )
     (set_local $6
      (get_local $3)
     )
     (block $label$7
      (br_if $label$7
       (i32.le_u
        (tee_local $8
         (i32.shr_u
          (i32.add
           (i32.add
            (tee_local $5
             (i32.and
              (i32.add
               (get_local $7)
               (i32.const 7)
              )
              (i32.const -8)
             )
            )
            (get_local $3)
           )
           (i32.const 65535)
          )
          (i32.const 16)
         )
        )
        (tee_local $4
         (current_memory)
        )
       )
      )
      (drop
       (grow_memory
        (i32.sub
         (get_local $8)
         (get_local $4)
        )
       )
      )
      (br_if $label$3
       (i32.ne
        (get_local $8)
        (current_memory)
       )
      )
      (set_local $6
       (i32.load offset=8936
        (i32.const 0)
       )
      )
     )
     (i32.store offset=8936
      (i32.const 0)
      (i32.add
       (get_local $6)
       (get_local $5)
      )
     )
     (br_if $label$3
      (i32.eq
       (get_local $3)
       (i32.const -1)
      )
     )
     (br_if $label$2
      (i32.eq
       (i32.add
        (tee_local $6
         (i32.load
          (i32.add
           (tee_local $1
            (i32.add
             (get_local $0)
             (i32.mul
              (get_local $1)
              (i32.const 12)
             )
            )
           )
           (i32.const 8196)
          )
         )
        )
        (tee_local $8
         (i32.load
          (get_local $2)
         )
        )
       )
       (get_local $3)
      )
     )
     (block $label$8
      (br_if $label$8
       (i32.eq
        (get_local $8)
        (tee_local $1
         (i32.load
          (tee_local $5
           (i32.add
            (get_local $1)
            (i32.const 8200)
           )
          )
         )
        )
       )
      )
      (i32.store
       (tee_local $6
        (i32.add
         (get_local $6)
         (get_local $1)
        )
       )
       (i32.or
        (i32.and
         (i32.load
          (get_local $6)
         )
         (i32.const -2147483648)
        )
        (i32.add
         (i32.sub
          (i32.const -4)
          (get_local $1)
         )
         (get_local $8)
        )
       )
      )
      (i32.store
       (get_local $5)
       (i32.load
        (get_local $2)
       )
      )
      (i32.store
       (get_local $6)
       (i32.and
        (i32.load
         (get_local $6)
        )
        (i32.const 2147483647)
       )
      )
     )
     (i32.store
      (tee_local $2
       (i32.add
        (get_local $0)
        (i32.const 8388)
       )
      )
      (tee_local $2
       (i32.add
        (i32.load
         (get_local $2)
        )
        (i32.const 1)
       )
      )
     )
     (i32.store
      (i32.add
       (tee_local $0
        (i32.add
         (get_local $0)
         (i32.mul
          (get_local $2)
          (i32.const 12)
         )
        )
       )
       (i32.const 8196)
      )
      (get_local $3)
     )
     (i32.store
      (tee_local $8
       (i32.add
        (get_local $0)
        (i32.const 8192)
       )
      )
      (get_local $7)
     )
    )
    (return
     (get_local $8)
    )
   )
   (block $label$9
    (br_if $label$9
     (i32.eq
      (tee_local $8
       (i32.load
        (get_local $2)
       )
      )
      (tee_local $7
       (i32.load
        (tee_local $1
         (i32.add
          (tee_local $3
           (i32.add
            (get_local $0)
            (i32.mul
             (get_local $1)
             (i32.const 12)
            )
           )
          )
          (i32.const 8200)
         )
        )
       )
      )
     )
    )
    (i32.store
     (tee_local $3
      (i32.add
       (i32.load
        (i32.add
         (get_local $3)
         (i32.const 8196)
        )
       )
       (get_local $7)
      )
     )
     (i32.or
      (i32.and
       (i32.load
        (get_local $3)
       )
       (i32.const -2147483648)
      )
      (i32.add
       (i32.sub
        (i32.const -4)
        (get_local $7)
       )
       (get_local $8)
      )
     )
    )
    (i32.store
     (get_local $1)
     (i32.load
      (get_local $2)
     )
    )
    (i32.store
     (get_local $3)
     (i32.and
      (i32.load
       (get_local $3)
      )
      (i32.const 2147483647)
     )
    )
   )
   (i32.store offset=8384
    (get_local $0)
    (tee_local $3
     (i32.add
      (i32.load
       (tee_local $7
        (i32.add
         (get_local $0)
         (i32.const 8388)
        )
       )
      )
      (i32.const 1)
     )
    )
   )
   (i32.store
    (get_local $7)
    (get_local $3)
   )
   (return
    (i32.const 0)
   )
  )
  (i32.store
   (get_local $2)
   (i32.add
    (get_local $8)
    (get_local $7)
   )
  )
  (get_local $2)
 )
 (func $free (param $0 i32)
  (local $1 i32)
  (local $2 i32)
  (local $3 i32)
  (block $label$0
   (block $label$1
    (br_if $label$1
     (i32.eqz
      (get_local $0)
     )
    )
    (br_if $label$1
     (i32.lt_s
      (tee_local $2
       (i32.load offset=8824
        (i32.const 0)
       )
      )
      (i32.const 1)
     )
    )
    (set_local $3
     (i32.const 8632)
    )
    (set_local $1
     (i32.add
      (i32.mul
       (get_local $2)
       (i32.const 12)
      )
      (i32.const 8632)
     )
    )
    (loop $label$2
     (br_if $label$1
      (i32.eqz
       (tee_local $2
        (i32.load
         (i32.add
          (get_local $3)
          (i32.const 4)
         )
        )
       )
      )
     )
     (block $label$3
      (br_if $label$3
       (i32.gt_u
        (i32.add
         (get_local $2)
         (i32.const 4)
        )
        (get_local $0)
       )
      )
      (br_if $label$0
       (i32.gt_u
        (i32.add
         (get_local $2)
         (i32.load
          (get_local $3)
         )
        )
        (get_local $0)
       )
      )
     )
     (br_if $label$2
      (i32.lt_u
       (tee_local $3
        (i32.add
         (get_local $3)
         (i32.const 12)
        )
       )
       (get_local $1)
      )
     )
    )
   )
   (return)
  )
  (i32.store
   (tee_local $3
    (i32.add
     (get_local $0)
     (i32.const -4)
    )
   )
   (i32.and
    (i32.load
     (get_local $3)
    )
    (i32.const 2147483647)
   )
  )
 )
 (func $__wasm_nullptr (type $FUNCSIG$v)
  (unreachable)
 )
)
