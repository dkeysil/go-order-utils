"""
Generate V2 CTF Exchange order signing vectors using py-clob-client-v2 as the
oracle. The output is pasted as golden constants into
pkg/builder/exchange_order_builder_impl_test.go.

Usage:
    python3 -m venv .venv
    .venv/bin/pip install py-clob-client-v2==1.0.0
    .venv/bin/python scripts/gen_v2_vectors.py
"""
from py_clob_client_v2.constants import BYTES32_ZERO
from py_clob_client_v2.order_utils.exchange_order_builder_v2 import (
    ExchangeOrderBuilderV2,
)
from py_clob_client_v2.order_utils.model.order_data_v2 import OrderDataV2
from py_clob_client_v2.order_utils.model.signature_type_v2 import SignatureTypeV2
from py_clob_client_v2.signer import Signer


CHAIN_ID = 80002
EXCHANGE_V2 = "0xE111180000d2663C0091e4f400237545B87B996B"
NEG_RISK_EXCHANGE_V2 = "0xe2222d279d744050d28e00520010520000310F59"

# Anvil default account #0: private key is publicly known, fine for tests.
PRIV_KEY = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
MAKER = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"

SALT = 479249096354
TIMESTAMP = "1700000000000"


def vector(label, contract_addr, sig_type, metadata=BYTES32_ZERO, builder_code=BYTES32_ZERO):
    signer = Signer(private_key=PRIV_KEY, chain_id=CHAIN_ID)
    b = ExchangeOrderBuilderV2(
        contract_address=contract_addr,
        chain_id=CHAIN_ID,
        signer=signer,
        generate_salt=lambda: SALT,
    )

    data = OrderDataV2(
        maker=MAKER,
        tokenId="1234",
        makerAmount="100000000",
        takerAmount="50000000",
        side=0,  # BUY
        signer=MAKER,
        signatureType=sig_type,
        timestamp=TIMESTAMP,
        metadata=metadata,
        builder=builder_code,
        expiration="0",
    )

    order = b.build_order(data)
    typed = b.build_order_typed_data(order)
    order_hash = b.build_order_hash(typed)
    signed = b.build_signed_order(data)

    print(f"// ---- {label} ----")
    print(f"//   contract       = {contract_addr}")
    print(f"//   signatureType  = {int(sig_type)}")
    print(f"//   metadata       = {metadata}")
    print(f"//   builder        = {builder_code}")
    print(f"//   salt           = {SALT}")
    print(f"//   timestamp      = {TIMESTAMP}")
    print(f"expectedOrderHash = {order_hash}")
    print(f"expectedSignature = {signed.signature}")
    print()


if __name__ == "__main__":
    vector("CTFExchange V2, EOA", EXCHANGE_V2, SignatureTypeV2.EOA)
    vector("NegRiskCTFExchange V2, EOA", NEG_RISK_EXCHANGE_V2, SignatureTypeV2.EOA)
    vector("CTFExchange V2, POLY_GNOSIS_SAFE", EXCHANGE_V2, SignatureTypeV2.POLY_GNOSIS_SAFE)
    vector(
        "CTFExchange V2, EOA, with non-zero metadata + builder",
        EXCHANGE_V2,
        SignatureTypeV2.EOA,
        metadata="0x1111111111111111111111111111111111111111111111111111111111111111",
        builder_code="0x2222222222222222222222222222222222222222222222222222222222222222",
    )
