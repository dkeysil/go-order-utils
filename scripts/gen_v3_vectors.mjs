// Generate ExchangeV3 (combos) order signing vectors using viem as the
// oracle. viem's hashTypedData under this domain was verified byte-exact
// against on-chain OrderFilled topic1 (Polygon tx 0xdfe5c6f0e78e59da197f
// fe1ef52598438549b89b26cd742711df088b46c9e8c5). The output is pasted as
// golden constants into pkg/builder/exchange_order_builder_impl_test.go.
//
// Usage:
//     npm install viem
//     node scripts/gen_v3_vectors.mjs
import { privateKeyToAccount } from 'viem/accounts';
import { hashTypedData } from 'viem';

const CHAIN_ID = 137; // ExchangeV3 is Polygon-mainnet only
const EXCHANGE_V3 = '0xe3333700cA9d93003F00f0F71f8515005F6c00Aa';

// Anvil default account #0: private key is publicly known, fine for tests.
const account = privateKeyToAccount('0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80');
const MAKER = '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266';

const SALT = 479249096354n;
const TIMESTAMP = 1700000000000n;
const ZERO32 = '0x' + '0'.repeat(64);

const domain = { name: 'Polymarket CTF Exchange', version: '3', chainId: CHAIN_ID, verifyingContract: EXCHANGE_V3 };
const types = {
  Order: [
    { name: 'salt', type: 'uint256' },
    { name: 'maker', type: 'address' },
    { name: 'signer', type: 'address' },
    { name: 'tokenId', type: 'uint256' },
    { name: 'makerAmount', type: 'uint256' },
    { name: 'takerAmount', type: 'uint256' },
    { name: 'side', type: 'uint8' },
    { name: 'signatureType', type: 'uint8' },
    { name: 'timestamp', type: 'uint256' },
    { name: 'metadata', type: 'bytes32' },
    { name: 'builder', type: 'bytes32' },
  ],
};

async function vector(label, signatureType, metadata = ZERO32, builder = ZERO32) {
  const message = {
    salt: SALT,
    maker: MAKER,
    signer: MAKER,
    tokenId: 1234n,
    makerAmount: 100000000n,
    takerAmount: 50000000n,
    side: 0, // BUY
    signatureType,
    timestamp: TIMESTAMP,
    metadata,
    builder,
  };
  const hash = hashTypedData({ domain, types, primaryType: 'Order', message });
  const signature = await account.signTypedData({ domain, types, primaryType: 'Order', message });
  console.log(`// ---- ${label} ----`);
  console.log(`//   signatureType  = ${signatureType}`);
  console.log(`//   metadata       = ${metadata}`);
  console.log(`//   builder        = ${builder}`);
  console.log(`expectedOrderHash = ${hash}`);
  console.log(`expectedSignature = ${signature}`);
  console.log();
}

await vector('ExchangeV3, EOA', 0);
await vector('ExchangeV3, POLY_GNOSIS_SAFE', 2);
await vector(
  'ExchangeV3, EOA, with non-zero metadata + builder',
  0,
  '0x1111111111111111111111111111111111111111111111111111111111111111',
  '0x2222222222222222222222222222222222222222222222222222222222222222',
);
