# Contributing to GoTe

We welcome contributions to the GoTe project! This document outlines how you can contribute to the development of this template for generating Go microservices with opinionated packages.

## Purpose

This template repository is designed to provide a starting point for building Go microservices within Adroit Group, offering pre-configured structures and opinionated packages to streamline development.

## Scope

This repository focuses on providing a robust and opinionated template for generating Go microservices. Contributions should be aligned with this goal.

**In Scope:**

- Bug fixes related to the template structure, generated code, or included packages.
- Improvements to the existing opinionated packages.
- Adding new, opinionated packages that enhance the microservice development experience in Go.
- Improvements to the documentation of the template or included packages.
- Examples demonstrating how to use the generated microservices or specific packages.

**Out of Scope:**

- Extending the template to generate services in languages other than Go.
- Adding frontend or user interface components.
- Contributions that significantly deviate from the established opinionated patterns.

## How to Contribute

We follow a workflow based on GitHub Issues and Pull Requests.

1.  **Open an Issue:** If you find a bug, have a feature request, or propose a significant change, please open an issue first. We have [issue templates](link to your issue templates if applicable) available to help you provide the necessary information. This allows for discussion and helps avoid duplicate work.
2.  **Fork the Repository:** Fork the GoTe repository to your own GitHub account.
3.  **Create a Branch:** Create a new branch in your forked repository for your contribution. Use a descriptive branch name related to the issue or feature.
4.  **Make Your Changes:** Implement your bug fix or feature in your new branch.
5.  **Test Your Changes:** Ensure your changes are covered by unit tests. All provided linters must pass.
6.  **Commit Your Changes:** Write clear and concise commit messages following the [Conventional Commits](https://www.conventionalcommits.org/en/latest/) specification.
7.  **Open a Pull Request:** Open a pull request from your branch in your forked repository to the `main` branch in the main repository.

## Pull Request Review

All pull requests are subject to review.

1.  **CI Checks:** Your pull request must pass all Continuous Integration (CI) checks. This includes running linters and tests.
2.  **Code Review:** A member of the Adroit team will review your code. They may provide feedback and request changes.

## Style Guidelines

We strive for consistency in our codebase. Please adhere to the following:

- All provided linters must pass. These linters are in place to guide code style and identify potential issues.
- Follow the [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md) for general Go coding practices.

## Development Environment

While there are no strict prerequisites beyond having Go installed, we highly recommend using [`mise`](https://mise.jdx.dev/) to manage dependencies and ensure you have the correct tools installed for development. Refer to the project's documentation for specific mise configuration.

## Documentation

- **Godoc:** All exported functions, types, and variables in library code should be documented using Godoc comments.
- **General Documentation:** For significant features or changes to the template itself, update the relevant section to reflect your contribution in the README.

## Getting Help

If you have questions or need assistance, please open a GitHub issue.

## License

This project is licensed under the Apache License 2.0. See the `LICENSE` file for more details.

## Code of Conduct

While we don't have a formal Code of Conduct document at this time, we expect all contributors to interact respectfully and professionally.
