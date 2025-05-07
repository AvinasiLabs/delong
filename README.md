# 🔐 Delong - Privacy-Preserving Computation Platform

## 📌 Overview

Delong is a trustless privacy-preserving computation platform specifically designed for longevity and biomedical research. It leverages blockchain governance combined with Trusted Execution Environment (TEE) technology, enabling researchers to analyze sensitive biomedical data securely without compromising privacy. Currently, the platform operates on [Phala Network](https://phala.network/), utilizing its extensive network of TEE nodes and on-chain DAO governance for algorithm approval and execution.

This repository contains the fully open-sourced code running inside the Delong platform's TEE instances. The primary purpose of this open-source effort is to allow users and security auditors to verify that the code running within Phala’s Confidential Virtual Machine (CVM) aligns exactly with this publicly accessible source code, adhering to a true "trust but verify" principle.

## 💡 Motivation

Biomedical and longevity research often involves highly sensitive personal health data, requiring strong privacy protections. Traditional centralized data handling models expose these data to potential breaches and unauthorized access, creating significant ethical and security risks. Additionally, researchers need assurance that their analytical algorithms are executed faithfully, without manipulation or unauthorized inspection.

Delong aims to solve these issues by providing a secure, decentralized, and transparent computational environment where sensitive data remains fully protected and computation integrity is guaranteed. By integrating blockchain governance with TEE technologies, Delong ensures robust privacy, verifiable trustworthiness, and community-driven oversight.

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

For detailed insights and complete smart contract implementations, please visit our [Avinasi Contracts repository](https://github.com/AvinasiLabs/Avinasi_contracts).

## 🔍 Open Source & Audit

Transparency is a foundational principle of the Delong platform. By open-sourcing all TEE-related code, the platform enables thorough auditing by security experts and community members. Users can independently inspect and confirm the integrity and security of algorithms executed within the TEE environment, thus significantly enhancing trust and reliability.

Additionally, through remote attestation provided by TEE technology, users can confidently verify the consistency of running code with the publicly audited source code, reinforcing the platform’s core "trust but verify" philosophy.

We warmly welcome security researchers and community members to review our code, suggest improvements, and participate actively in the continuous security enhancement of the Delong platform. 🌟
