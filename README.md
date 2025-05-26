# 🧬 Delong - Privacy-Preserving Computation Platform for Biomedical Data

## 📌 Overview

Delong is a trustless privacy-preserving computation platform specifically designed for longevity and biomedical research. It leverages blockchain governance combined with Trusted Execution Environment (TEE) technology, enabling researchers to analyze sensitive biomedical data securely without compromising privacy. Currently, the platform operates on [Phala Network](https://phala.network/), utilizing its extensive network of TEE nodes and on-chain DAO governance for algorithm approval and execution.

This repository contains the fully open-sourced code running inside the Delong platform's TEE instances. The primary purpose of this open-source effort is to allow users and security auditors to verify that the code running within Phala’s Confidential Virtual Machine (CVM) aligns exactly with this publicly accessible source code, adhering to a true "trust but verify" principle.

## 💡 Motivation

Biomedical and longevity research relies on sensitive personal data, such as genomic profiles and diagnostic records. While Trusted Execution Environments (TEEs) ensure secure, privacy-preserving computation, transparency and incentive alignment require a broader trust model.

Delong combines blockchain with TEE to achieve:

- 🔐 Secure and encrypted storage of user-contributed bio-data
- 🧪 Confidential execution of algorithms without exposing raw data
- 🧾 Immutable on-chain records of all data usage and algorithm executions
- 🧑‍⚖️ Community-driven audit of algorithm safety via decentralized governance
- 💎 Tokenized incentives: data contributors are rewarded with Dataset Tokens and evolution NFTs based on actual usage of their data
- 🎓 Verifiable scientific impact: contributors can trace how their data enabled real research outcomes

By recording every data contribution and algorithm usage on-chain, Delong enables **verifiable attribution**, **trustless governance**, and **sustainable token-driven participation**, empowering both researchers and citizen contributors in an open, auditable, and privacy-first scientific ecosystem.

## 🚀 Key Features

* **🔒 Data Privacy**: Data is encrypted at all times, decrypted only briefly within the TEE for computation, protecting sensitive information from any external party.
* **✅ Trustworthy Execution**: All computations are isolated within secure TEE hardware, ensuring algorithms executed are pre-approved by decentralized blockchain governance.
* **👀 Transparency and Auditability**: Every critical operation, from data submission to algorithm execution, is logged transparently on-chain, facilitating community oversight.
* **⚡ Scalable & Efficient**: Leveraging thousands of globally distributed TEE nodes provided by Phala Network ensures scalable performance with robust parallel computation capabilities.

## 🛠️ Architectural Overview

Delong adopts a comprehensive blockchain + TEE hybrid architecture that combines secure, transparent on-chain contract operations with privacy-preserving off-chain TEE computations. Our architecture follows an event-driven approach, minimizing on-chain state storage and maximizing transparency, efficiency, and security.

### Core Components:

* **AlgorithmReview Contract**: Handles the governance and auditing of algorithms submitted by researchers. A controlled committee initially reviews algorithms for potential data leakage. Auditing results are transparently emitted as blockchain events. This system plans to transition from centralized committees to a decentralized, token-based governance model involving relevant stakeholders and AI-assisted auditing.

* **DataContribution Contract**: Records user data contributions and usage behaviors transparently on-chain through events. Data contributions are managed off-chain with TEE-backed systems, ensuring secure, private data processing.

* **Off-Chain System**: Executes approved algorithms securely off-chain within TEE nodes, providing robust data confidentiality and execution integrity.

For detailed insights and complete smart contract implementations, please visit our [delong-contract](https://github.com/AvinasiLabs/delong-contract) repository.

## 📅 Strategic Roadmap

## ✅ Q2 2025 — MVP Closed Alpha Launch

**Milestone:** Complete the first trustless end-to-end demo in a controlled environment

- ✅ Chain-based record of data submission and algorithm approval
- ✅ TEE-based algorithm execution with Docker container isolation
- ✅ Basic Web UI for dataset upload and algorithm interaction
- ✅ Initial backend modules: API Service, Chainsync, Runtime

---

## 🚀 Q3 2025 — Public Launch and Contribution Scoring Framework

**Milestone:** Delong protocol goes live on public mainnet with first real users and foundation for token rewards

- 🔄 Smart contracts deployed on Ethereum mainnet (or other production-grade L1)
- 🧠 Auditing system transition: from curator committee → token-gated governance voting
- 🌐 First batch of real users onboarded (scientists submitting algorithms, data contributors uploading samples)
- 🎯 Design contribution-weighted reward system for future token distribution

## 🛠 Q4 2025 — Scientific Developer Ecosystem

**Milestone:** Empower researchers to easily integrate Delong with their workflows

- 🧰 Release SDK for dataset loading, algorithm packaging, result validation and tools for auditing + diagnostic helpers
- 📸 Standardize data output formats (models, images, numerical results)
- 🧑‍🔬 Begin pilot collaborations with 2–3 research labs
- 📈 Expand open-access dataset registry to >10 representative sets (e.g. methylation, transcriptomics, aging markers)
- 🧬 Begin onboarding real longitudinal samples from contributors (~1,000 profiles target)
- 🧪 Upgrade voting UX + token staking for audit participation

---

## 🧱 Q1 2026 — Reliability and Scaling Infrastructure

**Milestone:** Prepare for growth and long-term sustainability

- 🗂️ Migrate off-chain storage to distributed FS (e.g. JuiceFS) with hot-swap TEE capability
- 🔁 Enable snapshot + TEE node migration + resilience testing
- 🔐 Implement tiered dataset access policy (free / staked / permissioned)
- 🌍 Begin exploring multi-chain or cross-chain data referencing model

---

## ⚖️ Q2 2026 — Token Launch and Community Governance

**Milestone:** Transition platform ownership to token holders

- 🎉 Official TGE (Token Generation Event)
- ⚖️ Launch Delong DAO
  - Define proposal lifecycle, quorum, voting weights
- 🧠 Enable algorithm audit market (open review + reward system)
- 📊 Publish real-world data usage reports and reward history

---

## 🔭 Long-Term Vision (2026+)

- Enable global DataDAO governance for dataset curation
- Standardize Delong protocol for integration into longevity journals and registries

## 🔍 Open Source & Audit

Transparency is a foundational principle of the Delong platform. By open-sourcing all TEE-related code, the platform enables thorough auditing by security experts and community members. Users can independently inspect and confirm the integrity and security of algorithms executed within the TEE environment, thus significantly enhancing trust and reliability.

Additionally, through remote attestation provided by TEE technology, users can confidently verify the consistency of running code with the publicly audited source code, reinforcing the platform’s core "trust but verify" philosophy.

We warmly welcome security researchers, privacy advocates, and community contributors to audit our codebase, propose enhancements, and actively participate in building a secure and transparent future for biomedical research through Delong. 🌟
