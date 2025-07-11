## Adapter (Structural Pattern)

### Context:

You have two interfaces that don’t match:
1. **Target:** the interface your code (the client) expects.
2. **Adaptee:** an existing interface that provides the functionality you need, but with a different API.

### Problem:

Your client code depends on `Target`, yet you need to reuse `Adaptee`’s functionality without changing its API.

### Solution:

Introduce an **Adapter** - a struct that:
1. **Implements** the `Target` interface.
2. Holds a reference to the `Adaptee`.
3. In each `Target` method, **delegates** the call to the corresponding `Adaptee` method.
