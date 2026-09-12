Yes. Since this repository contains **all three projects**, the README should present it as your **Go Projects Portfolio**, with each project clearly separated.

# Go Projects

A collection of practical Go projects developed and extended from algorithmic programming exercises during my software engineering training.

This repository demonstrates my progression in **Go programming, algorithmic problem solving, data processing, software engineering practices, and security-focused development**.

## Projects

### 🔐 1. Security Toolkit

A lightweight command-line toolkit for text and number analysis.

**Key capabilities:**

* Prime factorization
* Hidden sequence detection
* Character intersection
* Character union
* Alternating data selection

**Concepts demonstrated:**

* Algorithms and problem solving
* String processing
* Integer factorization
* Input validation
* CLI development
* Unit testing

**Security relevance:**
The project explores foundational concepts that can be extended toward security applications such as pattern detection, sensitive-data handling, input validation, and security-data analysis.

📁 `security-toolkit/`

---

### 🛡️ 2. PathGuard

A path validation engine based on exact-jump reachability.

Given a sequence of allowed movements, PathGuard determines whether execution can reach the final position without leaving the permitted range.

Example:

```text
[2 3 1 1 4]

0 → 2 → 3 → 4

VALID
```

**Key capabilities:**

* Path validation
* Reachability checking
* Path reconstruction
* Boundary validation
* Cycle detection
* Unit testing

**Concepts demonstrated:**

* Algorithm design
* State tracking
* Slices and maps
* Edge-case handling
* Package organization
* Testing

**Security relevance:**
The same principles can be applied to validating state transitions, execution paths, workflows, and security policies.

📁 `pathguard/`

---

### ⚙️ 3. Go Data Processing Engine

A reusable collection of data-processing utilities built around Go slices and strings.

**Key capabilities:**

* Data chunking
* Slice concatenation
* Data interleaving
* Text normalization

**Concepts demonstrated:**

* Slice manipulation
* Data transformation
* Input validation
* Reusable Go packages
* CLI application design
* Unit testing

**Security/AI relevance:**
These operations provide foundational building blocks for processing datasets, logs, batch workloads, and data used in security and AI pipelines.

📁 `go-data-engine/`

---

## Technologies

* Go
* Git
* GitHub
* Command-Line Interfaces
* Go Testing
* GitHub Actions

## Engineering Practices

Across these projects, I focus on:

* Writing reusable functions and packages
* Handling invalid and edge-case input
* Testing core functionality
* Organizing code into maintainable structures
* Using version control effectively
* Documenting technical decisions
* Building foundations that can be extended into security and AI applications

## Portfolio Direction

These projects represent part of my progression toward **Software Engineering and AI Security**.

My goal is to build beyond isolated programming exercises by transforming algorithmic foundations into practical tools that can eventually support:

```text
Software Engineering
        ↓
Data Processing
        ↓
Security Engineering
        ↓
AI Security
```

## Repository Structure

```text
Go-projects/
│
├── security-toolkit/
│   ├── main.go
│   ├── main_test.go
│   └── README.md
│
├── pathguard/
│   ├── pathguard/
│   ├── cmd/
│   ├── README.md
│   └── ...
│
└── go-data-engine/
    ├── processor/
    ├── cmd/
    ├── README.md
    └── ...
```

## Getting Started

Each project is independently structured and contains its own README with installation, usage, examples, and testing instructions.

Clone the repository:

```bash
git clone https://github.com/Fakepoj/Go-projects.git
cd Go-projects
```

Navigate into any project and follow its README.

## Author

**Kenneth Fatore**

Go programming trainee with a focus on software engineering, cybersecurity, and AI security.

---

> These projects started from algorithmic programming exercises and have been developed into practical software projects to demonstrate engineering ability, problem-solving skills, and continuous technical growth.

This is ready to paste directly into the **`README.md` at the root of your `Go-projects` repository**.
