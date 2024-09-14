## Ice Cream Shop Simulator

# Overview
This Go program simulates an ice cream store where users can order different ice cream flavors with various toppings. It features user input for selecting ice cream types, adding toppings, and processing orders, including a discount on the first order. The application uses several design patterns to manage the complexity of the system.

## Features
- **Three Classic Ice Cream Flavors**: Chocolate, Creamy, and Strawberry.
- **Toppings**: Users can add toppings such as nuts, chocolate chips, fruits, and maraschino cherries.
- **First-Order Discount**: A 30% discount is applied to the user's first purchase.
- **Remaining Balance**: The user's remaining balance is displayed after each purchase.
- **Multiple Orders**: Users can purchase multiple ice creams until their balance runs out.

## Design Patterns
- **Factory Pattern**: Used to create different types of ice cream (Chocolate, Creamy, Strawberry).
- **Decorator Pattern**: Allows users to add multiple toppings to their ice cream.
- **Singleton Pattern**: Ensures the first-order discount is applied only once.
- **Facade Pattern**: Simplifies the process of ordering and purchasing ice cream by handling the interactions between the system components.

## Installation
1. Make sure you have Go installed. You can download it from [here](https://go.dev/dl/).
2. Clone the repository:
    ```bash
    git clone <repository-url>
    cd ice-cream-shop
    ```
3. Run the program:
    ```bash
    go run main.go
    ```

## How to Use
1. **Start the Program**: The user will be prompted to input the amount of money they have to spend. You need a minimum of 200 units to proceed.
   
2. **Menu Options**:
   - **Assemble Ice Cream**: Choose your base ice cream flavor and add toppings manually.
   - **Classic Ice Creams**: Choose from predefined combinations of ice cream and toppings (e.g., Chocolate Ice Cream with Nuts).
   - **Exit**: Exit the program.

3. **Ice Cream Options**:
   - **Flavors**:
     - Chocolate Ice Cream
     - Creamy Ice Cream
     - Strawberry Ice Cream
   - **Toppings**:
     - Nuts
     - Chocolate Chips
     - Fruit
     - Maraschino Cherry

4. **First-Order Discount**: A 30% discount is automatically applied on the total cost of your first order.

5. **Remaining Balance**: After each purchase, you will be shown your remaining balance and can decide whether to buy another ice cream.

