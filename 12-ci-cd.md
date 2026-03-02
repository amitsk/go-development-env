# 12. CI/CD (Automation on GitHub)

In this chapter, we'll learn how to automate our testing and building process using **CI/CD**.

[&larr; Back to [TOC](../README.md#table-of-contents)] | [&larr; [11-static-analysis-linting.md](./11-static-analysis-linting.md)] | [&rarr; [13-advanced-topics.md](./13-advanced-topics.md)]

## What is CI/CD?

In professional software development, you're constantly changing code. To make sure you don't accidentally break things, we use automation tools to run our tests and builds every single time we save our work to GitHub.

- **Continuous Integration (CI)**: Automatically testing and linting your code every time you push it. This ensures that new changes don't \"break\" the existing application.
- **Continuous Delivery/Deployment (CD)**: Automatically building and \"releasing\" your application to your users.

---

## Introducing GitHub Actions

[**GitHub Actions**](https://github.com/features/actions) is a powerful tool built directly into GitHub that handles CI/CD for you. You define your automation steps in a YAML file inside your repository.

### Example: A Simple CI Workflow

Create a file at `.github/workflows/ci.yml`:

```yaml
name: CI

# 1. When should this run?
on: [push, pull_request]

jobs:
  test:
    # 2. What operating system should it run on?
    runs-on: ubuntu-latest

    steps:
    # 3. Get the code from the repository
    - uses: actions/checkout@v4

    # 4. Set up the environment using mise
    # This ensures the CI server uses the same Go version as you!
    - uses: jdx/mise-action@v2

    # 5. Run the automation steps from our Makefile
    - name: Run Linting
      run: make lint

    - name: Run Tests
      run: make test

    - name: Run Build
      run: make build
```

---

## Why Use `mise` in CI?

One of the biggest problems in software development is when code works on your laptop but fails on the server. By using the `jdx/mise-action`, you guarantee that the CI server is using the **exact same version of Go** (and other tools) that you defined in your `.mise.toml` file.

---

## The Benefits

- **Peace of Mind**: You'll get a green checkmark next to your code on GitHub if all tests pass.
- **Safe Collaboration**: If you're working with others, you can be sure that their changes won't break your work (and vice versa).
- **Faster Feedback**: You don't have to wait for someone to manually test your code.

## Next Step

We've covered a lot! Let's take a quick look at some more advanced topics you'll encounter as you grow as a Go developer.

[13-advanced-topics.md &rarr;](./13-advanced-topics.md)

[&larr; Back to [TOC](../README.md#table-of-contents)]
