---
sidebar_position: 1
---

# Introduction

Zog is a schema builder for runtime value parsing and validation. Define a schema, transform a value to match, assert the shape of an existing value, or both. Zog schemas are extremely expressive and allow modeling complex, interdependent validations, or value transformations.

Killer Features:

- Concise yet expressive schema interface, equipped to model simple to complex data models
- **[Zod](https://github.com/colinhacks/zod)-like API**, use method chaining to build schemas in a typesafe manner
- **Extensible**: add your own Tests and Schemas
- **Rich errors** with detailed context, make debugging a breeze
- **Fast**: Zog is one of the fastest Go validation libraries. We are just behind the goplayground/validator for most of the [govalidbench](https://github.com/Oudwins/govalidbench/tree/master) benchmarks.
- **Built-in coercion** support for most types
- Zero dependencies!
- **Four Helper Packages**
  - **zenv**: parse environment variables
  - **zhttp**: parse http forms & query params
  - **zjson**: parse json
  - **i18n**: Opinionated solution to good i18n zog errors

> **API Stability:**
>
> - I will consider the API stable when we reach v1.0.0
> - However, I believe very little API changes will happen from the current implementation. The APIs most likely to change are the **data providers** (please don't make your own if possible use the helpers whose APIs will not change meaningfully) and the z.Ctx most other APIs should remain the same. I could be wrong but I don't expect many breaking changes.
> - Although we want to keep breaking changes to a minimum, Zog is still in version 0 and will have breaking changes in the minor versions as per semver. So please be careful when upgrading minor versions.
