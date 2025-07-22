## Factory Method (Creational Pattern)

### Description:

A **Factory Method** design pattern is a creational design pattern that defines an interface for creating objects 
but lets subclasses decide which class to instantiate.

### Structure:

1. `Product`: declares the interface that all concrete products must implement.
2. `ConcreteProduct`: implements the `Product` interface, representing a specific variant of the product.
3. `Creator`: declares the factory method, which returns a `Product`. 
4. `ConcreteCreator`: a class with its onw fields that follow `Creator` methods.

### When to use?

When a class has some general processing logic, but the specific subclass can only be determined at runtime, the 
**Factory Method** pattern can be used.
