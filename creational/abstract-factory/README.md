## Abstract Factory (Creational Pattern)

### Description:

An **Abstract Factory** is a factory of factories. It provides a way to create related or dependent objects without
specifying their concrete classes.

### Structure:

1. `AbstractFactory`  
   Declares a set of factory methods, one for each abstract product type (e.g. `CreateButton() Button`, `CreateCheckbox() Checkbox`).
2. `ConcreteFactory1`  
   Implements the `AbstractFactory` interface to produce a specific family of concrete products (e.g. `WinButton`, `WinCheckbox`).
3. `ConcreteFactory2`  
   Another implementation of `AbstractFactory` producing an alternative product family (e.g. `MacButton`, `MacCheckbox`).
4. `AbstractProductA`  
   Declares the interface for the first type of product (e.g. `Button`).
5. `AbstractProductB`  
   Declares the interface for the second type of product (e.g. `Checkbox`).
6. `ConcreteProductA1`, `ConcreteProductA2`  
   Concrete implementations of `AbstractProductA` corresponding to each family.
7. `ConcreteProductB1`, `ConcreteProductB2`  
   Concrete implementations of `AbstractProductB` corresponding to each family.

### When to use?

When you need to create a set of related or dependent objects and want to implement this through a unified interface,
you can use the abstract factory pattern.