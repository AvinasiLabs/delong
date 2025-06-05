# Delong Docker Setup with Sepolia Testnet

This guide explains how to set up the Delong application using Sepolia testnet instead of a local Anvil node.

## Prerequisites

1. **Docker and Docker Compose** installed
2. **Sepolia Test ETH** - Get from faucets (see below)

## Quick Start

1. **Generate a private key:**
   ```bash
   # Install Foundry if you haven't already
   curl -L https://foundry.paradigm.xyz | bash
   foundryup

   # Generate a new wallet (shows both address and private key)
   cast wallet new
   ```

2. **Create environment file:**
   ```bash
   cd delong/deploy/docker
   cp .env.template .env
   ```

3. **Edit .env file:**
   ```env
   ETH_HTTP_URL=https://ethereum-sepolia-rpc.publicnode.com
   ETH_WS_URL=wss://ethereum-sepolia-rpc.publicnode.com
   OFFICIAL_ACCOUNT_PRIVATE_KEY=0x[your_private_key_from_step_1]
   JWT_SECRET=your_secure_jwt_secret_here
   ```

4. **Get test ETH:**
   - Use the address shown by `cast wallet new` command
   - Get Sepolia ETH from [sepoliafaucet.com](https://sepoliafaucet.com/)

5. **Start services:**
   ```bash
   docker-compose -f docker-compose.local.yml up -d
   ```

That's it! The application will automatically deploy contracts on first run.

## Detailed Setup Instructions

### 1. Create Environment File

Copy the template file and fill in your values:

```bash
./scripts/setup_env.sh
```

### 2. Configure RPC Endpoints

You can use any Sepolia RPC provider. Here are some options:

- **Public RPC** (free, no signup): `https://ethereum-sepolia-rpc.publicnode.com`
- **Infura** (more reliable): `https://sepolia.infura.io/v3/YOUR_PROJECT_ID`
- **Alchemy** (more reliable): `https://eth-sepolia.g.alchemy.com/v2/YOUR_API_KEY`

Simply set the URLs directly in your `.env` file.

### 3. Set up Test Account

You need a private key for an Ethereum account that will be used to deploy contracts and fund other accounts.

#### Option A: Using Foundry's `cast` (Recommended - shows both address and private key)
```bash
# Install Foundry if you haven't already
curl -L https://foundry.paradigm.xyz | bash
foundryup

# Generate a new wallet
cast wallet new
```

#### Option B: Using MetaMask (User-friendly)
1. Install MetaMask browser extension
2. Create a new account
3. Go to Account Details → Export Private Key
4. Copy the private key (already includes 0x prefix)

#### Option C: Using OpenSSL (Simple, but you need to calculate address separately)
```bash
# Generate a private key
openssl rand -hex 32

# The output is your private key (add 0x prefix when using it)
# Example output: a1b2c3d4e5f6789012345678901234567890abcdef1234567890abcdef123456
# Use as: 0xa1b2c3d4e5f6789012345678901234567890abcdef1234567890abcdef123456
```

#### Option D: Use Existing Test Account
If you already have a test account, export the private key from MetaMask or your wallet.

**⚠️ SECURITY WARNING**: Only use test accounts! Never use mainnet private keys!

### 4. Get Sepolia Test ETH

Your test account needs Sepolia ETH to deploy contracts and pay for gas. Get some from these faucets:

- [Sepolia Faucet](https://sepoliafaucet.com/)
- [Alchemy Sepolia Faucet](https://sepoliafaucet.com/)
- [Infura Sepolia Faucet](https://www.infura.io/faucet/sepolia)

You'll need at least 0.1 ETH for contract deployment and operations.

### 5. Complete Environment Configuration

Your `.env` file should look like this:

```env
# RPC Configuration - Use any Sepolia RPC endpoint
ETH_HTTP_URL=https://ethereum-sepolia-rpc.publicnode.com
ETH_WS_URL=wss://ethereum-sepolia-rpc.publicnode.com

# Account Configuration
OFFICIAL_ACCOUNT_PRIVATE_KEY=0x1234567890abcdef...
JWT_SECRET=your_secure_jwt_secret_here

# Optional: Gas Configuration
# GAS_PRICE_GWEI=20
# FUNDING_THRESHOLD_ETH=0.01
# FUNDING_TOPUP_ETH=0.1
```

### 6. Start the Services

```bash
# Start all services
docker-compose -f docker-compose.local.yml up -d

# Check logs
docker-compose -f docker-compose.local.yml logs -f delong

# Stop services
docker-compose -f docker-compose.local.yml down
```

## Service Overview

The setup includes these services:

- **mysql**: Database for application data
- **ipfs**: IPFS node for decentralized storage
- **dstack-simulator**: Simulation environment
- **delong**: Main application connecting to Sepolia

## Network Configuration

- **Network**: Sepolia Testnet
- **Chain ID**: 11155111
- **RPC**: Infura endpoints
- **Explorer**: [sepolia.etherscan.io](https://sepolia.etherscan.io/)

## Troubleshooting

### Common Issues

1. **"insufficient funds for gas * price + value"**
   - Your account needs more Sepolia ETH
   - Use the faucets listed above

2. **"connection refused" or RPC errors**
   - Check your ETH_HTTP_URL and ETH_WS_URL are correct
   - Try switching to a different RPC provider (public RPCs can be unreliable)
   - Check your internet connection

3. **"invalid private key"**
   - Ensure private key starts with "0x"
   - Private key should be 64 characters (32 bytes) in hex

4. **Contract deployment fails**
   - Check account has sufficient ETH (>0.05 ETH recommended)
   - Verify network connectivity to Infura

### Verify Setup

Check your account balance:
```bash
# Replace with your account address and RPC URL
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"eth_getBalance","params":["0xYourAccountAddress","latest"],"id":1}' \
  https://ethereum-sepolia-rpc.publicnode.com
```

### Logs and Debugging

```bash
# View application logs
docker-compose -f docker-compose.local.yml logs delong

# View all service logs
docker-compose -f docker-compose.local.yml logs

# Restart specific service
docker-compose -f docker-compose.local.yml restart delong
```

## Security Notes

- Never commit `.env` file to version control
- Only use test accounts and test ETH
- Monitor your Infura usage to avoid rate limits
- Keep your private keys secure and never share them

## Migration from Anvil

If you're migrating from the previous Anvil setup:

1. Contracts will need to be redeployed on Sepolia
2. All transaction hashes and addresses will be different
3. Transaction confirmation times will be longer (real network)
4. Gas costs will be real (though minimal on testnet)

The application will automatically deploy contracts on first run if they don't exist in the database.
